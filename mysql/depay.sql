CREATE
DATABASE  IF NOT EXISTS `depay`;
USE
`depay`;
--
-- GTID state at the beginning of the backup 
--

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) DEFAULT NULL,
  `pwd` varchar(20) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `enable` int DEFAULT NULL,
  `merchant_id` bigint DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `block_height`
--

DROP TABLE IF EXISTS `block_height`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `block_height` (
  `id` int NOT NULL AUTO_INCREMENT,
  `height` bigint DEFAULT NULL,
  `chain` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=110723 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `merchant`
--

DROP TABLE IF EXISTS `merchant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `merchant` (
  `id` int NOT NULL AUTO_INCREMENT,
  `merchant_id` bigint DEFAULT NULL,
  `merchant_name` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `secret_key` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `public_key` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `update_time` datetime DEFAULT NULL,
  `web_hook` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_merchant_id` (`merchant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `merchant_address`
--

DROP TABLE IF EXISTS `merchant_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `merchant_address` (
  `id` int NOT NULL AUTO_INCREMENT,
  `merchant_id` bigint DEFAULT NULL,
  `coin` varchar(10) DEFAULT NULL,
  `chain` varchar(15) DEFAULT NULL,
  `address` varchar(80) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_coin` (`merchant_id`,`chain`,`address`,`coin`)
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `pay_order`
--

DROP TABLE IF EXISTS `pay_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pay_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` bigint DEFAULT NULL,
  `usdt_amount` decimal(30,18) DEFAULT NULL,
  `swap_amount` decimal(30,18) DEFAULT NULL,
  `token_amount` decimal(30,18) DEFAULT NULL,
  `token_address` varchar(50) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `merchant_address` varchar(50) DEFAULT NULL,
  `merchant_id` bigint DEFAULT NULL,
  `user_address` varchar(50) DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `payed_usdt` decimal(30,18) DEFAULT NULL,
  `success_url` varchar(255) DEFAULT NULL,
  `cancel_url` varchar(255) DEFAULT NULL,
  `notifyed` int DEFAULT '0',
  `chain` varchar(10) DEFAULT NULL,
  `merchant_order` varchar(50) DEFAULT NULL,
  `tx_id` varchar(80) DEFAULT NULL,
  `rec_address` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=177 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `request_log`
--

DROP TABLE IF EXISTS `request_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `request_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `interface` varchar(255) DEFAULT NULL,
  `method` varchar(10) DEFAULT NULL,
  `req` text,
  `rsp` text,
  `create_time` datetime DEFAULT NULL,
  `merchant_id` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  `merchant_order` varchar(50) DEFAULT NULL,
  `direct` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1960096 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sub_order`
--

DROP TABLE IF EXISTS `sub_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sub_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` bigint DEFAULT NULL,
  `amount` decimal(30,18) DEFAULT NULL,
  `token_address` varchar(50) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `merchant_address` varchar(50) DEFAULT NULL,
  `user_address` varchar(50) DEFAULT NULL,
  `merchant_id` int DEFAULT NULL,
  `create_time` date DEFAULT NULL,
  `update_time` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subscribes`
--

DROP TABLE IF EXISTS `subscribes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `subscribes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `merchant_address` varchar(50) DEFAULT NULL,
  `user_address` varchar(50) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-02-01 21:46:20
