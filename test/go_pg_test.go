package test

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"testing"
)

type User struct {
	Id   int64 `pg:"id,pk"`
	Name string
}

func TestPostgres(t *testing.T) {
	db := pg.Connect(&pg.Options{
		Addr:     "139.198.165.102:30981",
		User:     "root",
		Password: "msdnmm",
		Database: "root",
	})

	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			t.Log("关闭数据库连接失败")
		}
	}(db)

	_, err0 := db.Exec("CREATE DATABASE tests")
	if err0 != nil {
		t.Errorf("创建数据库失败:%s", err0)
	}

	user1 := &User{
		Id:   123,
		Name: "test name",
	}

	// 创建表
	err1 := db.Model((*User)(nil)).CreateTable(&orm.CreateTableOptions{})
	if err1 != nil {
		t.Errorf("创建表失败:%s", err1)
	}

	// 插入
	result, err := db.Model(user1).Insert()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", result)

	var users []User
	err2 := db.Model(&users).Select()
	if err2 != nil {
		t.Error(err2)
	}

	user1.Name = "newname"
	_, err3 := db.Model(user1).WherePK().Update()
	if err3 != nil {
		t.Error(err3)
	}

	usersUp := []*User{
		{
			Id:   123,
			Name: "asd",
		},
		{
			Id:   12223,
			Name: "asd2",
		},
	}

	_, err4 := db.Model(&usersUp).Update()
	if err3 != nil {
		t.Error(err4)
	}

	_, err = db.Exec("DROP DATABASE tests")

}
