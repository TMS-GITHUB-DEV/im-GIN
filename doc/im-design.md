## 接口

[前提：因为有token，在解构token时可获得发送请求用户的uid，因此发送请求时可不传入uid]

### account

#### 注册

帐号密码注册[POST]

> 输入手机号与密码
>
> 手机号11位校验
>
> 密码长度>=6 & <=20 & 仅由数字、字母或下划线组成
>
> 
>
> 错误：帐号已被注册



#### 登录

帐号密码登录[POST]

> 输入手机号与密码
>
> 手机号11位校验
>
> 密码长度>=6 & <=20 & 仅由数字、字母或下划线组成
>
> 通过上面的检查后查询数据库
>
> 存在则生成token并返回
>
> 
>
> 错误：查询数据库不存在返回“帐号或密码错误”



### user

#### 个人信息修改

1. 头像修改[PUT]

   > 上传图片得到返回url
   >
   > 将图片url和uid一并发送

2. 其他信息修改[PUT]

   > 将要修改的信息和uid一并发送



### friend

#### 好友搜索

手机号/uid搜索[GET]

> 判断长度是属于手机号搜索还是uid搜索
>
> 存在则返回好友个人信息（未定）
>
> 不存在则返回空



#### 添加好友

[POST]

> 传入自己的uid和好友的uid
>
> 添加一条待确认记录



#### 删除好友

[DELETE]

> 传入自己的uid和好友的uid
>
> 



#### 好友申请列表

#### 好友列表

#### 删除好友

#### 设置好友备注