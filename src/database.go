package conways

import (
	"fmt"
	"github.com/go-xorm/xorm"

	_ "github.com/lib/pq"
	"log"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "pass"
	DB_NAME     = "postgres"
)

type Conways struct{
	Uid int64  `xorm:"pk not null autoincr"`
	XAxis int64
	YAxis int64
	Grid string
}


func connect_database() *xorm.Engine{
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST,DB_PORT,DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	en, err := xorm.NewEngine("postgres", dbinfo)
	if err!=nil{
		log.Println("engine creation failed", err)
	}

	//defer post.db.Close()

	err =en.Ping()
	if err !=nil{
		panic(err)
	}

	log.Println("Successfully connected")
	return en
}

func sync_tables(en *xorm.Engine){
	err:=en.Sync(new(Conways))
	if err!=nil{
		log.Println("creation error",err)
		return
	}
	log.Println("Successfully synced")
}

func insert_data(en *xorm.Engine){
	u:=new(Conways)
	u.Uid=3
	u.XAxis=3
	u.YAxis=3
	u.Grid="***...***"
	affected, _ :=en.Insert(u)

	log.Println("Inserted user id:",u.Uid,":: affected:",affected)
}


func query_single_data(en *xorm.Engine){
	user:=Conways{Uid: 1}
	has,_:=en.Get(&user)

	log.Println("-----------",has)
	log.Println("-----------",user)

	results, err := en.Query("select * from conways_uid_seq")
	fmt.Print("-------------2 ", results, err)

	//Another way
	var user2 Conways
	has2, _ :=en.Get(&user2)    //Primary key
	log.Println("-----------",has2)
	log.Println("-----------", user2)
}
