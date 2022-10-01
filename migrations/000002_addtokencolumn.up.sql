CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY, 
    first_name VARCHAR(50), 
    last_name VARCHAR(50), 
    email VARCHAR(50), 
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

ALTER TABLE note_table
    ADD COLUMN user_id UUID;

ALTER TABLE note_table
    ADD CONSTRAINT usersfk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;