package go_path_to_port_handler

import (
	"log"
	"net/http"
	"testing"
)

// In real life it should be another app
func anotherAppIndextHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("anotherAppIndextHandler")
	w.Write([]byte("anotherAppIndextHandler"))
}

// In real life it should be another app
func anotherAppInit() {
	log.Println("Emulate app on port 8010")
	anotherApp := http.NewServeMux()
	anotherApp.HandleFunc("/", anotherAppIndextHandler)
	go http.ListenAndServe(":8010", anotherApp)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler")
	w.Write([]byte("IndexHandler"))
}

func main() {
	anotherAppInit()

	// open http://localhost:8000/test/
	log.Println("Listen port 8000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc(NewPathToPortHandler("/test/", 8010))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Test(t *testing.T) {
	main()
}
