package main

import (
        "web"
        "io/ioutil"
        "flag"
	"dml"
)

func globalServe(val string) string {
        file, err := ioutil.ReadFile("/" + val)
        if err != nil {
		return "404: Global file not found: dml-g/" + val
        }
        return dml.ParseDoc(string(file))
}

func contextServe(val string) string {
	if len(val) == 0 || (len(val) == 1 && val == "/") {
                val = "index.dml"
        }
        file, err := ioutil.ReadFile(val)
        if err != nil {
		return "404: File not found: dml/" + val
        }
        return dml.ParseDoc(string(file))
}

func main() {
        global := flag.Bool("global", false, "Allow the server to access any files that the user running it has access to.")
	flag.Parse()
        web.Get("/dml/(.*)", contextServe)
        web.Get("/dml", contextServe)
        if *global {
		web.Get("/dml-g/(.*)", globalServe)
		web.Get("/dml-g", globalServe)
	}
        web.Run("0.0.0.0:8080")
}

