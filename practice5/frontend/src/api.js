import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080', // Указываем адрес вашего бэкенда
});

// Добавляем интерцептор для JWT
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
}, (error) => {
    return Promise.reject(error);
});

export default api;