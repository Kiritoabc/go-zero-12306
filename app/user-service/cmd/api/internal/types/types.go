// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id           int64  `json:"id"`           // ID
	Username     string `json:"userName"`     // 用户名
	RealName     string `json:"realName"`     // 真实姓名
	Region       string `json:"region"`       // 国家/地区
	IdType       int64  `json:"idType"`       // 证件类型
	IdCard       string `json:"idCard"`       // 证件号
	Phone        string `json:"phone"`        // 手机号
	Telephone    string `json:"telephone"`    // 固定电话
	Mail         string `json:"mail"`         // 邮箱
	UserType     int64  `json:"userType"`     // 旅客类型
	VerifyStatus int64  `json:"verifyStatus"` // 审核状态
	PostCode     string `json:"postCode"`     // 邮编
	Address      string `json:"address"`      // 地址
}

type LoginReq struct {
	UserNameOrMailOrPhone string `json:"userNameOrMailOrPhone"`
	Password              string `json:"password"`
}

type LoginResp struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	RealName    string `json:"realName"`
	AccessToken string `json:"accessToken"`
}

type CheckLoginReq struct {
	AccessToken string `form:"accessToken"`
}

type CheckLoginResp struct {
	UserId      string `json:"userId"`      // 用户id
	UserName    string `json:"userName"`    // 用户名
	RealName    string `json:"realName"`    // 真实姓名
	AccessToken string `json:"accessToken"` // Token
}

type AccessToken struct {
	AccessToken string `form:"accessToken"`
}

type LogoutReq struct {
	AccessToken string `form:"accessToken"`
}

type LogoutResp struct {
}

type RegisterReq struct {
	UserName    string `json:"userName"`    // 用户命
	Password    string `json:"password"`    // 密码
	RealName    string `json:"realName"`    // 真实姓名
	IdType      int64  `json:"idType"`      // 证件类型
	IdCard      string `json:"idCard"`      //  证件号
	Phone       string `json:"phone"`       // 手机号
	Mail        string `json:"mail"`        // 邮箱
	UserType    int64  `json:"userType"`    // 旅客类型
	VerifyState int64  `json:"verifyState"` // 审核状态
	PostCode    string `json:"postCode"`    // 邮编
	Address     string `json:"address"`     // 地址
	Region      string `json:"region"`      // 国家地区
	Telephone   string `json:"telephone"`   // 固定电话
}

type RegisterResp struct {
	UserName string `json:"userName"` // 用户名
	RealName string `json:"realName"` // 真实姓名
	Phone    string `json:"phone"`    // 手机号
}

type QueryUserReq struct {
	UserName string `form:"userName"`
}

type QueryUserResp struct {
	UserName    string `json:"userName"`    // 用户命
	RealName    string `json:"realName"`    // 真实姓名
	IdType      int64  `json:"idType"`      // 证件类型
	IdCard      string `json:"idCard"`      //  证件号
	Phone       string `json:"phone"`       // 手机号
	Mail        string `json:"mail"`        // 邮箱
	UserType    int64  `json:"userType"`    // 旅客类型
	VerifyState int64  `json:"verifyState"` // 审核状态
	PostCode    string `json:"postCode"`    // 邮编
	Address     string `json:"address"`     // 地址
	Region      string `json:"region"`      // 国家地区
	Telephone   string `json:"telephone"`   // 固定电话
}

type UserQueryActualReq struct {
	UserName string `form:"userName"`
}

type UserQueryActualResp struct {
	UserName    string `json:"userName"`    // 用户命
	RealName    string `json:"realName"`    // 真实姓名
	IdType      int64  `json:"idType"`      // 证件类型
	IdCard      string `json:"idCard"`      //  证件号
	Phone       string `json:"phone"`       // 手机号
	Mail        string `json:"mail"`        // 邮箱
	UserType    int64  `json:"userType"`    // 旅客类型
	VerifyState int64  `json:"verifyState"` // 审核状态
	PostCode    string `json:"postCode"`    // 邮编
	Address     string `json:"address"`     // 地址
	Region      string `json:"region"`      // 国家地区
	Telephone   string `json:"telephone"`   // 固定电话
}

type UserHasUsernameReq struct {
	UserName string `form:"userName"`
}

type UserHasUsernameResp struct {
	Bool bool `json:"bool"`
}

type UserUpdateReq struct {
	Id       string `json:"id"`       // 用户ID
	UserName string `json:"userName"` // 用户名
	Mail     string `json:"mail"`     // 邮箱
	UserType int64  `json:"userType"` // 旅客类型
	PostCode string `json:"postCode"` // 邮编
	Address  string `json:"address"`  // 地址
}

type UserUpdateResp struct {
}

type UserDeleteReq struct {
	UserName string `json:"userName"` // 用户名
}

type UserDeleteResp struct {
}

type PassengerReq struct {
}

type PassengerRespDTO struct {
	Id           string `json:"id"`           // 乘车人id
	UserName     string `json:"userName"`     // 用户名
	RealName     string `json:"realName"`     // 真实姓名
	IdType       int64  `json:"idType"`       // 证件类型
	IdCard       string `json:"idCard"`       // 证件号类型
	ActualIdCard string `json:"actualIdCard"` // 真实证件号码
	DiscountType int64  `json:"discountType"` // 优惠类型
	Phone        string `json:"phone"`        // 真实手机号
	CreateDate   string `json:"createDate"`   // 添加日期
	VerifyStatus int64  `json:"verifyStatus"` // 审核状态
}

type PassengerResp struct {
	List []PassengerRespDTO `json:"list"`
}

type PassengerQueryReq struct {
	UserName string  `form:"userName"`
	Ids      []int64 `form:"ids"`
}

type PassengerActualRespDTO struct {
	Id           string `json:"id"`           // 乘车人id
	UserName     string `json:"userName"`     // 用户名
	RealName     string `json:"realName"`     // 真实姓名
	IdType       int64  `json:"idType"`       // 证件类型
	IdCard       string `json:"idCard"`       // 证件号码
	DiscountType int64  `json:"discountType"` // 优惠类型
	Phone        string `json:"phone"`        // 手机号
	CreateDate   string `json:"createDate"`   //添加日期
	VerifyStatus int64  `json:"verifyStatus"` // 审核状态
}

type PassengerQueryResp struct {
	List []PassengerActualRespDTO `json:"list"`
}

type SavePassengerReq struct {
	Id           string `json:"id"`           // 乘车人id
	RealName     string `json:"realName"`     // 真实姓名
	IdType       int64  `json:"idType"`       // 证件类型
	IdCard       string `json:"idCard"`       // 证件号类型
	DiscountType int64  `json:"discountType"` // 优惠类型
	Phone        string `json:"phone"`        // 真实手机号
}

type SavePassengerResp struct {
}

type UpdatePassengerReq struct {
	Id           string `json:"id"`           // 乘车人id
	RealName     string `json:"realName"`     // 真实姓名
	IdType       int64  `json:"idType"`       // 证件类型
	IdCard       string `json:"idCard"`       // 证件号类型
	DiscountType int64  `json:"discountType"` // 优惠类型
	Phone        string `json:"phone"`        // 真实手机号
}

type UpdatePassengerResp struct {
}

type RemovePassengerReq struct {
	Id string `json:"id"`
}

type RemovePassengerResp struct {
}
