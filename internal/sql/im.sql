create
    database im_db;

use
    im_db;

create table tb_user
(
    id         bigint unsigned primary key comment 'id',
    phone      char(11) unique  not null default '' comment '电话号码',
    pwd        varchar(20)      not null default '' comment '密码',
    nickname   varchar(10)      not null default '' comment '用户昵称',
    sex        int(2) unsigned  not null default 0 comment '性别 0：保密 1：男 2：女',
    email      varchar(50)      not null default '' comment '邮箱',
    avatar_url char(100)        not null default '' comment '头像url',
    birthday   int(10) unsigned not null default 0 comment '出生日期',
    region     varchar(50)      not null default '' comment '地区',
    created_at int(10) unsigned not null default 0 comment '创建时间',
    updated_at int(10) unsigned not null default 0 comment '更新时间',
    deleted_at int(10) unsigned not null default 0 comment '删除时间'
) comment '用户信息表';

create table tb_friend
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    user1_id   bigint unsigned  not null default 0 comment '用户1 id',
    user2_id   bigint unsigned  not null default 0 comment '用户2 id',
    remark1    varchar(10)      not null default '' comment '用户1对用户2的备注',
    remark2    varchar(10)      not null default '' comment '用户2对用户1的备注',
    created_at int(10) unsigned not null default 0 comment '创建时间',
    updated_at int(10) unsigned not null default 0 comment '更新时间',
    deleted_at int(10) unsigned not null default 0 comment '删除时间',

    unique key idx_unique_friend_id (user1_id, user2_id)
) comment '好友关系表';

create table tb_chat_message
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    sender_uid bigint unsigned  not null default 0 comment '发送者uid',
    friend_id  bigint unsigned  not null default 0 comment '好友关系id',
    group_id   bigint unsigned  not null default 0 comment '群聊关系id',
    content    varchar(2000)    not null default '' comment '消息内容，如是图片消息则设为‘[图片]’',
    `type`     int unsigned     not null default 0 comment '消息类型，0：文本消息 1：图片消息',
    `status`   int unsigned     not null default 0b00000 comment '消息状态，从右往左已读、撤回、接收方删除、发送方删除；留一位扩展位',
    img_url    char(100)        not null default '' comment '图片消息url', # todo 这里路径还没确定下来
    path_url   varchar(100)     not null default '' comment '路径url(文件|卡片)',
    created_at int(10) unsigned not null default 0 comment '创建时间',
    updated_at int(10) unsigned not null default 0 comment '更新时间',
    deleted_at int(10) unsigned not null default 0 comment '删除时间'
) comment '聊天消息表';

create table tb_file
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    file_name  varchar(100)     not null default '' comment '文件名',
    file_url   char(100)        not null default '' comment '文件url',
    `hash`     char(128)        not null default '' comment '文件hash',
    size       int unsigned     not null default 0 comment '文件大小(kb)',
    `type`     varchar(10)      not null default 0 comment '文件类型',
    created_at int(10) unsigned not null default 0 comment '创建时间',
    updated_at int(10) unsigned not null default 0 comment '更新时间',
    deleted_at int(10) unsigned not null default 0 comment '删除时间',

    unique key idx_unique_hash (`hash`)
) comment '文件表';
