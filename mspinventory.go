package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func webHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s request\n", r.URL.Path)
  t, err := template.ParseFiles("web/templ/base.html",
    "web/templ/header.html",
    "web/templ/footer.html")
  if r.URL.Path == "/" {
    _, err = t.ParseFiles("web/index.html")
  }

  if strings.Contains(r.URL.Path, ".html") {
    _, err = t.ParseFiles(fmt.Sprintf("web%s", r.URL.Path))
  }

  if err != nil {
    log.Printf("Error parsing template: %s\n", err)
    w.WriteHeader(http.StatusNotFound)
    return
  }

  t.Execute(w, nil)
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    panic("PORT env var missing")
  }
 
  http.FileServer(http.Dir("web"))
  server := fmt.Sprintf("127.0.0.1:%s", port)
  http.HandleFunc("/", webHandler)
  log.Printf("Starting webserver on port %s", port)
  if err := http.ListenAndServe(server, nil); err != nil {
    log.Printf("Error starting server: %s", err)
  }
}
