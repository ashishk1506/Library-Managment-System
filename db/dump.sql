-- MySQL dump 10.13  Distrib 8.1.0, for macos13.3 (arm64)
--
-- Host: localhost    Database: mydb
-- ------------------------------------------------------
-- Server version	8.1.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `mydb`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `mydb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `mydb`;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  `author` varchar(45) DEFAULT NULL,
  `count` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (13,'Tom and Jerry','Warner Bros',1),(14,'Pubgs','Tencent',1),(15,'Amazing spider man','Andreq Garfield',1),(16,'call of duty','sledgehammer',1),(17,'Top Gun','Tom Cruise',1),(18,'Avengers','Hydra',1),(20,'Fantastic 4','Tom cu',1),(28,'Battlefield4','FrostBite',1),(30,'Sherlock Holmnes','stark',1),(31,'Sherlock Holmnes 123','stark',1),(32,'Sherlock Holmnes 1234','stark',1),(33,'Sherlock Holmnes 1234','stark',1),(34,'Sherlock Holmnes 12345','stark',1);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment`
--

DROP TABLE IF EXISTS `payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payment` (
  `pid` int NOT NULL AUTO_INCREMENT,
  `transid` int DEFAULT NULL,
  PRIMARY KEY (`pid`),
  KEY `fk_payment_1_idx` (`transid`),
  CONSTRAINT `fk_payment_1` FOREIGN KEY (`transid`) REFERENCES `transaction` (`tid`) ON DELETE SET NULL ON UPDATE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment`
--

LOCK TABLES `payment` WRITE;
/*!40000 ALTER TABLE `payment` DISABLE KEYS */;
INSERT INTO `payment` VALUES (4,1),(1,18),(2,22),(5,27),(3,28);
/*!40000 ALTER TABLE `payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reader`
--

DROP TABLE IF EXISTS `reader`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reader` (
  `rid` int NOT NULL AUTO_INCREMENT,
  `rname` varchar(45) DEFAULT NULL,
  `wallet` int DEFAULT '100',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reader`
--

LOCK TABLES `reader` WRITE;
/*!40000 ALTER TABLE `reader` DISABLE KEYS */;
INSERT INTO `reader` VALUES (1,'Ashish',4258);
/*!40000 ALTER TABLE `reader` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction` (
  `tid` int NOT NULL AUTO_INCREMENT,
  `readerid` int DEFAULT NULL,
  `bookid` int DEFAULT NULL,
  `issuedate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `returndate` timestamp NULL DEFAULT NULL,
  `amount` int DEFAULT '0',
  PRIMARY KEY (`tid`),
  KEY `fk_transaction_1_idx` (`readerid`),
  KEY `fk_transaction_2_idx` (`bookid`),
  CONSTRAINT `fk_transaction_1` FOREIGN KEY (`bookid`) REFERENCES `books` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT,
  CONSTRAINT `fk_transaction_2` FOREIGN KEY (`readerid`) REFERENCES `reader` (`rid`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction`
--

LOCK TABLES `transaction` WRITE;
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
INSERT INTO `transaction` VALUES (1,1,NULL,'2023-03-13 07:15:46','2023-03-13 09:53:42',9476),(2,1,13,'2023-03-13 07:15:46','2023-03-13 09:54:36',9530),(3,1,14,'2023-03-13 07:15:46','2023-03-13 09:40:16',8670),(4,1,15,'2023-03-13 07:17:11','2023-03-13 07:17:44',33),(5,1,16,'2023-03-13 07:17:11','2023-03-13 07:17:49',38),(6,1,17,'2023-03-13 07:17:11','2023-03-13 07:17:51',40),(15,1,15,'2023-03-13 07:47:16','2023-03-13 07:47:46',30),(16,1,16,'2023-03-13 07:47:16','2023-03-13 07:47:49',33),(17,1,17,'2023-03-13 07:47:16','2023-03-13 07:47:52',36),(18,1,18,'2023-03-13 07:47:16','2023-03-13 07:47:56',40),(19,1,15,'2023-03-13 07:48:00','2023-03-13 09:39:05',6665),(20,1,16,'2023-03-13 07:48:00','2023-03-13 09:39:05',6665),(21,1,17,'2023-03-13 07:48:00','2023-03-13 09:39:05',6665),(22,1,18,'2023-03-13 07:48:00','2023-03-13 09:39:05',6665),(23,1,18,'2023-03-13 11:31:31','2023-03-17 12:02:26',347455),(24,1,13,'2023-03-13 11:32:07','2023-03-13 11:32:26',19),(25,1,14,'2023-03-13 11:32:07','2023-03-13 11:32:26',19),(26,1,15,'2023-03-13 11:32:07','2023-03-13 11:32:26',19),(27,1,NULL,'2023-03-17 11:16:57','2023-03-17 12:52:39',5742),(28,1,28,'2023-03-17 12:08:17','2023-03-17 12:08:30',13),(29,1,13,'2023-03-17 12:55:30','2023-03-17 12:56:31',61),(30,1,14,'2023-03-17 12:55:30','2023-03-17 12:56:31',61),(31,1,15,'2023-03-17 12:55:30','2023-03-17 12:56:31',61),(32,1,NULL,'2023-03-17 13:00:44','2023-03-17 13:01:26',42);
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (9,'2023-03-20 11:22:02','2023-03-20 11:22:02',NULL,'1234','$2a$10$7y5bqNSwnEyWWn0p09uUXu3O7rAfUNFc.xzv87m1rJaJrOi4P6R9u'),(10,'2023-03-20 11:24:42','2023-03-20 11:24:42',NULL,'abcd','$2a$10$PWrNG18ycve50F/wgjdBh.6Y8Sev2WIirnHCJjcjn9K671rWIN8MG'),(11,'2023-03-20 11:27:01','2023-03-20 11:27:01',NULL,'ABC','$2a$10$URyhe46.ir7lrMIMiRZuWeYfYRved29BFR/TjTJ0rOzZ7tn6bQxw6'),(12,'2023-03-20 11:29:04','2023-03-20 11:29:04',NULL,'ddfdf','$2a$10$D979pDTeHy.KlRCTp6Ga.uj73EFkDcXWggMwt8oSQzPeiVSNU9sQu'),(13,'2023-03-20 11:37:35','2023-03-20 11:37:35',NULL,'abcdef','$2a$10$3NzSKUquGbZxsNi03zpQQeMZhq0wpMwrqp35kbfukyvhuAWkHixZu'),(14,'2023-03-20 11:38:36','2023-03-20 11:38:36',NULL,'abcdevasvcdasf','$2a$10$Z8meq6m0aGegoAanmtnNXe5056/JKo5.JJxm3dUN1sv8KrJtef9ZK');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-02-09 16:25:53
