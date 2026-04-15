CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);