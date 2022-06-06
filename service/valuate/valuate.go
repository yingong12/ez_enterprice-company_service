package valuate

import (
	"company_service/logger"
	"company_service/model"
	repository "company_service/repository/valuate"
	"company_service/utils"
)

func Create(data model.ValuateMuttable) (err error) {
	/*
		生成valuate_id
			1. 写KAFKA
			2. 写DB更新为state=1
	*/
	taskID := utils.GenerateValuateID()
	err = repository.Create(data, taskID)
	if err != nil {
		return
	}
	partition, offset, err := repository.ProduceTaskMessage(taskID, data.FormData)
	if err != nil {
		//TODO: 失败日志落盘。 离线任务滚动生产。
		return
	}
	//log
	logger.Info("message produced.  partition:%d offset:%d", partition, offset)
	return
}
func Search(appID string, page, pageSize int) (res []model.Valuate, err error) {
	//search
	return repository.Search(appID, page, pageSize)
}
