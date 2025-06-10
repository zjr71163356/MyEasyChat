-- 用户表 (users)
-- 该表存储了应用程序的用户信息。
--
-- 字段说明:
--   id: VARCHAR(24) - 用户唯一标识符，主键。通常是一个生成的唯一ID。
--   avatar: VARCHAR(191) - 用户头像的URL或路径。默认为空字符串。
--   name: VARCHAR(24) - 用户名或昵称。不能为空。
--   phone: VARCHAR(20) - 用户的手机号码。不能为空。
--   password: VARCHAR(191) - 用户密码（通常是哈希后的值）。可以为空，例如通过第三方登录的用户。
--   status: INT(10) - 用户状态。例如：0-正常，1-禁用等。可以为空。
--   created_at: TIMESTAMP - 记录创建时间。可以为空，通常在创建记录时自动填充。
--   updated_at: TIMESTAMP - 记录最后更新时间。可以为空，通常在更新记录时自动填充。
--
-- 引擎: InnoDB
-- 字符集: utf8mb4
-- 排序规则: utf8mb4_unicode_ci

CREATE TABLE `users` (
     `id` varchar(24) COLLATE utf8mb4_unicode_ci  NOT NULL ,
     `avatar` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
     `name` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL,
     `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
     `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
     `status` int(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
     `created_at` timestamp NULL DEFAULT NULL,
     `updated_at` timestamp NULL DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  goctl model mysql ddl --src user.sql --dir "./models/" -c