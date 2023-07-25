package test

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

var (
	HOST     = "139.198.165.102"
	USER     = "postgres"
	PASSWORD = "msdnmm,."
	PORT     = "30000"
	DATABASE = "postgres"
	SslMode  = "disable"
	TimeZone = "Asia/Shanghai"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGORM(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASSWORD, DATABASE, PORT, SslMode, TimeZone)
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: false,
		}))

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
