CREATE TABLE `command` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `command_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `protocal` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL ,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `http_command` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `command_id` int NOT NULL,
  `header_id` int NOT NULL,
  `is_insert` boolean NOT NULL,
  `insert_template_id` int,
  `request_type` varchar(255) NOT NULL,
  `url_name` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL ,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `http_headers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `auth_type` varchar(255) NOT NULL,
  `content_type` varchar(255) NOT NULL,
  `accept` varchar(255) NOT NULL DEFAULT '*/*',
  `accept_encoding` varchar(255) NOT NULL DEFAULT 'gzip, deflate, br',
  `connection` varchar(255) NOT NULL DEFAULT 'keep-alive',
  `created_at` datetime NOT NULL ,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `basic_auth` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `header_id` int NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL ,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `bearer_token` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `header_id` int NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL ,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `http_command` ADD FOREIGN KEY (`command_id`) REFERENCES `command` (`id`);

ALTER TABLE `http_command` ADD FOREIGN KEY (`header_id`) REFERENCES `http_headers` (`id`);

ALTER TABLE `basic_auth` ADD FOREIGN KEY (`header_id`) REFERENCES `http_headers` (`id`);

ALTER TABLE `bearer_token` ADD FOREIGN KEY (`header_id`) REFERENCES `http_headers` (`id`);
