CREATE TABLE `spring_blogs_2025`.`posts`
(
    `id`         VARCHAR(36) NOT NULL,
    `heading`    VARCHAR(45) NULL,
    `content`    VARCHAR(2048) NULL,
    `author`     VARCHAR(45) NULL,
    `likes`      INT,
    `created_at` DATETIME NULL,
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    PRIMARY KEY (`id`)
);