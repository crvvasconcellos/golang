package main

import (
	"fmt"
	"testing"
)

func Ola() string {
	return "Olá, mundo"
}

func main() {
	fmt.Println(Ola())
}

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
