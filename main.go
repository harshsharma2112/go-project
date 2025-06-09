// this is case 1 and basiscs of creation fo rest api
// here we have used hello world and http server package

/*package main

import (
	"fmt"
	"log"
	"net/http"
)

// Root URL is the actual adress
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s with token:%s\n", r.URL.Path, r.URL.Query().Get("token"))
	})
	// get query is to show dynamic request made to the server

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
 // to serve the rqst whenever the url conatin static
	log.Println("web server started")
	http.ListenAndServe(":80", nil)
} */

// case 2 using mutex more advance
/*package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter() // creating a new router and registring it with the request handler
//{title and page are placeholder}
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r) // above i using nil to use the default server of hhtp serveer package but here creating my own
} */

//case 3 for my service now

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/harshsharma2112/go-project.git/details" // here details is the alias of details package
	//reason for importing it
	// 	“Hey, this whole folder is called github.com/harshsharma2112/go-project.git, and I will always refer to folders inside it using this name.”
	// So Go doesn't care where the file physically is. It cares about the name you told it in go.mod.
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking applictaion health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response) // taken this line from json
	//Imagine w is like the outbox of your server. When a client makes a request, you use w to place a response (HTML, JSON, text, etc.) back into the outbox.
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving homepage")
	w.WriteHeader(http.StatusOK)                   //This line sets the HTTP status code for the response you're sending back from your server to the client.
	fmt.Fprintf(w, "Aplication is running and up") //fmt.Fprintf is a Go function that formats a string and writes it to a specific output, like a file, buffer, or an HTTP response.
}

func detailhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname() // using alias to use the func from the folder
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	response := map[string]string{
		"hostname": hostname,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(response) // niche phle print kra rhe the normally but now giving json resposne
	//fmt.Println(hostname, IP)
	//fmt.Fprintf(w, "Hostname: %s", hostname)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/", roothandler)
	r.HandleFunc("/detail", detailhandler)
	log.Println("server has started")
	log.Fatal(http.ListenAndServe(":80", r))

}
