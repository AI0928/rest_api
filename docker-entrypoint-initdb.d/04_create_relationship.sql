DROP TABLE IF EXISTS relationship;

CREATE TABLE relationship (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT,
  `follow_user_id` INT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO relationship (follow_user_id, user_id) VALUES (1, 2);
INSERT INTO relationship (follow_user_id, user_id) VALUES (2, 1);