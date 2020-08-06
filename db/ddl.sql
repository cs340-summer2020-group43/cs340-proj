create table Products (
    `id` varchar(20) not null,
    `product_class` varchar(10) not null,
    `type` varchar(10) not null,
    `cell` int not null,
    primary key (`id`),
    foreign key (`cell`) references Cells (`id`)
);

create table Cells (
    `id` int not null,
    `description` varchar(20),
    primary key (`id`)
);

create table Errors (
    `id` int not null,
    `description` varchar(255),
    primary key (`id`)
);

create table Testers (
    `id` int not null,
    `description` varchar(255),
    `cell` int not null,
    primary key (`id`),
    foreign key (`cell`) references Cells (`id`)
);

create table Users (
    `id` int not null auto_increment,
    `name` varchar(255) not null,
    primary key (`id`)
);

create table Tests (
    `id` int not null auto_increment,
    `product_id` varchar(20) not null,
    `serial_num` varchar(31) not null,
    `cell` int not null,
    `start_time` datetime not null,
    `end_time` datetime not null,
    `error_time` datetime,
    `error_type` int,
    `result_1` float,
    `result_2` float,
    `result_3` float,
    `result_4` float,
    `result_5` float,
    `result_6` float,
    `pass` bool not null,
    `tester` int not null,
    `user` int not null,
    primary key (`id`),
    foreign key (`product_id`) references Products (`id`),
    foreign key (`cell`) references Cells (`id`),
    foreign key (`error_type`) references Errors (`id`),
    foreign key (`tester`) references Testers (`id`),
    foreign key (`user`) references Users (`id`)
);

create table Testers_Users (
    `tester` int not null,
    `user` int not null,
    primary key (`tester`, `user`),
    foreign key (`tester`) references Testers (`id`),
    foreign key (`user`) references Users (`id`)
);

create table Testers_Products (
    `tester` int not null,
    `product` varchar(20) not null,
    primary key (`tester`, `product`),
    foreign key (`tester`) references Testers (`id`),
    foreign key (`product`) references Products (`id`)
);
