package model

import (
	"fmt"
	"time"

	"github.com/Kurt-Liang/blog-service/global"
	"github.com/Kurt-Liang/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Error == nil {
		nowTime := time.Now().Unix()

		if createTimeField, ok := db.Statement.Schema.FieldsByName["CreatedOn"]; ok {
			if createTimeField != nil {
				createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
			}
		}

		if modifyTimeField, ok := db.Statement.Schema.FieldsByName["ModifiedOn"]; ok {
			if modifyTimeField != nil {
				modifyTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if _, ok := db.Statement.Get("gorm:update_column"); !ok {
		db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		var extraOption string
		if str, ok := db.Statement.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField := db.Statement.Schema.LookUpField("DeletedOn")
		isDelField := db.Statement.Schema.LookUpField("IsDel")

		if db.Statement.Unscoped && deletedOnField != nil && isDelField != nil {
			now := time.Now().Unix()
			db.Exec(
				fmt.Sprintf(
					"UPDATE %v SET %v=%v,%v=%v%v%v",
					db.Statement.Table,
					deletedOnField.DBName,
					now,
					isDelField.DBName,
					1,
					addExtraSpaceIfExist(db.Statement.SQL.String()),
					addExtraSpaceIfExist(extraOption),
				),
			)
		} else {
			db.Exec(
				fmt.Sprintf(
					"DELETE FROM %v%v%v",
					db.Statement.Table,
					addExtraSpaceIfExist(db.Statement.SQL.String()),
					addExtraSpaceIfExist(extraOption),
				),
			)
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
