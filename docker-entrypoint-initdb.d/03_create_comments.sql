DROP TABLE IF EXISTS comments;

CREATE TABLE comments (
  `id` INT NOT NULL AUTO_INCREMENT,
  `comment` VARCHAR(255),
  `post_id` INT,
  `user_id` INT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO comments (comment, post_id, user_id) VALUES ("Hello", 1, 2);
INSERT INTO comments (comment, post_id, user_id) VALUES ("hi", 2, 1);