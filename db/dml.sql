-- Each ':' character denotes the start of a variable, which is followed by a
-- space and a type description

-- adding
insert into Products (`id`, `product_class`, `type`, `cell`) values(
    :id string,
    :product_class string,
    :type string,
    :cell int
);

insert into Cells (`id`, `description`) values(
    :id int,
    :description string
);

insert into Errors (`id`, `description`) values(
    :id int,
    :description string
);

insert into Testers (`id`, `description`, `cell`, `product_id`) values(
    :id int,
    :description string,
    :cell int,
    :product_id string
);

insert into Users (`name`) values(
    :name string
);

insert into Tests (
    `product_id`,
    `serial_num`,
    `cell`,
    `start_time`,
    `end_time`,
    `error_time`,
    `error_type`,
    `result_1`,
    `result_2`,
    `result_3`,
    `result_4`,
    `result_5`,
    `result_6`,
    `pass`,
    `tester`,
    `user`
) values(
    :product_id string,
    :serial_num string,
    :cell int,
    :start_time string,
    :end_time string,
    :error_time string,
    :error_type int,
    :result_1 float,
    :result_2 float,
    :result_3 float,
    :result_4 float,
    :result_5 float,
    :result_6 float,
    :pass int,
    :tester int,
    :user int
);

insert into Testers_Users (`tester`, `user`) values(
    :tester int,
    :user int
);

insert into Testers_Products (`tester`, `product`) values(
    :tester int,
    :product string
);

-- updating
update Products
set
    `id` = :id,
    `product_class` = :product_class,
    `type` = :type,
    `cell` = :cell,
where
    `id` = :id;

update Cells
set
    `id` = :id,
    `description` = :description,
where
    `id` = :id;

update Errors
set
    `id` = :id,
    `description` = :description,
where
    `id` = :id;

update Testers
set
    `id` = :id,
    `description` = :description,
    `cell` = :cell,
    `product_id` = :product_id,
where
    `id` = :id;

update Users
set
    `id` = :id,
    `name` = :name,
where
    `id` = :id;

update Tests
set
    `id` = :id,
    `product_id` = :product_id,
    `serial_num` = :serial_num,
    `cell` = :cell,
    `start_time` = :start_time,
    `end_time` = :end_time,
    `error_time` = :error_time,
    `error_type` = :error_type,
    `result_1` = :result_1,
    `result_2` = :result_2,
    `result_3` = :result_3,
    `result_4` = :result_4,
    `result_5` = :result_5,
    `result_6` = :result_6,
    `pass` = :pass,
    `tester` = :tester,
    `user` = :user,
where
    `id` = :id;

-- showing
select * from Products;

select * from Cells;

select * from Errors;

select * from Testers;

select * from Users;

select * from Tests;

select * from Testers_Products;  -- potentially change this to be more readable


-- delete
delete from Tests where `id` = :id int;
delete from Testers_Users where `tester` = :tester and `user` = :user int;


-- search only for failed tests
select * from Tests where `pass` = false;
