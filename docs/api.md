## 草莓壁纸接口文档 V0.1

### 接口通用返回格式定义
```
{
    code:0,     // 0-成功
    data:[]|{}|'', //返回的数据
    message:'请求成功' //返回的提示信息
}
```

###### body参数类型
> 'Content-Type': 'application/x-www-form-urlencoded'


### 接口基地址 http://strawberry.wangkaibo.com

### 草莓壁纸app使用接口

#### 接口说明： 第一次注册的时候保留注册信息
- ##### 请求URL
    /register
- ##### 请求方式
    POST
- ##### 请求参数
    请求参数|参数类型|参数说明
    --|--|--|
    platform|String|系统类型
    platformVersion|String|系统版本
    version|String|安装的软件版本
    username|String|安装用户的电脑用户名
    uid|String|生成的软件唯一ID,生成方式 md5_32(`${os}${ovVersion}${userName}${resTime}`)

- ##### 返回参数
    默认

- ##### 举个🌰
    默认

#### 接口说明： 统计正在使用的用户
- ##### 请求URL
    /active
- ##### 请求方式
    POST
- ##### 请求参数
    请求参数|参数类型|参数说明
    --|--|--|
    uid|String|软件ID

- ##### 返回参数
    {
        "code": 400,
        "data": {},
        "message": "传入的UID不存在"
    }

    {
        "code": 0,
        "data": {},
        "message": "success"
    }
- ##### 举个🌰
    默认


### 后台管理接口地址
后台管理提供数据查看、统计分析的相关页面，采用token进行用户鉴权
暂不提供开放注册


#### 接口说明： 获得后台管理的统计数据
- ##### 请求URL
    /statistic
- ##### 请求方式
    GET
- ##### 请求参数
    无
- ##### 返回参数
    返回参数|参数类型|参数说明
    --|--|--|
    totalNum|String|已安装的总数量
    activeNum|String|最近一天的活跃量
    installAdd|Array|近60天的安装量统计
    activeDate|Array|近60天内的日活
    osSta|Object|安装平台统计


- ##### 举个🌰
```js
{
    code:0,
    data:{
        totalNum:'12',
        activeNum:'12',
        installAdd:[
            {num:'5',date:'2018-1-12'},
            {num:'6',date:'2018-1-13'},
        ],
        activeDate:[
            {num:'5',date:'2018-1-12'},
            {num:'6',date:'2018-1-13'},
        ],
        osSta:{
            mac:'12',
            win:'34'
        }

    },
    message:'成功'
}
```


#### 接口说明： 获取已发布的未过期公告
- ##### 请求URL
    /notice
- ##### 请求方式
    GET
- ##### 请求参数
    无
- ##### 返回参数
    如下

- ##### 举个🌰
```js
{
    code:0,
    data:[
        {
            time:'', // 公告时间
            content:'', //公告内容
        }
    ],
    message:'成功'
}
```

#### 接口说明： 获取所有公告
- ##### 请求URL
    /notice_list
- ##### 请求方式
    GET
- ##### 请求参数
    无
- ##### 返回参数
    如下

- ##### 举个🌰
```js
{
    code:0,
    data:[],
    message:'成功'
}
```

#### 接口说明： 删除公告
- ##### 请求URL
    /notice/:id
- ##### 请求方式
    DELETE
- ##### 请求参数
    请求参数|参数类型|参数说明
    --|--|--|
    id|String|通知id
- ##### 返回参数
    如下

- ##### 举个🌰
```js
{
    code:0,
    data:{},
    message:'成功'
}
```

#### 接口说明： 发布公共
- ##### 请求URL
    /notice/:id
- ##### 请求方式
    PATCH
- ##### 请求参数
    请求参数|参数类型|参数说明
    --|--|--|
    id|String|通知id
    status|String|0-不发布 1-发布 参数放到body中
- ##### 返回参数
    如下

- ##### 举个🌰
```js
{
    code:0,
    data:{},
    message:'成功'
}
```

#### 接口说明： 编辑或者新增公告
- ##### 请求URL
    /notice
- ##### 请求方式
    POST
- ##### 请求参数
    请求参数|参数类型|参数说明
    --|--|--|
    id|String?|通知id 有id则编辑公告，id放到params中
    content|String|公共内容
    expire_time|时间戳|过期时间
- ##### 返回参数
    如下

- ##### 举个🌰
```js
{
    code:0,
    data:{},
    message:'成功'
}
```

#### 接口说明： 登录
- ##### 请求URL
    /login
- ##### 请求方式
    POST
- ##### 请求参数
    请求参数|参数类型|参数说明
    --|--|--|
    username|String|用户名
    password|String|密码
- ##### 返回参数
    如下

- ##### 举个🌰
```js
{
    code:0,
    data:{
        token:''
    },
    message:'成功'
}
```
