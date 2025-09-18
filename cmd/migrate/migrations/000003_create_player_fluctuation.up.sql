CREATE TABLE player_fluctuation (
    id SERIAL PRIMARY KEY,
    player_id INT NOT NULL REFERENCES player(id) ON DELETE CASCADE,
    value NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);