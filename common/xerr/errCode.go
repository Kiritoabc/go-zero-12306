package xerr

// 成功返回
const OK = "200"

/**(前3位代表业务模块,后三位代表具体功能)**/

// 全局错误码 100 开头
const SERVER_COMMON_ERROR = "100001"
const REUQEST_PARAM_ERROR = "100002"
const TOKEN_EXPIRE_ERROR = "100003"
const TOKEN_GENERATE_ERROR = "100004"
const DB_ERROR = "100005"
const DB_UPDATE_AFFECTED_ZERO_ERROR = "100006"

// 用户注册错误
const USER_REGISTER_FAIL = "A006000" //("A006000", "用户注册失败"),

const USER_NAME_NOTNULL = "A006001" //("A006001", "用户名不能为空"),

const PASSWORD_NOTNULL = "A006002" //("A006002", "密码不能为空"),

const PHONE_NOTNULL = "A006003" //("A006003", "手机号不能为空"),

const ID_TYPE_NOTNULL = "A006004" //("A006004", "证件类型不能为空"),

const ID_CARD_NOTNULL = "A006005" //("A006005", "证件号不能为空"),

const HAS_USERNAME_NOTNULL = "A006006" //("A006006", "用户名已存在"),

const PHONE_REGISTERED = "A006007" //("A006007", "手机号已被占用"),

const MAIL_REGISTERED = "A006008" //("A006008", "邮箱已被占用"),

const MAIL_NOTNULL = "A006009" //("A006009", "邮箱不能为空"),

const USER_TYPE_NOTNULL = "A006010" //("A006010", "旅客类型不能为空"),

const POST_CODE_NOTNULL = "A006011" //("A006011", "邮编不能为空"),

const ADDRESS_NOTNULL = "A006012" //("A006012", "地址不能为空"),

const REGION_NOTNULL = "A006012" //("A006012", "国家/地区不能为空"),

const TELEPHONE_NOTNULL = "A006013" //("A006013", "固定电话不能为空"),

const VERIFY_STATE_NOTNULL = "A006014" //("A006014", "审核状态不能为空"),

const REAL_NAME_NOTNULL = "A006015" //("A006015", "真实姓名不能为空");

// 用户登录错误
const LOGIN_MAIL_NOT_EXIST = "A006001" //用户登录时邮箱不存在
