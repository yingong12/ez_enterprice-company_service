package valuate

import (
	val "company_service/http/request/valuate"
	"company_service/model"
	repository "company_service/repository/valuate"
	"company_service/utils"
	"encoding/json"
	"log"
)

func Create(data val.Create) (err error) {
	taskID := utils.GenerateValuateID()
	j, err := json.Marshal(data)
	if err != nil {
		return
	}
	en := model.ValuateMuttable{
		AppID:    data.AppID,
		State:    data.State,
		FormData: string(j),
	}
	err = repository.Create(en, taskID)
	if err != nil {
		return
	}
	choicesFomatted := parseChoices(data.Choices)
	jj, err := json.Marshal(choicesFomatted)
	if err != nil {
		return err
	}
	//json转化为数组
	partition, offset, err := repository.ProduceTaskMessage(taskID, string(jj), data.ProfitData)
	if err != nil {
		//TODO: 失败日志落盘。 离线任务滚动生产。
		return
	}
	//log
	log.Printf("message produced.  partition:%d offset:%d", partition, offset)
	return
}
func Search(appID string, page, pageSize int) (res []model.Valuate, err error) {
	//search
	return repository.Search(appID, page, pageSize)
}

func parseChoices(choices val.Choices) (res [][]int) {
	res = [][]int{
		choices.CompetitiveLandscape,
		choices.BarrierIndustry,
		choices.IndustryHook,
		choices.Threat,
		choices.PolicySupportPower,
		choices.PolicySupportType,
		choices.Founder,
		choices.ManagementExp,
		choices.AverageWorkingYears,
		choices.UndergraduatesRadio,
		choices.CertificateRadio,
		choices.IsTraining,
		choices.BusinessArea,
		choices.Loyalty,
		choices.IntellectualPropertyNumbers,
		choices.RdInvestmentRadio,
		choices.PrdType,
		choices.CoreTurnoverRadio,
		choices.IsInternalControlSystem,
		choices.IsDishonestyRecord,
		choices.IsRegularPhysicalExamination,
		choices.ExitStrategy,
	}
	for k, v := range res {
		cur := make([]int, utils.ValChoiceToLength[k])
		for _, v1 := range v {
			realV := v1 - 1
			//过滤不合格的
			if realV >= len(cur) || realV < 0 {
				continue
			}
			cur[realV] = 1
		}
		res[k] = cur
	}
	return

}
