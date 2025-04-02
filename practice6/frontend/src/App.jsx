import { useState, useEffect } from "react";
import axios from "axios";

axios.defaults.baseURL = "http://localhost:8080";
axios.defaults.withCredentials = true;

export default function App() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [nickname, setNickname] = useState("");
    const [loggedIn, setLoggedIn] = useState(false);
    const [profile, setProfile] = useState(null);
    const [theme, setTheme] = useState(
        localStorage.getItem("theme") || "light"
    );
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(false);
    const [message, setMessage] = useState("");

    useEffect(() => {
        document.body.className = theme;
        localStorage.setItem("theme", theme);
        checkAuth();
    }, [theme]);

    const checkAuth = async () => {
        try {
            const res = await axios.get("/profile");
            setProfile(res.data);
            setLoggedIn(true);
        } catch {
            setLoggedIn(false);
        }
    };

    const toggleTheme = () => {
        setTheme(prev => prev === "light" ? "dark" : "light");
    };

    const showMessage = (msg) => {
        setMessage(msg);
        setTimeout(() => setMessage(""), 3000);
    };

    const register = async () => {
        try {
            await axios.post("/register", {
                username,
                password,
                nickname
            });
            showMessage("Регистрация успешна");
            setUsername("");
            setPassword("");
            setNickname("");
        } catch {
            showMessage("Ошибка регистрации");
        }
    };

    const login = async () => {
        try {
            await axios.post("/login", { username, password });
            await checkAuth();
            showMessage("Вход выполнен");
            setUsername("");
            setPassword("");
        } catch {
            showMessage("Ошибка входа");
        }
    };

    const logout = async () => {
        try {
            await axios.post("/logout");
            setLoggedIn(false);
            setProfile(null);
            showMessage("Выход выполнен");
        } catch {
            showMessage("Ошибка выхода");
        }
    };

    const fetchData = async () => {
        setLoading(true);
        try {
            const res = await axios.get("/data");
            console.log("Ответ сервера:", res.data); // Добавляем логирование
            if (res.data.status === "ok") {
                setData(res.data.data);
                showMessage("Данные получены");
            } else {
                showMessage("Ошибка формата данных");
            }
        } catch (error) {
            console.error("Ошибка получения данных:", error);
            showMessage("Ошибка получения данных");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className={`app ${theme}`}>
            <div className="header">
                <button
                    className="theme-toggle"
                    onClick={toggleTheme}
                >
                    {theme === "light" ? "🌙 Тёмная" : "☀️ Светлая"} тема
                </button>
                {message && <div className="message">{message}</div>}
            </div>

            {!loggedIn ? (
                <div className="auth-form">
                    <input
                        className="input"
                        placeholder="Логин"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        className="input"
                        type="password"
                        placeholder="Пароль"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <input
                        className="input"
                        placeholder="Никнейм"
                        value={nickname}
                        onChange={(e) => setNickname(e.target.value)}
                    />
                    <div className="buttons">
                        <button className="btn" onClick={register}>
                            Регистрация
                        </button>
                        <button className="btn" onClick={login}>
                            Вход
                        </button>
                    </div>
                </div>
            ) : (
                <div className="content">
                    <div className="profile-info">
                        <h2>Добро пожаловать, {profile?.nickname}!</h2>
                        <p>Ваш логин: {profile?.username}</p>
                    </div>

                    <div className="buttons">
                        <button
                            className="btn"
                            onClick={fetchData}
                            disabled={loading}
                        >
                            {loading ? "Загрузка..." : "Получить данные"}
                        </button>
                        <button className="btn" onClick={logout}>
                            Выход
                        </button>
                    </div>

                    {data ? (
                        <div className="data-box">
                            <h3>Актуальные данные:</h3>
                            <p>{data}</p>
                        </div>
                    ) : (
                        <div className="data-placeholder">
                            Данные еще не загружены
                        </div>
                    )}
                </div>
            )}
        </div>
    );
}