package valuate

import (
	val "company_service/http/request/valuate"
	"company_service/model"
	"company_service/providers"
	enterprise "company_service/repository"
	repository "company_service/repository/valuate"
	"company_service/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sync"
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
	//写估值表
	err = repository.Create(en, taskID)
	if err != nil {
		return
	}
	//查询企业信息,并添加到kafka msg中
	res, err := enterprise.GetEnterpriseByAppIDs([]string{data.AppID})
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("没找到相关企业信息")
		return
	}
	//TODO:插入企业信息
	enter := res[0]
	enData := map[string]interface{}{
		"district":             "",
		"industry":             "",
		"register_number":      enter.RegistrationNumber,
		"legal_representative": enter.LegalRepresentative,
		"business_scope":       enter.BusinessScope,
	}
	d := utils.DFSDistrict(&providers.DisrictDict, enter.District)
	if d != nil {
		enData["district"] = d.Label
	}
	i := utils.DFSIndustry(&providers.IndustryDict, enter.Industry)
	if i != nil {
		enData["industry"] = i.Label
	}
	shInfo := enter.ShareHolders
	choicesFomatted := parseChoices(data.Choices)
	jj, err := json.Marshal(choicesFomatted)
	if err != nil {
		return err
	}
	//json转化为数组
	partition, offset, err := repository.ProduceTaskMessage(taskID, string(jj), data.ProfitData, enData, shInfo)
	if err != nil {
		//TODO: 失败日志落盘。 离线任务滚动生产。
		return
	}
	//log
	log.Printf("val message produced.  partition:%d offset:%d", partition, offset)
	return
}
func getStaticFileAsString(taskID string) (res string, err error) {
	rsp, err := providers.HttpClientStatic.Get(providers.HttpClientStatic.BaseURL + taskID + ".json")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(rsp.Body)
	res = string(b)
	log.Println(res)
	return
}
func Search(appID string, page, pageSize int) (res []model.Valuate, err error) {
	//search
	list, err := repository.Search(appID, page, pageSize)
	if err != nil || len(list) == 0 {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(list))
	//
	for i := 0; i < len(list); i++ {
		go func(fn string, idx int) {
			var err error
			var fileStr string
			defer func() {
				if err != nil {
					log.Println("估值结果查询失败", err)
				}
				list[idx].Result = fileStr
				wg.Done()
			}()
			fileStr, err = getStaticFileAsString(fn)
			return
		}(list[i].ValuateID, i)
	}
	wg.Wait()
	res = list
	return
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
		choices.MarketingMethod,
		choices.BusinessProcessing,
		choices.Loyalty,
		choices.IntellectualPropertyNumbers,
		choices.RdInvestmentRadio,
		choices.PrdType,
		choices.InnovationLevel,
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
