CREATE TABLE IF NOT EXISTS task(
   id serial PRIMARY KEY,
   name VARCHAR (500) NOT NULL,
   description TEXT,
   position INTEGER NOT NULL,
   column_id INTEGER NOT NULL,
   FOREIGN KEY (column_id) REFERENCES "column" (id) ON DELETE CASCADE
);