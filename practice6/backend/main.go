package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"` // Добавляем поле никнейма
}

type CacheItem struct {
	Data      []byte
	ExpiresAt time.Time
}

var (
	store   *session.Store
	users   = make(map[string]User)
	cache   = make(map[string]*CacheItem)
	storage = memory.New()
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	store = session.New(session.Config{
		Storage:        storage,
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
		CookieSameSite: "Lax",
		CookieHTTPOnly: true,
	})

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "N8aW4XwBBAFvpEF1s4XrW9q4dRkq6JjYwJvK7iP5tbE=",
	}))

	app.Post("/register", register)
	app.Post("/login", login)
	app.Get("/profile", authMiddleware, profile)
	app.Post("/logout", logout)
	app.Get("/data", cacheMiddleware, getData)

	app.Listen(":8080")
}

func register(c *fiber.Ctx) error {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if _, exists := users[user.Username]; exists {
		return c.Status(fiber.StatusConflict).SendString("User already exists")
	}

	if user.Nickname == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Nickname is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	users[user.Username] = User{
		Username: user.Username,
		Password: string(hashedPassword),
		Nickname: user.Nickname,
	}

	return c.SendStatus(fiber.StatusCreated)
}

func login(c *fiber.Ctx) error {
	var creds User
	if err := c.BodyParser(&creds); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, exists := users[creds.Username]
	if !exists {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sess.Set("username", user.Username)
	if err := sess.Save(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func profile(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	username := sess.Get("username").(string)
	user, exists := users[username]
	if !exists {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(fiber.Map{
		"username": user.Username,
		"nickname": user.Nickname,
	})
}

func logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := sess.Destroy(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func getData(c *fiber.Ctx) error {
	cacheDir := "./cache"
	cacheFile := filepath.Join(cacheDir, "data_cache.txt")

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка создания директории кэша")
	}

	fileInfo, err := os.Stat(cacheFile)
	if err == nil {
		if time.Since(fileInfo.ModTime()) < time.Minute {
			data, err := ioutil.ReadFile(cacheFile)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Ошибка чтения кэша")
			}
			return c.JSON(fiber.Map{
				"status": "ok",
				"data":   string(data),
			})
		}
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	newData := fmt.Sprintf("Данные созданы: %s", currentTime)

	if err := ioutil.WriteFile(cacheFile, []byte(newData), 0644); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка записи кэша")
	}

	if err = os.Chtimes(cacheFile, time.Now(), time.Now()); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка обновления времени файла")
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"data":   newData,
	})
}

func authMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if sess.Get("username") == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func cacheMiddleware(c *fiber.Ctx) error {
	cacheKey := sha256.Sum256([]byte(c.OriginalURL()))
	key := fmt.Sprintf("%x", cacheKey)

	if item, exists := cache[key]; exists && time.Now().Before(item.ExpiresAt) {
		return c.Status(fiber.StatusOK).Send(item.Data)
	}

	err := c.Next()

	if c.Response().StatusCode() == fiber.StatusOK {
		cacheDir := "./cache"
		if err := os.MkdirAll(cacheDir, 0755); err != nil {
			return err
		}

		cacheFile := filepath.Join(cacheDir, key)
		data := c.Response().Body()
		if err := ioutil.WriteFile(cacheFile, data, 0644); err != nil {
			return err
		}

		cache[key] = &CacheItem{
			Data:      data,
			ExpiresAt: time.Now().Add(1 * time.Minute),
		}
	}

	return err
}
