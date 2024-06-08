CREATE TABLE IF NOT EXISTS bids
(
    id         UUID PRIMARY KEY,
    item_id    UUID    NOT NULL,
    bidder_id  UUID    NOT NULL,
    auction_id UUID    NOT NULL,
    amount     INTEGER NOT NULL,
    time       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (item_id) REFERENCES arts (id),
    FOREIGN KEY (bidder_id) REFERENCES users (id),
    FOREIGN KEY (auction_Id) REFERENCES auctions (id)
);
