CREATE TABLE `golang_projects_2021`.`posts`
(
    `id`         INT           NOT NULL AUTO_INCREMENT,
    `heading`    VARCHAR(80)   NULL,
    `author`     VARCHAR(45)   NULL,
    `content`    VARCHAR(1024) NULL,
    `likes`      INT           NULL DEFAULT 0,
    `created_at` DATETIME      NULL,
    PRIMARY KEY (`id`)
);