CREATE TABLE IF NOT EXISTS art_images
(
    id        SERIAL PRIMARY KEY,
    art_id    UUID NOT NULL,
    image_url VARCHAR(255),
    FOREIGN KEY (art_id) REFERENCES arts (id)
);