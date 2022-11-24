DROP TABLE IF EXISTS workingDays CASCADE;
DROP TABLE IF EXISTS employee CASCADE;
DROP TABLE IF EXISTS customer CASCADE;
DROP TABLE IF EXISTS item CASCADE;
DROP TABLE IF EXISTS item_info CASCADE;
DROP TABLE IF EXISTS item_note CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS outlet CASCADE;
DROP TABLE IF EXISTS mainStock CASCADE;
DROP TABLE IF EXISTS outletStock CASCADE;
DROP TABLE IF EXISTS ordersList CASCADE;
DROP TABLE IF EXISTS orderToStock CASCADE;
DROP TABLE IF EXISTS provider CASCADE;
DROP TABLE IF EXISTS ksa CASCADE;
DROP TABLE IF EXISTS receipt CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS taxes CASCADE;
DROP TABLE IF EXISTS order_to_provider CASCADE;
DROP TABLE IF EXISTS item_category CASCADE;




CREATE TABLE employee(
                         id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                         firstName varchar(25),
                         lastName varchar(25),
                         login varchar(25),
                         passwd varchar(100),
                         phone varchar(25),
                         workTime varchar(10),
                         salary decimal,
                         PRIMARY KEY (id)
);
INSERT INTO employee (firstName,lastName,login,passwd,phone,workTime,salary)
VALUES
    ('Dexter','Roberts','Dexter','3A400F99-5646-9FBA-7445-EB9ACCF79397','(375) 371-1106','пн-пт',200),
    ('Tatiana','Bolton','Tatiana','1829DB78-7C70-6A1A-B6D7-34AE42E3DB16','(415) 473-0673','пн-пт',225),
    ('Yuri','Roth','Yuri','CC181B04-24BD-65E3-74EC-A64685A873E4','(575) 825-9877','пн-пт',250),
    ('Sheila','Padilla','Sheila','19869AC8-36D3-D9ED-6973-E5A42B469D4B','(360) 485-0822','пн-пт',275),
    ('Jackson','Curtis','Jackson','AC3D8018-B34E-9792-CA4D-D18C32D8325C','1-457-742-9850','пн-пт',300);

CREATE TABLE workingDays(
                            id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                            date date,
                            employee_id int,
                            FOREIGN KEY (employee_id)REFERENCES employee(id),
                            PRIMARY KEY (id)

);
INSERT INTO workingDays(date, employee_id)
VALUES ('Nov 3, 2021',2),
       ('Nov 3, 2021',2),
       ('Nov 7, 2021',1),
       ('Nov 3, 2021',1),
       ('Nov 3, 2021',5),
       ('Nov 2, 2021',4),
       ('Nov 7, 2021',4),
       ('Nov 5, 2021',5),
       ('Nov 7, 2021',2),
       ('Nov 2, 2021',4),
       ('Nov 7, 2021',3),
       ('Nov 6, 2021',2),
       ('Nov 4, 2021',1),
       ('Nov 2, 2021',2),
       ('Nov 3, 2021',3),
       ('Nov 6, 2021',4),
       ('Nov 1, 2021',1),
       ('Nov 7, 2021',5);



CREATE TABLE outlet(
                       id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                       address varchar(50),
                       tax decimal,
                       employee_id int,
                       FOREIGN KEY (employee_id)REFERENCES employee(id),
                       PRIMARY KEY (id)
);
INSERT INTO outlet(address, tax, employee_id)
VALUES
    ('P.O. Box 419, 8851 Vitae St.',15.0,1),
    ('8403 Proin Rd.',13.0,2),
    ('Ap #760-2926 Et Avenue',14.0,3),
    ('982-9846 Commodo Street',16.0,4),
    ('P.O. Box 204, 1339 Nascetur Rd.',10.0,5);

CREATE TABLE item_note(
                          id int GENERATED ALWAYS AS IDENTITY UNIQUE ,
                          date_of_delivery date,
                          price decimal,
                          firstTax int,
                          PRIMARY KEY (id)
);
INSERT INTO item_note(date_of_delivery, price, firstTax)
VALUES ('Mar 15, 2023',249,15),
       ('Apr 6, 2023',229,15),
       ('Jul 23, 2023',319,15),
       ('Jul 25, 2023',459,15),
       ('May 20, 2023',321,15),
       ('Dec 21, 2022',434,15),
       ('Jul 11, 2023',288,15),
       ('Apr 14, 2022',123,51),
       ('Sep 1, 2022',123,61),
       ('Sep 2, 2022',124,53);

CREATE TABLE item_category(
    id int GENERATED ALWAYS AS IDENTITY UNIQUE,
    _category VARCHAR(255)
);
INSERT INTO item_category(_category)
VALUES ('Smartphone'),
       ('Laptop'),
       ('Mouse'),
       ('Keyboard');


CREATE TABLE item_info(
                          id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                          itemname varchar(255),
                          itemimage varchar(255),
                          itemDescription text,
                          garantia int
);
INSERT INTO item_info(itemname, itemDescription, garantia,itemimage)
VALUES ('Смартфон Huawei nova 10 NCO-LX1 8GB/128GB (мерцающий серебристый)','Android, экран 6.67" OLED (1080x2400), Qualcomm Snapdragon 778G, ОЗУ 8 ГБ, флэш-память 128 ГБ, камера 50 Мп, аккумулятор 4000 мАч, 2 SIM',12,'https://content2.onliner.by/catalog/device/header/493e85704849c1901776af1c8cdecba8.jpeg'),
       ('Смартфон HONOR X7 4GB/128GB (титановый серебристый)','Android, экран 6.74" IPS (720x1600), Qualcomm Snapdragon 680, ОЗУ 4 ГБ, флэш-память 128 ГБ, карты памяти, камера 48 Мп, аккумулятор 5000 мАч, 2 SIM',12,'https://content2.onliner.by/catalog/device/header/3b307bf8affc58c8b7055f10e36d3eef.jpeg'),
       ('Смартфон Tecno Camon 19 Pro 8GB/128GB (эко черный)','Android, экран 6.8" IPS (1080x2460), Mediatek Helio G96, ОЗУ 8 ГБ, флэш-память 128 ГБ, камера 64 Мп, аккумулятор 5000 мАч, 2 SIM',12,'https://content2.onliner.by/catalog/device/header/f257877735c6afc711cdcc04878edf9b.jpeg'),
       ('Смартфон Apple iPhone 11 64GB Воcстановленный by Breezy, грейд B (фиолетовый)','Apple iOS, экран 6.1" IPS (828x1792), Apple A13 Bionic, ОЗУ 4 ГБ, флэш-память 64 ГБ, камера 12 Мп, аккумулятор 3046 мАч, 1 SIM, влагозащита IP68',12,'https://content2.onliner.by/catalog/device/header/4e5c922c2e7e7250e7ba90ba06330ab1.jpeg'),
       ('Смартфон POCO X4 GT 8GB/128GB международная версия (серебристый)','Android, экран 6.6" IPS (1080x2460), Mediatek Dimensity 8100, ОЗУ 8 ГБ, флэш-память 128 ГБ, камера 64 Мп, аккумулятор 5080 мАч, 2 SIM',12,'https://content2.onliner.by/catalog/device/header/1f52e804d859b6dc0bfc95de5dd82040.jpeg'),
       ('Смартфон Xiaomi Redmi 9A 2GB/32GB международная версия (серый)','Android, экран 6.53" IPS (720x1600), Mediatek Helio G25, ОЗУ 2 ГБ, флэш-память 32 ГБ, карты памяти, камера 13 Мп, аккумулятор 5000 мАч, 2 SIM',12,'https://content2.onliner.by/catalog/device/header/47f365ecf82722c2a2c62089016b38e1.jpeg'),
       ('Смартфон POCO M4 Pro 5G 4GB/64GB международная версия (голубой)','Android, экран 6.6" IPS (1080x2400), Mediatek Dimensity 810, ОЗУ 4 ГБ, флэш-память 64 ГБ, карты памяти, камера 50 Мп, аккумулятор 5000 мАч, 2 SIM, влагозащита IP53',12,'https://content2.onliner.by/catalog/device/header/37042e091e48e8cd17f8e21aa12e3a60.jpeg'),
       ('Смартфон POCO M4 Pro 5G 4GB/64GB международная версия (голубой)','Android, экран 6.6" IPS (1080x2400), Mediatek Dimensity 810, ОЗУ 4 ГБ, флэш-память 64 ГБ, карты памяти, камера 50 Мп, аккумулятор 5000 мАч, 2 SIM, влагозащита IP53',12,'https://content2.onliner.by/catalog/device/header/37042e091e48e8cd17f8e21aa12e3a60.jpeg'),
       ('Смартфон POCO M4 Pro 5G 4GB/64GB международная версия (голубой)','Android, экран 6.6" IPS (1080x2400), Mediatek Dimensity 810, ОЗУ 4 ГБ, флэш-память 64 ГБ, карты памяти, камера 50 Мп, аккумулятор 5000 мАч, 2 SIM, влагозащита IP53',12,'https://content2.onliner.by/catalog/device/header/37042e091e48e8cd17f8e21aa12e3a60.jpeg'),
       ('Смартфон POCO M4 Pro 5G 4GB/64GB международная версия (голубой)','Android, экран 6.6" IPS (1080x2400), Mediatek Dimensity 810, ОЗУ 4 ГБ, флэш-память 64 ГБ, карты памяти, камера 50 Мп, аккумулятор 5000 мАч, 2 SIM, влагозащита IP53',12,'https://content2.onliner.by/catalog/device/header/37042e091e48e8cd17f8e21aa12e3a60.jpeg'),
       ('Смартфон POCO M4 Pro 5G 4GB/64GB международная версия (голубой)','Android, экран 6.6" IPS (1080x2400), Mediatek Dimensity 810, ОЗУ 4 ГБ, флэш-память 64 ГБ, карты памяти, камера 50 Мп, аккумулятор 5000 мАч, 2 SIM, влагозащита IP53',12,'https://content2.onliner.by/catalog/device/header/37042e091e48e8cd17f8e21aa12e3a60.jpeg');


CREATE TABLE provider(
                         id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                         name varchar(50),
                         phone varchar(12)

);
INSERT INTO provider(name, phone)
VALUES ('Blaupunkt','+37521231234'),
       ('Витязь','+37541232232'),
       ('Xiaomi','+37524144142'),
       ('Xiaomi','+37521213222'),
       ('Xiaomi','+37521451234'),
       ('Xiaomi','+37522322234'),
       ('Haylou','+37521234234'),
       ('Canyon','+37523132134'),
       ('Realme','+37511222234'),
       ('Apple','+37521231234');

CREATE TABLE item(
                     id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                     note_id int REFERENCES item_note(id),
                     info_id int REFERENCES item_info(id),
                     provider_id int REFERENCES provider(id),
                     category_id int REFERENCES item_category(id),
                     PRIMARY KEY (id)
);
INSERT INTO item(note_id, info_id, provider_id, category_id)
VALUES   (1,1,1,1),
         (2,2,2,1),
         (3,3,3,1),
         (4,4,4,1),
         (5,5,5,1),
         (6,6,6,1),
         (7,7,7,1),
         (8,8,8,2),
         (9,9,9,2),
         (10,10,10,2);



CREATE TABLE mainStock(
                          item_id int REFERENCES item(id),
                          itemCount int
);
INSERT INTO mainStock(item_id, itemCount)
VALUES
       (8,55),
       (9,11),
       (10,12);

CREATE TABLE customer(
                         id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                         firstName varchar(25),
                         lastName varchar(25),
                         phone varchar(13)
);
INSERT INTO customer(firstName, lastName, phone)
VALUES     ('Beau','Mitchell','+37533488634'),
           ('Hilda','Carr','+37562131142'),
           ('Laurel','Maldonado','+37563796950'),
           ('Madison','Mccoy','+37569708123'),
           ('Quail','Soto','+37525941586');


CREATE TABLE taxes(
                      id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                      nds int,
                      nrt int
);
INSERT INTO taxes(nds, nrt)
VALUES (15,5);
CREATE TABLE ksa
(
    id        int GENERATED ALWAYS AS IDENTITY UNIQUE,
    outlet_id int,
    ksa_limit int,
    PRIMARY KEY (id),
    FOREIGN KEY (outlet_id) REFERENCES outlet (id)
);
INSERT INTO ksa(outlet_id, ksa_limit)
VALUES (1,500),
       (2,750),
       (3,655),
       (4,1000),
       (5,980);
CREATE TABLE orders(
                       id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                       date date,
                       price decimal,
                       cash boolean,
                       taxes int REFERENCES taxes(id),
                       item_id int REFERENCES item(id),
                       customer_id int REFERENCES customer(id),
                       employee_id int REFERENCES employee(id),
                       ksa_id int REFERENCES ksa(id),
                       PRIMARY KEY (id)
);
INSERT INTO orders(date, price,  cash, taxes, item_id, customer_id, employee_id, ksa_id)
VALUES  ('Nov 2, 2021',200,true,1,1,1,1,1),
        ('Mar 4, 2023',500,false,1,2,2,2,2),
        ('Mar 3, 2023',1000,false,1,2,2,2,2);


CREATE TABLE orderToStock(
                           id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                           date timestamp,
                           item_id int REFERENCES item(id),
                           employee_id int REFERENCES employee(id)
);
INSERT INTO orderToStock(date, item_id, employee_id)
VALUES   ('2022-11-2',3,3),
         ('2022-10-21',4,4),
         ('2022-10-22',5,5),
         ('2022-10-23',6,5),
         ('2022-10-24',7,5);


CREATE TABLE outletStock(
                            item_id int REFERENCES item(id),
                            outlet_id int REFERENCES outlet(id),
                            itemCount int
);
INSERT INTO outletStock(item_id, outlet_id, itemCount)
VALUES (1,1,12),
       (2,2,11),
       (3,3,43),
       (4,4,23),
       (5,5,22),
       (7,5,22);




CREATE TABLE receipt(
                        id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                        employee_id int REFERENCES employee(id),
                        order_id int REFERENCES orders(id),
                        customer_id int REFERENCES customer(id),
                        PRIMARY KEY (id),
                        FOREIGN KEY (employee_id) REFERENCES employee(id)
);
INSERT INTO receipt(employee_id, order_id,customer_id)
VALUES (1,1,1),
       (2,2,2);

CREATE TABLE users(
                     id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                      username varchar(255),
                      email varchar (255),
                      password_hash varchar(255)
);

CREATE TABLE order_to_provider(
    id int GENERATED ALWAYS AS IDENTITY UNIQUE,
    date time,
    item int REFERENCES item(id)
)












































