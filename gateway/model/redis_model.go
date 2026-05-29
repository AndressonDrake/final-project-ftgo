package model

type RedisProducer struct {
	TrxType string      `json:"trx_type"`
	SubType string      `json:"sub_type"`
	Data    interface{} `json:"data"`
	Email   string      `json:"email"`
}
