/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50648
 Source Host           : localhost:3306
 Source Schema         : vip

 Target Server Type    : MySQL
 Target Server Version : 50648
 File Encoding         : 65001

 Date: 15/06/2020 11:28:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for adminInfo
-- ----------------------------
DROP TABLE IF EXISTS `adminInfo`;
CREATE TABLE `adminInfo` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Sex` tinyint(4) DEFAULT '0',
  `Mobile` char(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `CreTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `LastLogin` datetime DEFAULT NULL,
  `IsDelete` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  UNIQUE KEY `UserName` (`UserName`),
  UNIQUE KEY `Mobile` (`Mobile`),
  KEY `ID_2` (`ID`,`UserName`,`Mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for jiFenRecord
-- ----------------------------
DROP TABLE IF EXISTS `jiFenRecord`;
CREATE TABLE `jiFenRecord` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `VipInfoID` int(11) DEFAULT NULL,
  `AdminID` int(11) DEFAULT NULL,
  `CreTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `TypeOP` tinyint(4) DEFAULT NULL,
  `Num` float DEFAULT NULL,
  `IsDelete` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  KEY `ID_2` (`ID`,`VipInfoID`,`AdminID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `AdminID` int(11) DEFAULT NULL,
  `TypeOp` tinyint(4) DEFAULT NULL,
  `Comment` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `CreTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `IsDelete` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `SessionID` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `UserName` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `AdminID` int(11) NOT NULL,
  PRIMARY KEY (`SessionID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for vipInfo
-- ----------------------------
DROP TABLE IF EXISTS `vipInfo`;
CREATE TABLE `vipInfo` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` char(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Sex` tinyint(4) DEFAULT '0',
  `Mobile` char(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `JiFen` float DEFAULT NULL,
  `JiFenCount` float DEFAULT NULL,
  `Status` tinyint(4) DEFAULT '0',
  `CreTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `IsDelete` tinyint(4) DEFAULT '0',
  `Belong` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  UNIQUE KEY `Name` (`Name`),
  UNIQUE KEY `Mobile` (`Mobile`),
  KEY `ID_2` (`ID`,`Name`,`Mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
