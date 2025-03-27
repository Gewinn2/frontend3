import React, { useState } from 'react';
import {
    Container,
    Paper,
    Typography,
    TextField,
    Button,
    Box,
    Tabs,
    Tab
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import api from '../api';

const AuthPage = () => {
    const [tabValue, setTabValue] = useState(0);
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleTabChange = (event, newValue) => {
        setTabValue(newValue);
        setError('');
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        try {
            const url = tabValue === 0 ? '/login' : '/signup';
            const response = await api.post(url, { email, password });

            if (tabValue === 0) {
                localStorage.setItem('token', response.data.token);
                navigate('/dashboard');
            } else {
                alert('Регистрация прошла успешно! Пожалуйста, войдите.');
                setTabValue(0);
            }
        } catch (err) {
            setError(err.response?.data?.error || 'Произошла ошибка');
        }
    };

    return (
        <Box
            sx={{
                minHeight: '100vh',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                backgroundColor: '#f5f5f5',
                p: 2
            }}
        >
            <Paper elevation={3} sx={{ p: 4, width: 350 }}>
                <Tabs
                    value={tabValue}
                    onChange={handleTabChange}
                    variant="fullWidth"
                    textColor="primary"
                    indicatorColor="primary"
                >
                    <Tab label="Вход" />
                    <Tab label="Регистрация" />
                </Tabs>

                <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
                    <TextField
                        label="Email"
                        fullWidth
                        margin="normal"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <TextField
                        label="Пароль"
                        type="password"
                        fullWidth
                        margin="normal"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    {error && (
                        <Typography color="error" variant="body2" sx={{ mt: 1 }}>
                            {error}
                        </Typography>
                    )}
                    <Button
                        type="submit"
                        variant="contained"
                        fullWidth
                        sx={{ mt: 2, backgroundColor: '#1976d2', color: '#fff' }}
                    >
                        {tabValue === 0 ? 'Войти' : 'Зарегистрироваться'}
                    </Button>
                </Box>
            </Paper>
        </Box>
    );
};

export default AuthPage;
