// Copyright 2015 Philip Schlump
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/pschlump/css/scanner"
)

var opts struct {
	CssFileName string `short:"s" long:"cssFile"    description:"Path to css file"        default:""`
}

func main() {

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		fmt.Printf("Invalid Command Line: %s\n", err)
		os.Exit(1)
	}

	if opts.CssFileName == "" {
		fmt.Printf("Usage test1 -s File.css\n")
		os.Exit(1)
	}

	myCSS, err := ioutil.ReadFile(opts.CssFileName)
	if err != nil {
		fmt.Printf("Errror: %s\n", err)
		os.Exit(1)
	}

	s := scanner.New(string(myCSS))
	for {
		token := s.Next()
		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		// Do something with the token...
		fmt.Printf("%+v\n", token)
	}

}
