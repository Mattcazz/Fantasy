CREATE TABLE players (
    id        SERIAL PRIMARY KEY,
    team_id   INT NOT NULL REFERENCES teams(web_id) ON DELETE CASCADE,
    name      VARCHAR(100) NOT NULL,
    position  VARCHAR(20) NOT NULL CHECK (position IN ('Goalkeeper', 'Defender', 'Midfielder', 'Attacker')),
    points    INT DEFAULT 0,
    value     FLOAT  NOT NULL,
    avg       FLOAT  NOT NULL,
    img_url   TEXT,
    web_id    INT UNIQUE NOT NULL,
    status    VARCHAR(20) NOT NULL CHECK (status IN ('ok', 'injured', 'doubt', 'sanctioned'))
);