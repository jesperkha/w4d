PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `recipes` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`description` text);
CREATE TABLE `ingredients` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`unit` text,`calories` integer,`carbohydrates` integer,`sugars` integer,`protein` integer,`fats` integer,`sodium` integer);
CREATE TABLE `recipe_ingredient` (`recipe_id` integer,`ingredient_id` integer,PRIMARY KEY (`recipe_id`,`ingredient_id`),CONSTRAINT `fk_recipe_ingredient_recipe` FOREIGN KEY (`recipe_id`) REFERENCES `recipes`(`id`),CONSTRAINT `fk_recipe_ingredient_ingredient` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients`(`id`));
DELETE FROM sqlite_sequence;
COMMIT;
