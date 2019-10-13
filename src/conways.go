package conways

import (
	"fmt"
	"github.com/bmizerany/pat"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func PostGrids(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
	
}

func ListenAndServe() {
	en:=connect_database()

	sync_tables(en)
	insert_data(en)

	query_single_data(en)

	fmt.Println("start serving...")

	m := pat.New()
	m.Get("/grids", http.HandlerFunc(HelloServer))
	m.Post("/hello/:name", http.HandlerFunc(PostGrids))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
