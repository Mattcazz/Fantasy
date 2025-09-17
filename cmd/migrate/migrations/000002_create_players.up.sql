CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    nationality TEXT NOT NULL,
    position TEXT NOT NULL,
    points INT DEFAULT 0,
    price NUMERIC(10,2),
    fluctuation NUMERIC(10,2),
    img_url TEXT,
    team_id INT REFERENCES teams(id) ON DELETE SET NULL
);