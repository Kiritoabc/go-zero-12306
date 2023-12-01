package xerr

var message map[string]string

func init() {
	message = make(map[string]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"

	//用户注册错误信息
	message[USER_REGISTER_FAIL] = "用户注册失败"
	message[USER_NAME_NOTNULL] = "用户名不能为空"
	message[PASSWORD_NOTNULL] = "密码不能为空"
	message[PHONE_NOTNULL] = "手机号不能为空"
	message[ID_TYPE_NOTNULL] = "证件类型不能为空"
	message[ID_CARD_NOTNULL] = "证件号不能为空"
	message[HAS_USERNAME_NOTNULL] = "用户名已存在"
	message[PHONE_REGISTERED] = "手机号已被占用"
	message[MAIL_REGISTERED] = "邮箱已被占用"
	message[MAIL_NOTNULL] = "邮箱不能为空"
	message[USER_TYPE_NOTNULL] = "旅客类型不能为空"
	message[POST_CODE_NOTNULL] = "邮编不能为空"
	message[ADDRESS_NOTNULL] = "地址不能为空"
	message[REGION_NOTNULL] = "国家/地区不能为空"
	message[TELEPHONE_NOTNULL] = "固定电话不能为空"
	message[VERIFY_STATE_NOTNULL] = "审核状态不能为空"
	message[REAL_NAME_NOTNULL] = "真实姓名不能为空"

	// 用户登录错误
	message[LOGIN_MAIL_NOT_EXIST] = "用户名/手机号/邮箱不存在"
	message[NAME_ERROR] = "用户名错误"
}

func MapErrMsg(errcode string) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode string) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
