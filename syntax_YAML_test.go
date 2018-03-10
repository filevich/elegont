package main

import (
    "testing"
    "strings"
)

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

func Test_VariantDATA_got(t *testing.T) {
    var (
        vDT = &VariantDATA{ Definition:"hi,HOWareyou?ImFINE,Thanks..." }
        got = vDT.got("HOW")
        expected = true
        oops = expected != got
    )

    if oops {
        t.Error(
            "\nFOR:", vDT.Definition,
            "\nGOT:", got,
            "\nEXPECTED:", expected,
        )
    }

}


func Test_VariantDATA_got2(t *testing.T) {
    var (
        vDT = &VariantDATA{ Definition:"hi,HOWareyou?ImFINE,Thanks..." }
        got = vDT.got("ThanksJJ")
        expected = false
        oops = expected != got
    )

    if oops {
        t.Error(
            "\nFOR:", vDT.Definition,
            "\nGOT:", got,
            "\nEXPECTED:", expected,
        )
    }

}


func Test_VariantDATA_got3(t *testing.T) {
    var (
        vDT = &VariantDATA{ Definition:"hi,HOWareyou?ImFINE,Thanks..." }
        got = vDT.got("HOW","are","you")
        expected = true
        oops = expected != got
    )

    if oops {
        t.Error(
            "\nFOR:", vDT.Definition,
            "\nGOT:", got,
            "\nEXPECTED:", expected,
        )
    }

}


func Test_VariantDATA_got4(t *testing.T) {
    var (
        vDT = &VariantDATA{ Definition:"hi,HOWareyou?ImFINE,Thanks..." }
        got = vDT.got("HOW","areNOT","you")
        expected = false
        oops = expected != got
    )

    if oops {
        t.Error(
            "\nFOR:", vDT.Definition,
            "\nGOT:", got,
            "\nEXPECTED:", expected,
        )
    }

}

// case sensitive test
func Test_VariantDATA_got5(t *testing.T) {
    var (
        vDT = &VariantDATA{ Definition:"hi,HOWareyou?ImFINE,Thanks..." }
        got = vDT.got("HOW","ARE","you")
        expected = false // (?)
        oops = expected != got
    )

    if oops {
        t.Error(
            "\nFOR:", vDT.Definition,
            "\nGOT:", got,
            "\nEXPECTED:", expected,
        )
    }

    // no erros. yep, case sesitive.

}
