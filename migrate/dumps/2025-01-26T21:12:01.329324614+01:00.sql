PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `recipes` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`description` text);
CREATE TABLE `recipe_ingredient` (`recipe_id` integer,`ingredient_id` integer,PRIMARY KEY (`recipe_id`,`ingredient_id`),CONSTRAINT `fk_recipe_ingredient_recipe` FOREIGN KEY (`recipe_id`) REFERENCES `recipes`(`id`),CONSTRAINT `fk_recipe_ingredient_ingredient` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients`(`id`));
CREATE TABLE IF NOT EXISTS "ingredients"  (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`unit` text,`calories` integer,`carbohydrates` real,`sugars` real,`protein` real,`fats` real,`sodium` real,`milliliters` real,`grams` real,`fiber` real);
INSERT INTO ingredients VALUES(1,'Onion',NULL,41,9.5,4.4000000953674316406,1.2999999523162841796,0.20000000298023223876,2.7999999523162841796,0.0,93.999999999999999996,1.2999999523162841796);
DELETE FROM sqlite_sequence;
INSERT INTO sqlite_sequence VALUES('ingredients',1);
COMMIT;
