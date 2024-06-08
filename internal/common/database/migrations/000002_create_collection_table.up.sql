CREATE TABLE IF NOT EXISTS collections
(
    id          UUID PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    slug        VARCHAR(255) NOT NULL,
    author_id   UUID         NOT NULL,
    description TEXT         NULL,
    FOREIGN KEY (author_id) REFERENCES users (id)
);