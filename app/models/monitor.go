package models

import (
	"github.com/goravel/framework/database/orm"

	"github.com/TheTNB/panel/pkg/tools"
)

type Monitor struct {
	orm.Model
	Info tools.MonitoringInfo `gorm:"not null;serializer:json" json:"info"`
}
