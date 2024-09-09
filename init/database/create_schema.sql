-- initdb/init.sql

CREATE DATABASE IF NOT EXISTS `engineer-country`;

USE `engineer-country`;

CREATE TABLE IF NOT EXISTS `country` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `country_name` VARCHAR(255),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `engineer` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `first_name` VARCHAR(255),
    `last_name` VARCHAR(255),
    `gender` SMALLINT,
    `country_id` BIGINT,
    `title` VARCHAR(255),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`country_id`) REFERENCES `country`(`id`)
);

CREATE TABLE IF NOT EXISTS `country_counts` (
    `key` VARCHAR(255) PRIMARY KEY,
    `value` BIGINT
)
