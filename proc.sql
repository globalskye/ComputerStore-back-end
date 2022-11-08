/*DROP FUNCTION IF EXISTS DeleteEmployee(emp_id int);
DROP TRIGGER IF EXISTS BadEmployeeWhoDontWorking ON employee;

CREATE FUNCTION DeleteEmployee(emp_id int)
    RETURNS trigger
AS $$
    BEGIN

DECLARE days int;

BEGIN
    SELECT count(date) INTO days FROM workingdays WHERE employee_id=emp_id;

    IF row(NEW.*) IS DISTINCT FROM row(OLD.*)

    THEN DELETE FROM employee WHERE id = emp_id;
        RETURN new;

    ELSE
        RETURN old;
    END IF;
    end;
    end;

$$
LANGUAGE 'plpgsql';

SELECT * FROM DeleteEmployee(1);



CREATE TRIGGER BadEmployeeWhoDontWorking
    BEFORE DELETE ON employee
    FOR EACH ROW
    EXECUTE FUNCTION DeleteEmployee(id);



CREATE OR REPLACE FUNCTION keep_create_update_info_correct()
    RETURNS TRIGGER AS $$
BEGIN
    IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
        NEW.updated_at = now();
        NEW.created_at := OLD.created_at;
        RETURN NEW;
    ELSE
        RETURN OLD;
    END IF;
END;
$$ language 'plpgsql';
*/






DROP TRIGGER IF EXISTS emp_stamp on employee;
DROP TRIGGER IF EXISTS emp_audit_del on employee;
DROP TRIGGER IF EXISTS emp_audit_upd on employee;
DROP TRIGGER IF EXISTS emp_audit_ins on employee;
DROP FUNCTION IF EXISTS emp_stamp();
DROP FUNCTION IF EXISTS process_emp_audit();



CREATE FUNCTION emp_stamp() RETURNS trigger AS $emp_stamp$
BEGIN

    IF NEW.firstname IS NULL THEN
        RAISE EXCEPTION 'employee name  cannot be null';
    END IF;
    IF NEW.salary IS NULL THEN
        RAISE EXCEPTION '% cannot have null salary', NEW.firstname;
    END IF;


    IF NEW.salary < 0 THEN
        RAISE EXCEPTION '% cannot have a - salary', NEW.firstname;
    END IF;


    RETURN NEW;
END;
$emp_stamp$ LANGUAGE plpgsql;

CREATE TRIGGER emp_stamp BEFORE INSERT OR UPDATE ON employee
    FOR EACH ROW EXECUTE FUNCTION emp_stamp();

insert into employee(firstname, salary) VALUES
                            ('david',-5);




CREATE TABLE emp_audit(
                          operation         char(1)   NOT NULL,
                          stamp             timestamp NOT NULL,
                          userid            text      NOT NULL,
                          empname           text      NOT NULL,
                          salary integer
);

CREATE OR REPLACE FUNCTION process_emp_audit() RETURNS TRIGGER AS $employee$
BEGIN
    --
    -- Create rows in emp_audit to reflect the operations performed on emp,
    -- making use of the special variable TG_OP to work out the operation.
    --
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO employee
        SELECT 'DELETE', now(), user, o.* FROM old_table o;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO employee
        SELECT 'UPDATE', now(), user, n.* FROM new_table n;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO employee
        SELECT 'INSERT', now(), user, n.* FROM new_table n;
    END IF;
    RETURN NULL;
END;
$employee$ LANGUAGE plpgsql;

CREATE TRIGGER emp_audit_ins
    AFTER INSERT ON employee
    REFERENCING NEW TABLE AS new_table
    FOR EACH STATEMENT EXECUTE FUNCTION process_emp_audit();
CREATE TRIGGER emp_audit_upd
    AFTER UPDATE ON employee
    REFERENCING OLD TABLE AS old_table NEW TABLE AS new_table
    FOR EACH STATEMENT EXECUTE FUNCTION process_emp_audit();
CREATE TRIGGER emp_audit_del
    AFTER DELETE ON employee
    REFERENCING OLD TABLE AS old_table
    FOR EACH STATEMENT EXECUTE FUNCTION process_emp_audit();

insert into employee(firstname, lastname, login, passwd, phone, worktime, salary) VALUES
('test','test','test','test','test','test',50)