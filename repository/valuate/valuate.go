package valuate

import (
	"company_service/library/env"
	"company_service/model"
	"company_service/providers"
	"encoding/json"

	"github.com/Shopify/sarama"
)

func Search(appID string, page, pageSize int) (res []model.Valuate, err error) {
	res = make([]model.Valuate, 0)
	tx := providers.DBenterprise.Table(model.GetValuateTable())
	tx.
		Where("app_id", appID).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&res)
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

func ProduceTaskMessage(taskID, formData string) (partition int32, offset int64, err error) {
	msgData := map[string]string{
		"taskID": taskID,
		"data":   formData,
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