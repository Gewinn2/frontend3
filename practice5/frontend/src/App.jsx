import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import AuthPage from './pages/AuthPage';
import DashboardPage from './pages/DashboardPage';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/auth" element={<AuthPage />} />
                <Route
                    path="/dashboard"
                    element={
                            <DashboardPage />
                    }
                />
                <Route path="*" element={<AuthPage />} />
            </Routes>
        </Router>
    );
}

export default App;