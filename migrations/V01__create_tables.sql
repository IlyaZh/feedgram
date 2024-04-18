CREATE DATABASE IF NOT EXISTS feedgram;

CREATE TABLE `sources` (
  `id` bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `url` text NOT NULL,
  `title` text,
  `link` text NOT NULL,
  `description` text NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime
) ENGINE='InnoDB';

ALTER TABLE `sources`
ADD UNIQUE `url` (`url`),
ADD INDEX `is_active` (`is_active`),
ADD INDEX `created_at` (`created_at`),
ADD INDEX `updated_at` (`updated_at`),
ADD INDEX `deleted_at` (`deleted_at`);


CREATE TABLE `posts` (
  `id` bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `source_id` bigint(20) NOT NULL,
  `title` text NOT NULL,
  `body` text NOT NULL,
  `link` text NOT NULL,
  `author` text NOT NULL,
  `has_read` tinyint(1) NOT NULL DEFAULT '0',
  `additional_info` json,
  `posted_at` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`source_id`) REFERENCES `sources` (`id`)
) ENGINE='InnoDB';


ALTER TABLE `posts`
ADD INDEX `link` (`link`),
ADD INDEX `has_read` (`has_read`),
ADD INDEX `created_at` (`created_at`),
ADD INDEX `posted_at` (`posted_at`);
