package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMysql(t *testing.T) {
	const HOST = "host"
	const USER = "root"
	const PORT = "3306"
	const PASSWORD = "pass"
	const DATABASE = "db"
	const CHARSET = "utf8mb4"
	const PARSE_TIME = "True"
	const LOC = "Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		USER, PASSWORD, HOST, PORT, DATABASE, CHARSET, PARSE_TIME, LOC,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		t.Errorf("启动失败:%s", err)
	}

	// 建表, 默认表名为结构体小写+s
	err = db.AutoMigrate(&Product{})
	if err != nil {
		t.Errorf("迁移失败:%s", err)
	}

	// 创建数据
	db.Create(&Product{
		Code:  "1231",
		Price: 10,
	})

	// 批量插入
	db.Create([]Product{{
		Code:  "1",
		Price: 110,
	}, {
		Code:  "2",
		Price: 120,
	},
	})

	// 单条更新
	var p1 Product
	db.First(&p1, 1).
		Model(&p1).
		Update("Code", 22)

	// 批量更新
	db.Model(&p1).
		Where("Code > ?", 1).
		Update("Code", 22)

	// 单条查询
	db.First(&p1, 1)
	t.Logf("单条查询:%v", p1)

	var prods []Product

	// 条件查询
	db.Find(&prods, "Code > ?", 2)
	t.Logf("prods:%v", prods)

	// 全部查询
	db.Find(&prods)
	t.Logf("prods:%v", prods)

	// 单个条件删除
	db.First(&p1, 1).Delete(&p1)

	// 批量删除
	db.Where("Code =", 2).Delete(&Product{})

	//defer func() {
	//	// 删表
	//	err = db.Migrator().DropTable(&Product{})
	//	if err != nil {
	//		t.Errorf("删表失败%s", err)
	//	}
	//}()
}
