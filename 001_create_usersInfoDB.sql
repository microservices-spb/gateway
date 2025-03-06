CREATE TABLE IF NOT EXISTS usersInfo (
    id SERIAL PRIMARY KEY,
    "username" TEXT UNIQUE,
    "password" TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW () 
);

DROP TABLE IF EXISTS usersInfo;