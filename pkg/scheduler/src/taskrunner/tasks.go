package taskrunner

import (
	"os"
	"log"
	"errors"
	"sync"
	"todogo/pkg/scheduler/src/dbops"
)

func DeleteVideo(vid string) error{
	err := os.Remove(VIDEO_PATH + vid)

	if err != nil && !os.IsNotExist(err){ //如果没有移动
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}


func VideoClearDispatcher(dc dataChan) error{
	res, err := dbops.ReadVideoDeletionRecord(3)

	if err!=nil {
		log.Printf("Video clear dispatcher error: %v", err)
	}
	if len(res)==0 {
		return errors.New("All tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error{
	errMap := &sync.Map{}
	var err error

	// 做个标记，无限循环
	forloop:
		for {
			select {
			case vid :=<- dc:
				go func(id interface{}) {
					if err := DeleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
			default:
				break forloop
			}
		}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error) //将v 强制转换成error
		if err != nil {
			return false
		}
		return true
	})

	return err

}

