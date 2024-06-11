CREATE TABLE IF NOT EXISTS auctions
(
    id            UUID PRIMARY KEY,
    item_id       UUID                     NOT NULL,
    type          SMALLINT                 NOT NULL,
    initial_price INTEGER                  NOT NULL,
    current_price INTEGER                  NOT NULL,
    status        SMALLINT                 NOT NULL,
    start_date    TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date      TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (item_id) REFERENCES arts (id)
);