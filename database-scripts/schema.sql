

create table if not exists person(
    id  serial,
    name varchar(500),
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);


insert into person(name) values ('José da silva'), ('Maria da silva'), ('Marcela José'), ('Marcelo oliveira');