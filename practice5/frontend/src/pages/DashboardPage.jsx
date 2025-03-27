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
    const navigate = useNavigate();

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    navigate('/auth');
                    return;
                }

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
                <Box sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    mb: 4,
                    width: '100%'
                }}>
                    <Typography variant="h4">Добро пожаловать</Typography>
                    <Button
                        variant="contained"
                        color="error"
                        onClick={handleLogout}
                    >
                        Выйти
                    </Button>
                </Box>

                <Paper elevation={3} sx={{ p: 4, width: '100%' }}>
                    <Typography variant="h6" gutterBottom>
                        Информация об аккаунте
                    </Typography>
                    <Typography>ID пользователя: {userData.id}</Typography>
                    <Typography sx={{ mt: 2 }}>
                        Вы успешно авторизованы!
                    </Typography>
                </Paper>

                {error && (
                    <Typography color="error" sx={{ mt: 2, textAlign: 'center' }}>
                        {error}
                    </Typography>
                )}
            </Container>
        </Box>
    );
};

export default DashboardPage;