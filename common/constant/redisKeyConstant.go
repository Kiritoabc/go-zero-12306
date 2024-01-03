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

/**
 * 站点查询，Key Prefix + 起始城市_终点城市_日期
 */

const REGION_TRAIN_STATION = "index12306-ticket-service:region_train_station:%s_%s"

/**
 * 列车基本信息，Key Prefix + 列车ID
 */

const TRAIN_INFO = "index12306-ticket-service:train_info:"

/**
 * 地区与站点映射查询
 */

const REGION_TRAIN_STATION_MAPPING = "index12306-ticket-service:region_train_station_mapping"

/**
 * 站点查询分布式锁 Key
 */

const LOCK_REGION_TRAIN_STATION_MAPPING = "index12306-ticket-service:lock:region_train_station_mapping"

/**
 * 站点查询分布式锁 Key
 */

const LOCK_REGION_TRAIN_STATION = "index12306-ticket-service:lock:region_train_station"

/**
 * 列车站点座位价格查询，Key Prefix + 列车ID_起始城市_终点城市
 */

const TRAIN_STATION_PRICE = "index12306-ticket-service:train_station_price:%s_%s_%s"

/**
 * 列车站点座位价格查询分布式锁 Key
 */

const LOCK_TRAIN_STATION_PRICE = "index12306-ticket-service:lock:train_station_price"

/**
 * 地区以及车站查询，Key Prefix + ( 车站名称 or 查询方式 )
 */

const REGION_STATION = "index12306-ticket-service:region-station:"

/**
 * 站点余票查询，Key Prefix + 列车ID_起始站点_终点
 */

const TRAIN_STATION_REMAINING_TICKET = "index12306-ticket-service:train_station_remaining_ticket:"

/**
 * 列车车厢查询，Key Prefix + 列车ID
 */

const TRAIN_CARRIAGE = "index12306-ticket-service:train_carriage:"

/**
 * 车厢余票查询，Key Prefix + 列车ID_起始站点_终点
 */

const TRAIN_STATION_CARRIAGE_REMAINING_TICKET = "index12306-ticket-service:train_station_carriage_remaining_ticket:"

/**
 * 站点详细信息查询，Key Prefix + 列车ID_起始站点_终点
 */

const TRAIN_STATION_DETAIL = "index12306-ticket-service:train_station_detail:"

/**
 * 列车路线信息查询，Key Prefix + 列车ID
 */

const TRAIN_STATION_STOPOVER_DETAIL = "index12306-ticket-service:train_station_stopover_detail:"

/**
 * 列车站点缓存
 */

const STATION_ALL = "index12306-ticket-service:all_station"

/**
 * 列车车厢状态， Key Prefix + 列车 ID + 起始站点 + 目的站点 + 车厢编号
 */

const TRAIN_CARRIAGE_SEAT_STATUS = "index12306-ticket-service:train_carriage_seat_status:"

/**
 * 用户购票分布式锁 Key
 */

const LOCK_PURCHASE_TICKETS = "${unique-name:}index12306-ticket-service:lock:purchase_tickets_%s"

/**
 * 用户购票分布式锁 Key v2
 */

const LOCK_PURCHASE_TICKETS_V2 = "${unique-name:}index12306-ticket-service:lock:purchase_tickets_%s_%d"

/**
 * 获取全部地点集合 Key
 */

const QUERY_ALL_REGION_LIST = "index12306-ticket-service:query_all_region_list"

/**
 * 列车购买令牌桶，Key Prefix + 列车ID
 */

const TICKET_AVAILABILITY_TOKEN_BUCKET = "index12306-ticket-service:ticket_availability_token_bucket:"

/**
 * 获取全部地点集合分布式锁 Key
 */

const LOCK_QUERY_ALL_REGION_LIST = "index12306-ticket-service:lock:query_all_region_list"

/**
 * 获取列车车厢数量集合分布式锁 Key
 */

const LOCK_QUERY_CARRIAGE_NUMBER_LIST = "index12306-ticket-service:lock:query_carriage_number_list_%s"

/**
 * 获取地区以及站点集合分布式锁 Key
 */

const LOCK_QUERY_REGION_STATION_LIST = "index12306-ticket-service:lock:query_region_station_list_%s"

/**
 * 获取相邻座位余票分布式锁 Key
 */

const LOCK_SAFE_LOAD_SEAT_MARGIN_GET = "index12306-ticket-service:lock:safe_load_seat_margin_%s"

/**
 * 列车购买令牌桶加载数据 Key
 */

const LOCK_TICKET_AVAILABILITY_TOKEN_BUCKET = "index12306-ticket-service:lock:ticket_availability_token_bucket:%s"
