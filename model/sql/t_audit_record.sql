CREATE TABLE t_audit_record (
	`id` BIGINT UNSIGNED NOT NULL  COMMENT '自增id',
	`audit_id` varchar(64) NOT NULL DEFAULT '审批id' COMMENT '用户提交时后台自动生成的唯一id'
	`app_id` varchar(64) NOT NULL DEFAULT '' COMMENT '企业id或者机构id',
	`app_type` TINYINT UNSIGNED NOT NULL DEFAULT 0  COMMENT '0-企业 1-机构',
	`state` tinyint NOT NULL DEFAULT 0 COMMENT '0-通过 1-审核中 2-未通过',
	`form_data`  varchar(3000) DEFAULT COMMENT '数据表单信息,json形式',
	`requested_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '审批提交时间',
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	primary key (`id`),
	unique key `app_id`,
	key `state` COMMENT 'O端查询用',
	key `requested_at` COMMENT 'O端查询用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='审核表明细表'