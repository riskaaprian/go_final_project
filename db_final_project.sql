/*
SQLyog Community v12.4.1 (64 bit)
MySQL - 10.1.21-MariaDB : Database - go_db
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`go_db` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `go_db`;

/*Table structure for table `articles` */

DROP TABLE IF EXISTS `articles`;

CREATE TABLE `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(500) DEFAULT NULL,
  `isi` text,
  `status` varchar(100) DEFAULT 'unpublish',
  `id_user` int(11) DEFAULT NULL,
  `image` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Data for the table `articles` */

insert  into `articles`(`id`,`title`,`isi`,`status`,`id_user`,`image`) values 
(1,'test artikel 1','Seorang warga China di Bandung dilarikan ke Rumah Sakit Hasan Sadikin karena diduga terjangkit virus corona. Ia merupakan salah satu kontraktor proyek kereta cepat di Bandung, CERC.\r\n\r\nSebelumnya beredar informasi bahwa WN China berusia 31 tahun itu karyawan PT Kereta Cepat Indonesia China (KCIC). Namun PT KCIC membantah kabar tersebut.\r\n\r\n\"Dia tidak bekerja di PT KCIC, namun karyawan salah satu kontraktor Kereta Cepat Jakarta-Bandung,\" kata Humas PT KCIC Denny Yusdiana ketika dikonfirmasi, Senin (27/1).\r\n\r\nIa menjelaskan, ketika mengetahui pekerja itu menunjukkan tanda-tanda sakit, PT KCIC langsung melarikan ia ke RSHS Bandung.\r\n\r\n\"Ini adalah prosedur preventif, jadi semua pekerja yang menunjukkan gejala demam dan flu,\" jelas dia.','publish',1,'gambar.jpg'),
(2,'test artikel 2 ','isinya apa aja..,,,,aaa','publish',2,'gbr.jpg');

/*Table structure for table `contact` */

DROP TABLE IF EXISTS `contact`;

CREATE TABLE `contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(200) DEFAULT NULL,
  `pesan` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `contact` */

insert  into `contact`(`id`,`email`,`pesan`) values 
(1,'test@yahoo.com','test error'),
(2,'yee@yahoo.com','test 2');

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `first_name` varchar(200) DEFAULT NULL,
  `last_name` varchar(150) DEFAULT NULL,
  `password` varchar(150) DEFAULT NULL,
  `status` int(11) DEFAULT '2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Data for the table `users` */

insert  into `users`(`id`,`username`,`first_name`,`last_name`,`password`,`status`) values 
(1,'admin@yahoo.com','admin','Administrator','$2a$10$A04qaD8olEVahk5/i899yeUvLUayHLuS3IPKtqjSYhmck4tb9R8Fq',1),
(2,'ian@yahoo.com','ian','ian','$2a$10$A04qaD8olEVahk5/i899yeUvLUayHLuS3IPKtqjSYhmck4tb9R8Fq',2),
(3,'','','',NULL,0),
(4,'makan@yahoo.com','makan','nasi','$2a$10$9rUUNnuczSPa0nqPtbrsg.vfKkTl1sJvneP0xDG4KCSnJJL2zLT1.',2),
(5,'ujang@yahoo.com','ujang','ujang','$2a$10$kTq07Kwd5cC0I9vXUWpe7.uz95vbNNk1ChgU4rvMIJUGPryO2o/7i',2);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
