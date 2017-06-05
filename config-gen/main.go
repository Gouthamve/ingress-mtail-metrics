package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"text/template"

	yaml "gopkg.in/yaml.v1"
)

type config struct {
	Backends []struct {
		Host  string `yaml:"host"`
		Paths []struct {
			Name   string `yaml:"name"`
			Regexp string `yaml:"regexp"`
		} `yaml:"paths"`
	} `yaml:"backends"`
}

func main() {
	cfgStr, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	c := config{}
	err = yaml.Unmarshal(cfgStr, &c)
	if err != nil {
		log.Fatalln(err)
	}

	mtailProg := bytes.NewBuffer(nil)
	t := template.Must(template.New("header").Parse(header))
	err = t.Execute(mtailProg, nil)
	if err != nil {
		log.Fatalln(err)
	}

	funcMap := template.FuncMap{
		"escRegexp": regexp.QuoteMeta,
	}

	t = template.Must(template.New("body").Funcs(funcMap).Parse(upstream))

	for _, s := range c.Backends {
		for _, path := range s.Paths {
			err := t.Execute(mtailProg, struct {
				Host       string
				PathName   string
				PathRegexp string
			}{s.Host, path.Name, path.Regexp})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.Write(mtailProg.Bytes())
	if err != nil {
		log.Fatalln(err)
	}
}
