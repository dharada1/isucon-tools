-- INDEXなしのテーブル定義
CREATE TABLE `sample_table` (
  `user_id` int(10) NOT NULL,
  `partner_id` int(10) NOT NULL,
  `active_status` tinyint(3) NOT NULL,
  `description` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- PRIMARY KEYの追加
ALTER TABLE `sample_table` ADD PRIMARY KEY(`user_id`);
-- INDEXの追加
ALTER TABLE `sample_table` ADD INDEX `idx_partner_id` (`partner_id`);
-- 複合INDEXの追加
ALTER TABLE `sample_table` ADD INDEX `idx_user_id_partner_id` (`user_id`, `partner_id`);
-- 複合INDEXの追加
ALTER TABLE `sample_table` ADD INDEX `idx_user_id_partner_id_active_status` (`user_id`, `partner_id`, `active_status`);

-- INDEXを張った結果
SHOW CREATE TABLE `sample_table`;
/*
CREATE TABLE `sample_table` (
  `user_id` int(10) NOT NULL,
  `partner_id` int(10) NOT NULL,
  `active_status` tinyint(3) NOT NULL,
  `description` varchar(100) NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `idx_partner_id` (`partner_id`),
  KEY `idx_user_id_partner_id` (`user_id`,`partner_id`),
  KEY `idx_user_id_partner_id_active_status` (`user_id`,`partner_id`,`active_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
