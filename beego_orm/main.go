package main

import (
	_ "Go_learn/beego_orm/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost/orm_test?disable")
	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	//profile := new(models.Profile)
	//profile.Age = 30
	//
	//user := new(models.User)
	//user.Profile = profile
	//user.Name = "slene"

	//fmt.Println(o.Insert(profile))
	//fmt.Println(o.Insert(user))

	//o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")
	fmt.Print(r)
}
