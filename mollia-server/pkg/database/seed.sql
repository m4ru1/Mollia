 -- 创建初始管理员用户
-- 用户名: admin
-- 密码: password
INSERT INTO `users` (`username`, `password_hash`, `bio`) VALUES
('admin', '$2a$10$tO7FCer6BHLToyuTWKvk8e1TXlxMxxO9RpkfPl4G4WL7vZc/l5DrK', '网站管理员');
