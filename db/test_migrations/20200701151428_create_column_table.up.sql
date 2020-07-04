CREATE TABLE IF NOT EXISTS "column"(
   id serial PRIMARY KEY,
   name VARCHAR (255) NOT NULL,
   position INTEGER NOT NULL,
   project_id INTEGER NOT NULL,
   FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
);
INSERT INTO "column" (name, position, project_id) values ('test', 1, 1);