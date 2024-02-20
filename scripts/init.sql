CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO tasks (id, title, done) VALUES (1, 'First Task', FALSE);
INSERT INTO tasks (id, title, done) VALUES (3, 'Second Task', FALSE);
INSERT INTO tasks (id, title, done) VALUES (4, 'Third Task', FALSE);
INSERT INTO tasks (id, title, done) VALUES (5, 'Fourth Task', FALSE);
INSERT INTO tasks (id, title, done) VALUES (6, ' Task', FALSE);

