CREATE TABLE IF NOT EXISTS comment(
   id serial PRIMARY KEY,
   text TEXT NOT NULL,
   created_at TIMESTAMP NOT NULL,
   task_id INTEGER NOT NULL,
   FOREIGN KEY (task_id) REFERENCES task (id) ON DELETE CASCADE
);
INSERT INTO comment (text, created_at, task_id) values ('test', current_timestamp, 1);