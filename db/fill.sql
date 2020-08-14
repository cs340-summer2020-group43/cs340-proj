insert into Cells (`id`, `description`) values
(1, "H800"),
(2, "H500"),
(3, "PX3");

insert into Errors (`id`,`description`) values
(1, "The first kind of error"), 
(2, "Failed Post Test"),
(3, "Mensor Set Error");

insert into Users (`name`) values
("John Cena"),
("Scartlett Johnson"),
("Natalie Portman");

insert into Products (`id`, `product_class`, `type`, `cell`) values
("CDLSHTD1", "CDL", "1CAL", 1),
("CW2LA2D", "CW2", "1C", 2),
("LVCT00102D", "S3", "S3", 3);

insert into Testers (`id`, `description`, `cell`) values
(1, "Communication Board Tester EM000041", 1),
(2, "TX/TR Tester EM000054", 2),
(3, "Hx21/Hx21SP Tester EM000478", 3);

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
) values
(
    "CDLSHTD1",
    "S3434H5566",
    1,
    "2020-07-26 14:55:00",
    "2020-07-26 14:58:00",
    "2020-07-26 14:58:00",
    1,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    FALSE,
    1,
    2
),
(
    "CW2LA2D",
    "h3555YY373",
    2,
    "2020-07-25 11:55:00",
    "2020-07-26 11:58:00",
    NULL,
    NULL,
    3.33,
    4.56,
    6.77,
    101.22,
    134.55,
    155.66,
    TRUE,
    2,
    1
);

insert into Testers_Users (`tester`, `user`) values
(1, 1),
(1, 2),
(1, 3),
(2, 1),
(2, 2),
(3, 1);

insert into Testers_Products (`tester`, `product`) values
(1,"CDLSHTD1"),
(2,"CW2LA2D"),
(3,"LVCT00102D");

-- showing
select * from Products;

select * from Cells;

select * from Errors;

select * from Testers;

select * from Users;

select * from Tests;

select * from Testers_Products;  -- potentially change this to be more readable

-- search only for failed tests
select * from Tests where `pass` = true;
