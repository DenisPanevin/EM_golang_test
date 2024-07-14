CREATE TABLE IF NOT EXISTS jobs (
                                    id bigserial PRIMARY KEY NOT NULL,
                                    user_id INTEGER NOT NULL,
                                    task_id INTEGER NOT NULL,
                                    started timestamptz,
                                    stopped timestamptz,

                                    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                                    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);
