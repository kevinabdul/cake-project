-- +goose Up
insert into `cakes`(title, description, rating, image) 
values("black forest", "its a cake", 8.4, "blackforesturl"),
("peanut butter", "its a cake", 8.5, "peanutbutterurl");


-- +goose Down
delete from `cakes`;