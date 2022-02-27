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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 52, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 35, 10, 8, 13, 8,
	14, 8, 36, 3, 9, 3, 9, 7, 9, 41, 10, 9, 12, 9, 14, 9, 44, 11, 9, 3, 10,
	6, 10, 47, 10, 10, 13, 10, 14, 10, 48, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 6, 3, 2, 50,
	59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 54, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2,
	2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 29,
	3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 34, 3, 2, 2, 2, 17, 38, 3, 2, 2, 2,
	19, 46, 3, 2, 2, 2, 21, 22, 7, 45, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7,
	47, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 44, 2, 2, 26, 8, 3, 2, 2, 2, 27,
	28, 7, 49, 2, 2, 28, 10, 3, 2, 2, 2, 29, 30, 7, 42, 2, 2, 30, 12, 3, 2,
	2, 2, 31, 32, 7, 43, 2, 2, 32, 14, 3, 2, 2, 2, 33, 35, 9, 2, 2, 2, 34,
	33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2,
	2, 37, 16, 3, 2, 2, 2, 38, 42, 9, 3, 2, 2, 39, 41, 9, 4, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 18, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 47, 9, 5, 2, 2, 46, 45, 3,
	2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49,
	50, 3, 2, 2, 2, 50, 51, 8, 10, 2, 2, 51, 20, 3, 2, 2, 2, 6, 2, 36, 42,
	48, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "NUM", "ID", "WHITESPACE",
}

var lexerRuleNames = []string{
	"MAS", "MEN", "POR", "DIV", "PARI", "PARD", "NUM", "ID", "WHITESPACE",
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
	GramaticaLexerNUM        = 7
	GramaticaLexerID         = 8
	GramaticaLexerWHITESPACE = 9
)
