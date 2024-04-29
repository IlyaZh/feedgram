SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE `feedgram`;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `sources`;
CREATE TABLE `sources` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `url` text NOT NULL,
  `title` text DEFAULT NULL,
  `link` text NOT NULL,
  `description` text NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `last_sync_at` datetime DEFAULT NULL,
  `last_posted_at` datetime DEFAULT NULL,
  `last_post_link` text DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`url`) USING HASH,
  KEY `is_active` (`is_active`),
  KEY `last_posted_at` (`last_posted_at`),
  KEY `created_at` (`created_at`),
  KEY `updated_at` (`updated_at`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `sources` (`id`, `url`, `title`, `link`, `description`, `is_active`, `last_sync_at`, `last_posted_at`, `last_post_link`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1,	'https://www.atlassian.com/feed',	'Atlassian Engineering',	'https://www.atlassian.com/engineering/feed',	'Unleashing the potential of all teams with tips, tools, and practices',	1,	'2024-04-29 16:21:43',	'2024-01-16 11:28:05',	'https://www.atlassian.com/engineering/the-future-of-automation-at-atlassian-generating-confluence-automation-rules-with-large-language-models',	'2024-04-29 16:05:31',	'2024-04-29 16:21:43',	NULL),
(2,	'https://netflixtechblog.com/feed',	'Netflix TechBlog - Medium',	'https://netflixtechblog.com/feed',	'Learn about Netflix’s world class engineering efforts, company culture, product developments and more. - Medium',	1,	'2024-04-29 16:21:43',	'2024-04-09 22:12:32',	'https://netflixtechblog.com/the-making-of-ves-the-cosmos-microservice-for-netflix-video-encoding-946b9b3cd300?source=rss----2615bd06b42e---4',	'2024-04-29 16:05:51',	'2024-04-29 16:21:43',	NULL),
(3,	'',	'SoundCloud Backstage Blog',	'https://developers.soundcloud.com/blog.rss',	'SoundCloud\'s developer blog.',	1,	'2024-04-29 16:27:49',	'2024-04-23 00:00:00',	'https://developers.soundcloud.com/blog/oauth-migration',	'2024-04-29 16:06:38',	'2024-04-29 16:27:49',	NULL),
(7,	'https://developers.500px.com/feed',	'500px Engineering Blog - Medium',	'https://developers.500px.com/feed',	'Welcome to the 500px Engineering Blog! This is where we, the engineers at 500px, share and discuss the challenges and interesting problems we solve in our day-to-day lives. 500px is always hiring: https://jobs.500px.com. - Medium',	1,	'2024-04-29 16:21:43',	'2018-03-14 14:07:38',	'https://developers.500px.com/understanding-rendering-in-react-redux-7044c6402a75?source=rss----5d9282daaaa1---4',	'2024-04-29 16:13:49',	'2024-04-29 16:21:43',	NULL),
(8,	'https://engineering.atspotify.com/feed/',	'Spotify Engineering',	'https://engineering.atspotify.com/feed',	'Spotify’s official technology blog',	1,	'2024-04-29 16:21:43',	'2024-04-02 21:36:32',	'https://engineering.atspotify.com/2024/04/data-platform-explained/',	'2024-04-29 16:15:43',	'2024-04-29 16:21:43',	NULL),
(10,	'https://tech.instacart.com/feed',	'tech-at-instacart - Medium',	'https://tech.instacart.com/feed',	'Instacart Engineering - Medium',	1,	'2024-04-29 16:21:43',	'2024-04-15 21:07:57',	'https://tech.instacart.com/optimizing-at-the-edge-using-regression-discontinuity-designs-to-power-decision-making-51e296615046?source=rss----587883b5d2ee---4',	'2024-04-29 16:16:45',	'2024-04-29 16:21:43',	NULL),
(11,	'https://engineering.indeedblog.com/feed/',	'Indeed Engineering Blog',	'https://engineering.indeedblog.com/feed/',	'We help people get jobs.',	1,	'2024-04-29 16:21:43',	'2024-01-31 21:57:58',	'https://engineering.indeedblog.com/blog/2024/01/composite-web-performance-metric/',	'2024-04-29 16:20:20',	'2024-04-29 16:21:43',	NULL),
(13,	'https://open.nytimes.com/feed',	'NYT Open - Medium',	'https://open.nytimes.com/feed',	'How we design and build digital products at The New York Times. - Medium',	1,	'2024-04-29 16:22:44',	'2024-03-26 14:17:56',	'https://open.nytimes.com/milestones-on-our-journey-to-standardize-experimentation-at-the-new-york-times-2c6d32db0281?source=rss----51e1d1745b32---4',	'2024-04-29 16:22:42',	'2024-04-29 16:22:44',	NULL);

-- 2024-04-29 16:27:51

