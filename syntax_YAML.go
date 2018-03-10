package main

import (
	"regexp"
	"strings"
)

/* Parsing. What a wonderful world! /s */

type SyntaxDATA map[string]([]VariantDATA)

type VariantDATA struct {
	Definition string
	Delimiters []string
}

func (this *VariantDATA) got(attrs ...string) bool {
	var count int = 0
	for _, attr := range(attrs) {
		isThere := strings.Contains(this.Definition, attr)
		if isThere {
			count++
		}
	}
	return len(attrs) == count
}

func (this *VariantDATA) gotDelimiters() bool {
	return this.Delimiters == []
}

func (this *VariantDATA) getDelimiters() []Delimiter {
	var dels []Delimiter

	if len(v.Delimiters) > 0 {
		for _, d := range v.Delimiters {
			dels = append(dels, Delimiters[d])
		}
	}

	return dels
}

func NewVariant(component string, vDT VariantDATA) IVariant {
	var variant IVariant = nil
	var gotDelimiters bool = vDT.gotDelimiters
	var dels = vDT.getDelimiters()

	switch component {
	case "import":
		if gotDelimiters {
			variant = &BImport{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LImport{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "if":
		if gotDelimiters {
			variant = &BIf{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LIf{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "while":
		if gotDelimiters {
			variant = &BWhile{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LWhile{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "package":
		if gotDelimiters {
			variant = &BPackage{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LPackage{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "type":		
		variant = &LType{
			Definition: regexp.MustCompile(v.Definition),
			Delimiters: dels,
		}
	case "struct":
		if gotDelimiters {
			variant = &BImport{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LImport{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "variable":
		if gotDelimiters {
			variant = &BVariable{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LVariable{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "for":
		if gotDelimiters {
			variant = &BFor{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LFor{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "func":
		if gotDelimiters {
			variant = &BFunc{
				Definition: regexp.MustCompile(v.Definition),
				Delimiters: dels,
			}
		}
		else {
			variant = &LIfunc{
				Definition: regexp.MustCompile(v.Definition),
			}
		}
	case "inc":
		variant = &LInc{
			Definition: regexp.MustCompile(v.Definition),
			Delimiters: dels,
		}
		
	default:
		panic(true)
	}

	return variant
}
