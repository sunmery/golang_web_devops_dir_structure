package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var MysqlDB *gorm.DB // pg 客户端
var MysqlErr error

// InitMysqlConnect 初始化数据库服务
func InitMysqlConnect() {
	HOST := os.Getenv("MYSQL_HOST")            // 主机
	USER := os.Getenv("MYSQL_USER")            // 用户, 默认为 postgres
	PORT := os.Getenv("MYSQL_PORT")            // 密码
	PASSWORD := os.Getenv("MYSQL_PASSWORD")    // 端口
	DATABASE := os.Getenv("MYSQL_DATABASE")    // 数据库
	CHARSET := os.Getenv("MYSQL_CHARSET")      // 编码
	ParseTime := os.Getenv("MYSQL_PARSE_TIME") // 是否解析时间
	LOC := os.Getenv("MYSQL_LOC")              // 时区

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		USER, PASSWORD, HOST, PORT, DATABASE, CHARSET, ParseTime, LOC,
	)
	MysqlDB, MysqlErr = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if MysqlErr != nil {
		panic("启动 Mysql 服务失败!, 错误信息:" + MysqlErr.Error())
	}
}
