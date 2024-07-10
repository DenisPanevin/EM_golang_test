create table IF NOT EXISTS jobs(

    id bigserial PRIMARY KEY not null ,
    user_id INT NOT NULL,
    task_id INT NOT NULL,
    started timestamptz,
    stopped timestamptz,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (task_id) REFERENCES tasks (id)

);

