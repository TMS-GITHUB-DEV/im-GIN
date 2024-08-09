create table tb_user
(
    id         bigint unsigned primary key comment 'id',
    phone      char(11) not null default '' comment '电话号码',
    email      varchar(50)     not null default '' comment '邮箱',
    `password`        varchar(20)     not null default '' comment '密码',
    nickname   varchar(10)     not null default '一个小可爱' comment '用户昵称',
    avatar_url char(100)       not null default '' comment '头像url',
    sex        char(2)      not null default 0 comment '性别：保密/男/女',
    region     varchar(50)     not null default '' comment '地区',
    birth_at  int unsigned not null default 0 comment '出生日期',
    created_at int unsigned not null default 0 comment '创建时间',
    updated_at int unsigned not null default 0 comment '更新时间',
    deleted_at int unsigned not null default 0 comment '删除时间',

    unique key idx_unique_phone (phone)
) comment '用户信息表';

create table tb_friend
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    active_uid   bigint unsigned not null default 0	comment '主动者uid',
    passive_uid   bigint unsigned not null default 0 comment '被动者uid',
    active_remark   varchar(10) not null default '' comment '主动者对被动者的备注',
    passive_remark    varchar(10) not null default '' comment '被动者对主动者的备注',
    created_at int unsigned not null default 0 comment '创建时间',
    updated_at int unsigned not null default 0 comment '更新时间',
    deleted_at int unsigned not null default 0 comment '删除时间',

    unique key idx_unique_friend_id (active_uid, passive_uid)
) comment '好友关系表';

create table tb_chat_message
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    sender_uid bigint unsigned not null default 0 comment '发送者uid',
    friend_id  bigint unsigned not null default 0 comment '好友关系id',
    content    varchar(2000) not null default '' comment '消息内容，如是图片消息则设为‘[图片]’',
    `type`     int unsigned not null default 0 comment '消息类型，0：文本消息 1：图片消息',
    `status`   int unsigned not null default 0b00000 comment '消息状态，从右往左已读、撤回、接收方删除、发送方删除；留一位扩展位',
    path_url varchar(200) not null default '' comment '路径url',
    created_at int unsigned not null default 0 comment '创建时间',
    updated_at int unsigned not null default 0 comment '更新时间',
    deleted_at int unsigned not null default 0 comment '删除时间'
) comment '聊天消息表';

create table tb_file
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    `hash`     char(64)    not null default '' comment '文件hash',
    file_url   char(200)    not null default '' comment '文件url',
    size       int unsigned not null default 0 comment '文件大小(kb)',
    `type`     varchar(10)  not null default 0 comment '文件类型',
    created_at int unsigned not null default 0 comment '创建时间',
    updated_at int unsigned not null default 0 comment '更新时间',
    deleted_at int unsigned not null default 0 comment '删除时间',

    unique key idx_unique_hash (`hash`)
) comment '文件表';

create table tb_file_report
(
    id         bigint unsigned primary key auto_increment comment '自增id',
    file_id bigint unsigned not null default 0 comment '文件id',
    uid bigint unsigned not null default 0 comment '引用者uid',
    file_name varchar(150) not null default '' comment '文件名',
    created_at int unsigned not null default 0 comment '创建时间',
    updated_at int unsigned not null default 0 comment '更新时间',
    deleted_at int unsigned not null default 0 comment '删除时间'
) comment '文件引用记录表';
