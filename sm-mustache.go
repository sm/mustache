/*

  NAME

  sm-mustache - Mustache template command line

  SYNOPSIS

    sm-mustache [options]

  OPTIONS

    -data="{key1=value1}:~{key2=value2}": key=value data pair, overrides json file data
    -json="{string}": json data string
    -output="{file}": path to output file
    -template="{file}": path to template file

    -data|--data "{key1=value1:~key2=value2}

      Assign a value to a given key for the template rendering

    -json|--json "{string}"

      To be implemented

    -output|--output "{file}"

      specify the output file to write the rendered template to

    -template|--template "{file}"

      specify the template file to use

  DESCRIPTION

  sm-mustache will read the given template file, and using the given data render
  it to the specified output file location.

*/

package main

import (
	"flag"
	"github.com/hoisie/mustache"
	"log"
	"os"
	"strings"
)

var json = flag.String("json", "{string}", "json data string")
var template = flag.String("template", "{file}", "path to template file")
var output = flag.String("output", "{file}", "path to output file")
var data = flag.String("data", "{key1=value1}:~{key2=value2}", "key=value data pair, overrides json file data")

var mapped_data map[string]string

func main() {

	mapped_data = make(map[string]string)

	flag.Parse()

	//Error checking on arguements passed in
	if *template == "{file}" {
		log.Fatal("ERROR: A template file location must be specified with --template={{path to template file}}")
	} else {
		file, err := os.Open(*template)
		if file == nil && err != nil {
			log.Fatalf("ERROR: Unable to open template file '%s'", *template)
		}
	}

	//Split up the data set passed in
	each_set := strings.Split(*data, ":~")

	for i := range each_set {
		split_data := strings.Split(each_set[i], "=")
		if len(split_data) > 1 {
			mapped_data[split_data[0]] = split_data[1]
		}
	}

	file_data := mustache.RenderFile(*template, mapped_data)

	if *output != "{file}" {
		out_file, err := os.Create(*output)
		if err != nil {
			log.Fatalf("ERROR: Unable to create output file %s", *output)
		}
		out_file.WriteString(file_data)
	}
}
