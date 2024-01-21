CREATE TABLE `gorm`.`user_profiles`
(
    `id`      int(20) NOT NULL AUTO_INCREMENT,
    `sex`     tinyint(4) NULL DEFAULT NULL,
    `age`     int(10) NULL DEFAULT NULL,
    `user_id` int(20) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4;