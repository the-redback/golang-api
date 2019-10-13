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

func insert_data(en *xorm.Engine, data *Conways){
	affected, _ :=en.Insert(data)
	log.Println("Inserted user id:",data.ID,":: affected:",affected,"data")
}

func update_data(en *xorm.Engine, data *Conways){
	affected,_ :=en.Id(data.ID).Update(&data)

	log.Println("affected row: ", affected, data)

}

func query_single_data(en *xorm.Engine){
	user:=Conways{ID: 1}
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
