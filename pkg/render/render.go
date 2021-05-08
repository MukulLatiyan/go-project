package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateCache, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("error getting template cache: ", err)
	}
	fmt.Println("actual value:", templateCache[tmpl])
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error occurred while parsing template:", err)
		return
	}
}

// RenderTemplateTest renders all template
func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently :", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		fmt.Println("ts ==", ts)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		fmt.Println(matches)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = template.ParseGlob("./templates/*.layout.html")
			fmt.Println("Inside len(matches)", ts)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
