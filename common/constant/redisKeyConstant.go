package constant

/**
* 用户注册锁，Key Prefix + 用户名
 */

const LOCK_USER_REGISTER = "index12306-user-service:lock:user-register:"

/**
* 用户注销锁，Key Prefix + 用户名
 */

const USER_DELETION = "index12306-user-service:user-deletion:"

/**
* 用户注册可复用用户名分片，Key Prefix + Idx
 */

const USER_REGISTER_REUSE_SHARDING = "index12306-user-service:user-reuse:"

/**
* 用户乘车人列表，Key Prefix + 用户名
 */

const USER_PASSENGER_LIST = "index12306-user-service:user-passenger-list:"
