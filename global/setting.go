package global

import (
	"github.com/Kurt-Liang/blog-service/pkg/logger"
	"github.com/Kurt-Liang/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
