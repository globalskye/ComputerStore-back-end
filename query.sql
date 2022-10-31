/*Вывести перечень товаров, находящихся в указанном торговом зале*/
SELECT ii.itemname,ii.iteminfo,ii.garantia , itemcount from outletstock
JOIN item i on i.id = outletstock.item_id
JOIN item_info ii on ii.id = i.info_id;

/*Вывести состояние центрального склада фирмы по состоянию на заданное число.*/
SELECT item_id, itemcount  from mainstock
JOIN item i on i.id = mainstock.item_id
JOIN item_note itn on i.note_id = itn.id
WHERE itn.date_of_delivery < now() ;/*TODO передавать дату через API c фронта*/

/*Вывести «общий» склад фирмы.*/
SELECT mainstock.item_id,mainstock.itemcount, o.item_id, o.itemcount  FROM mainstock
JOIN outletstock o on mainstock.item_id = o.item_id;

/*Перечень сделок, оформленных конкретным продавцом с делением на «розницу» и «безнал».*/
SELECT * FROM orders WHERE employee_id = 1 and cash = true; /*false = безнал*/

/*Отчет о продажах продавцов (общая сумма сделок) за заданный период времени.*/
SELECT employee_id, sum(price) FROM orders  /*TODO передавать дату через API с фронта*/
WHERE  orders.date BETWEEN 'Nov 2, 2020' and 'Nov 2, 2026'
GROUP BY employee_id;

/*Вывести суммарный доход фирмы за указанный период времени, отдельно учитывая «наличные» и
«безналичные» сделки.*/

SELECT sum(price) FROM orders WHERE cash = true;

SELECT sum(price) FROM orders WHERE cash = false;


SELECT  outlet_id ,outletstock.item_id FROM outletstock
LEFT  JOIN orders o on outletstock.item_id = o.item_id
WHERE (o.date BETWEEN 'Nov 2, 2020' AND 'Nov 2, 2026') AND (o.item_id IS NULL);

SELECT * FROM ordertostock
WHERE( date BETWEEN NOW()::DATE-EXTRACT(DOW FROM NOW())::INTEGER-7
          AND NOW()::DATE-EXTRACT(DOW from NOW())::INTEGER
)AND (employee_id=3);


SELECT * FROM orders