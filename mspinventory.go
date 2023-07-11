package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)


func webHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s request\n", r.URL.Path)
  t, err := template.ParseFiles("web/templ/base.html", "web/templ/header.html", "web/templ/footer.html")

  var d interface{}
  switch r.URL.Path {
    case "/", "/index.html":
      _, err = t.ParseFiles("web/index.html")
      d = alerts
    case "/inventory.html":
      d = items
    case "/alert.html":
      query := r.URL.Query().Get("id")
      id, err := strconv.Atoi(query)
      if err != nil {
        log.Printf("Could not find id %s", query)
      }

      d = alerts[id]
    case "/providers.html":
      d = providers
    case "/addprovider":
      if r.Method == http.MethodPost {
        t, _ := strconv.Atoi(r.FormValue("type"))
        p := provider {
          id: len(providers) + 1,
          ptype: t,
          key: r.FormValue("apikey"),
        }

        providers[p.id] = p
      }

      http.Redirect(w, r, "/providers.html", http.StatusSeeOther)
  }

  if strings.Contains(r.URL.Path, ".html") {
    _, err = t.ParseFiles(fmt.Sprintf("web%s", r.URL.Path))
  }

  if err != nil {
    log.Printf("Error parsing template: %s\n", err)
    w.WriteHeader(http.StatusNotFound)
    return
  }

  t.Execute(w, d)
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    panic("PORT env var missing")
  }

  alerts = make(map[int]alert)
  providers = make(map[int]provider)
  //test data
  testAlerts()

  http.FileServer(http.Dir("web"))
  server := fmt.Sprintf("127.0.0.1:%s", port)
  http.HandleFunc("/", webHandler)
  log.Printf("Starting webserver on port %s", port)
  if err := http.ListenAndServe(server, nil); err != nil {
    log.Printf("Error starting server: %s", err)
  }
}
