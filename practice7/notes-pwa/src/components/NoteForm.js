import React, { useState, useEffect } from 'react';

function NoteForm({ onAdd, onUpdate, editingNote }) {
    const [text, setText] = useState('');

    useEffect(() => {
        if (editingNote) {
            setText(editingNote.text);
        }
    }, [editingNote]);

    const handleSubmit = (e) => {
        e.preventDefault();
        if (text.trim()) {
            if (editingNote) {
                onUpdate({ ...editingNote, text });
            } else {
                onAdd(text);
            }
            setText('');
        }
    };

    return (
        <form onSubmit={handleSubmit} className="note-form">
            <input
                type="text"
                value={text}
                onChange={(e) => setText(e.target.value)}
                placeholder="Введите заметку..."
                required
            />
            <button type="submit">
                {editingNote ? 'Обновить' : 'Добавить'}
            </button>
            {editingNote && (
                <button
                    type="button"
                    onClick={() => {
                        setText('');
                        onUpdate(null);
                    }}
                    className="cancel-btn"
                >
                    Отмена
                </button>
            )}
        </form>
    );
}

export default NoteForm;