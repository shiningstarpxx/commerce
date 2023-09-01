CREATE TABLE `user` (
                        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                        `username` VARCHAR(45) NOT NULL,
                        `password` VARCHAR(255) NOT NULL,
                        `email` VARCHAR(255) NOT NULL,
                        `created_at` DATETIME NOT NULL,
                        `updated_at` DATETIME NOT NULL,
                        PRIMARY KEY  (`id`),
                        UNIQUE INDEX `username_UNIQUE` (`username` ASC),
                        UNIQUE INDEX `email_UNIQUE` (`email` ASC)
) ENGINE=InnoDB;