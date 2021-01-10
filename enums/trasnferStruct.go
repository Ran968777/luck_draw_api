package enums

import (
	"luck_draw/model"
	"time"
)

//微信登录参数
type WxMiniLoginData struct {
	Iv string `form:"iv" json:"iv" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
	EncryptedData string `form:"encrypted_data" json:"encrypted_data" binding:"required"`
}

//活动创建参数
type ActivityCreateParam struct {
	Name 			string 		`form:"name" json:"name" binding:"required"`
	GiftId 			int64 		`form:"gift_id" json:"gift_id" binding:"required"`
	LimitJoin 		int32 	 	`form:"limit_join" json:"limit_join" binding:"required"` 			//是否限制参加人数
	JoinLimitNum 	float32 	`form:"join_limit_num" json:"join_limit_num" binding:"required"` 	//限制参加人数
	ReceiveLimit 	float32 	`form:"receive_limit" json:"receive_limit" binding:"required"` 		//每人限领数量
	Des 			string      `form:"des" json:"des" binding:"required"`
	Attachments 	string   	`form:"attachments" json:"attachments" binding:"required"`
	StartAt 		string    	`form:"start_at" json:"start_at" binding:"required"`				//活动开始时间
	EndAt 			string      `form:"end_at" json:"end_at" binding:"required"`					//活动截止时间
	RunAt 			string      `form:"run_at" json:"run_at" binding:"required"`					//开奖时间
	ShareTitle 		string    	`form:"share_title" json:"share_title"` 							//分享标题
	ShareImage 		string    	`form:"share_image" json:"share_image"` 							//分享图片
}

//活动详情返回参数
type ActivityDetailFormat struct {
	ID        		uint
	Name 			string
	GiftId 			int64
	Type 			int8
	FromType 		int8
	JoinNum 		int32
	LimitJoin 		int32
	JoinLimitNum 	float32
	Des 			string
	Attachments 	string
	Status 			int8
	ShareTitle 		string
	ShareImage 		string
	CreatedAt 		time.Time
	Gift      		*GiftDetail
}

//活动分页
type  ActivityPageFormat struct {
	ID        		uint
	Name 			string
	GiftId 			int64
	Type 			int8   		 	//活动类型
	FromType 		int32   		 //发布活动的用户类型
	JoinNum 		int32 		   	//已参加人数
	JoinLimitNum 	float32 	 	//限制参加人数
	//Attachments 	string
	Status 			int8		 	//活动状态
	Gift			*model.Gift
}

type GiftParam struct {
	Name 		string 		`form:"name" json:"name"`
	Num 		float32 	`form:"num" json:"num"`
	Type 		int8   		`form:"type" json:"type"` 				//奖品类型，1=红包，2=商品，3=话费
	FROM        int8   		`form:"from" json:"from"`   			//奖品来源，1=平台，2=用户
	STATUS      int8   		`form:"status" json:"status"` 			//奖品状态，1=上架，2=下架，下架不可用
	Des    		string      `form:"describe" json:"des"`
	Attachments string  	`form:"attachment" json:"attachment"`
}

type GiftDetail struct {
	ID			uint
	Name 		string
	UserId 		int
	Num 		float32
	Type 		int8
	Des    		string
	Attachments string
}