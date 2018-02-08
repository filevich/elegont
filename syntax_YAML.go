package main

import (
	"regexp"
)

/* Parsing. What a wonderful world! /s */

type SyntaxDATA map[string]([]VariantDATA)

type VariantDATA struct {
	Variant, Definition string
	Delimiters          []string
}

func NewVariant(v VariantDATA) IVariant {
	var variant IVariant = nil

	switch v.Variant {
	case "inLine":
		variant = &inLine{
			Definition: regexp.MustCompile(v.Definition),
		}
	case "inBlock":
		var dels []Delimiter

		if len(v.Delimiters) > 0 {
			for _, d := range v.Delimiters {
				dels = append(dels, Delimiters[d])
			}
		}
		variant = &inBlock{
			Definition: regexp.MustCompile(v.Definition),
			Delimiters: dels,
		}
	default:
		panic(true)
	}

	return variant
}
