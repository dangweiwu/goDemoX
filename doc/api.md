---
theme: orange
---
```run
window.SetConfig(
    "http://127.0.0.1:8009",
    {"Authorization":""}
)
```
# xx系统管理

- 版本v0.0.1

- 系统管理 2024年 12月 1日
# 1 用户管理

系统用户管理 增删查改

## 1.1 创建用户

> 基础信息

- **PATH: /api/admin**
- **METHOD: POST**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|account|账号|string|required|admin||
|phone|手机号|string|max=11|123456789||
|name|姓名|string|max=100|张三||
|status|状态|string|oneof=0 1|1|0 无效 1 有效|
|password|密码|string|max=100,required|123456||
|memo|备注|string|max=300|||
|email|Email|string|omitempty,email|||
|is_super_admin|是否超级管理员|string|oneof=0 1|1|0:否 1:是|
|role|角色ID|string|max=100|||

```button
var req = {

    Url:"/api/admin",
    Method:"POST",
    Header:{
        Authorization:"",
    },
    Form:{
        account:"admin",
        phone:"123456789",
        name:"张三",
        status:"1",
        password:"123456",
        memo:"",
        email:"",
        is_super_admin:"1",
        role:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 1.2 修改用户

> 基础信息

- **PATH: /api/admin/:id**
- **METHOD: PUT**
> URL 参数 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|id|用户ID|int|1||
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|phone|手机号|string|max=11|123456789||
|name|姓名|string|max=100|张三||
|status|状态|string|oneof=0 1|1|0 无效 1 有效|
|memo|备注|string|max=300|||
|email|Email|string|omitempty,email|||
|is_super_admin|是否超级管理员|string|oneof=0 1|1|0:否 1:是|
|role|角色ID|string|max=100|||

```button
var req = {

    Url:"/api/admin/:id",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        phone:"123456789",
        name:"张三",
        status:"1",
        memo:"",
        email:"",
        is_super_admin:"1",
        role:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 1.3 修改密码

> 基础信息

- **PATH: /api/admin/resetpwd/:id**
- **METHOD: PUT**
> URL 参数 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|id|用户ID|int|1||
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|

```button
var req = {

    Url:"/api/admin/resetpwd/:id",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data|新密码|string||数字与字母组合的随机6位密码|



---

## 1.4 查询用户

> 基础信息

- **PATH: /api/admin**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|
>  QUERY 参数 



| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- | -- | -- |
|limit|条数|string|10||
|current|页码|string|1||
|account|账号|string|admin||
|phone|手机号|string|12345678911||
|email|email|string|||
|name|姓名|string|||

```button
var req = {

    Url:"/api/admin",
    Method:"GET",
    Header:{
        Authorization:"tokenstring",
    },
    Query:{
        limit:"10",
        current:"1",
        account:"admin",
        phone:"12345678911",
        email:"",
        name:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|page|分页数据|||参考Page定义|
|data|数据|any||参考Data定义|

>  Page定义 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|limit|条数|int|||
|current|当前页码|int|||
|total|总数|int|||

>  []Data 定义 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|account|账号|string|admin||
|phone|手机号|string|123456789||
|name|姓名|string|张三||
|status|状态|string|1|0 无效 1 有效|
|memo|备注|string|||
|email|Email|string|||
|is_super_admin|是否超级管理员|string|1|0:否 1:是|
|role|角色ID|string|||

---

## 1.5 删除用户

> 基础信息

- **PATH: /api/admin/:id**
- **METHOD: DELETE**
> URL 参数 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|id|用户ID|int|1||
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|

```button
var req = {

    Url:"/api/admin/:id",
    Method:"DELETE",
    Header:{
        Authorization:"tokenstring",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||type|ok|成功|



>  500 RESPONSE 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string||code msg|
|data|代码|string||kind=code时候 data指示异常代码|
|msg|消息|string||kind=msg时候 msg标识异常信息|

>  msg 数据 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|msg|||禁止删除自己||
|msg|||记录不存在||



---

## 1.6 删除用户

> 基础信息

- **PATH: /api/role/:id**
- **METHOD: DELETE**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|

```button
var req = {

    Url:"/api/role/:id",
    Method:"DELETE",
    Header:{
        Authorization:"tokenstring",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||type|ok|成功|



>  500 RESPONSE 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string||code msg|
|data|代码|string||kind=code时候 data指示异常代码|
|msg|消息|string||kind=msg时候 msg标识异常信息|

>  msg 数据 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|msg|||记录不存在||



---
# 2 系统我的

包括基本信息获取修改 登录登出 token刷新

## 2.1 登录

> 基础信息

- **PATH: /api/login**
- **METHOD: POST**
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|account|账号|string||admin||
|password|密码|string||12345||

```button
var req = {

    Url:"/api/login",
    Method:"POST",
    Form:{
        account:"admin",
        password:"12345",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|access_token|鉴权token|string||header头Authorization参数|
|refresh_at|刷新时间戳|int64||到达时间戳时间进行token刷新|
|refresh_token|刷新token|string||刷新token时所用参数|

>  400 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string||code msg|
|data|代码|string||kind=code时候 data指示异常代码|
|msg|消息|string||kind=msg时候 msg标识异常信息|

>  Msg 数据 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|msg|密码错误||||
|msg|账号不存在||||
|msg|账号被禁用||||



---

## 2.2 登出

> 基础信息

- **PATH: /api/logout**
- **METHOD: POST**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|

```button
var req = {

    Url:"/api/logout",
    Method:"POST",
    Header:{
        Authorization:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 2.3 我的详情

> 基础信息

- **PATH: /api/my**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|

```button
var req = {

    Url:"/api/my",
    Method:"GET",
    Header:{
        Authorization:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|account|账号|string|||
|phone|电话|string|||
|name|姓名|string|||
|memo|备注|string|||
|email|Email|string|||
|is_super_admin|是否超级管理员|string||0不是 1是|
|role|角色ID|string|||

---

## 2.4 修改我的信息

> 基础信息

- **PATH: /api/my**
- **METHOD: PUT**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|phone|电话|string|max=11|12312312312||
|name|姓名|string|max=100|张三||
|memo|备注|string|max=300|这是备注||
|email|Email|string|omitempty,email|2@qq.com||

```button
var req = {

    Url:"/api/my",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        phone:"12312312312",
        name:"张三",
        memo:"这是备注",
        email:"2@qq.com",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 2.5 修改密码

> 基础信息

- **PATH: /api/admin/my/password**
- **METHOD: PUT**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|password|原始密码|string|required|||
|new_password|新密码|string|required|||

```button
var req = {

    Url:"/api/admin/my/password",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        password:"",
        new_password:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 2.6 刷新token

> 基础信息

- **PATH: /api/token/refresh**
- **METHOD: POST**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|refresh_token|刷新token|string|required|||

```button
var req = {

    Url:"/api/token/refresh",
    Method:"POST",
    Header:{
        Authorization:"",
    },
    Form:{
        refresh_token:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|access_token|鉴权token|string||header头Authorization参数|
|refresh_at|刷新时间戳|int64||到达时间戳时间进行token刷新|
|refresh_token|刷新token|string||刷新token时所用参数|

>  401 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string||code msg|
|data|代码|string||kind=code时候 data指示异常代码|
|msg|消息|string||kind=msg时候 msg标识异常信息|

>  Msg 数据 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|msg|refreshtoken已失效||||



---

## 2.7 获取权限

> 基础信息

- **PATH: /api/my-auth**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|

```button
var req = {

    Url:"/api/my-auth",
    Method:"GET",
    Header:{
        Authorization:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|data||any||响应数据 参考Data定义或说明|

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data|角色|[]string||角色数组|



>  400 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string||code msg|
|data|代码|string||kind=code时候 data指示异常代码|
|msg|消息|string||kind=msg时候 msg标识异常信息|

>  msg 数据 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|msg|||角色已被禁用||
|msg|||角色不存在||



---
# 3 系统设置

包括链路追踪,指标采集

## 3.1 运行状态

> 基础信息

- **PATH: /api/sys**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|权限|type||bascAuth base64(name:password)|

```button
var req = {

    Url:"/api/sys",
    Method:"GET",
    Header:{
        Authorization:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|run_time|运行时间|string||日期格式化字符串|
|start_time|开始时间|string||日期格式化字符串|
|open_trace|链路追踪开关|string|||
|open_metric|指标采集开关|string|||

---

## 3.2 设定开关

> 基础信息

- **PATH: /api/sys**
- **METHOD: PUT**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|权限|type||bascAuth base64(name:password)|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|name|名称|string|oneof=trace metric|||
|act|开关|string|oneof=0 1|||

```button
var req = {

    Url:"/api/sys",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        name:"",
        act:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---
# 4 权限管理



## 4.1 创建权限

> 基础信息

- **PATH: /api/auth**
- **METHOD: POST**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|name|权限名称|string|max=100,required|||
|code|权限ID|string|required,max=100|||
|order_num|排序|int||||
|api|API|string|max=200|||
|method|方法|string|max=50|GET||
|kind|类型|string|oneof=0 1||0:api 1:菜单|
|parent_id|上级ID|int||||

```button
var req = {

    Url:"/api/auth",
    Method:"POST",
    Header:{
        Authorization:"",
    },
    Form:{
        name:"",
        code:"",
        order_num:,
        api:"",
        method:"GET",
        kind:"",
        parent_id:,
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 4.2 修改权限

> 基础信息

- **PATH: /api/auth/:id**
- **METHOD: PUT**
> URL 参数 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|id|权限ID|int|1||
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|name|权限名称|string|max=100,required|||
|order_num|排序|int||||
|api|API|string|max=200|||
|method|方法|string|max=50|GET||
|kind|类型|string|oneof=0 1||0:api 1:菜单|
|parent_id|上级ID|int||||

```button
var req = {

    Url:"/api/auth/:id",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        name:"",
        order_num:,
        api:"",
        method:"GET",
        kind:"",
        parent_id:,
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 4.3 权限查询

> 基础信息

- **PATH: /api/auth**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|
>  QUERY 参数 



| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string|0|0:按钮 1:页面|

```button
var req = {

    Url:"/api/auth",
    Method:"GET",
    Header:{
        Authorization:"tokenstring",
    },
    Query:{
        kind:"0",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|data||any||响应数据 参考Data定义或说明|

>  Data定义 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|name|权限名称|string|||
|code|权限ID|string|||
|order_num|排序|int|||
|api|API|string|||
|method|方法|string|GET||
|kind|类型|string||0:api 1:菜单|
|parent_id|上级ID|int|||
|children|子集|[]self|||

---

## 4.4 删除用户

> 基础信息

- **PATH: /api/auth/:id**
- **METHOD: DELETE**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|

```button
var req = {

    Url:"/api/auth/:id",
    Method:"DELETE",
    Header:{
        Authorization:"tokenstring",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||type|ok|成功|



>  500 RESPONSE 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|kind|类型|string||code msg|
|data|代码|string||kind=code时候 data指示异常代码|
|msg|消息|string||kind=msg时候 msg标识异常信息|

>  msg 数据 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|msg|||该权限下包含其他权限，禁止删除！||
|msg|||记录不存在||



---

## 4.5 获取所有URL

创建修改权限时url参数从这获取

> 基础信息

- **PATH: /api/auth**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|

```button
var req = {

    Url:"/api/auth",
    Method:"GET",
    Header:{
        Authorization:"tokenstring",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data|权限列表|[]string|['/api/admin']|列表数据|



---
# 5 角色管理



## 5.1 创建角色

> 基础信息

- **PATH: /api/role**
- **METHOD: POST**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|code|编码|string|required,max=100|||
|name|名称|string|max=100|||
|order_num|权限代码|int|||6位编码12顶级菜单34当前菜单56接口编码|
|status|状态|string||||
|memo|备注|string|max=300|||

```button
var req = {

    Url:"/api/role",
    Method:"POST",
    Header:{
        Authorization:"",
    },
    Form:{
        code:"",
        name:"",
        order_num:,
        status:"",
        memo:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 5.2 修改角色

> 基础信息

- **PATH: /api/role/:id**
- **METHOD: PUT**
> URL 参数 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|id|角色ID|int|1||
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|name|名称|string|max=100|||
|order_num|权限代码|int|||6位编码12顶级菜单34当前菜单56接口编码|
|status|状态|string|||0:禁用1:启用|
|memo|备注|string|max=300|||

```button
var req = {

    Url:"/api/role/:id",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        name:"",
        order_num:,
        status:"",
        memo:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 5.3 修改角色

> 基础信息

- **PATH: /role/auth/:id**
- **METHOD: PUT**
> URL 参数 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|id|角色ID|int|1||
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string||鉴权|
>  FORM 参数 

| 参数 | 说明 | 类型 | 校验 | 示例 | 备注 |
| --- | --- | -- | -- | -- |  --  |
|auth|权限列表|[]string||[auth1,auth2...]||

```button
var req = {

    Url:"/role/auth/:id",
    Method:"PUT",
    Header:{
        Authorization:"",
    },
    Form:{
        auth:[auth1,auth2...],
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- |-- | -- |
|data||string|ok|成功|



---

## 5.4 角色查询

> 基础信息

- **PATH: /api/role**
- **METHOD: GET**
>  HEADER 参数 



| 参数 | 说明| 类型 | 示例 | 备注 |
| --- | --- | --- | --- | --- |
|Authorization|token|string|tokenstring|鉴权|
>  QUERY 参数 



| 参数 | 说明 |类型| 示例 | 备注 |
| --- | --- | -- | -- | -- |
|limit|条数|string|10||
|current|页码|string|1||
|code|角色编码|string|||
|name|角色名称|string|||

```button
var req = {

    Url:"/api/role",
    Method:"GET",
    Header:{
        Authorization:"tokenstring",
    },
    Query:{
        limit:"10",
        current:"1",
        code:"",
        name:"",
    },
}
window.Fetch(req);
```

>  200 Response 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|page|分页数据|||参考Page定义|
|data|数据|any||参考Data定义|

>  Page定义 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|limit|条数|int|||
|current|当前页码|int|||
|total|总数|int|||

>  []Data 定义 

| 参数 | 说明 | 类型 | 示例 | 备注 |
| --- | --- | -- | -- | -- |
|code|编码|string|||
|name|名称|string|||
|order_num|排序|int|||
|status|状态|string||0:禁用 1:启用|
|memo|备注|string|||
|auth|权限ID|[]string|||

---