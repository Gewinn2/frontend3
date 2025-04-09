import React from 'react';

function Note({ note, onDelete, onEdit }) {
    return (
        <div className="note">
            <div className="note-content">
                <p>{note.text}</p>
                <small>{note.date}</small>
            </div>
            <div className="note-actions">
                <button
                    onClick={() => onEdit(note)}
                    className="edit-btn"
                >
                    Редактировать
                </button>
                <button
                    onClick={() => onDelete(note.id)}
                    className="delete-btn"
                >
                    Удалить
                </button>
            </div>
        </div>
    );
}

export default Note;