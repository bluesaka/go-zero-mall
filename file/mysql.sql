CREATE database if not exists go-zero-mall;
CREATE TABLE if not exists `user` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
	`gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
	`mobile` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
	`password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `idx_mobile_unique` (`mobile`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE if not exists `product` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(255)  NOT NULL DEFAULT '' COMMENT '产品名称',
	`desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '产品描述',
	`stock`  int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '产品库存',
	`amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '产品金额',
	`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '产品状态',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE if not exists `order` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
	`pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '产品ID',
	`amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '订单金额',
	`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `idx_uid` (`uid`),
	KEY `idx_pid` (`pid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE if not exists `pay` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
	`oid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单ID',
	`amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '产品金额',
	`source` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付方式',
	`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `idx_uid` (`uid`),
	KEY `idx_oid` (`oid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE database if not exists dtm_barrier;
CREATE TABLE if not exists dtm_barrier.barrier(
  id bigint(22) PRIMARY KEY AUTO_INCREMENT,
  trans_type varchar(45) default '',
  gid varchar(128) default '',
  branch_id varchar(128) default '',
  op varchar(45) default '',
  barrier_id varchar(45) default '',
  reason varchar(45) default '' comment 'the branch type who insert this record',
  create_time datetime DEFAULT now(),
  update_time datetime DEFAULT now(),
  key(create_time),
  key(update_time),
  UNIQUE key(gid, branch_id, op, barrier_id)
);