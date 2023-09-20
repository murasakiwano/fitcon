package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	people := []FitConner{
		{
			Name: "Monkey D. Luffy",
			Meta1: Metas{
				FatPercentage: "*",
				LeanMass:      "*",
			},
			Meta2: Metas{
				VisceralFat:   "*",
				FatPercentage: "*",
				LeanMass:      "*",
			},
		},
		{
			Name: "Roronoa Zoro",
			Meta1: Metas{
				FatPercentage: "*",
				LeanMass:      "*",
			},
			Meta2: Metas{
				VisceralFat:   "*",
				FatPercentage: "*",
				LeanMass:      "*",
			},
		},
	}

	// tmplFile := "./templates/layout.html"
	tmpl, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var f *os.File
	for _, person := range people {
		f, err = os.Create(fmt.Sprintf("%s.html", person.Name))
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		err = tmpl.Execute(f, person)
	}

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// fs := http.FileServer(http.Dir("./frontend"))
	// http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))
	//
	// http.HandleFunc("/", serveTemplate)
	//
	// log.Print("Listening on :5656...")
	// err = http.ListenAndServe(":5656", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
