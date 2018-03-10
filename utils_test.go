package main

import (
	"testing"
)

func Test_Surroundings(t *testing.T) {

	var (
		str      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		loc      = 10
		radius   = 6
		expected = "EFGHIJKLMNOPQ"
		got, err = surroundings(str, loc, radius)
		oops     = expected != got || err != nil
	)

	if oops {
		t.Error(
			"\nFOR:", "str:", str, "loc:", loc, "radius:", radius,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
			"\nERROR:", err,
		)
	}
}

func Test_Surroundings2(t *testing.T) {

	var (
		str      = "ABC"
		loc      = 2
		radius   = 2
		expected = "ABC"
		got, err = surroundings(str, loc, radius)
		oops     = expected != got || err != nil
	)

	if oops {
		t.Error(
			"\nFOR:", "str:", str, "loc:", loc, "radius:", radius,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
			"\nERROR:", err,
		)
	}
}

func Test_Surroundings3(t *testing.T) {

	var (
		str      = "ABC"
		loc      = 0
		radius   = 2
		expected = "ABC"
		got, err = surroundings(str, loc, radius)
		oops     = expected != got || err != nil
	)

	if oops {
		t.Error(
			"\nFOR:", "str:", str, "loc:", loc, "radius:", radius,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
			"\nERROR:", err,
		)
	}
}

func Test_Surroundings4(t *testing.T) {

	var (
		str      = "ABC"
		loc      = -1
		radius   = 2
		expected = ""
		got, err = surroundings(str, loc, radius)
		oops     = expected != got || err == nil
	)

	if oops {
		t.Error(
			"\nFOR:", "str:", str, "loc:", loc, "radius:", radius,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
			"\nERROR:", err,
		)
	}
}

func Test_Surroundings5(t *testing.T) {

	var (
		str      = "ABC"
		loc      = 1
		expected = "ABC"
		got, err = surroundings(str, loc)
		oops     = expected != got || err != nil
	)

	if oops {
		t.Error(
			"\nFOR:", "str:", str, "loc:", loc,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
			"\nERROR:", err,
		)
	}
}

func Test_Surroundings6(t *testing.T) {

	var (
		str      = "A"
		loc      = 0
		expected = "A"
		got, err = surroundings(str, loc)
		oops     = expected != got || err != nil
	)

	if oops {
		t.Error(
			"\nFOR:", "str:", str, "loc:", loc,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
			"\nERROR:", err,
		)
	}
}


func Test_FindValueIndex(t *testing.T) {

	var (
		arr      = []string{"ABC", "DEF", "GHI"}
		got      = findIndex("GHI", arr)
		expected = 2
		oops     = expected != got
	)

	if oops {
		t.Error(
			"\nFOR:", arr,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
		)
	}

	
	arr      = []string{"ABC", "DEF", "GHI"}
	got      = findIndex("ABCD", arr)
	expected = -1
	oops     = expected != got


	if oops {
		t.Error(
			"\nFOR:", arr,
			"\nEXPECTED:", expected,
			"\nGOT:", got,
		)
	}
}