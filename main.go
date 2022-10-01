package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"text/template"
)

type Api struct {
	Id   int
	Name string
}

func main() {
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("error: ", err)
	}

	paths := []string{
		"tmpl.json",
	}
	t, err := template.New("tmpl.json").ParseFiles(paths...)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	var wg sync.WaitGroup
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go process(i, &wg, t)

	}
	wg.Wait()

}

func process(id int, wg *sync.WaitGroup, t *template.Template) {
	defer wg.Done()
	api := Api{id, fmt.Sprintf("Test templates %d", id)}
	path := fmt.Sprintf("files/newfile_%d.json", id)
	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
	}
	defer f.Close()

	err = t.Execute(f, api)
	if err != nil {
		log.Println("error: ", err)
	}
}
