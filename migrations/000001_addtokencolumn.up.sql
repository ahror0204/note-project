CREATE TABLE IF NOT EXISTS note_table(
    id uuid NOT NULL PRIMARY KEY,
    title TEXT,
    body TEXT,
    exp_time VARCHAR(50),
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);