// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 6, 16, 63, 10, 16, 13, 16, 14, 16, 64, 3, 16,
	3, 16, 2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2, 4, 3, 2, 50,
	59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 68, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 35, 3, 2, 2, 2,
	7, 37, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 43, 3, 2,
	2, 2, 15, 45, 3, 2, 2, 2, 17, 47, 3, 2, 2, 2, 19, 49, 3, 2, 2, 2, 21, 51,
	3, 2, 2, 2, 23, 53, 3, 2, 2, 2, 25, 55, 3, 2, 2, 2, 27, 57, 3, 2, 2, 2,
	29, 59, 3, 2, 2, 2, 31, 62, 3, 2, 2, 2, 33, 34, 7, 45, 2, 2, 34, 4, 3,
	2, 2, 2, 35, 36, 7, 47, 2, 2, 36, 6, 3, 2, 2, 2, 37, 38, 7, 44, 2, 2, 38,
	8, 3, 2, 2, 2, 39, 40, 7, 49, 2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7, 42,
	2, 2, 42, 12, 3, 2, 2, 2, 43, 44, 7, 43, 2, 2, 44, 14, 3, 2, 2, 2, 45,
	46, 7, 48, 2, 2, 46, 16, 3, 2, 2, 2, 47, 48, 9, 2, 2, 2, 48, 18, 3, 2,
	2, 2, 49, 50, 7, 67, 2, 2, 50, 20, 3, 2, 2, 2, 51, 52, 7, 68, 2, 2, 52,
	22, 3, 2, 2, 2, 53, 54, 7, 69, 2, 2, 54, 24, 3, 2, 2, 2, 55, 56, 7, 70,
	2, 2, 56, 26, 3, 2, 2, 2, 57, 58, 7, 71, 2, 2, 58, 28, 3, 2, 2, 2, 59,
	60, 7, 72, 2, 2, 60, 30, 3, 2, 2, 2, 61, 63, 9, 3, 2, 2, 62, 61, 3, 2,
	2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66,
	3, 2, 2, 2, 66, 67, 8, 16, 2, 2, 67, 32, 3, 2, 2, 2, 4, 2, 64, 3, 8, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'.'", "", "'A'", "'B'",
	"'C'", "'D'", "'E'", "'F'",
}

var lexerSymbolicNames = []string{
	"", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "PUNTO", "NUM", "LA", "LB",
	"LC", "LD", "LE", "LF", "WHITESPACE",
}

var lexerRuleNames = []string{
	"MAS", "MEN", "POR", "DIV", "PARI", "PARD", "PUNTO", "NUM", "LA", "LB",
	"LC", "LD", "LE", "LF", "WHITESPACE",
}

type GramaticaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewGramaticaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *GramaticaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaLexer(input antlr.CharStream) *GramaticaLexer {
	l := new(GramaticaLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GramaticaLexer tokens.
const (
	GramaticaLexerMAS        = 1
	GramaticaLexerMEN        = 2
	GramaticaLexerPOR        = 3
	GramaticaLexerDIV        = 4
	GramaticaLexerPARI       = 5
	GramaticaLexerPARD       = 6
	GramaticaLexerPUNTO      = 7
	GramaticaLexerNUM        = 8
	GramaticaLexerLA         = 9
	GramaticaLexerLB         = 10
	GramaticaLexerLC         = 11
	GramaticaLexerLD         = 12
	GramaticaLexerLE         = 13
	GramaticaLexerLF         = 14
	GramaticaLexerWHITESPACE = 15
)
