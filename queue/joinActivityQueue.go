package queue

import (
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"luck_draw/enums"
	"luck_draw/model"
	"luck_draw/service"
	"luck_draw/util"
	"sync"
	"time"
)

func AttemptJoin(db *gorm.DB,id string)  {
	finish := 0
	msg := "参加失败，请重试"

	defer func() {
		db.Close()
		service.SocketNotify(finish,msg)
	}()

	tx := db.Begin()
	joinLog := &model.JoinLog{}
	err := joinLog.LockById(tx,id)
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		util.ErrDetail(enums.ACTIVITY_DEAL_QUEUE_NOT_FOUND,enums.ActivityQueueNotFound.Error(),id)
		return
	}

	if joinLog.Status != model.JOIN_LOG_STATUS_QUEUE {
		tx.Rollback()
		return
	}

	activity := &model.Activity{}
	err = activity.LockById(tx,joinLog.ActivityId)
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		util.ErrDetail(enums.ACTIVITY_DEAL_QUEUE_A_NOT_FOUND,enums.ActivityQueueANotFound.Error(),id)
		return
	}

	if float32(activity.JoinNum) >= activity.JoinLimitNum {
		data := make(map[string]interface{})
		data["remark"] = "人数已满"
		data["status"] = model.JOIN_LOG_STATUS_FAIL
		err := joinLog.Update(tx,joinLog.ID,data)
		msg = "人数已满，下次抓紧机会啦"
		if err != nil {
			tx.Rollback()
			util.ErrDetail(enums.ACTIVITY_DEAL_QUEUE_UPDATE_LOG_ERR,enums.ActivityJoinLogUpdateFailErr.Error(),id)
			return
		}
	}

	data := make(map[string]interface{})
	data["remark"] = "加入成功"
	data["status"] = model.JOIN_LOG_STATUS_SUCCESS
	err = joinLog.Update(tx,joinLog.ID,data)
	if err != nil {
		tx.Rollback()
		util.ErrDetail(enums.ACTIVITY_DEAL_QUEUE_UPDATE_LOG_ERR,enums.ActivityJoinLogUpdateFailErr.Error(),id)
		return
	}

	activityData := make(map[string]interface{})
	activityData["join_num"] = activity.JoinNum+1
	err = activity.Update(tx,activity.ID,activityData)
	if err != nil {
		tx.Rollback()
		util.ErrDetail(enums.ACTIVITY_DEAL_QUEUE_UPDATE_A_ERR,enums.ActivityUpdateJoinNumFailErr.Error(),id)
		return
	}

	tx.Commit()

	finish = 1
	msg = "加入成功"

	return
}

/**
 * 监听redis
 */
func ListenAttemptJoin(wg *sync.WaitGroup)  {
	redis := util.NewRedis()
	t := time.Second * 59

	queue := enums.ACTIVITY_QUEUE
	redis.OnQueue(wg,queue,t, func(result *redis2.StringSliceCmd, e error) {
		if len(result.Val()) > 0 {
			db,connectErr := model.Connect()
			if connectErr != nil {
				//丢到重试
				return
			}

			util.Info(fmt.Sprintf("取出加入活动的log ID：%v",result.Val()[1]));
			AttemptJoin(db,result.Val()[1])
		}
	})
}

func Listen()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go ListenAttemptJoin(&wg)
	wg.Wait()
	//程序退出，需要通知开发人员
}
