// Copyright 2015 Philip Schlump
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Pack css Remove extra blank lines.

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
	OutFileName string `short:"o" long:"output"    description:"Path packed output to css file"        default:"out.css"`
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

	o := ""

	s := scanner.New(string(myCSS))
	for {
		token := s.Next()
		if token.Type == scanner.TokenError {
			fmt.Printf("Syntax Error: %s\n", token)
			os.Exit(2)
		} else if token.Type == scanner.TokenEOF {
			break
		}
		// Do something with the token...
		// fmt.Printf("%+v\n", token)

		switch token.Type {

		case scanner.TokenComment:
		case scanner.TokenS:

		case scanner.TokenIdent:
			fallthrough
		case scanner.TokenAtKeyword:
			fallthrough
		case scanner.TokenString:
			fallthrough
		case scanner.TokenHash:
			fallthrough
		case scanner.TokenNumber:
			fallthrough
		case scanner.TokenPercentage:
			fallthrough
		case scanner.TokenDimension:
			fallthrough
		case scanner.TokenURI:
			fallthrough
		case scanner.TokenUnicodeRange:
			fallthrough
		case scanner.TokenCDO:
			fallthrough
		case scanner.TokenCDC:
			fallthrough
		case scanner.TokenFunction:
			fallthrough
		case scanner.TokenIncludes:
			fallthrough
		case scanner.TokenDashMatch:
			fallthrough
		case scanner.TokenPrefixMatch:
			fallthrough
		case scanner.TokenSuffixMatch:
			fallthrough
		case scanner.TokenSubstringMatch:
			fallthrough
		case scanner.TokenBOM:
			o += token.Value
		case scanner.TokenChar:
			if token.Value == "}" {
				o += "}\n"
			} else {
				o += token.Value
			}
		}
	}
	// fmt.Printf("s=%s\n", o)
	ioutil.WriteFile(opts.OutFileName, []byte(o), 0644)

}
