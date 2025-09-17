CREATE TABLE player_fluctuations (
    id SERIAL PRIMARY KEY,
    player_id INT NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    value NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);