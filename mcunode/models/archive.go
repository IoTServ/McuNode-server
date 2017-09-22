package models

import (
	_"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_"time"
)

type Event struct {
	User      string
	Content   []byte
}


// Model Struct
//type DEVICE struct {
//	Id   int
//	MCUID string `orm:"index"`
//	DATA string `orm:"size(10240)"`
//	DATE time.Time
//}
//
//func init() {
//	// set default database
//	orm.RegisterDataBase("default", "mysql", "root:root@/mcunode?charset=utf8", 30)
//
//	// register model
//	orm.RegisterModel(new(DEVICE))
//
//	// create table
//	orm.RunSyncdb("default", false, true)
//
//	//o := orm.NewOrm()
//	//
//	//user := DEVICE{MCUID: "slene",DATA:"wode",DATE:time.Now()}
//	//
//	//// insert
//	//id, _ := o.Insert(&user)
//	//fmt.Printf(string(id))
//}