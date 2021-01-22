package enums

const ACTIVITY_QUEUE				= "luck_activity_queue"  	 //参加活动队列
const ACTIVITY_QUEUE_TRY			= "luck_activity_queue_try"  //参加活动重试队列
const SOCKET_USER_TOKEN  			= "luck_api_socket_token"	 //用户socket token

//业务不相关错误
const (
	SUCCESS					= 0
	FAIL 					= 1
	DB_CONNECT_ERR 			= 2
	READ_CONFIG_ERR			= 3
	DECODE_ARR_ERR			= 4
)

//授权相关 1000 ~ 1999
//参数错误
const (
	AUTH_PARAMS_ERROR 				= 1000		//登录参数错误
	AUTH_LOGIN_TYPE_ERR				= 1001 		//登录类型错误
	AUTH_REQUEST_SESSION_ERR		= 1002		//请求微信session出错
	AUTH_REQUEST_SESSION_RESP_ERR	= 1003		//请求微信session返回错误异常
	AUTH_PARSE_JSON_ERR				= 1004		//解析用户json失败
	AUTH_USER_QUERY_ERR				= 1005		//用户查询错误
	AUTH_USER_SAVE_ERR				= 1006		//新增用户数据库异常
	AUTH_USER_UPDATE_ERR			= 1007		//用户数据更新失败
	AUTH_USER_PARSE_JWT_ERR			= 1008		//解析json失败
	AUTH_USER_CREATE_JWT_ERR		= 1009		//生成jwt异常
	AUTH_TOKEN_EXPIRED				= 1010		//token已过有效期
	AUTH_TOKEN_NULL					= 1011		//token为空
	AUTH_NOT_LOGIN					= 1012		//未登录
	Auth_TRANS_UID_ERR				= 1013		//userId类型转化失败
	AUTH_USER_NOT_FOUND				= 1014		//用户不存在
)

//活动相关 2000 ~ 2999
const (
	ACTIVITY_PARAM_ERR 					= 2000 		//参数错误
	ACTIVITY_START_DATE_ERR 			= 2001 		//活动开始如期解析错误
	ACTIVITY_END_DATE_ERR 				= 2002 		//活动截止日期解析错误
	ACTIVITY_RUN_DATE_ERR 				= 2003 		//活动开奖日期解析错误
	ACTIVITY_SAVE_ERR 					= 2004 		//活动保存失败
	ACTIVITY_PAGE_ERR					= 2005		//分页查询错误
	ACTIVITY_DETAIL_PARAM_ERR			= 2006		//详情id不能为空
	ACTIVITY_DETAIL_QUERY_ERR			= 2007		//详情查询错误
	ACTIVITY_DETAIL_NOT_FOUND			= 2008		//详情不存在
	ACTIVITY_JOIN_PARAM_ERR				= 2009		//参团参数失败，id为空
	ACTIVITY_JOIN_LIMIT					= 2010		//活动参与人数达到限制啦
	ACTIVITY_JOIN_SAVE_LOG_FAIL			= 2011		//参加活动失败
	ACTIVITY_JOIN_REPEAT				= 2012		//您已参加该活动，不可重复参加
	ACTIVITY_JOIN_QUERY_ERR				= 2013		//查询参与日志出错
	ACTIVITY_PUSH_QUEUE_ERR				= 2014		//参加活动写入队列失败
	ACTIVITY_DEAL_QUEUE_NOT_FOUND		= 2015		//处理参加活动队列的记录不存在
	ACTIVITY_DEAL_QUEUE_A_NOT_FOUND		= 2016		//处理参加活动队列的活动记录不存在
	ACTIVITY_DEAL_QUEUE_UPDATE_LOG_ERR	= 2017		//更新活动参与记录因为加入活动因为人数已满失败出错
	ACTIVITY_DEAL_QUEUE_UPDATE_A_ERR	= 2018		//更新活动参与人数出错
	ACTIVITY_STATUS_NOT_RUNNING	    	= 2019		//活动不是可参加状态
	ACTIVITY_MEMBER_ENOUTH		    	= 2020		//活动参加人数已满
)

//礼品相关 3000 ~ 3999
const (
	GIFT_SAVE_ERR					= 3000 		//礼品保存失败
	GIFT_FIRST_ERR					= 3001 		//礼品查询出错
	GIFT_NOT_FOUND					= 3002 		//礼品不存在
	GIFT_GET_DETAIL_ERR				= 3003 		//礼品详情查询错误
)

//socket相关 4000 ~ 4999
const (
	SOCKET_ENCRYPE_ERR				= 4000		//生成签名出错
	SOCKET_SIGN_ENCODE_ERR			= 4001		//sign转成json异常
	SOCKET_POST_SIGN_ERR			= 4002		//请求授权网络出错
	SOCKET_AUTH_ERR					= 4003		//请求授权返回失败
)


