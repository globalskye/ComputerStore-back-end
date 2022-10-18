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
DROP TABLE IF EXISTS provider CASCADE;
DROP TABLE IF EXISTS ksa CASCADE;
DROP TABLE IF EXISTS receipt CASCADE;
DROP TABLE IF EXISTS owner CASCADE;
DROP TABLE IF EXISTS taxes CASCADE;

CREATE TABLE employee(
                         id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                         firstName varchar(25),
                         lastName varchar(25),
                         login varchar(25),
                         passwd varchar(100),
                         phone varchar(25),
                         workingDays varchar(10),
                         salary decimal,
                         PRIMARY KEY (id)
);
INSERT INTO employee (firstName,lastName,login,passwd,phone,workingDays,salary)
VALUES
    ('Dexter','Roberts','Dexter','3A400F99-5646-9FBA-7445-EB9ACCF79397','(375) 371-1106','пн-пт',200),
    ('Tatiana','Bolton','Tatiana','1829DB78-7C70-6A1A-B6D7-34AE42E3DB16','(415) 473-0673','пн-пт',225),
    ('Yuri','Roth','Yuri','CC181B04-24BD-65E3-74EC-A64685A873E4','(575) 825-9877','пн-пт',250),
    ('Sheila','Padilla','Sheila','19869AC8-36D3-D9ED-6973-E5A42B469D4B','(360) 485-0822','пн-пт',275),
    ('Jackson','Curtis','Jackson','AC3D8018-B34E-9792-CA4D-D18C32D8325C','1-457-742-9850','пн-пт',300);


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
                          firstPrice decimal,
                          firstTax int,
                          PRIMARY KEY (id)
);
INSERT INTO item_note(date_of_delivery, firstPrice, firstTax)
VALUES ('Mar 15, 2023',125,15),
       ('Apr 6, 2023',324,15),
       ('Jul 23, 2023',234,15),
       ('Jul 25, 2023',244,15),
       ('May 20, 2023',321,15),
       ('Dec 21, 2022',434,15),
       ('Jul 11, 2023',233,15),
       ('Apr 14, 2022',241,51),
       ('Sep 16, 2023',123,61),
       ('May 31, 2023',124,53);



CREATE TABLE item_info(
                          id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                          itemname varchar(255),
                          iteminfo text,
                          garantia int
);
INSERT INTO item_info(itemname, iteminfo, garantia)
VALUES ('Телевизор Blaupunkt 32HB5000T','
Диагональ	32" (80 см)
Разрешение	1366x768 (HD)
Частота обновления	60 Гц
Smart TV	есть
Операционная система	Android TV (9.0)
Беспроводные интерфейсы	Bluetooth (5.0), Wi-Fi (802.11a/b/g/n)
Количество HDMI	2
Толщина панели	80 мм',12),

('Телевизор Витязь 24LH1207','
Диагональ	24" (61 см)
Разрешение	1366x768 (HD)
Smart TV	есть
Операционная система	Яндекс.ТВ (Android 11)
Беспроводные интерфейсы	Wi-Fi
Количество HDMI	2
Толщина панели	72 мм',12),

('Телевизор Xiaomi Mi TV P1 43 (L43M6-6ARG) / ELA4624GL','
Диагональ	43" (108 см)
Разрешение	3840x2160 (4K UHD)
Частота обновления	60 Гц
Расширенный динамический диапазон (HDR)	есть
Smart TV	есть
Операционная система	Android TV (10)
Беспроводные интерфейсы	Bluetooth (5.0), Wi-Fi (2,4 ГГц/5 ГГц)
Количество HDMI	3
Толщина панели	84.8',12),

('Смартфон Xiaomi Redmi 9C 4GB/128GB без NFC (серый)','
Диагональ экрана	6.53 ″
Оперативная память	4 Гб
Постоянная память	128 Гб
Версия операционной системы	Android 10.0
Дополнительный модуль камеры	есть, датчик глубины, макрообъектив
Разрешение камеры	13 Мп
Кол-во SIM-карт	2
Емкость аккумулятора	5000 мАч
Процесс зарядки	быстрая зарядка, стандартная зарядка',12),

('Смартфон Xiaomi Redmi 9A 2GB/32GB (серый)','
Диагональ экрана	6.53 ″
Оперативная память	2 Гб
Постоянная память	32 Гб
Версия операционной системы	Android 10.0
Разрешение камеры	13 Мп
Кол-во SIM-карт	2
Емкость аккумулятора	5000 мАч
Процесс зарядки	быстрая зарядка, стандартная зарядка',12),

('Смартфон Xiaomi Redmi 9A 2GB/32GB (зеленый)','
Диагональ экрана	6.53 ″
Оперативная память	2 Гб
Постоянная память	32 Гб
Версия операционной системы	Android 10.0
Разрешение камеры	13 Мп
Кол-во SIM-карт	2
Емкость аккумулятора	5000 мАч
Процесс зарядки	быстрая зарядка, стандартная зарядка',12),

('Умные часы Haylou LS02 (черный)','
Вид часов	фитнес-часы
Вибросигнал	есть
Размер экрана	1.4 ″
Класс защиты	IP68
Время автономной работы	480 ч
Дистанционное управление	музыка',6),

('Умные часы Canyon CNS-SW74BL','
Вид часов	часы-компаньон
Вибросигнал	есть
Функции	управление воспроизведением, фонарик
Размер экрана	1.3 ″
Класс защиты	IP67
Время автономной работы	120 ч (5 дней активной работы)
Дистанционное управление	камера смартфона, музыка',6),

('Умные часы Realme Watch 2 / RMW2008 (черный)','
Вид часов	фитнес-часы
Вибросигнал	есть
Размер экрана	1.4 ″ (3.5)
Класс защиты	IP68
Время автономной работы	288 ч (теоретическое время работы 12 дней)
Дистанционное управление	камера смартфона',12),

('Умные часы Apple Watch Series 7 GPS 41mm / MKMY3 (алюминий белый/сияющая звезда)
ПРЕМИУМ','
Вид часов	часы-компаньон
Программная платформа	Watch OS
Вибросигнал	есть
Функции	управление воспроизведением, экстренный вызов SOS
Размер экрана	1.69 ″
Влагозащита	50 м
Класс защиты	IP6X
Время автономной работы	18 ч
Дистанционное управление	музыка',12);

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
                     PRIMARY KEY (id)
);
INSERT INTO item(note_id, info_id, provider_id)
VALUES   (1,1,1),
         (2,2,2),
         (3,3,3),
         (4,4,4),
         (5,5,5),
         (6,6,6),
         (7,7,7),
         (8,8,8),
         (9,9,9),
         (10,10,10);



CREATE TABLE mainStock(
                          item_id int REFERENCES item(id),
                          itemCount int
);
INSERT INTO mainStock(item_id, itemCount)
VALUES (1,44),
       (2,22),
       (3,32),
       (4,42),
       (5,12),
       (6,414),
       (7,23),
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
    ksa_limit decimal,
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
INSERT INTO orders(date,  cash, taxes, item_id, customer_id, employee_id, ksa_id)
VALUES  ('Nov 2, 2021',true,1,1,1,1,1),
        ('Mar 4, 2023',false,1,2,2,2,2);

CREATE TABLE orderToStock(
                           id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                           date date,
                           item_id int REFERENCES item(id),
                           employee_id int REFERENCES employee(id)
);
INSERT INTO orderToStock(date, item_id, employee_id)
VALUES   ('Dec 8, 2021',3,3),
         ('Jun 9, 2022',4,4),
         ('Jun 23, 2023',5,5),
         ('Sep 16, 2022',6,5),
         ('Jun 20, 2023',7,5);


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
                        employee_id int,
                        order_id int REFERENCES orders(id),
                        PRIMARY KEY (id),
                        FOREIGN KEY (employee_id) REFERENCES employee(id)
);

CREATE TABLE owner(
                      username varchar(50),
                      passwd varchar(50)

);

/*INSERT INTO taxes (nds, nrt) VALUES (20,13);*/
/*
SELECT e.firstName, e.lastName FROM outlet
JOIN employee e on e.id = outlet.employee_id;



*/








































