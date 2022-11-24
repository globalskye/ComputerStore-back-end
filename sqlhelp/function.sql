DROP FUNCTION IF EXISTS GetPassword(userId integer);
DROP FUNCTION IF EXISTS GetPriceWithNds(orderId integer);

CREATE FUNCTION GetPassword(userId integer)
    RETURNS varchar(255)
AS $$BEGIN
    RETURN( SELECT password_hash FROM users
            WHERE users.id = userId
    );
end;
$$
LANGUAGE 'plpgsql';

SELECT GetPassword(1);


CREATE FUNCTION GetPriceWithNds(orderId integer)
    RETURNS decimal
AS $$
    BEGIN
        DECLARE price decimal;
        DECLARE Tax decimal;
        DECLARE lastPrice decimal;
    BEGIN
            SELECT orders.price FROM orders
            WHERE orders.id = orderId
            INTO price;
            SELECT taxes.nds FROM taxes
            INTO Tax;
            lastPrice := price- price%Tax;
            RETURN (lastPrice);

        end;
end;
$$
LANGUAGE 'plpgsql';

SELECT GetPriceWithNds(1);



