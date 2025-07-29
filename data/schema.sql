--temporary for testing (will update in future)
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

INSERT INTO users (name) VALUES ('Alice');
INSERT INTO users (name) VALUES ('Bob');