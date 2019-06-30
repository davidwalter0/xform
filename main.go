package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

// VERSION is what is returned by the `-v` flag
var Version = "development"

var names = map[string]struct{}{"yaml2json": {}, "json2hcl": {}, "json2yaml": {}, "yaml2hcl": {}, "hcl2json": {}, "hcl2yaml": {}}

func isIn(name string) bool {
	_, ok := names[name]
	return ok
}

func main() {
	version := flag.Bool("version", false, "Prints current app version")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		return
	}

	name := os.Args[0]
	if !isIn(name) && len(os.Args) > 1 {
		name = os.Args[1]
	}
	name = path.Base(name)
	switch name {
	case "yaml2hcl":
		yaml2hcl()
	case "yaml2json":
		yaml2json()
	case "json2yaml":
		json2yaml()
	case "hcl2yaml":
		hcl2yaml()
	case "hcl2json":
		hcl2json()
	case "json2hcl":
		json2hcl()
	default:
		fmt.Fprintf(os.Stderr, "Program name %s or %v is unknown\n", name, os.Args)
		fmt.Fprintf(os.Stderr, `This program can be run by linking to the each of the call names and running one of the inputs

yaml2json
yaml2hcl
json2hcl
json2yaml
hcl2json
hcl2yaml

as the name for this multi call binary or calling as 
xlate [yaml2json, json2hcl, json2yaml, yaml2hcl, hcl2json, hcl2yaml]

e.g.

xlate yaml2json < yamlfile > jsonfile

`)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var raw []byte
var err error
var v interface{}

func read() ([]byte, error) {
	raw, err = ioutil.ReadAll(os.Stdin)
	return raw, err
}

func yaml2hcl() error {
	if raw, err = read(); err != nil {
		return err
	}

	err = yaml.Unmarshal(raw, &v)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	raw, err = json.Marshal(v)
	if err != nil {
		return err
	}

	ast, err := jsonParser.Parse([]byte(raw))
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}
	return nil
}

func hcl2yaml() error {
	if raw, err = read(); err != nil {
		return err
	}

	err = hcl.Unmarshal(raw, &v)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	raw, err = yaml.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Println(string(raw))
	return nil
}

func yaml2json() error {

	if raw, err = read(); err != nil {
		return err
	}

	err = yaml.Unmarshal(raw, &v)
	if err != nil {
		return err
	}

	raw, err = json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(raw))
	return nil
}

func json2yaml() error {

	if raw, err = read(); err != nil {
		return err
	}

	err = json.Unmarshal(raw, &v)
	if err != nil {
		return err
	}

	raw, err = yaml.Marshal(v)
	if err != nil {
		return err
	}

	fmt.Println(string(raw))
	return nil
}

func json2hcl() error {
	return toHCL()
}

func hcl2json() error {
	return toJSON()
}

func toJSON() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	var v interface{}
	err = hcl.Unmarshal(input, &v)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal json: %s", err)
	}
	fmt.Println(string(json))
	return nil
}

func toHCL() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	ast, err := jsonParser.Parse([]byte(input))
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}
	return nil
}
