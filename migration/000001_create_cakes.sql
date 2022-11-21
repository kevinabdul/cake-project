-- +goose Up
CREATE TABLE `cakes` (
    id int auto_increment,
    title varchar(100) not null,
    description varchar(1000) not null,
    rating float,
    image varchar(1000),
    created_at datetime default now(),
    updated_at datetime default now(),
    constraint primary key(id)
);

-- +goose Down
DROP TABLE `cakes`;