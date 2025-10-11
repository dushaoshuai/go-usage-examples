package gorm_and_json

// 表定义如下：
//
// CREATE TABLE `json_test` (
//  `pk` bigint NOT NULL AUTO_INCREMENT COMMENT 'primary key',
//  `json_default_mysql_NULL` json DEFAULT NULL,
//  `json_default_json_null` json NOT NULL DEFAULT (_utf8mb4'null'),
//  `json_default_json_object` json NOT NULL DEFAULT (json_object()),
//  PRIMARY KEY (`pk`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='测试、学习 json';
//
// 这张表主要是学习 MySQL 中的 JSON, 以及使用 Gorm 来操作 JSON.
//
// 表中的数据如下：
//
// mysql> SELECT * FROM `json_test`;
// +----+-------------------------+------------------------+--------------------------+
// | pk | json_default_mysql_NULL | json_default_json_null | json_default_json_object |
// +----+-------------------------+------------------------+--------------------------+
// |  1 | NULL                    | null                   | {}                       |
// +----+-------------------------+------------------------+--------------------------+
// 1 row in set (0.00 sec)
