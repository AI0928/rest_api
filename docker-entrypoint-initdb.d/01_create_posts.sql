DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
  `id` INT NOT NULL AUTO_INCREMENT,
  `content` VARCHAR(255),
  `user_id` INT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO posts (content, user_id) VALUES ("Hello World", 1);
INSERT INTO posts (content, user_id) VALUES ("second", 2);