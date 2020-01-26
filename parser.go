package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

// ParseFileData : parse and combine the HCL from all the input files
func ParseFileData(filedata map[string][]byte) Root {
	var root Root
	var combined Root

	parser := hclparse.NewParser()

	// check for .tf.json extension and parse TF JSON
	for filename, filebytes := range filedata {
		f, parseDiags := parser.ParseHCL(filebytes, filename)
		if parseDiags.HasErrors() {
			log.Fatal(parseDiags.Error())
		}

		diags := gohcl.DecodeBody(f.Body, nil, &root)
		if len(diags) != 0 {
			for _, diag := range diags {
				fmt.Printf("decoding - %s\n", diag)
			}
		}

		for _, local := range root.Locals {
			combined.Locals = append(combined.Locals, local)
		}
		for _, variable := range root.Variables {
			combined.Variables = append(combined.Variables, variable)
		}
		for _, resource := range root.Resources {
			combined.Resources = append(combined.Resources, resource)
		}
		for _, data := range root.Data {
			combined.Data = append(combined.Data, data)
		}
		for _, module := range root.Modules {
			combined.Modules = append(combined.Modules, module)
		}
		for _, output := range root.Outputs {
			combined.Outputs = append(combined.Outputs, output)
		}
		for _, provider := range root.Providers {
			combined.Providers = append(combined.Providers, provider)
		}
		for _, terraform := range root.Terraform {
			combined.Terraform = append(combined.Terraform, terraform)
		}
	}

	return combined
}

// ParseDirectory : parses a directory of HCL2 format files with .tf as an extension
func ParseDirectory(path string, verbose bool) (Root, map[string][]byte) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var combined Root
	var filedata map[string][]byte
	filedata = make(map[string][]byte)

	for _, file := range files {
		filename := fmt.Sprintf("%s/%s", path, file.Name())
		if strings.HasSuffix(filename, ".tf") {
			file, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			filebytes, err := ioutil.ReadAll(file)

			filedata[filename] = filebytes

			combined = ParseFileData(filedata)
		} else {
			if verbose {
				fmt.Printf("Ignoring %s\n", filename)
			}
		}
	}

	return combined, filedata
}
