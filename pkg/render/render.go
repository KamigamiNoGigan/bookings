package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/KamigamiNoGigan/booking/pkg/config"

	"github.com/KamigamiNoGigan/booking/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplates(w http.ResponseWriter, name string, td *models.DataStruct) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateCache()
	}
	t, ok := tc[name]
	if !ok {
		log.Fatal("error template name")
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, td)

	buf.WriteTo(w)
}

func CreateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = t
	}
	return myCache, nil
}
