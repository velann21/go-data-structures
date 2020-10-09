package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"text/template"
)

type Joke struct {
	Who string
	Punchline string
}

func main() {
	buckets := Buckets{
		Name:"velann211212",
	UnMarshalYaml:func(b Buckets)(map[string]interface{}){
		yamlFile, err := ioutil.ReadFile("/Users/singaravelannandakumar/go/src/awesomeProject3/go_templating/templates/values.yaml")
		if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
		err = yaml.Unmarshal(yamlFile, &b.Values)
		if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
		return  b.Values
	},
	PrintArgsfunc :func(string2 interface{})(map[string]interface{}){
		fmt.Println(string2)
		return nil
	},
	GetBucket: func(b Buckets) Buckets {
		return b
	},
	}


	// process a template file
	resultB := ProcessFile("/Users/singaravelannandakumar/go/src/awesomeProject3/go_templating/templates/dep.yaml", buckets)
	fmt.Println(resultB)
}

func ProcessFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)

	if err != nil {
		panic(err)
	}
	return process(tmpl, vars)
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}


type Buckets struct {
	Name string
	Values map[string]interface{}
	UnMarshalYaml func(b Buckets) (map[string]interface{})
	PrintArgsfunc func(string2 interface{})(map[string]interface{})
	GetBucket func(b Buckets)Buckets
}
