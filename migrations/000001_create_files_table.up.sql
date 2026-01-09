CREATE TABLE files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    encrypted INTEGER NOT NULL, -- 1/0 boolean value
    size_bytes INTEGER NOT NULL,
    alias TEXT NOT NULL UNIQUE,
    expires TIMESTAMP,
    highlighting_language TEXT,

    storage_provider TEXT NOT NULL,
    storage_key      TEXT NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
