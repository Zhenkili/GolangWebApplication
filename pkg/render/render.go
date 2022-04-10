package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

//把tmpl写进w里作为应答
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	templatecache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	template, ok := templatecache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	//将template解析后写入到bug里
	_ = template.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error in write out the template:", err)
	}
}

//create a template map
func CreateTemplateCache() (map[string]*template.Template, error) {

	//要哪个名字就直接给出template的map
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		fmt.Println("error glob templates files:", err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Print("current is the page for ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//抓取base网页
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Println("error glob templates files:", err)
			return myCache, err
		}

		if len(matches) > 0 {
			//用matche到的layout加持
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
