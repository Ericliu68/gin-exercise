package models

import (
	"time"
)

// user 表
type Userlogin struct {
	Id          int        `gorm:"column:id;primaryKey"`
	Name        string     `gorm:"column:name"`
	Age         int        `gorm:"column:age"`
	Create_time *time.Time `gorm:"column:create_time"`
	Password    string     `gorm:"column:password"`
}

// yaml 解析
type SettingYaml struct {
	Info   InfoYaml   `yaml:"infra"`
	Common CommonYaml `yaml:"common"`
}
type InfoYaml struct {
	MongoUri            string   `yaml:"MONGO_URI"`
	KafkaHosts          []string `yaml:"KAFKA_HOSTS"`
	EsHosts             string   `yaml:"ES_HOSTS"`
	RedisPushWebsite    string   `yaml:"REDIS_PUSH_WEBSITE"`
	RedisPushFeature    string   `yaml:"REDIS_PUSH_FEATURE"`
	RedisPushStatistics string   `yaml:"REDIS_PUSH_STATISTICS"`
	CeleryBrokerUrl     string   `yaml:"CELERY_BROKER_URL"`
	CeleryResultBackend string   `yaml:"CELERY_RESULT_BACKEND"`
}

type CommonYaml struct {
	ServiceUser            string `yaml:"SERVICE_USER"`
	KafkaHitsTopic         string `yaml:"KAFKA_HITS_TOPIC"`
	KafkaGroupId           string `yaml:"KAFKA_GROUP_ID"`
	KafkaStatisticsGroupId string `yaml:"KAFKA_STATISTICS_GROUP_ID"`
	EsHitsIndex            string `yaml:"ES_HITS_INDEX"`
	MongodbName            string `yaml:"MONGO_DB_NAME"`
	EsIndex                string `yaml:"ES_INDEX"`
	EsResultIndex          string `yaml:"ES_RESUlT_INDEX"`
	PaginationNum          string `yaml:"PAGINATION_NUM"`
	SchedulerTimezone      string `yaml:"SCHEDULER_TIMEZONE"`
	ServiceFeature         string `yaml:"SERVICE_FEATURE"`
	EsTimeFrommat          string `yaml:"ES_TIME_FROMMAT"`
	EsHttpLogIndex         string `yaml:"ES_HTTP_LOG_INDEX"`
	EsHitsResultIndex      string `yaml:"ES_HITS_RESULT_INDEX"`
	EsScanResultIndex      string `yaml:"ES_SCAN_RESULT_INDEX"`
	KafkaHttpLogDelimiter  string `yaml:"KAFKA_HTTPLOG_DELIMITER"`
	KafkaHitsResultTopic   string `yaml:"KAFKA_HITS_RESULT_TOPIC"`
	DirSend                string `yaml:"DIR_SEND"`
	DirReceive             string `yaml:"DIR_RECEIVE"`
	KafkaGroupIdStat       string `yaml:"KAFKA_GROUP_ID_STAT"`
	KafkaAlarmTopic        string `yaml:"KAFKA_ALARM_TOPIC"`
	KafkaAlarmGroupId      string `yaml:"KAFKA_ALARM_GROUP_ID"`
	KafkaBlockTopic        string `yaml:"KAFKA_BLOCK_TOPIC"`
	KafkaBlockGroupId      string `yaml:"KAFKA_BLOCK_GROUP_ID"`
}
