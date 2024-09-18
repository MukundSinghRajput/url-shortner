CREATE TABLE IF NOT EXISTS tiny (
    origin text PRIMARY KEY,
    short text UNIQUE NOT NULL
);
