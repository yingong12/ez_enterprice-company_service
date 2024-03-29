package valuate

import (
	"company_service/library/env"
	"company_service/model"
	"company_service/providers"
	"encoding/json"

	"github.com/Shopify/sarama"
	"gorm.io/gorm"
)

func Search(tx *gorm.DB, appID string, page, pageSize int) (res []model.Valuate, err error) {
	res = make([]model.Valuate, 0)
	tx = tx.Table(model.GetValuateTable()).
		Where("app_id", appID).
		Where("state", 1).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&res)
	err = tx.Error
	return
}
func Total(tx *gorm.DB, appID string, page, pageSize int) (total int64, err error) {
	tx = tx.Table(model.GetValuateTable()).
		Where("app_id", appID).
		Where("state", 1).
		Count(&total)
	err = tx.Error
	return
}
func Create(data model.ValuateMuttable, valID string) (err error) {
	tx := providers.DBenterprise.Table(model.GetValuateTable())
	en := model.Valuate{}
	en.ValuateID = valID
	en.ValuateMuttable = data
	tx.Create(en)
	err = tx.Error
	return
}

func ProduceTaskMessage(taskID string, choices string, data string, enInfo map[string]interface{}, shareholders string) (partition int32, offset int64, err error) {
	msgData := map[string]interface{}{
		"task_id":       taskID,
		"data":          data,
		"choice":        choices,
		"en_data":       enInfo,
		"share_holders": shareholders,
	}
	j, err := json.Marshal(msgData)
	if err != nil {
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: env.GetStringVal("KAFKA_TOPIC_VALUATE"),
		Value: sarama.ByteEncoder(j),
	}
	partition, offset, err = providers.ValProducer.SendMessage(msg)
	return
}
