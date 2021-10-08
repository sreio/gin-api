package models

import (
	"fmt"
	"gin-api/pkg/setting"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct{
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `gorm:"column:created_on" json:"created_on"`
	ModifiedOn int `gorm:"column:modified_on" json:"modified_on"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Minute)

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok { // 判断是否存在这个字段
			if createTimeField.IsBlank { // 判断字段值是否为空
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("modifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok { // 查找等于gorm:update_column的字段属性
		scope.SetColumn("ModifiedOn", time.Now().Unix()) // 如果没有指定gorm:update_column 则设置ModifiedOn
	}
}