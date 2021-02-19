package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	TerraformRawTemplater()
}

func terraformTemplating(){
	saSource := ServiceAccountStruct{
		Project: "dev",
		Region: "us-central-1",
		SaNAME: "go-template-sa",
		Description: "Templating test",
	}
	resultB := ProcessFile("/Users/singaravelan.nandhakumar/go/src/github.com/velann21/go-data-structures/go_templating/templates/main.tf", saSource)
	fmt.Println(resultB)
	CreateFolder()
	CreateFile(resultB)
}

func TerraformRawTemplater(){
	saSource := ServiceAccountStruct{
		Project: "dev",
		Region: "us-central-1",
		SaNAME: "go-template-sa",
		Description: "Templating test",
	}
	result := ProcessRawTemplate(saSource)
	fmt.Println(result)
	CreateFolder()
	CreateFile(result)
}

func ProcessRawTemplate(sa ServiceAccountStruct)string{
	var tmplBytes bytes.Buffer
	t := template.New("Person template")
	t, err := t.Parse(ServiceAccount)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}
	err = t.Execute(&tmplBytes, sa)
	if err != nil {
		log.Fatal("Execute: ", err)
		return ""
	}
	return tmplBytes.String()

}

func CreateFolder(){
	_, err := os.Stat("/Users/singaravelan.nandhakumar/go/src/github.com/velann21/go-data-structures/go_templating/templates/service_account")
	isNotExist := os.IsNotExist(err)
	if isNotExist{
		err := os.Mkdir("/Users/singaravelan.nandhakumar/go/src/github.com/velann21/go-data-structures/go_templating/templates/service_account", 0777)
		if err != nil{
			fmt.Println("error CreateFolder", err)
		}
		fmt.Println("File created")
	}
}

func CreateFile(result string){
	file, err := os.Create("/Users/singaravelan.nandhakumar/go/src/github.com/velann21/go-data-structures/go_templating/templates/service_account/main.tf")
	if err != nil{
		fmt.Println("error CreateFile" ,err)
	}
	status, err := file.Write([]byte(result))
	if err != nil{
		fmt.Println("error CreateFile Write ", err)
	}
	fmt.Println(status)
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

type ServiceAccountStruct struct {
	Project string
	Region string
	SaNAME string
	Description string
}

const ServiceAccount = `
terraform {
  # The modules used in this example have been updated with 0.12 syntax, which means the example is no longer
  # compatible with any versions below 0.12.
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.52.0"
    }
  }
  required_version = ">= 0.12.6"
  backend "gcs" {
    bucket = "gotfbucket"
    prefix = "devops-guidelines/goterraform-test"
  }
}

provider "google" {
  project = "{{.Project}}"
  region = "{{.Region}}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE A CUSTOM SERVICE ACCOUNT FOR THE CD PIPELINES
# ---------------------------------------------------------------------------------------------------------------------

module "cd_serviceaccount" {
  source = "git@git.build.ingka.ikea.com:/terraform/gcp-service-account.git?ref=v0.0.1"
  name                  = "{{.SaNAME}}"
  project               = "{{.Project}}"
  description           = "{{.Description}}"
}

`