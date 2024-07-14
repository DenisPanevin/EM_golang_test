

create table IF NOT EXISTS tasks(

                                   id bigserial  primary key,
                                   name varchar not null unique

);

insert into tasks (id,name) values (0 , 'idle')