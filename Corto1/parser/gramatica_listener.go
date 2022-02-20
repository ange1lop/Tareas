// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterListas is called when entering the listas production.
	EnterListas(c *ListasContext)

	// EnterListaA is called when entering the listaA production.
	EnterListaA(c *ListaAContext)

	// EnterDA is called when entering the dA production.
	EnterDA(c *DAContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitListas is called when exiting the listas production.
	ExitListas(c *ListasContext)

	// ExitListaA is called when exiting the listaA production.
	ExitListaA(c *ListaAContext)

	// ExitDA is called when exiting the dA production.
	ExitDA(c *DAContext)
}
