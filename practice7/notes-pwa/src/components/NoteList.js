import React from 'react';
import Note from './Note';

function NoteList({ notes, onDelete, onEdit }) {
    return (
        <div className="note-list">
            {notes.length === 0 ? (
                <p>Нет заметок</p>
            ) : (
                notes.map(note => (
                    <Note
                        key={note.id}
                        note={note}
                        onDelete={onDelete}
                        onEdit={onEdit}
                    />
                ))
            )}
        </div>
    );
}

export default NoteList;