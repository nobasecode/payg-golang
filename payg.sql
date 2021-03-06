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
INSERT INTO `credit` VALUES ('13','1831656.984'),('11','1633.280'),('12','13495045.328');
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
  `storage` varchar(10) DEFAULT NULL,
  `credit` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES ('c3','12','2637413','2','2018-07-16 11:37:31','2704','4','1000','0'),('c1','11','3762591','2','2018-07-16 11:37:31','4096','2','500','0'),('c3','12','2637415','2','2018-07-16 11:37:33','2704','4','1000','0'),('c1','11','3762593','2','2018-07-16 11:37:33','4096','2','500','0'),('c3','12','2637417','2','2018-07-16 11:37:35','2704','4','1000','0'),('c1','11','3762595','2','2018-07-16 11:37:35','4096','2','500','0'),('c3','12','2637419','2','2018-07-16 11:37:37','2704','4','1000','0'),('c1','11','3762597','2','2018-07-16 11:37:37','4096','2','500','0'),('c3','12','2637421','2','2018-07-16 11:37:39','2704','4','1000','0'),('c1','11','3762599','2','2018-07-16 11:37:39','4096','2','500','0'),('c3','12','2637423','2','2018-07-16 11:37:41','2704','4','1000','0'),('c1','11','3762601','2','2018-07-16 11:37:41','4096','2','500','0'),('c3','12','2637425','2','2018-07-16 11:37:43','2704','4','1000','0'),('c1','11','3762603','2','2018-07-16 11:37:43','4096','2','500','0'),('c3','12','2637427','2','2018-07-16 11:37:45','2704','4','1000','0'),('c1','11','3762605','2','2018-07-16 11:37:45','4096','2','500','0'),('c3','12','2637429','2','2018-07-16 11:37:47','2704','4','1000','0'),('c1','11','3762607','2','2018-07-16 11:37:47','4096','2','500','0'),('c3','12','2637431','2','2018-07-16 11:37:49','2704','4','1000','0'),('c1','11','3762609','2','2018-07-16 11:37:49','4096','2','500','0'),('c3','12','2637433','2','2018-07-16 11:37:51','2704','4','1000','0'),('c1','11','3762611','2','2018-07-16 11:37:51','4096','2','500','0'),('c3','12','2637435','2','2018-07-16 11:37:53','2704','4','1000','0'),('c1','11','3762613','2','2018-07-16 11:37:53','4096','2','500','0'),('c3','12','2637437','2','2018-07-16 11:37:55','2704','4','1000','0'),('c1','11','3762615','2','2018-07-16 11:37:55','4096','2','500','0'),('c3','12','2639759','2322','2018-07-16 12:16:37','2704','4','1000','0'),('c3','12','2639761','2','2018-07-16 12:16:39','2704','4','1000','0'),('c3','12','2812246','172485','2018-07-18 12:11:24','2704','4','1000','0'),('c3','12','2812248','2','2018-07-18 12:11:26','2704','4','1000','0'),('c3','12','3238409','426161','2018-07-23 10:34:07','2704','4','1000','0'),('c3','12','3238411','2','2018-07-23 10:34:09','2704','4','1000','0'),('c3','12','3291253','1528','2018-07-23 14:32:15','2704','4','1000','0'),('c3','12','3291255','2','2018-07-23 14:32:17','2704','4','1000','0'),('c3','12','3291257','2','2018-07-23 14:32:19','2704','4','1000','0'),('c3','12','3291259','2','2018-07-23 14:32:21','2704','4','1000','0'),('c3','12','3291261','2','2018-07-23 14:32:23','2704','4','1000','0'),('c3','12','3291263','2','2018-07-23 14:32:25','2704','4','1000','0'),('c3','12','3291265','2','2018-07-23 14:32:27','2704','4','1000','0'),('c3','12','3291267','2','2018-07-23 14:32:29','2704','4','1000','0'),('c3','12','3291269','2','2018-07-23 14:32:31','2704','4','1000','0'),('c3','12','3291271','2','2018-07-23 14:32:33','2704','4','1000','0'),('c3','12','3291273','2','2018-07-23 14:32:35','2704','4','1000','0'),('c3','12','3291275','2','2018-07-23 14:32:37','2704','4','1000','0'),('c3','12','3291332','57','2018-07-23 14:33:34','2704','4','1000','0'),('c3','12','3291334','2','2018-07-23 14:33:36','2704','4','1000','0'),('c3','12','3291336','2','2018-07-23 14:33:38','2704','4','1000','0'),('c3','12','3291338','2','2018-07-23 14:33:40','2704','4','1000','0'),('c3','12','3291340','2','2018-07-23 14:33:42','2704','4','1000','0'),('c3','12','3291342','2','2018-07-23 14:33:44','2704','4','1000','0'),('c3','12','3291344','2','2018-07-23 14:33:46','2704','4','1000','0'),('c3','12','3291346','2','2018-07-23 14:33:48','2704','4','1000','0'),('c3','12','3291350','4','2018-07-23 14:33:56','2704','4','1000','0'),('c3','12','3291361','11','2018-07-23 14:35:12','2704','4','1000','0'),('c3','12','3291363','2','2018-07-23 14:35:14','2704','4','1000','0'),('c3','12','3291365','2','2018-07-23 14:35:16','2704','4','1000','0'),('c3','12','3291367','2','2018-07-23 14:35:18','2704','4','1000','0'),('c3','12','3291369','2','2018-07-23 14:35:20','2704','4','1000','0'),('c3','12','3291371','2','2018-07-23 14:35:22','2704','4','1000','0'),('c3','12','3291373','2','2018-07-23 14:35:24','2704','4','1000','0'),('c3','12','3291375','2','2018-07-23 14:35:26','2704','4','1000','0'),('c3','12','3291377','2','2018-07-23 14:35:28','2704','4','1000','0'),('c3','12','3291379','2','2018-07-23 14:35:30','2704','4','1000','0'),('c3','12','3291381','2','2018-07-23 14:35:32','2704','4','1000','0'),('c3','12','3291383','2','2018-07-23 14:35:34','2704','4','1000','0'),('c3','12','3291385','2','2018-07-23 14:35:36','2704','4','1000','0'),('c3','12','3291387','2','2018-07-23 14:35:38','2704','4','1000','0'),('c3','12','3291389','2','2018-07-23 14:35:40','2704','4','1000','0'),('c3','12','3291391','2','2018-07-23 14:35:42','2704','4','1000','0'),('c3','12','3291393','2','2018-07-23 14:35:44','2704','4','1000','0'),('c3','12','3291395','2','2018-07-23 14:35:46','2704','4','1000','0'),('c3','12','3291397','2','2018-07-23 14:35:48','2704','4','1000','0'),('c3','12','3291399','2','2018-07-23 14:35:50','2704','4','1000','0'),('c3','12','3291401','2','2018-07-23 14:35:52','2704','4','1000','0'),('c3','12','3291403','2','2018-07-23 14:35:54','2704','4','1000','0'),('c3','12','3291405','2','2018-07-23 14:35:56','2704','4','1000','0'),('c3','12','3291418','13','2018-07-23 14:36:16','2704','4','1000','0'),('c2','13','52957245','14','2018-07-23 14:37:16','8780','8','800','0'),('c2','13','52957247','2','2018-07-23 14:37:18','8780','8','800','0'),('c2','13','52957249','2','2018-07-23 14:37:20','8780','8','800','0'),('c2','13','52957914','665','2018-07-23 14:48:25','8780','8','800','6381.340'),('c2','13','52957916','2','2018-07-23 14:48:27','8780','8','800','19.192'),('c2','13','52957918','2','2018-07-23 14:48:29','8780','8','800','19.192'),('c2','13','52957920','2','2018-07-23 14:48:31','8780','8','800','19.192'),('c2','13','52957922','2','2018-07-23 14:48:33','8780','8','800','19.192'),('c2','13','52957924','2','2018-07-23 14:48:35','8780','8','800','19.192'),('c2','13','52957931','7','2018-07-23 14:48:48','8780','8','800','67.172'),('c3','12','3291426','8','2018-07-23 14:56:18','2704','4','1000','29.696'),('c3','12','3291428','2','2018-07-23 14:56:20','2704','4','1000','7.424'),('c3','12','3291430','2','2018-07-23 14:56:22','2704','4','1000','7.424'),('c3','12','3291432','2','2018-07-23 14:56:24','2704','4','1000','7.424'),('c3','12','3291434','2','2018-07-23 14:56:26','2704','4','1000','7.424'),('c3','12','3291495','61','2018-07-23 14:57:27','2704','4','1000','226.432'),('c3','12','3291497','2','2018-07-23 14:57:29','2704','4','1000','7.424'),('c3','12','3291499','2','2018-07-23 14:57:31','2704','4','1000','7.424'),('c3','12','3291501','2','2018-07-23 14:57:33','2704','4','1000','7.424'),('c3','12','3291503','2','2018-07-23 14:57:35','2704','4','1000','7.424'),('c3','12','3291505','2','2018-07-23 14:57:37','2704','4','1000','7.424'),('c3','12','3291507','2','2018-07-23 14:57:39','2704','4','1000','7.424');
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
INSERT INTO `watchlist` VALUES ('c2','52957931','2018-07-23 14:57:39','13'),('c1','3764929','2018-07-23 14:57:39','11'),('c3','3291507','2018-07-23 14:57:39','12');
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

-- Dump completed on 2018-07-23 14:59:28
