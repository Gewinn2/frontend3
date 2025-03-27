import React, { useEffect, useState } from 'react';
import {
    Container,
    Typography,
    Button,
    Box,
    Paper,
    CircularProgress
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import api from '../api';

const DashboardPage = () => {
    const [userData, setUserData] = useState(null);
    const [error, setError] = useState('');
    const [token, setToken] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const storedToken = localStorage.getItem('token');
                if (!storedToken) {
                    navigate('/auth');
                    return;
                }

                setToken(storedToken); // сохраняем токен

                const response = await api.get('/auth/check');
                setUserData(response.data);
            } catch (err) {
                localStorage.removeItem('token');
                navigate('/auth');
                setError(err.response?.data?.error || 'Сессия истекла');
            }
        };

        checkAuth();
    }, [navigate]);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/auth');
    };

    if (!userData) {
        return (
            <Box
                sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                    minHeight: '100vh',
                    backgroundColor: '#fff'
                }}
            >
                <CircularProgress />
            </Box>
        );
    }

    return (
        <Box
            sx={{
                minHeight: '100vh',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                backgroundColor: '#fff',
                p: 3
            }}
        >
            <Container maxWidth="md">
                <Paper elevation={3} sx={{ p: 4 }}>
                    <Typography variant="h5" gutterBottom>
                        Добро пожаловать!
                    </Typography>
                    <Typography variant="body1" sx={{ mb: 2 }}>
                        <strong>ID пользователя:</strong> {userData.id}
                    </Typography>
                    <Typography variant="body2" sx={{ wordBreak: 'break-all', mb: 2 }}>
                        <strong>JWT токен:</strong> {token}
                    </Typography>
                    <Button variant="contained" color="primary" onClick={handleLogout}>
                        Выйти
                    </Button>
                </Paper>
            </Container>
        </Box>
    );
};

export default DashboardPage;
