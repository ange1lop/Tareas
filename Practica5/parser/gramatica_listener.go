// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExp2 is called when entering the exp2 production.
	EnterExp2(c *Exp2Context)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExp2 is called when exiting the exp2 production.
	ExitExp2(c *Exp2Context)
}
