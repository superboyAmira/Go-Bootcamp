package anomaly

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// интерфейс поддерживающий TxExecutor из postgresql.go
type Loader interface {
	LoadDb(db *gorm.DB) error
}

type AnomalyModel struct {
	gorm.Model
	SessionId string `gorm:"primaryKey"`
	Frequency float64
	Mean      float64
	STD       float64
	Time      *timestamppb.Timestamp
}

func (AnomalyModel) TableName() string {
	return "anomalies"
}

func NewAnomalyModel(
	session string,
	frequency,
	mean,
	std float64,
	time *timestamppb.Timestamp) *AnomalyModel {
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
