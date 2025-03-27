package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var users = make(map[string]User)

func SignUp(c *fiber.Ctx) error {
	var u User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if _, ok := users[u.Email]; ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user with this email already exists"})
	}

	u.ID = len(users) + 1

	users[u.Email] = u

	return c.Status(fiber.StatusOK).JSON(u)
}

func Login(c *fiber.Ctx) error {
	var req LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	u := users[req.Email]
	if req.Password != u.Password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong password"})
	}

	token, err := GenerateToken(u.ID, "12345678")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(LoginUserResponse{ID: u.ID, Token: token})
}

func CheckLogin(c *fiber.Ctx) error {
	id := c.Locals("id")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "you are logged in", "id": id})
}

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
}

func WithJWTAuth(c *fiber.Ctx, signingKey string) error {
	header := c.Get("Authorization")

	if header == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing auth token"})
	}

	tokenString := strings.Split(header, " ")

	if len(tokenString) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid auth header"})
	}

	id, err := ParseToken(tokenString[1], signingKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	// Записываем id в контекст, чтобы в дальнейшем использовать в других функциях
	c.Locals("id", id)
	return c.Next()
}

func GenerateToken(id int, signingKey string) (string, error) {
	claims := &tokenClaims{
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(724 * time.Hour).Unix(),
			"IssuedAr":  time.Now().Unix(),
		},
		id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}

func ParseToken(tokenString string, signingKey string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, err
	}

	if time.Now().Unix() > int64(claims.MapClaims["ExpiresAt"].(float64)) {
		return 0, errors.New("Token has expired")
	}

	return claims.UserId, nil
}

type User struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	ID int `json:"id"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}
