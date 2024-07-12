

create table IF NOT EXISTS tasks(

                                   id bigserial not null primary key,
                                   name varchar not null unique

);