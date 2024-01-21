CREATE TABLE `users`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `username`   varchar(30)  NOT NULL COMMENT '账号',
    `password`   varchar(100) NOT NULL COMMENT '密码',
    `createtime` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4