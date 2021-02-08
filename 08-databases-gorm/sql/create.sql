CREATE DATABASE  IF NOT EXISTS `gorm_projects` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `gorm_projects`;
-- MySQL dump 10.13  Distrib 8.0.16, for Win64 (x86_64)
--
-- Host: localhost    Database: gorm_projects
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `companies`
--

DROP TABLE IF EXISTS `companies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `companies` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_companies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `companies`
--

LOCK TABLES `companies` WRITE;
/*!40000 ALTER TABLE `companies` DISABLE KEYS */;
INSERT INTO `companies` VALUES (1,'2021-02-08 19:48:58.256','2021-02-08 19:48:58.256',NULL,'Linux Foundation'),(2,'2021-02-08 19:48:58.256','2021-02-08 19:48:58.256',NULL,'Sun Microsystems'),(3,'2021-02-08 19:48:58.256','2021-02-08 19:48:58.256',NULL,'Google'),(4,'2021-02-08 19:48:58.256','2021-02-08 19:48:58.256',NULL,'Docker Inc.');
/*!40000 ALTER TABLE `companies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `projects` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `description` longtext,
  `budget` double DEFAULT NULL,
  `start_date` datetime(3) DEFAULT NULL,
  `finished` tinyint(1) DEFAULT NULL,
  `company_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_projects_deleted_at` (`deleted_at`),
  KEY `fk_companies_projects` (`company_id`),
  CONSTRAINT `fk_companies_projects` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects`
--

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` VALUES (1,'2021-02-08 19:50:07.271','2021-02-08 19:50:07.271',NULL,'tux','Linux mascot project',1000,'1991-01-01 00:00:00.000',1,1),(2,'2021-02-08 19:50:07.271','2021-02-08 19:50:07.271',NULL,'duke','Java mascot project',2000,'1996-01-01 00:00:00.000',1,2),(3,'2021-02-08 19:50:07.271','2021-02-08 19:50:07.271',NULL,'gopher','Linux mascot project',1000,'2009-01-01 00:00:00.000',1,3),(4,'2021-02-08 19:50:07.271','2021-02-08 19:50:07.271',NULL,'moby dock','Docker mascot project',1500,'2013-01-01 00:00:00.000',1,4);
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projects_users`
--

DROP TABLE IF EXISTS `projects_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `projects_users` (
  `project_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`project_id`,`user_id`),
  KEY `fk_projects_users_user` (`user_id`),
  CONSTRAINT `fk_projects_users_project` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`),
  CONSTRAINT `fk_projects_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects_users`
--

LOCK TABLES `projects_users` WRITE;
/*!40000 ALTER TABLE `projects_users` DISABLE KEYS */;
INSERT INTO `projects_users` VALUES (1,1),(2,1),(3,1),(4,1),(1,2),(2,2),(3,2),(4,2),(1,3),(2,3),(3,3),(4,3),(1,4),(2,4),(3,4),(4,4);
/*!40000 ALTER TABLE `projects_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext,
  `last_name` longtext,
  `email` longtext,
  `username` longtext,
  `password` longtext,
  `active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2021-02-08 19:48:53.903','2021-02-08 19:48:53.903',NULL,'Linus','Torvalds','linus@linux.com','linus','{bcrypt}$2a$05$MGNalQBW2CwUn0VX0mUOS.tIeQB14Q1l1meLR1gbw3MQ12jsdT1/a',1),(2,'2021-02-08 19:48:53.903','2021-02-08 19:48:53.903',NULL,'James','Gosling','gosling@java.com','james','{bcrypt}$2a$05$0xQResCzqUgHL27btMhqtuAPh02ZxuvJcs4EaQmGPcrzeXq1FXmuW',1),(3,'2021-02-08 19:48:53.903','2021-02-08 19:48:53.903',NULL,'Rob','Pike','pike@golang.com','rob','{bcrypt}$2a$05$YPmCEEr7zut1o9GQb6zW6OJLOd0ewbTBxLAU3cNlMuSMVw/WOm8uy',1),(4,'2021-02-08 19:48:53.903','2021-02-08 19:48:53.903',NULL,'Kamel','Founadi','kamel@docker.com','kamel','{bcrypt}$2a$05$votrVfUT1fTHDcF7tgztWevFYFWN8r8qfl/SU/COOA8Zwu3P3G9qu',1);
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

-- Dump completed on 2021-02-08 19:52:55
