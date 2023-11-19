### 1. "login"

1. route definition

- Url: /api/user-service/v1/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	UserNameOrMailOrPhone string `json:"userNameOrMailOrPhone"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	RealName string `json:"realName"`
	AccessToken string `json:"accessToken"`
}
```

### 2. "check login"

1. route definition

- Url: /api/user-service/check-login
- Method: GET
- Request: `CheckLoginReq`
- Response: `CheckLoginResp`

2. request definition



```golang
type CheckLoginReq struct {
	AccessToken string `form:"accessToken"`
}
```


3. response definition



```golang
type CheckLoginResp struct {
	UserId string `json:"userId"` // 用户id
	UserName string `json:"userName"` // 用户名
	RealName string `json:"realName"` // 真实姓名
	AccessToken string `json:"accessToken"` // Token
}
```

### 3. "logout"

1. route definition

- Url: /api/user-service/logout
- Method: GET
- Request: `LogoutReq`
- Response: `LogoutResp`

2. request definition



```golang
type LogoutReq struct {
	AccessToken string `form:"accessToken"`
}
```


3. response definition



```golang
type LogoutResp struct {
}
```

### 4. "根据用户名查询用户信息"

1. route definition

- Url: /api/user-service/query
- Method: GET
- Request: `QueryUserReq`
- Response: `QueryUserResp`

2. request definition



```golang
type QueryUserReq struct {
	UserName string `form:"userName"`
}
```


3. response definition



```golang
type QueryUserResp struct {
	UserName string `json:"userName"` // 用户命
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` //  证件号
	Phone string `json:"phone"` // 手机号
	Mail string `json:"mail"` // 邮箱
	UserType int64 `json:"userType"` // 旅客类型
	VerifyState int64 `json:"verifyState"` // 审核状态
	PostCode string `json:"postCode"` // 邮编
	Address string `json:"address"` // 地址
	Region string `json:"region"` // 国家地区
	Telephone string `json:"telephone"` // 固定电话
}
```

### 5. "根据用户名查询用户无脱敏信息"

1. route definition

- Url: /api/user-service/actual/query
- Method: GET
- Request: `UserQueryActualReq`
- Response: `UserQueryActualResp`

2. request definition



```golang
type UserQueryActualReq struct {
	UserName string `form:"userName"`
}
```


3. response definition



```golang
type UserQueryActualResp struct {
	UserName string `json:"userName"` // 用户命
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` //  证件号
	Phone string `json:"phone"` // 手机号
	Mail string `json:"mail"` // 邮箱
	UserType int64 `json:"userType"` // 旅客类型
	VerifyState int64 `json:"verifyState"` // 审核状态
	PostCode string `json:"postCode"` // 邮编
	Address string `json:"address"` // 地址
	Region string `json:"region"` // 国家地区
	Telephone string `json:"telephone"` // 固定电话
}
```

### 6. "检查用户名是否已存在"

1. route definition

- Url: /api/user-service/has-username
- Method: GET
- Request: `UserHasUsernameReq`
- Response: `UserHasUsernameResp`

2. request definition



```golang
type UserHasUsernameReq struct {
	UserName string `form:"userName"`
}
```


3. response definition



```golang
type UserHasUsernameResp struct {
	Bool bool `json:"bool"`
}
```

### 7. "修改用户"

1. route definition

- Url: /api/user-service/update
- Method: POST
- Request: `UserUpdateReq`
- Response: `UserUpdateResp`

2. request definition



```golang
type UserUpdateReq struct {
	Id string `json:"id"` // 用户ID
	UserName string `json:"userName"` // 用户名
	Mail string `json:"mail"` // 邮箱
	UserType int64 `json:"userType"` // 旅客类型
	PostCode string `json:"postCode"` // 邮编
	Address string `json:"address"` // 地址
}
```


3. response definition



```golang
type UserUpdateResp struct {
}
```

### 8. "注销用户"

1. route definition

- Url: /api/user-service/deletion
- Method: POST
- Request: `UserDeleteReq`
- Response: `UserDeleteResp`

2. request definition



```golang
type UserDeleteReq struct {
	UserName string `json:"userName"` // 用户名
}
```


3. response definition



```golang
type UserDeleteResp struct {
}
```

### 9. "register"

1. route definition

- Url: /api/user-service/v1/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterResp`

2. request definition



```golang
type RegisterReq struct {
	UserName string `json:"userName"` // 用户命
	Password string `json:"password"` // 密码
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` //  证件号
	Phone string `json:"phone"` // 手机号
	Mail string `json:"mail"` // 邮箱
	UserType int64 `json:"userType"` // 旅客类型
	VerifyState int64 `json:"verifyState"` // 审核状态
	PostCode string `json:"postCode"` // 邮编
	Address string `json:"address"` // 地址
	Region string `json:"region"` // 国家地区
	Telephone string `json:"telephone"` // 固定电话
}
```


3. response definition



```golang
type RegisterResp struct {
	UserName string `json:"userName"` // 用户名
	RealName string `json:"realName"` // 真实姓名
	Phone string `json:"phone"` // 手机号
}
```

### 10. "根据用户名查询乘车人列表"

1. route definition

- Url: /api/user-service/passenger/query
- Method: GET
- Request: `PassengerReq`
- Response: `PassengerRespDTO`

2. request definition



```golang
type PassengerReq struct {
}
```


3. response definition



```golang
type PassengerRespDTO struct {
	Id string `json:"id"` // 乘车人id
	UserName string `json:"userName"` // 用户名
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` // 证件号类型
	ActualIdCard string `json:"actualIdCard"` // 真实证件号码
	DiscountType int64 `json:"discountType"` // 优惠类型
	Phone string `json:"phone"` // 真实手机号
	CreateDate string `json:"createDate"` // 添加日期
	VerifyStatus int64 `json:"verifyStatus"` // 审核状态
}
```

### 11. "根据乘车人 ID 集合查询乘车人列表"

1. route definition

- Url: /api/user-service/inner/passenger/actual/query/ids
- Method: GET
- Request: `PassengerQueryReq`
- Response: `PassengerActualRespDTO`

2. request definition



```golang
type PassengerQueryReq struct {
	UserName string `form:"userName"`
	Ids []int64 `form:"ids"`
}
```


3. response definition



```golang
type PassengerActualRespDTO struct {
	Id string `json:"id"` // 乘车人id
	UserName string `json:"userName"` // 用户名
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` // 证件号码
	DiscountType int64 `json:"discountType"` // 优惠类型
	Phone string `json:"phone"` // 手机号
	CreateDate string `json:"createDate"` //添加日期
	VerifyStatus int64 `json:"verifyStatus"` // 审核状态
}
```

### 12. "新增乘车人"

1. route definition

- Url: /api/user-service/passenger/save
- Method: POST
- Request: `SavePassengerReq`
- Response: `SavePassengerResp`

2. request definition



```golang
type SavePassengerReq struct {
	Id string `json:"id"` // 乘车人id
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` // 证件号类型
	DiscountType int64 `json:"discountType"` // 优惠类型
	Phone string `json:"phone"` // 真实手机号
}
```


3. response definition



```golang
type SavePassengerResp struct {
}
```

### 13. "修改乘车人"

1. route definition

- Url: /api/user-service/passenger/update
- Method: POST
- Request: `UpdatePassengerReq`
- Response: `UpdatePassengerResp`

2. request definition



```golang
type UpdatePassengerReq struct {
	Id string `json:"id"` // 乘车人id
	RealName string `json:"realName"` // 真实姓名
	IdType int64 `json:"idType"` // 证件类型
	IdCard string `json:"idCard"` // 证件号类型
	DiscountType int64 `json:"discountType"` // 优惠类型
	Phone string `json:"phone"` // 真实手机号
}
```


3. response definition



```golang
type UpdatePassengerResp struct {
}
```

### 14. "移除乘车人"

1. route definition

- Url: /api/user-service/passenger/remove
- Method: POST
- Request: `RemovePassengerReq`
- Response: `RemovePassengerResp`

2. request definition



```golang
type RemovePassengerReq struct {
	Id string `json:"id"`
}
```


3. response definition



```golang
type RemovePassengerResp struct {
}
```

