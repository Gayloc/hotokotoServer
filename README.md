# 模仿一言接口
## 鉴权
- 使用`JWT-Bearer`
- 密钥、用户名、密码、Token过期时间均在`auth.go`中配置
### 鉴权方法
在请求头中添加`Authorization:Bearer <token>`
## 获取Token
接口:`http://host:8080/auth`<br/>
方法:`POST`<br/>
请求体类型:`application/json`<br/>
#### 请求示例
```json
{
    "username":"oldEight",
    "password":"olgglxdm"
}
```
#### 返回示例
```json
//失败
{
    "msg": "鉴权失败"
}
```
```json
//成功
{
    "data": {
        "token": "xxxxxxxx.xxxxxxxxxx.xxxxxxxxx"
    },
    "msg": "ok"
}
```
## 获取数据
### 获取全部数据
接口:`http://host:8080/`<br/>
方法:`GET`<br/>
#### 返回示例
```json
{
    "content": {
        "id": 0,
        "hitokoto": "",
        "hitokoto_type": "",
        "reviewer": 0,
        "from_who": "",
        "length": 0
    },
    "data": [
        {
            "id": 1,
            "hitokoto": "1",
            "hitokoto_type": "1",
            "reviewer": 1,
            "from_who": "1",
            "length": 1
        },
        {
            "id": 3,
            "hitokoto": "3",
            "hitokoto_type": "3",
            "reviewer": 3,
            "from_who": "3",
            "length": 3
        },
        {
            "id": 4,
            "hitokoto": "5",
            "hitokoto_type": "5",
            "reviewer": 5,
            "from_who": "5",
            "length": 4
        },
        {
            "id": 5,
            "hitokoto": "5",
            "hitokoto_type": "5",
            "reviewer": 0,
            "from_who": "5",
            "length": 1
        }
    ]
}
```
### 按类型获取数据
接口:`http://host:8080/`<br/>
参数:`type:string`要获取数据的类型<br/>
方法:`GET`<br/>
请求示例:`http://host:8080/?type=1`<br/>
#### 返回示例
```json
{
    "content": {
        "id": 0,
        "hitokoto": "",
        "hitokoto_type": "1",
        "reviewer": 0,
        "from_who": "",
        "length": 0
    },
    "data": [
        {
            "id": 1,
            "hitokoto": "1",
            "hitokoto_type": "1",
            "reviewer": 1,
            "from_who": "1",
            "length": 1
        }
    ]
}
```
## 添加数据
接口:`http://host:8080/`<br/>
方法:`POST`<br/>
请求体类型:`application/json`<br/>
#### 请求示例
```json
//必需项
{
    "hitokoto": "6",
    "hitokoto_type": "6",
    "from_who": "6"
}
```
```json
//全部
{
    "id": 5,              //默认自动生成(最后项目的id+1)
    "hitokoto": "5",
    "hitokoto_type": "5",
    "reviewer": 0,
    "from_who": "5",
    "length": 1           //默认自动生成
}
```
#### 返回示例
```json
//成功
{
    "message": "ok",
    "user":"oldEight"
}
```
```json
//失败
{
    "message": "提交内容不符合规范",
    "user":"oldEight"
}
```
## 删除数据
接口:`http://host:8080/`<br/>
参数:`id:int`要删除数据的id<br/>
方法:`DELETE`<br/>
请求体类型:`application/json`<br/>
> 手动设定的`id`重复则删除第一个
#### 返回示例
```json
//成功
{
    "message": "ok",
    "user":"oldEight"
}
```
```json
//失败
{
    "message": "删除失败",
    "user":"oldEight"
}
```
## 修改数据
接口:`http://host:8080/`<br/>
参数:`id:int`要修改数据的id<br/>
方法:`PUT`<br/>
请求体类型:`application/json`<br/>
> 手动设定的`id`重复则修改第一个
#### 请求示例
```json
//全部(可选)
{
    "hitokoto": "5",
    "hitokoto_type": "5",
    "reviewer": 0,
    "from_who": "5",
    "length": 1           //默认自动生成
}
```
#### 返回示例
```json
//成功
{
    "message": "ok",
    "user":"oldEight"
}
```
```json
//失败
{
    "message": "修改失败",
    "user":"oldEight"
}
```
## 说明
首次运行或`hitokoto.json`被删除时会根据`DefaultDatabase.json`文件自动生成默认的`hitokoto.json`文件。