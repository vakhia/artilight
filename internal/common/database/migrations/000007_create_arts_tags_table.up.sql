CREATE TABLE IF NOT EXISTS arts_tags
(
    art_id UUID NOT NULL,
    tag_id INT  NOT NULL,
    PRIMARY KEY (art_id, tag_id),
    FOREIGN KEY (art_id) REFERENCES arts (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
)