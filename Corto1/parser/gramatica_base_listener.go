// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type BaseGramaticaListener struct{}

var _ GramaticaListener = &BaseGramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseGramaticaListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGramaticaListener) ExitStart(ctx *StartContext) {}

// EnterListas is called when production listas is entered.
func (s *BaseGramaticaListener) EnterListas(ctx *ListasContext) {}

// ExitListas is called when production listas is exited.
func (s *BaseGramaticaListener) ExitListas(ctx *ListasContext) {}

// EnterListaA is called when production listaA is entered.
func (s *BaseGramaticaListener) EnterListaA(ctx *ListaAContext) {}

// ExitListaA is called when production listaA is exited.
func (s *BaseGramaticaListener) ExitListaA(ctx *ListaAContext) {}

// EnterDA is called when production dA is entered.
func (s *BaseGramaticaListener) EnterDA(ctx *DAContext) {}

// ExitDA is called when production dA is exited.
func (s *BaseGramaticaListener) ExitDA(ctx *DAContext) {}
