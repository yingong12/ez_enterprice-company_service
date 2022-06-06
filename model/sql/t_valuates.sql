CREATE TABLE t_valulates (
	`id` BIGINT UNSIGNED NOT NULL  AUTO_INCREMENT  COMMENT '自增id',
	`app_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '企业id',
	`valuate_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '估值id，链路唯一标识符',
	`state` tinyint NOT NULL DEFAULT 0 COMMENT '0-估值成功 1-估值中 2-估值失败 3-已取消',
	`form_data` varchar(3000)  NOT NULL DEFAULT '' COMMENT '表单数据',
	`result_path` VARCHAR(255)   NOT NULL DEFAULT '' COMMENT '估值结果路径',
	`requested_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	key `idx_app_id_state` (`app_id`,`state`),
	key `idx_requested_at` (`requested_at`),
	key `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='估值表'