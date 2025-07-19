CREATE TABLE IF NOT EXISTS `users_to_segment` (
  `id` INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT,
  `user_id` int NOT NULL,
  `segment_id` int NOT NULL
);
CREATE INDEX user_id_idx ON users_to_segment (user_id);
CREATE INDEX segment_id_idx ON users_to_segment (segment_id);