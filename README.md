# go_path_to_port_handler
For easy launching several go (golang) apps at one server

USAGE
```
import (
  gptp "github.com/nordorn/go_path_to_port_handler"
  "net/http"
  "log"
)

func main() {
  //apps should be launched on the ports
  http.HandleFunc(gptp.NewPathToPortHandler("/app1/", 8001))
  http.HandleFunc(gptp.NewPathToPortHandler("/app2/", 8002))
  http.HandleFunc(gptp.NewPathToPortHandler("/app3/", 8003))

  log.Println("Listen port 80")
  http.HandleFunc("/", indexHandler)
  log.Fatal(http.ListenAndServe(":80", nil))
}
```
