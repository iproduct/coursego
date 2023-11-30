CREATE SCHEMA IF NOT EXISTS `my-blogs-go2023`;

CREATE TABLE IF NOT EXISTS `my-blogs-go2023`.`posts` (
    `id` VARCHAR(36) NOT NULL,
    `title` VARCHAR(45) NULL,
    `author` VARCHAR(45) NULL,
    `content` VARCHAR(2048) NULL,
    `likes` INT NULL,
    `created_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`));