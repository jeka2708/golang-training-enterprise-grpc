--отделы
create table divisions
(
    id            serial       not null PRIMARY KEY,
    division_name varchar(100) not null
);
--должности
create table roles
(
    id          serial       not null PRIMARY KEY UNIQUE,
    name        varchar(150) not null,
    division_id int          not null,
    foreign key (division_id) references divisions (id)
);
--список работ
create table services
(
    id   serial       not null PRIMARY KEY UNIQUE,
    name varchar(255) not null,
    cost int          not null

);
--работники
create table workers
(
    id           serial      not null PRIMARY KEY UNIQUE,
    first_name   varchar(15) not null,
    last_name    varchar(15) not null,
    middle_name  varchar(15),
    phone_number varchar(13),
    role_id     int         not null,
    foreign key (role_id) references roles (id)
);
--клиенты
create table clients
(
    id           serial      not null PRIMARY KEY UNIQUE,
    first_name_c   varchar(15) not null,
    last_name_c    varchar(15) not null,
    middle_name_c  varchar(15),
    phone_number_c varchar(13)
);
--работы
create table works
(
    id         serial not null PRIMARY KEY UNIQUE,
    worker_id  int    not null,
    service_id int    not null,
    foreign key (worker_id) references workers (id),
    foreign key (service_id) references services (id)
);

create table work_clients
(
    id        serial not null PRIMARY KEY UNIQUE,
    client_id int    not null,
    work_id   int    not null,
    foreign key (client_id) references clients (id),
    foreign key (work_id) references works (id)
);

--add value for division
insert into divisions(division_name)
values ('Кадры');
insert into divisions(division_name)
values ('АСУ');
insert into divisions(division_name)
values ('Экономический');
insert into divisions(division_name)
values ('Транспортный');
insert into divisions(division_name)
values ('Ревизионный');
--add value for roles
insert into roles(name, division_id)
values ('Программист', 2);
insert into roles(name, division_id)
values ('Начальник кадров', 1);
insert into roles(name, division_id)
values ('Экономист', 3);
insert into roles(name, division_id)
values ('Водитель', 4);
insert into roles(name, division_id)
values ('Начальник отдела', 5);
--add value for services
insert into services(name, cost)
values ('Выпуск детали', 294);
insert into services(name, cost)
values ('Доставка до клиента', 344);
insert into services(name, cost)
values ('Создание ПО', 123);
insert into services(name, cost)
values ('Выпуск макета', 324);
insert into services(name, cost)
values ('Построение макета', 124);
--add value for worker
insert into workers(first_name, last_name, middle_name, phone_number, role_id)
values ('Петя', 'Иванов', 'Сергеевич', '375251234567', 1);
insert into workers(first_name, last_name, middle_name, phone_number, role_id)
values ('Вася', 'Кахов', 'Иванович', '375251234561', 2);
insert into workers(first_name, last_name, middle_name, phone_number, role_id)
values ('Иван', 'Петриков', 'Александрович', '375251234421', 3);
insert into workers(first_name, last_name, middle_name, phone_number, role_id)
values ('Дмитрий', 'Конрекевич', '', '375251265421', 4);
insert into workers(first_name, last_name, middle_name, phone_number, role_id)
values ('Александр', 'Конрекевич', '', '375251265421', 5);
--add value for clients
insert into clients(first_name_c, last_name_c, middle_name_c, phone_number_c)
values ('Дим', 'Иванович', 'Сергеевич', '375251264567');
insert into clients(first_name_c, last_name_c, middle_name_c, phone_number_c)
values ('Андрей', 'Кагортнич', 'Иванович', '375251221561');
insert into clients(first_name_c, last_name_c, middle_name_c, phone_number_c)
values ('Сергей', 'Песеков', 'Александрович', '375253234421');
insert into clients(first_name_c, last_name_c, middle_name_c, phone_number_c)
values ('Дмитрий', 'Конрекевич', '', '375253215421');
insert into clients(first_name_c, last_name_c, middle_name_c, phone_number_c)
values ('Александр', 'Конрекевич', '', '375251245421');
--add value for works
insert into works(worker_id, service_id)
values (1, 1);
insert into works(worker_id, service_id)
values (2, 2);
insert into works(worker_id, service_id)
values (3, 4);
insert into works(worker_id, service_id)
values (4, 4);
insert into works(worker_id, service_id)
values (5, 5);
--add value for clients_work
insert into work_clients(client_id, work_id)
values (1, 1);
insert into work_clients(client_id, work_id)
values (2, 2);
insert into work_clients(client_id, work_id)
values (3, 3);
insert into work_clients(client_id, work_id)
values (4, 4);
insert into work_clients(client_id, work_id)
values (5, 5);