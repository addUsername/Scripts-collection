package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("Provide more than one argument.")
	}

	var threads = os.Args[1]
	var orchestrator = os.Args[2]
	var fileNames = os.Args[3:]
	wkdir, _ := os.Getwd()
	switch orchestrator {

	case "1":
		//shaHash := sha512.New()
		var out []string
		for _, file := range fileNames {
			f, _ := ioutil.ReadFile(file)
			out = append(out, file+" "+fmt.Sprintf("%x", sha512.Sum512(f)))
		}
		a := strings.Join(out, "\n")
		//ioutil.WriteFile("D:\\go-single.txt", []byte(a), 0644)
		f, _ := os.Create(wkdir + "\\go-single.txt")
		defer f.Close()
		f.WriteString(a)

	case "2":
		if threads != "-1" {
			log.Fatal("nope")
		}
		var wg sync.WaitGroup
		wg.Add(len(fileNames))
		var out []string

		for _, file := range fileNames {
			go func(file string) {
				defer wg.Done()
				f, _ := ioutil.ReadFile(file)
				out = append(out, file+" "+fmt.Sprintf("%x", sha512.Sum512(f)))
			}(file)
		}
		wg.Wait()
		a := strings.Join(out, "\n")
		//ioutil.WriteFile("D:\\go-routines.txt", []byte(a), 0644)
		f, _ := os.Create(wkdir + "\\go-routines.txt")
		defer f.Close()
		f.WriteString(a)

	}
}
