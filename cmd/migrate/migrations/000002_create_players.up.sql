CREATE TABLE player (
    id SERIAL PRIMARY KEY,
    team_id INT REFERENCES team(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    points INT DEFAULT 0,
    price FLOAT DEFAULT 0,
    avg FLOAT DEFAULT 0,
    goals int DEFAULT 0, 
    assists int DEFAULT 0, 
    img_url TEXT DEFAULT '',
    web_id int UNIQUE NOT NULL
);