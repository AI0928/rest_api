DROP TABLE IF EXISTS likes;

CREATE TABLE likes (
  `id` INT NOT NULL AUTO_INCREMENT,
  `post_id` INT,
  `user_id` INT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO likes (post_id, user_id) VALUES (1, 2);
INSERT INTO likes (post_id, user_id) VALUES (2, 1);