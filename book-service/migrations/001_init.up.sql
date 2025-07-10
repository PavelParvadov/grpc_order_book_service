create table if not exists books(
    id serial primary key,
    name varchar(100) not null,
    author varchar(50) not null
);

insert into books (name, author)
values ('Золотая рыбка', 'А.С. Пушкин'),
       ('Наруто', 'В.В.Путин');