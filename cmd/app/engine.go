package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/template/html/v2"
)

func setupEngine() *html.Engine {
  engine := html.New("./views", ".html")
  engine.AddFunc("css",
    func() (res template.HTML) {
      filepath.Walk("public/assets", func(path string, info os.FileInfo, err error) error {
        if err != nil {
          return err
        }

        if strings.Contains(info.Name(), ".css") {
          res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
        }

        return nil
      })

      return
    },
  )

  return engine
}
