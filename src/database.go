package conways

import (
	"fmt"

	"github.com/go-xorm/xorm"

	"log"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "pass"
	DB_NAME     = "postgres"
)

func init() {
	en := connect_database()
	defer en.Close()
	sync_tables(en)
}

func connect_database() *xorm.Engine {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	en, err := xorm.NewEngine("postgres", dbinfo)
	if err != nil {
		log.Println("engine creation failed", err)
	}

	// whenever this engine is used,
	// defer post.db.Close()

	err = en.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected")
	return en
}

func sync_tables(en *xorm.Engine) {
	err := en.Sync(new(Conways))
	if err != nil {
		log.Println("creation error", err)
		return
	}
	log.Println("Successfully synced")
}

func query_data(en *xorm.Engine, data *Conways) (bool, error) {
	return en.Get(data)
}

func insert_data(en *xorm.Engine, data *Conways) {
	affected, _ := en.Insert(data)
	log.Println("Inserted user id:", data.ID, ":: affected:", affected, "data")
}

func update_data(en *xorm.Engine, data *Conways) {
	affected, _ := en.Id(data.ID).AllCols().Update(data)

	log.Println("affected row: ", affected, data)

}
