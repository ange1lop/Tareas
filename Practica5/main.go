/*
	go mod init Laboratorio5
	go get github.com/antlr/antlr4/runtime/Go/antlr
	antlr -o parser Gramatica.g4
	go run main.go
*/

package main

import (
	"Laboratorio5/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct {
	*parser.BaseGramaticaListener
}

func analizar(input string) {
	fmt.Println(input)

	// create input stream
	stream := antlr.NewInputStream(input)
	// create lexer
	lexer := parser.NewGramaticaLexer(stream)
	// create tokenStream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// create parser
	parser := parser.NewGramaticaParser(tokenStream)
	// get tree
	tree := parser.Start()

	// activate listener
	listen := Listener{}
	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)
}

func main() {
	// input := "2 + 5 * 3"
	input := "(2+5+19*6)*(3/5*3+2)"
	analizar(input)
}
