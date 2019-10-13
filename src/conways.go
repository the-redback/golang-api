package conways

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Conways struct {
	ID   int    `json:"id,omitempty" xorm:"'id' pk autoincr"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Data string `json:"data"`
}

func ListenAndServe() {
	fmt.Println("start serving...")

	m := pat.New()
	m.Post("/grids", http.HandlerFunc(PostGrids))
	m.Patch("/grids/:id", http.HandlerFunc(PatchGrids))
	m.Get("/grids/:id", http.HandlerFunc(GetGrids))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetGrids(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get(":id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("error during reading from ioutil: ", err)
		return
	}
	log.Println("received patch response for", id)

	var data Conways
	data.ID=id

	en := connect_database()
	defer en.Close()
	has, err := query_data(en, &data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("error during reading from database: ", err)
		return
	} else if !has {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("requested data not found"))
		log.Println("requested data not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	marshalledData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error during marshal data"))
		log.Println("error during marshal data: ", err)
		return
	}
	w.Write(marshalledData)
}

func PostGrids(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error during reading from ioutil"))
		log.Println("error during reading from ioutil: ", err)
		return
	}
	log.Println("received response", string(body))

	var data Conways
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("probably invalid input."))
		log.Println("error during unmarshal data: ", err)
		return
	}

	en := connect_database()
	defer en.Close()
	insert_data(en, &data)

	w.WriteHeader(http.StatusCreated)
	marshalledData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error during marshal data"))
		log.Println("error during marshal data: ", err)
		return
	}
	w.Write(marshalledData)
}

func PatchGrids(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get(":id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error during reading from ioutil"))
		log.Println("error during reading from ioutil: ", err)
		return
	}
	log.Println("received patch response for", id)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error during reading from ioutil"))
		log.Println("error during reading from ioutil: ", err)
		return
	}

	var data Conways
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("probably invalid input."))
		log.Println("error during unmarshal data: ", err)
		return
	}
	data.ID=id

	en := connect_database()
	defer en.Close()
	update_data(en, &data)

	w.WriteHeader(http.StatusCreated)
	marshalledData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error during marshal data"))
		log.Println("error during marshal data: ", err)
		return
	}
	w.Write(marshalledData)
}
