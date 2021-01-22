use mysql;
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'asdfQWER';

create database test;
use test;

CREATE TABLE `user_info`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `account_name` varchar(20) NOT NULL DEFAULT '' COMMENT '登录名称',
  `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '用户手机号，可用于登录',
  `nick_name` varchar(10) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(20) NOT NULL DEFAULT 'password' COMMENT '用户密码',
  `e_mail` varchar(30) NOT NULL DEFAULT '' COMMENT '用户邮箱',
  PRIMARY KEY (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='用户账号基本数据';