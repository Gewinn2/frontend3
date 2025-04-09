import React, { useState, useEffect } from 'react';
import NoteList from './components/NoteList';
import './App.css';

function App() {
    const [notes, setNotes] = useState([]);
    const [inputText, setInputText] = useState('');
    const [isOnline, setIsOnline] = useState(navigator.onLine);
    const [currentNoteId, setCurrentNoteId] = useState(null);
    const [installPrompt, setInstallPrompt] = useState(null);

    // Загрузка заметок и настройка онлайн-статуса
    useEffect(() => {
        const savedNotes = localStorage.getItem('notes');
        if (savedNotes) setNotes(JSON.parse(savedNotes));

        const handleStatusChange = () => {
            setIsOnline(navigator.onLine);
            if (navigator.onLine) {
                console.log('Back online, checking for updates...');
            }
        };

        window.addEventListener('online', handleStatusChange);
        window.addEventListener('offline', handleStatusChange);

        return () => {
            window.removeEventListener('online', handleStatusChange);
            window.removeEventListener('offline', handleStatusChange);
        };
    }, []);

    // Сохранение заметок
    useEffect(() => {
        localStorage.setItem('notes', JSON.stringify(notes));
    }, [notes]);

    const handleAddNote = () => {
        if (inputText.trim()) {
            setNotes([...notes, {
                id: Date.now(),
                text: inputText,
                date: new Date().toLocaleString()
            }]);
            setInputText('');
        }
    };

    const handleUpdateNote = () => {
        if (currentNoteId && inputText.trim()) {
            setNotes(notes.map(note =>
                note.id === currentNoteId
                    ? { ...note, text: inputText }
                    : note
            ));
            setCurrentNoteId(null);
            setInputText('');
        }
    };

    const handleDeleteNote = (id) => {
        setNotes(notes.filter(note => note.id !== id));
    };

    const handleEditClick = (note) => {
        setInputText(note.text);
        setCurrentNoteId(note.id);
    };

    return (
        <div className="app">
            <h1>Мои Заметки</h1>
            {!isOnline && <div className="offline-banner">Офлайн-режим</div>}

            <div className="note-input">
                <textarea
                    value={inputText}
                    onChange={(e) => setInputText(e.target.value)}
                    placeholder="Введите текст заметки..."
                />
                {currentNoteId ? (
                    <>
                        <button onClick={handleUpdateNote}>Обновить</button>
                        <button
                            onClick={() => {
                                setCurrentNoteId(null);
                                setInputText('');
                            }}
                            className="cancel-btn"
                        >
                            Отмена
                        </button>
                    </>
                ) : (
                    <button onClick={handleAddNote}>Добавить</button>
                )}
            </div>

            <NoteList
                notes={notes}
                onDelete={handleDeleteNote}
                onEdit={handleEditClick}
            />
        </div>
    );
}

export default App;