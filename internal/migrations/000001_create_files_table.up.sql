CREATE TABLE files (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    encrypted INTEGER NOT NULL, -- 1/0 boolean value
    size INTEGER NOT NULL,
    server_path TEXT NOT NULL,
    alias TEXT NOT NULL,
    domain TEXT NOT NULL,
    expires TIMESTAMP,
    display_as TEXT NOT NULL,
    highlighting_language TEXT,
    created_at TIMESTAMP NOT NULL
);
