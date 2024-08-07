package anomaly

import (
	"time"

	"gorm.io/gorm"
)

// интерфейс поддерживающий TxExecutor из postgresql.go
type Loader interface {
	LoadDb(db *gorm.DB) error
}

type AnomalyModel struct {
	SessionId string
	Frequency float64
	Mean      float64
	STD       float64
	Time      time.Time
}

func (AnomalyModel) TableName() string {
	return "anomalies"
}

func NewAnomalyModel(
	session string,
	frequency,
	mean,
	std float64,
	time time.Time) *AnomalyModel {
	return &AnomalyModel{
		SessionId: session,
		Frequency: frequency,
		Mean:      mean,
		STD:       std,
		Time:      time,
	}
}

func (r *AnomalyModel) LoadDb(db *gorm.DB) error {
	db.AutoMigrate(AnomalyModel{})
	if result := db.Create(r); result.Error != nil {
		return result.Error
	}
	return nil
}
