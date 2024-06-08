CREATE TABLE IF NOT EXISTS tags
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT
);