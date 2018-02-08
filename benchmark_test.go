package main

import (
    "testing"
)

type Blah struct {
    c complex128
    s string
    f float64
}

func (b *Blah) doPtr() {
}

func (b Blah) doCopy() {
}

func BenchmarkDoPtr(b *testing.B) {
    blah := Blah{}
    for i := 0; i < b.N; i++ {
        (&blah).doPtr()
    }
}

func BenchmarkDoCopy(b *testing.B) {
    blah := Blah{}
    for i := 0; i < b.N; i++ {
        blah.doCopy()
    }
}
