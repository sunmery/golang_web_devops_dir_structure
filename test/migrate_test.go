package test

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"path/filepath"
	"testing"
)

// 测试迁移 sql 数据
func TestMigrate(t *testing.T) {
	// 创建数据库连接
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", USER, PASSWORD, HOST, PORT, DATABASE, SslMode))
	if err != nil {
		log.Fatalf("Unable to open database connection: %v", err)
	}

	// 执行完成关闭数据库连接
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("关闭数据库连接失败 %v", err)
		}
	}(db)

	// 创建数据库
	_, err0 := db.Exec("CREATE DATABASE tests")
	if err0 != nil {
		t.Errorf("创建数据库失败:%s", err0)
	}

	// 获取 migrations 文件的绝对路径
	mPath, err := filepath.Abs("./migrations")
	fmt.Print(mPath)
	if err != nil {
		log.Fatalf("Unable to get absolute path: %v", err)
	}

	// 创建 postgres 驱动
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Unable to create postgres driver: %v", err)
	}

	// 创建迁移实例
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", mPath),
		"postgres", driver)
	if err != nil {
		log.Fatalf("Unable to create migration instance: %v", err)
	}

	// 执行迁移
	if err := m.Up(); err != nil {
		// 数据库迁移时，如果没有任何变化需要应用（也就是说，你的数据库已经处于最新状态）
		// m.Up() 会返回一个特殊的错误 migrate.ErrNoChange
		// 这里判断非 migrate.ErrNoChange 错误时才会终止程序
		if !errors.Is(migrate.ErrNoChange, err) {
			log.Fatalf("An error occurred while migrating: %v", err)
		}
	}

	log.Println("Migration completed successfully!")
}
