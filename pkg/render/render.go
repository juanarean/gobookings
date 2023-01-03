package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/juanarean/gobookings/pkg/config"
	"github.com/juanarean/gobookings/pkg/models"
)

// var tc = make(map[string]*template.Template)

// var functions = template.FuncMap {}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

func RenderTemplate(res http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}


	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cannot get template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(res)
	if err != nil {
		log.Println(err)
	}

	// parseTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.html")
	// err := parseTemplate.Execute(res, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template", err)
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(("./templates/*.page.html"))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(("./templates/*.layout.html"))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}


	return myCache, nil
}

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]

// 	if !inMap {
// 		fmt.Println("reading and creating")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println("error parsing")
// 		}
// 	} else {
// 		fmt.Println("using cached tmpl")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		fmt.Println("error executing")
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string {
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl
// 	return nil
// }