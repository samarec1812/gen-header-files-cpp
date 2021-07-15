package main

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	"fmt"
	"go/format"
	_ "io"
	"log"
	"os"
	"text/template"
)

type Data struct {
	CntrName string `json:"cntrName"`
	Public PublicData `json:"public"`
	Private PrivateData `json:"private"`
	// Fields []FieldParameters `json:"fields"`
	HeaderIncludes []string `json:"header_includes"`
	// Functions []Function `json:"functions"`
}
type PublicData struct {
	Fields []FieldParameters `json:"fields"`
	Functions []Function `json:"functions"`
}

type PrivateData struct {
	Fields []FieldParameters `json:"fields"`
	Functions []Function `json:"functions"`
}

type FieldParameters struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Default string `json:"default"`
}

type Function struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Parameters []FieldParameters `json:"parameters"`
	IdentAfterFunc []string `json:"ident_after_func"`
}



func readingJSON(name string, value interface{}) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(value)

}

func GenerateTemplate(fileName string, data Data) error {
	//info, err := ioutil.ReadFile("templates/Header2.tmpl")
	//if err != nil {
	//	return err
	//}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	//defer file.Close()
	//
	//_, err = io.Copy(file, bytes.NewReader(info))
	//if err != nil {
	//	return err
	//}
	tmpl, err := template.New("new_template").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}).ParseFiles("templates/Header2.tmpl")
	//tmpl, err := template.ParseFiles("templates/Header2.tmpl")
	if err != nil {
		return err
	}

	var tmplBytes bytes.Buffer
	err = tmpl.Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}).ExecuteTemplate(file, "Header2.tmpl", data)
	if err != nil {
		return err
	}
	Formatted, err := format.Source(tmplBytes.Bytes())
	if err != nil {
		return err
	}
	fmt.Println(string(Formatted))

	return err
}

func LastIndex(args ...interface{}) string {
	if len(args) == 0 { return ""}
	return fmt.Sprintf("%v", len(args) - 1)
}

func main() {

	data := Data{}
	err := readingJSON("parameters.json", &data)
	if err != nil {
		log.Println(err)
		return
	}
	err = GenerateTemplate("header.h", data)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("GOOD")

}
