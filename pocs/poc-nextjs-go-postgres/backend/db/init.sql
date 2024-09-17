CREATE TYPE status_enum AS ENUM ('todo', 'doing', 'done');

CREATE TABLE IF NOT EXISTS "todo" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status status_enum DEFAULT 'todo',
    due_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

