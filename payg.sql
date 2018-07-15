-- MySQL dump 10.14  Distrib 5.5.59-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: payg
-- ------------------------------------------------------
-- Server version	5.5.59-MariaDB-1ubuntu0.14.04.1

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
-- Table structure for table `configuration`
--

DROP TABLE IF EXISTS `configuration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `configuration` (
  `container` varchar(30) NOT NULL,
  `ram` varchar(6) NOT NULL,
  `cpu` varchar(5) NOT NULL,
  `storage` varchar(8) NOT NULL,
  `id_user` varchar(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `configuration`
--

LOCK TABLES `configuration` WRITE;
/*!40000 ALTER TABLE `configuration` DISABLE KEYS */;
INSERT INTO `configuration` VALUES ('c2','8780','8','800','13'),('c1','4096','2','500','11'),('c3','2704','4','1000','12');
/*!40000 ALTER TABLE `configuration` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `credit`
--

DROP TABLE IF EXISTS `credit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `credit` (
  `id_user` varchar(3) NOT NULL,
  `credit` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `credit`
--

LOCK TABLES `credit` WRITE;
/*!40000 ALTER TABLE `credit` DISABLE KEYS */;
INSERT INTO `credit` VALUES ('13','1838374.184'),('11','692576.280'),('12','2303499.268');
/*!40000 ALTER TABLE `credit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `container` varchar(30) DEFAULT NULL,
  `id_user` varchar(3) DEFAULT NULL,
  `using_s` varchar(20) DEFAULT NULL,
  `using_moment` varchar(10) DEFAULT NULL,
  `update_date` varchar(35) DEFAULT NULL,
  `ram` varchar(10) DEFAULT NULL,
  `cpu` varchar(10) DEFAULT NULL,
  `storage` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES ('c1','11','3613838','3407','2018-07-14 18:15:55','4096','2','500'),('c1','11','3613840','2','2018-07-14 18:15:57','4096','2','500'),('c1','11','3613842','2','2018-07-14 18:15:59','4096','2','500'),('c3','12','2636468','76','2018-07-14 18:17:15','2704','4','1000'),('c1','11','3613918','76','2018-07-14 18:17:15','4096','2','500'),('c3','12','2636470','2','2018-07-14 18:17:17','2704','4','1000'),('c1','11','3613920','2','2018-07-14 18:17:17','4096','2','500'),('c3','12','2636472','2','2018-07-14 18:17:19','2704','4','1000'),('c1','11','3613922','2','2018-07-14 18:17:19','4096','2','500'),('c3','12','2636474','2','2018-07-14 18:17:21','2704','4','1000'),('c1','11','3613924','2','2018-07-14 18:17:21','4096','2','500'),('c3','12','2636476','2','2018-07-14 18:17:23','2704','4','1000'),('c1','11','3613926','2','2018-07-14 18:17:23','4096','2','500'),('c3','12','2636478','2','2018-07-14 18:17:25','2704','4','1000'),('c1','11','3613928','2','2018-07-14 18:17:25','4096','2','500'),('c3','12','2636480','2','2018-07-14 18:17:27','2704','4','1000'),('c1','11','3613930','2','2018-07-14 18:17:27','4096','2','500'),('c3','12','2636482','2','2018-07-14 18:17:29','2704','4','1000'),('c1','11','3613932','2','2018-07-14 18:17:29','4096','2','500'),('c3','12','2636484','2','2018-07-14 18:17:31','2704','4','1000'),('c1','11','3613934','2','2018-07-14 18:17:31','4096','2','500'),('c3','12','2637269','785','2018-07-14 18:30:36','2704','4','1000'),('c1','11','3614719','785','2018-07-14 18:30:36','4096','2','500'),('c3','12','2637281','12','2018-07-14 18:30:48','2704','4','1000'),('c3','12','2637283','2','2018-07-14 18:30:50','2704','4','1000'),('c3','12','2637285','2','2018-07-14 18:30:52','2704','4','1000'),('c3','12','2637370','85','2018-07-14 18:32:17','2704','4','1000'),('c3','12','2637372','2','2018-07-14 18:32:19','2704','4','1000'),('c3','12','2637374','2','2018-07-14 18:32:21','2704','4','1000'),('c3','12','2637376','2','2018-07-14 18:32:23','2704','4','1000'),('c3','12','2637378','2','2018-07-14 18:32:25','2704','4','1000'),('c3','12','2637380','2','2018-07-14 18:32:27','2704','4','1000'),('c3','12','2637382','2','2018-07-14 18:32:29','2704','4','1000'),('c3','12','2637384','2','2018-07-14 18:32:31','2704','4','1000'),('c3','12','2637386','2','2018-07-14 18:32:33','2704','4','1000');
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `watchlist`
--

DROP TABLE IF EXISTS `watchlist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `watchlist` (
  `container` varchar(30) NOT NULL,
  `using_s` varchar(30) NOT NULL,
  `update_date` varchar(40) NOT NULL,
  `id_user` varchar(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `watchlist`
--

LOCK TABLES `watchlist` WRITE;
/*!40000 ALTER TABLE `watchlist` DISABLE KEYS */;
INSERT INTO `watchlist` VALUES ('c2','52957231','2018-07-14 18:32:33','13'),('c1','3614724','2018-07-14 18:32:33','11'),('c3','2637386','2018-07-14 18:32:33','12');
/*!40000 ALTER TABLE `watchlist` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-07-14 18:38:41
