CREATE TYPE  status AS ENUM ('1', '2', '3', '4');

create table IF NOT EXISTS tasks(

                                   id bigserial not null primary key,
                                   name varchar not null unique ,
                                   status status not null
);