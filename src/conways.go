package conways

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

type Conways struct {
	ID   int    `json:"id,omitempty" xorm:"'id' pk autoincr"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Data string `json:"data"`
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func PostGrids(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("Error during reading from ioutil: ", err)
	}
	log.Println("received response", string(body))

	var data Conways
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("probably invalid input."))
		log.Println("Error during unmarshal data: ", err)
	}

	en := connect_database()
	defer en.Close()
	insert_data(en, &data)

	w.WriteHeader(http.StatusCreated)
	marshalledData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("Error during marshal data: ", err)
	}
	w.Write(marshalledData)
}

func ListenAndServe() {
	fmt.Println("start serving...")

	m := pat.New()
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))
	m.Post("/grids", http.HandlerFunc(PostGrids))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
