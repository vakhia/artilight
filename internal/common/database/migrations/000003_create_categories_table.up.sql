CREATE TABLE IF NOT EXISTS categories
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT
);