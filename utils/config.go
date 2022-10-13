package utils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

/*
	GetYamlConfig() function aims to get the information from

the yaml file
*/
func GetYamlConfig(path string) map[interface{}]interface{} {
	data, err := ioutil.ReadFile(path)
	m := make(map[interface{}]interface{})
	if err != nil {
		log.Print(err.Error())
	}
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Print(err.Error())
	}
	return m
}

func GetXMLConfig(path string) map[string]string {
	var t xml.Token
	var err error

	Keylst := make([]string, 6)
	Valuelst := make([]string, 6)

	map1 := make(map[string]string)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err.Error())
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	i := 0
	j := 0
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {

		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			Keylst[i] = string(name)
			i = i + 1
		case xml.CharData:
			content1 := string([]byte(token))
			Valuelst[j] = content1
			j = j + 1
		}
	}
	for count := 0; count < len(Keylst); count++ {
		map1[Keylst[count]] = Valuelst[count]
	}

	return map1
}

func GetElement(key string, themap map[interface{}]interface{}) string {
	if value, ok := themap[key]; ok {

		return fmt.Sprint(value)
	}

	log.Print("can't find the config file")
	return ""
}
