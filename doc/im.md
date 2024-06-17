## 技术栈：

**服务器**：使用go搭建,使用gin+gorm

**前端**：使用uniapp搭建，使用vue2+viewUI，先做出基本功能编译成h5试运行

**数据库**：mysql

**缓存**：redis



## 需求：

### 基础功能：

#### 登录注册

简单的帐号密码登录注册

> 帐号为手机号，需要进行手机号校验
>
> 密码长度8～20位；仅由数字、大小写字母、'_'和'.'组成
>
> 密码使用sha-256加密



#### 用户个人信息编辑

**用户个人信息包括：**

手机号（暂不支持换绑，即仅显示，不支持更改）
昵称（2～10位|输入框）
性别（男/女|选择器）
头像（url|图片选择API）
出生日期（10位时间戳|年月日选择器）
地区（省份城市，看看能不能找到api|关联选择器）



#### 好友管理

- 好友搜索（帐号/昵称）
- 添加好友（需要验证）发一段文本
- 好友列表（最好能实现首字母跳转）
- 删除好友
- 设置好友备注



#### 聊天

- 列表页
  - 聊天会话以列表的形式展示

  - 展示内容包括头像、昵称（备注）、最新消息的内容、未读消息数

  - [底部tabbar设置总未读条数角标]

  - 下拉刷新、上拉加载更多会话

- 会话页

  - 支持发送文本消息与图片消息（不知道.gif能否有效果）[卡片消息] [文件]

  - 消息撤回（两分钟内）

  - 消息删除（隐藏）

  - 下拉获取历史消息

  - [聊天记录转发]



#### 文件管理







#### 群聊管理

- 创建群聊（设置群聊头像、群聊名）
- 邀请好友加入群聊[卡片消息]
- 加入群聊验证（管理员设置开启或关闭）
- '@'的消息
- 开启/关闭群消息
- 群公告





## 数据库表设计：

用户表、好友关系表、聊天信息表

> 约定：
>
> - 明确没有负数的数值类型使用`unsigned`
> - 除了`text`和`blob`类型，全部使用`not null deault`
> - 命名使用驼峰，用下划线`_`分隔
> - 所有字段和表都要使用`comment`说明
> - 不使用物理外键，使用逻辑外键
> - 时间使用10位时间戳（秒级）`int(10) unsigned`



### 用户表

| 字段名     | 类型        | 默认值       | 备注                   |
| ---------- | ----------- | ------------ | ---------------------- |
| id         | bigint      | 0            | 主键，`unsigned`雪花id |
|            |             |              |                        |
| phone      | char(11)    | ''           | 手机号，`unique`       |
| pwd        | varchar(20) | ''           | 密码                   |
| nickname   | varchar(10) | '一个小可爱' | 昵称                   |
| sex        | tinyint(1)  | 0            | 性别，0：男 1：女      |
| birthday   | bigint(10)  | 0            | 出生日期，`unsigned`   |
| region     | varchar(20) | ''           | 地区                   |
| avatar_url | varchar(50) | ''           | 头像url                |
| create_at  | int(10)     | 0            | 创建时间，`unsigned`   |
| update_at  | int(10)     | 0            | 更新时间，`unsigned`   |
| delete_at  | int(10)     | 0            | 删除时间，`unsigned`   |



### 好友关系表

| 字段      | 类型        | 默认值 | 备注                 |
| --------- | ----------- | ------ | -------------------- |
| id        | bigint      | 0      | 自增id               |
| user1_id  | bigint      | 0      | 用户1                |
| user2_id  | bigint      | 0      | 用户2                |
| beizhu1   | varchar(10) | ''     | user1称user2         |
| beizhu2   | varchar(10) | ''     | user2称user1         |
| create_at | int(10)     | 0      | 创建时间，`unsigned` |
| update_at | int(10)     | 0      | 更新时间，`unsigned` |
| delete_at | int(10)     | 0      | 删除时间，`unsigned` |



### 聊天消息表

| 字段       | 类型          | 默认值  | 备注                                                         |
| ---------- | ------------- | ------- | ------------------------------------------------------------ |
| id         | bigint        | 0       | 自增id                                                       |
| sender_uid | bigint        | 0       | 发送者                                                       |
| friend_id  | bigint        | 0       |                                                              |
| group_id   | bigint        | 0       |                                                              |
| content    | varchar(2000) | ''      | 消息内容，如是图片消息，默认是'[图片]'                       |
| type       | int           | 0       | 消息类型`unsigned`，0：文本消息 1：图片消息                  |
| status     | int           | 0b00000 | 消息状态`unsigned`，从右往左已读、撤回、接收方删除、发送方删除    0b01111 |
| send_at    | int(10)       | 0       | 发送时间，`unsigned`                                         |
| img_url    | char(100)     | ''      | 图片url                                                      |
| path_url   | varchar(100)  | ''      | 路径url（文件\|卡片）                                        |
| create_at  | int(10)       | 0       | 创建时间，`unsigned`                                         |
| update_at  | int(10)       | 0       | 更新时间，`unsigned`                                         |
| delete_at  | int(10)       | 0       | 删除时间，`unsigned`                                         |



### 文件表

| 字段      | 类型         | 默认值 | 备注                       |
| --------- | ------------ | ------ | -------------------------- |
| id        | bigint       | 0      | 自增id，`unsigned`         |
| file_name | varchar(100) | ''     | 文件名                     |
| hash      | char()       | ''     | 文件hash                   |
| file_url  | char()       | ''     | 文件url                    |
| size      | int          | 0      | 文件大小（kb），`unsigned` |
| type      | varchar()    |        |                            |
| create_at | int(10)      | 0      | 创建时间，`unsigned`       |
| update_at | int(10)      | 0      | 更新时间，`unsigned`       |
| delete_at | int(10)      | 0      | 删除时间，`unsigned`       |





### 群聊关系表

|           |      |      |      |
| --------- | ---- | ---- | ---- |
| id        |      |      |      |
| qun_id    |      |      |      |
| user_id   |      |      |      |
| timestamp |      |      |      |
|           |      |      |      |



## 设计







