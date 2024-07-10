create table IF NOT EXISTS users(
                                    id bigserial not null primary key,
                                    passportNumber varchar not null unique ,
                                    name varchar not null,
                                    surname varchar not null,
                                    patronymic varchar not null,
                                    address varchar not null
);

