CREATE TABLE IF NOT EXISTS cars (
    id serial not null,
    reg_number varchar(45) not null,
    mark varchar(45) not null,
    model varchar(45) not null,
    year int not null,
    owner_id int not null,
    PRIMARY KEY(id)
);