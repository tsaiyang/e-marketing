# cursors 表中插入数据
INSERT INTO `cursors` (name, offset) VALUES ('not_installed_0', 0);
INSERT INTO `cursors` (name, offset) VALUES ('not_installed_3', 0);
INSERT INTO `cursors` (name, offset) VALUES ('not_installed_7', 0);
INSERT INTO `cursors` (name, offset) VALUES ('not_installed_15', 0);
INSERT INTO `cursors` (name, offset) VALUES ('not_installed_30', 0);

# recipients 表中插入数据
INSERT INTO `recipients` (email, name) VALUES ('caiyang.young@gmail.com', 'Young Tsai');
INSERT INTO `recipients` (email, name) VALUES ('tsai.young@foxmail.com', 'Tsai Young');

# strategies 表中插入数据
INSERT INTO `strategies` (week, daily_limited) VALUES (1, 50);
INSERT INTO `strategies` (week, daily_limited) VALUES (2, 100);
INSERT INTO `strategies` (week, daily_limited) VALUES (3, 150);
INSERT INTO `strategies` (week, daily_limited) VALUES (4, 200);
INSERT INTO `strategies` (week, daily_limited) VALUES (5, 250);
INSERT INTO `strategies` (week, daily_limited) VALUES (6, 300);
INSERT INTO `strategies` (week, daily_limited) VALUES (7, 400);
INSERT INTO `strategies` (week, daily_limited) VALUES (8, 500);

# sender_strategies 表中插入数据
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 1);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 2);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 3);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 4);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 5);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 6);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 7);
INSERT INTO `sender_strategies` (sender_id,strategy_id) VALUES (1, 8);