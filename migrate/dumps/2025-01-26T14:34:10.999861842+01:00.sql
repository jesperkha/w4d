PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE recipes (id integer primary key, name text, `description` text);
INSERT INTO recipes VALUES(1,'Pasta','Some really good pasta');
INSERT INTO recipes VALUES(2,'Pasta','Some really good pasta');
COMMIT;
