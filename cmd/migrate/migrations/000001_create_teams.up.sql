CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    logo_url TEXT NOT NULL,
    web_id int UNIQUE NOT NULL 
);