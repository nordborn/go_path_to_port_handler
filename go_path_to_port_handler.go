// USAGE
// import (
//	gptp "github.com/nordorn/go_path_to_port_handler"
//	"log"
// )
//
// func main() {
//	// apps should be launched on the ports
//	http.HandleFunc(gptp.NewPathToPortHandler("/app1/", 8001))
//	http.HandleFunc(gptp.NewPathToPortHandler("/app2/", 8002))
//	http.HandleFunc(gptp.NewPathToPortHandler("/app3/", 8003))
//
//	log.Println("Listen port 80")
//	http.HandleFunc("/", indexHandler)
//	log.Fatal(http.ListenAndServe(":80", nil))
// }
package go_path_to_port_handler

import (
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

// For easy launching several apps at one server
// Path to handle in http.HandleFunc should be "/path/" (with trailing slash)
// Matching requests will be redirected from http://localhost/path/123/ to http://localhost:appPort/123/
func NewPathToPortHandler(path string, appPort int) (string, func(http.ResponseWriter, *http.Request)) {

	normalizedPath := strings.TrimSuffix(path, "/") + "/"
	return normalizedPath, func(w http.ResponseWriter, r *http.Request) {
		pathText := strings.Trim(path, "/")
		newUri := strings.Replace(strings.Replace(r.RequestURI, pathText, "", 1), "//", "/", -1)

		url := r.URL
		url.Path = newUri
		url.Host = strings.Split(r.Host, ":")[0] + ":" + strconv.Itoa(appPort)
		url.Scheme = "http"

		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	}
}
