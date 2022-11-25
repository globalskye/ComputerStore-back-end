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

/*«наличные»*/
SELECT sum(price) FROM orders WHERE cash = true;
/*«безналичные»*/
SELECT sum(price) FROM orders WHERE cash = false;


/*Заказы указанного продавца в течение текущей недели.*/
SELECT * FROM ordertostock
WHERE( date > now() - interval '7 days') AND (employee_id=3);


/*Отчет в виде графика работы конкретного продавца (дата, Ф.И.О., торговая точка) за месяц*/
SELECT date,e.lastname, e.firstname, o.address FROM workingdays
    JOIN employee e on e.id = workingdays.employee_id
    JOIN outlet o on e.id = o.employee_id
    WHERE date > now() - interval '30 days';

/*Торговый оборот по торговым точкам (сумма выручки за месяц).*/

SELECT sum(price), o.address FROM orders
    JOIN ksa k on orders.ksa_id = k.id
    JOIN outlet o on k.outlet_id = o.id
WHERE o.id=1
GROUP BY o.address;



/*Превышен лимит кассы (ksa)*/
SELECT o.address FROM orders
                          JOIN ksa k on orders.ksa_id = k.id
                          JOIN outlet o on k.outlet_id = o.id
         GROUP BY orders.id, k.ksa_limit, k.id, o.id
HAVING sum(orders.price) > k.ksa_limit;


SELECT item.id,ii.itemname,ii.image,ii.iteminfo,inote.firstprice,ii.garantia,ic.category FROM item
JOIN item_category ic on ic.id = item.category_id
JOIN item_info ii on ii.id = item.info_id
JOIN item_note inote on inote.id = item.note_id


