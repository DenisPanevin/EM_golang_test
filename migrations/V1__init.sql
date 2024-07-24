create table IF NOT EXISTS users(
                                    id bigserial not null primary key,
                                    passportNumber varchar not null unique ,
                                    name varchar not null,
                                    surname varchar not null,
                                    patronymic varchar not null,
                                    address varchar not null
);
create table IF NOT EXISTS tasks(

                                    id bigserial  primary key,
                                    name varchar not null unique

);

insert into tasks (id,name) values (0 , 'idle');
CREATE TABLE IF NOT EXISTS jobs (
                                    id bigserial PRIMARY KEY NOT NULL,
                                    user_id INTEGER NOT NULL,
                                    task_id INTEGER NOT NULL,
                                    started timestamptz,
                                    stopped timestamptz,

                                    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                                    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);
