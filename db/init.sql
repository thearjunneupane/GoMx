CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY NOT NULL,
    todo varchar(255),
    done boolean
);

-- INSERT INTO todos (todo, done) VALUES ('First Msg', false);