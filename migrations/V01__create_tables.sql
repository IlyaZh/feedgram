CREATE DATABASE IF NOT EXISTS feedgram;

USE feedgram;

CREATE TABLE `sources` (
  `id` bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `url` text NOT NULL,
  `title` text,
  `link` text NOT NULL,
  `description` text NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `last_sync_at` datetime,
  `last_posted_at` datetime,
  `last_post_link` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime
) ENGINE='InnoDB';

ALTER TABLE `sources`
ADD UNIQUE `url` (`url`(512)),
ADD INDEX `is_active` (`is_active`),
ADD INDEX `last_posted_at` (`last_posted_at`),
ADD INDEX `created_at` (`created_at`),
ADD INDEX `updated_at` (`updated_at`),
ADD INDEX `deleted_at` (`deleted_at`);
