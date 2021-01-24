-- MySQL dump 10.13  Distrib 5.7.32, for Linux (x86_64)
--
-- Host: localhost    Database: train
-- ------------------------------------------------------
-- Server version	5.7.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `booking`
--

DROP TABLE IF EXISTS `booking`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking` (
  `id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `schedule_id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `booking_code` varchar(30) CHARACTER SET latin1 NOT NULL,
  `customer_id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `departure_date` date NOT NULL,
  `qty` int(11) NOT NULL,
  `price` decimal(10,0) NOT NULL,
  `total` decimal(10,0) NOT NULL,
  `expired_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `booking_status` varchar(30) CHARACTER SET latin1 NOT NULL DEFAULT 'Booked',
  `booked_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `paid_at` timestamp NULL DEFAULT NULL,
  `expired_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `booking_FK` (`schedule_id`),
  KEY `booking_FK_1` (`customer_id`),
  CONSTRAINT `booking_FK` FOREIGN KEY (`schedule_id`) REFERENCES `schedule` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `booking_FK_1` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking`
--

LOCK TABLES `booking` WRITE;
/*!40000 ALTER TABLE `booking` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `customer_code` varchar(30) CHARACTER SET latin1 NOT NULL,
  `first_name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `last_name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `email` varchar(100) CHARACTER SET latin1 NOT NULL,
  `phone_number` varchar(50) CHARACTER SET latin1 NOT NULL,
  `gender` enum('m','f') CHARACTER SET latin1 NOT NULL,
  `birth_date` date NOT NULL,
  `activation_code` varchar(6) CHARACTER SET latin1 NOT NULL,
  `password` varchar(255) CHARACTER SET latin1 NOT NULL,
  `status_active` int(11) DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES ('40129f30-b5a3-4cdd-88cb-a1da278eb5cc','20210124807236','Ado','Pabianko','adopabianko@gmail.com','087874083220','m','1992-01-06','76767B','$2a$10$yaJuCYtrMoK4zUQo7yR7L.5zFhJKd19vDzhwUq.PYA7BxC9pmU6CC',1,'2021-01-23 22:47:35','2021-01-23 22:48:46');
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `passenger`
--

DROP TABLE IF EXISTS `passenger`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `passenger` (
  `id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `booking_id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `ticket_number` varchar(30) NOT NULL,
  `first_name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `last_name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `email` varchar(100) CHARACTER SET latin1 NOT NULL,
  `phone_number` varchar(50) CHARACTER SET latin1 NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `passenger_FK` (`booking_id`),
  CONSTRAINT `passenger_FK` FOREIGN KEY (`booking_id`) REFERENCES `booking` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `passenger`
--

LOCK TABLES `passenger` WRITE;
/*!40000 ALTER TABLE `passenger` DISABLE KEYS */;
/*!40000 ALTER TABLE `passenger` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule`
--

DROP TABLE IF EXISTS `schedule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schedule` (
  `id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `origin` varchar(30) CHARACTER SET latin1 NOT NULL,
  `destination` varchar(30) CHARACTER SET latin1 NOT NULL,
  `train_code` varchar(30) CHARACTER SET latin1 NOT NULL,
  `time` time NOT NULL,
  `quota` int(11) NOT NULL,
  `balance` int(11) NOT NULL,
  `price` decimal(10,0) NOT NULL,
  `status_active` int(11) DEFAULT '1',
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule`
--

LOCK TABLES `schedule` WRITE;
/*!40000 ALTER TABLE `schedule` DISABLE KEYS */;
INSERT INTO `schedule` VALUES ('1b4750d7-598b-11eb-b2cb-0242ac120003','PDL','GMR','KA002','15:00:00',100,100,75000,1,'2021-01-01','2022-01-01','2021-01-18 12:45:55','2021-01-18 12:45:55'),('56ca1c7d-598b-11eb-b2cb-0242ac120003','PDL','GMR','KA004','16:00:00',100,100,75000,1,'2021-01-01','2022-01-01','2021-01-18 12:47:35','2021-01-18 12:47:35'),('5a976b0a-598b-11eb-b2cb-0242ac120003','PDL','GMR','KA005','17:00:00',100,100,75000,1,'2021-01-01','2022-01-01','2021-01-18 12:47:42','2021-01-18 12:47:42'),('8c6ae736-597e-11eb-b2cb-0242ac120003','PDL','GMR','KA001','14:00:00',100,0,75000,1,'2021-01-01','2022-01-01','2021-01-18 11:16:02','2021-01-24 03:38:36');
/*!40000 ALTER TABLE `schedule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `station`
--

DROP TABLE IF EXISTS `station`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `station` (
  `id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `station_code` varchar(30) CHARACTER SET latin1 NOT NULL,
  `station_name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `station_city` varchar(100) CHARACTER SET latin1 DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `station`
--

LOCK TABLES `station` WRITE;
/*!40000 ALTER TABLE `station` DISABLE KEYS */;
INSERT INTO `station` VALUES ('69e51477-57e9-11eb-b939-0242ac120002','GMR','GAMBIR','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e51753-57e9-11eb-b939-0242ac120002','JAKK','JAKARTA KOTA','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e5186c-57e9-11eb-b939-0242ac120002','JNG','JATINEGARA','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e518e2-57e9-11eb-b939-0242ac120002','MRI','MANGGARAI','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e5194c-57e9-11eb-b939-0242ac120002','PSE','PASAR SENEN','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e519ad-57e9-11eb-b939-0242ac120002','THB','TANAH ABANG','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e51a0e-57e9-11eb-b939-0242ac120002','TPK','TANJUNG PRIUK','JAKARTA','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e51b3a-57e9-11eb-b939-0242ac120002','BD','BANDUNG','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e51b9e-57e9-11eb-b939-0242ac120002','CCL','CICALENGKA','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('69e51bf6-57e9-11eb-b939-0242ac120002','CD','CIKADONGDONG','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a041c5d-57e9-11eb-b939-0242ac120002','CMI','CIMAHI','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a041e5a-57e9-11eb-b939-0242ac120002','GDB','GEDEBAGE','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a041efa-57e9-11eb-b939-0242ac120002','HRP','HAURPUGUR','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a041f6d-57e9-11eb-b939-0242ac120002','KAC','KIARACONDONG','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a041fce-57e9-11eb-b939-0242ac120002','PDL','PADALARANG','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a04202c-57e9-11eb-b939-0242ac120002','RCK','RANCAEKEK','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58'),('6a042090-57e9-11eb-b939-0242ac120002','RH','RENDEH','BANDUNG','2021-01-16 10:55:58','2021-01-16 10:55:58');
/*!40000 ALTER TABLE `station` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `train`
--

DROP TABLE IF EXISTS `train`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `train` (
  `id` varchar(255) CHARACTER SET latin1 NOT NULL,
  `train_code` varchar(30) CHARACTER SET latin1 NOT NULL,
  `train_name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `train`
--

LOCK TABLES `train` WRITE;
/*!40000 ALTER TABLE `train` DISABLE KEYS */;
INSERT INTO `train` VALUES ('c0b9e9e3-57e9-11eb-b939-0242ac120002','KA001','Kereta Api 001','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ec27-57e9-11eb-b939-0242ac120002','KA002','Kereta Api 002','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ece4-57e9-11eb-b939-0242ac120002','KA003','Kereta Api 003','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ed2e-57e9-11eb-b939-0242ac120002','KA004','Kereta Api 004','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ed67-57e9-11eb-b939-0242ac120002','KA005','Kereta Api 005','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ed9f-57e9-11eb-b939-0242ac120002','KA006','Kereta Api 006','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9edd8-57e9-11eb-b939-0242ac120002','KA007','Kereta Api 007','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ee10-57e9-11eb-b939-0242ac120002','KA008','Kereta Api 008','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ee4c-57e9-11eb-b939-0242ac120002','KA009','Kereta Api 009','2021-01-16 10:58:23','2021-01-16 10:58:23'),('c0b9ee80-57e9-11eb-b939-0242ac120002','KA010','Kereta Api 010','2021-01-16 10:58:23','2021-01-16 10:58:23');
/*!40000 ALTER TABLE `train` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-24 10:47:11
