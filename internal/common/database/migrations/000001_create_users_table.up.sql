CREATE TABLE IF NOT EXISTS users
(
    id               UUID PRIMARY KEY,
    first_name       VARCHAR(255)                NOT NULL,
    last_name        VARCHAR(255)                NOT NULL,
    email            VARCHAR(255) UNIQUE         NOT NULL,
    avatar           VARCHAR(255)                NULL,
    cover            VARCHAR(255)                NULL,
    password         VARCHAR(255)                NOT NULL,
    position         VARCHAR(255)                NULL,
    location         VARCHAR(255)                NULL,
    bio              TEXT                        NULL,
    gender           VARCHAR(255)                NULL,
    currency         VARCHAR(255)                NULL,
    phone_number     VARCHAR(255)                NULL,
    address          VARCHAR(255)                NULL,
    created_at       TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at       TIMESTAMP WITHOUT TIME ZONE NOT NULL
);