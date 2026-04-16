CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE works (
   id UUID PRIMARY KEY,
   type TEXT NOT NULL,
   status TEXT NOT NULL,
   conversation_id UUID NULL,
   input JSONB NOT NULL,
   output JSONB NULL,
   error_message TEXT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   completed_at TIMESTAMPTZ NULL,
   CONSTRAINT fk_works_conversation
       FOREIGN KEY (conversation_id)
       REFERENCES conversations(id)
       ON DELETE SET NULL
);

CREATE INDEX idx_works_status ON works(status);
CREATE INDEX idx_works_created_at ON works(created_at DESC);
CREATE INDEX idx_works_conversation_id ON works(conversation_id);
