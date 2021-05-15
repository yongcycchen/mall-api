/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 5/13/2021 17:26:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for config_kv_store
-- ----------------------------
DROP TABLE IF EXISTS `config_kv_store`;
CREATE TABLE `config_kv_store` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'key',
  `config_key` varchar(255) NOT NULL COMMENT 'config key',
  `config_value` varchar(255) NOT NULL COMMENT 'conifg value',
  `prefix` varchar(255) NOT NULL COMMENT 'prefix',
  `suffix` varchar(255) NOT NULL COMMENT 'suffix',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT 'status 1yes 0no',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT 'is delete 1yes 0no',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_config_key` (`config_key`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COMMENT='config ';

-- ----------------------------
-- Table structure for verify_code_record
-- ----------------------------
DROP TABLE IF EXISTS `verify_code_record`;
CREATE TABLE `verify_code_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'increment id',
  `uid` bigint NOT NULL COMMENT 'UID',
  `business_type` tinyint DEFAULT NULL COMMENT '1-register/login，2-purchase',
  `verify_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'verify code',
  `expire` int DEFAULT NULL COMMENT 'expire unix',
  `country_code` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'country code',
  `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'phone number',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'email',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`),
  KEY `country_code_phone_index` (`country_code`,`phone`) USING BTREE COMMENT 'country code phone index',
  KEY `email_index` (`email`) USING BTREE COMMENT 'emial index',
  KEY `verify_code_index` (`verify_code`) USING BTREE COMMENT 'verify code index'
) ENGINE=InnoDB AUTO_INCREMENT=1095 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='verify code record';

SET FOREIGN_KEY_CHECKS = 1;
