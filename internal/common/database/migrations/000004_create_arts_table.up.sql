CREATE TABLE IF NOT EXISTS arts
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug          VARCHAR(255) UNIQUE        NOT NULL,
    title         VARCHAR(255)               NOT NULL,
    type          SMALLINT         DEFAULT 0 NULL,
    description   TEXT,
    price         NUMERIC(10, 2)             NOT NULL,
    min_bid       NUMERIC(10, 2)             NOT NULL,
    status        SMALLINT         DEFAULT 0 NULL,
    owner_id      UUID                       NOT NULL,
    category_id   INT                        NOT NULL,
    collection_id UUID                       NOT NULL,
    FOREIGN KEY (collection_id) REFERENCES collections (id),
    FOREIGN KEY (owner_id) REFERENCES users (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);