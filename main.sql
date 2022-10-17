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
                          note varchar(50),
                          PRIMARY KEY (id)
);
CREATE TABLE item_info(
                          id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                          itemname varchar(50),
                          iteminfo varchar(255),
                          garantia int,
                          firstPrice decimal,
                          firstTax int
);
CREATE TABLE provider(
                         id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                         name varchar(50),
                         phone varchar(12)

);

CREATE TABLE item(
                     id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                     note_id int REFERENCES item_note(id),
                     info_id int REFERENCES item_info(id),
                     provider_id int REFERENCES provider(id),
                     PRIMARY KEY (id)
);


CREATE TABLE mainStock(
                          item_id int REFERENCES item(id),
                          itemCount int
);
CREATE TABLE customer(
                         id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                         firstName varchar(25),
                         lastName varchar(25),
                         phone varchar(13)


);
CREATE TABLE taxes(
                      id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                      nds int,
                      nrt int
);
CREATE TABLE ksa
(
    id        int GENERATED ALWAYS AS IDENTITY UNIQUE,
    outlet_id int,
    ksa_limit decimal,
    PRIMARY KEY (id),
    FOREIGN KEY (outlet_id) REFERENCES outlet (id)
);
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
CREATE TABLE ordersList(
                           id int GENERATED ALWAYS AS IDENTITY UNIQUE,
                           date date,
                           item_id int REFERENCES item(id),
                           employee_id int REFERENCES employee(id)
);


CREATE TABLE outletStock(
                            item_id int REFERENCES item(id),
                            outlet_id int REFERENCES outlet(id),
                            itemCount int
);




/*insert into item_info (name, info, garantia, firstPrice, firstTax)
values (
        ('Keyboard Razer','RGB, 17 mehanical keys,',12,125.20,15),
        ('Keyboard g23 gaming','RGB, 25 mehanical keys,',12,125.20,15),
        ('Mouse Logitech','12000 dpi, superlight,',12,149.99,15),
        ('Mouse Bloody','8000 dpi, RGB,',12,125.99,15),
        ('Monitor Samsung 24`','240hz, 1920x1080, ultraRGB',24,325.44,20)
       );*/

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








































