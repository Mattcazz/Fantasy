CREATE TABLE player (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    nationality TEXT NOT NULL,
    position TEXT NOT NULL,
    points INT DEFAULT 0,
    price NUMERIC(10,2) DEFAULT 0,
    img_url TEXT DEFAULT '',
    team_id INT REFERENCES team(id) ON DELETE SET NULL
);