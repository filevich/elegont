package main

import (
	//"encoding/json"
	//"io/ioutil"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c, _ := NewConfig()
	if c.File_extension != ".ego" {
		t.Errorf("c.File_extension: want %v, got %v", ".ego", c.File_extension)
	}
	//t.Log(c)
}

func TestNewConfigCustom(t *testing.T) {
	c, err := NewConfig("./examples/config/test-custom-config-1.yaml")
	if err != nil {
		t.Error(err)
	}

	// if !ok {
	// 	t.Errorf("c.Var_declaration: want %v, got %v", [1]string{"var"}, c.Var_declaration)
	// }
	t.Log(c)
	t.Log(c.Syntax)
}

func TestNewConfigCustom2(t *testing.T) {
	c, err := NewConfig("./examples/config/Elegont.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Log(c)
	p := "as for asd"
	t.Log(c.Syntax["For"][0].Test(&p))

	q := c.Syntax["For"][0].(*inBlock)
	t.Log(q)
	t.Log(c.Syntax["For"][0].(*inBlock))
	t.Log(q.Test(&p))

	// t.Logf("%t", c.Syntax["For"].Definition)
}

func TestNewConfigCustom3(t *testing.T) {
	c, err := NewConfig("./examples/config/Elegont.yaml")
	check(err)

	t.Log(c)
}
