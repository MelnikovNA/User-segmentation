package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

//go:embed errors.go.tmpl
var tmplGoErrors string

//go:embed methods.go.tmpl
var tmplMethods string

const (
	errJson = "errors/errors.json"
	outPath = "codegen/go/apierrors"
	suffix  = "_gen.go"
	methods = "methods"
)

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type MapperTemplate struct {
	Package string
	Names   map[string]apiError
}

func main() {
	errJsonFile := flag.String("j", errJson, "path to errors.json")
	outPath := flag.String("o", outPath, "path to errors.json")
	packageName := flag.String("pkg", "apierrors", "package name")

	jsonBin, err := os.ReadFile(*errJsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mapper := &MapperTemplate{
		Package: *packageName,
	}

	err = json.Unmarshal(jsonBin, &mapper.Names)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := os.MkdirAll(*outPath, 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			fmt.Printf("%v %v", *outPath, err)
			os.Exit(1)
		}
	}

	outName := path.Join(*outPath, *packageName+suffix)
	t := template.Must(template.New("mapper").Parse(tmplGoErrors))
	fout, err := os.OpenFile(outName,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("%v %v", outName, err)
		os.Exit(1)
	}
	err = t.Execute(fout, mapper)
	if err != nil {
		log.Printf("%v %v", outName, err)
		os.Exit(1)
	}
	_ = fout.Close()

	outName = path.Join(*outPath, methods+suffix)
	s := strings.Split(tmplMethods, "\n")
	s[0] = fmt.Sprintf("package %v", *packageName)
	out := strings.Join(s, "\n")

	err = os.WriteFile(outName, []byte(out), 0644)
	if err != nil {
		log.Printf("%v %v", outName, err)
		os.Exit(1)
	}

}
