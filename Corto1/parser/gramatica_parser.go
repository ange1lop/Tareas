// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Laboratorio5/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 54, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 5, 2, 13, 10,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 23, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 33, 10, 4, 12, 4, 14,
	4, 36, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 52, 10, 5, 3, 5, 2, 3, 6, 6, 2, 4, 6,
	8, 2, 2, 2, 58, 2, 12, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2,
	8, 51, 3, 2, 2, 2, 10, 13, 5, 4, 3, 2, 11, 13, 7, 2, 2, 3, 12, 10, 3, 2,
	2, 2, 12, 11, 3, 2, 2, 2, 13, 3, 3, 2, 2, 2, 14, 15, 5, 6, 4, 2, 15, 16,
	7, 9, 2, 2, 16, 17, 5, 6, 4, 2, 17, 18, 8, 3, 1, 2, 18, 23, 3, 2, 2, 2,
	19, 20, 5, 6, 4, 2, 20, 21, 8, 3, 1, 2, 21, 23, 3, 2, 2, 2, 22, 14, 3,
	2, 2, 2, 22, 19, 3, 2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 25, 8, 4, 1, 2, 25,
	26, 5, 8, 5, 2, 26, 27, 8, 4, 1, 2, 27, 34, 3, 2, 2, 2, 28, 29, 12, 4,
	2, 2, 29, 30, 5, 8, 5, 2, 30, 31, 8, 4, 1, 2, 31, 33, 3, 2, 2, 2, 32, 28,
	3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 7, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 38, 7, 10, 2, 2, 38, 52, 8,
	5, 1, 2, 39, 40, 7, 11, 2, 2, 40, 52, 8, 5, 1, 2, 41, 42, 7, 12, 2, 2,
	42, 52, 8, 5, 1, 2, 43, 44, 7, 13, 2, 2, 44, 52, 8, 5, 1, 2, 45, 46, 7,
	14, 2, 2, 46, 52, 8, 5, 1, 2, 47, 48, 7, 15, 2, 2, 48, 52, 8, 5, 1, 2,
	49, 50, 7, 16, 2, 2, 50, 52, 8, 5, 1, 2, 51, 37, 3, 2, 2, 2, 51, 39, 3,
	2, 2, 2, 51, 41, 3, 2, 2, 2, 51, 43, 3, 2, 2, 2, 51, 45, 3, 2, 2, 2, 51,
	47, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 9, 3, 2, 2, 2, 6, 12, 22, 34, 51,
}
var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'.'", "", "'A'", "'B'",
	"'C'", "'D'", "'E'", "'F'",
}
var symbolicNames = []string{
	"", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "PUNTO", "NUM", "LA", "LB",
	"LC", "LD", "LE", "LF", "WHITESPACE",
}

var ruleNames = []string{
	"start", "listas", "listaA", "dA",
}

type GramaticaParser struct {
	*antlr.BaseParser
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GramaticaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

type Nodo struct {
	dir    float64
	tipo   float64
	cadena string
}

var desp int = 0

// var env []*entorno.Entorno
var sup *entorno.Entorno

// func push(e *entorno.Entorno) {
//     env = append(env, e)
// }

// func pop() *entorno.Entorno {
//     if len(env) < 1 {
//         panic("empty env")
//     }
//     result := env[len(env) - 1]
//     env = env[:len(env) - 1]
//     return result
// }

func gen(out string) {
	fmt.Println(out)
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserMAS        = 1
	GramaticaParserMEN        = 2
	GramaticaParserPOR        = 3
	GramaticaParserDIV        = 4
	GramaticaParserPARI       = 5
	GramaticaParserPARD       = 6
	GramaticaParserPUNTO      = 7
	GramaticaParserNUM        = 8
	GramaticaParserLA         = 9
	GramaticaParserLB         = 10
	GramaticaParserLC         = 11
	GramaticaParserLD         = 12
	GramaticaParserLE         = 13
	GramaticaParserLF         = 14
	GramaticaParserWHITESPACE = 15
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start  = 0
	GramaticaParserRULE_listas = 1
	GramaticaParserRULE_listaA = 2
	GramaticaParserRULE_dA     = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Listas() IListasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListasContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(10)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserNUM, GramaticaParserLA, GramaticaParserLB, GramaticaParserLC, GramaticaParserLD, GramaticaParserLE, GramaticaParserLF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.Listas()
		}

	case GramaticaParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Match(GramaticaParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListasContext is an interface to support dynamic dispatch.
type IListasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t rule contexts.
	GetT() IListaAContext

	// GetId returns the id rule contexts.
	GetId() IListaAContext

	// SetT sets the t rule contexts.
	SetT(IListaAContext)

	// SetId sets the id rule contexts.
	SetId(IListaAContext)

	// IsListasContext differentiates from other interfaces.
	IsListasContext()
}

type ListasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      IListaAContext
	id     IListaAContext
}

func NewEmptyListasContext() *ListasContext {
	var p = new(ListasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_listas
	return p
}

func (*ListasContext) IsListasContext() {}

func NewListasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListasContext {
	var p = new(ListasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_listas

	return p
}

func (s *ListasContext) GetParser() antlr.Parser { return s.parser }

func (s *ListasContext) GetT() IListaAContext { return s.t }

func (s *ListasContext) GetId() IListaAContext { return s.id }

func (s *ListasContext) SetT(v IListaAContext) { s.t = v }

func (s *ListasContext) SetId(v IListaAContext) { s.id = v }

func (s *ListasContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPUNTO, 0)
}

func (s *ListasContext) AllListaA() []IListaAContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListaAContext)(nil)).Elem())
	var tst = make([]IListaAContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListaAContext)
		}
	}

	return tst
}

func (s *ListasContext) ListaA(i int) IListaAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaAContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListaAContext)
}

func (s *ListasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterListas(s)
	}
}

func (s *ListasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitListas(s)
	}
}

func (p *GramaticaParser) Listas() (localctx IListasContext) {
	this := p
	_ = this

	localctx = NewListasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_listas)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)

			var _x = p.listaA(0)

			localctx.(*ListasContext).t = _x
		}
		{
			p.SetState(13)
			p.Match(GramaticaParserPUNTO)
		}
		{
			p.SetState(14)

			var _x = p.listaA(0)

			localctx.(*ListasContext).id = _x
		}

		sim := localctx.(*ListasContext).GetT().GetNodo().dir + localctx.(*ListasContext).GetId().GetNodo().tipo + 0.00
		cad := fmt.Sprint(sim)
		gen(cad)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)

			var _x = p.listaA(0)

			localctx.(*ListasContext).t = _x
		}

		cad := fmt.Sprintf("%f", localctx.(*ListasContext).GetT().GetNodo().dir)
		gen(cad)

	}

	return localctx
}

// IListaAContext is an interface to support dynamic dispatch.
type IListaAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IListaAContext

	// GetEl returns the el rule contexts.
	GetEl() IDAContext

	// SetL sets the l rule contexts.
	SetL(IListaAContext)

	// SetEl sets the el rule contexts.
	SetEl(IDAContext)

	// GetNodo returns the nodo attribute.
	GetNodo() *Nodo

	// SetNodo sets the nodo attribute.
	SetNodo(*Nodo)

	// IsListaAContext differentiates from other interfaces.
	IsListaAContext()
}

type ListaAContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	nodo   *Nodo
	l      IListaAContext
	el     IDAContext
}

func NewEmptyListaAContext() *ListaAContext {
	var p = new(ListaAContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_listaA
	return p
}

func (*ListaAContext) IsListaAContext() {}

func NewListaAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaAContext {
	var p = new(ListaAContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_listaA

	return p
}

func (s *ListaAContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaAContext) GetL() IListaAContext { return s.l }

func (s *ListaAContext) GetEl() IDAContext { return s.el }

func (s *ListaAContext) SetL(v IListaAContext) { s.l = v }

func (s *ListaAContext) SetEl(v IDAContext) { s.el = v }

func (s *ListaAContext) GetNodo() *Nodo { return s.nodo }

func (s *ListaAContext) SetNodo(v *Nodo) { s.nodo = v }

func (s *ListaAContext) DA() IDAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDAContext)
}

func (s *ListaAContext) ListaA() IListaAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaAContext)
}

func (s *ListaAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaAContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaAContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterListaA(s)
	}
}

func (s *ListaAContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitListaA(s)
	}
}

func (p *GramaticaParser) ListaA() (localctx IListaAContext) {
	return p.listaA(0)
}

func (p *GramaticaParser) listaA(_p int) (localctx IListaAContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaAContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaAContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GramaticaParserRULE_listaA, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)

		var _x = p.DA()

		localctx.(*ListaAContext).el = _x
	}

	localctx.(*ListaAContext).nodo = &Nodo{}
	din2 := 1 / 16.00
	localctx.(*ListaAContext).nodo.tipo = localctx.(*ListaAContext).GetEl().GetNodo().tipo * din2
	localctx.(*ListaAContext).nodo.dir = localctx.(*ListaAContext).GetEl().GetNodo().dir

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaAContext(p, _parentctx, _parentState)
			localctx.(*ListaAContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_listaA)
			p.SetState(26)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(27)

				var _x = p.DA()

				localctx.(*ListaAContext).el = _x
			}
			localctx.(*ListaAContext).nodo = &Nodo{}
			localctx.(*ListaAContext).nodo.dir = localctx.(*ListaAContext).GetL().GetNodo().dir*16 + localctx.(*ListaAContext).GetEl().GetNodo().dir
			divisor := localctx.(*ListaAContext).GetL().GetNodo().tipo / 16
			localctx.(*ListaAContext).nodo.tipo = localctx.(*ListaAContext).GetL().GetNodo().tipo + localctx.(*ListaAContext).GetEl().GetNodo().dir*divisor

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IDAContext is an interface to support dynamic dispatch.
type IDAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e token.
	GetE() antlr.Token

	// SetE sets the e token.
	SetE(antlr.Token)

	// GetNodo returns the nodo attribute.
	GetNodo() *Nodo

	// SetNodo sets the nodo attribute.
	SetNodo(*Nodo)

	// IsDAContext differentiates from other interfaces.
	IsDAContext()
}

type DAContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	nodo   *Nodo
	e      antlr.Token
}

func NewEmptyDAContext() *DAContext {
	var p = new(DAContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_dA
	return p
}

func (*DAContext) IsDAContext() {}

func NewDAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DAContext {
	var p = new(DAContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_dA

	return p
}

func (s *DAContext) GetParser() antlr.Parser { return s.parser }

func (s *DAContext) GetE() antlr.Token { return s.e }

func (s *DAContext) SetE(v antlr.Token) { s.e = v }

func (s *DAContext) GetNodo() *Nodo { return s.nodo }

func (s *DAContext) SetNodo(v *Nodo) { s.nodo = v }

func (s *DAContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *DAContext) LA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLA, 0)
}

func (s *DAContext) LB() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLB, 0)
}

func (s *DAContext) LC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLC, 0)
}

func (s *DAContext) LD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLD, 0)
}

func (s *DAContext) LE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLE, 0)
}

func (s *DAContext) LF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLF, 0)
}

func (s *DAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DAContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DAContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDA(s)
	}
}

func (s *DAContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDA(s)
	}
}

func (p *GramaticaParser) DA() (localctx IDAContext) {
	this := p
	_ = this

	localctx = NewDAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_dA)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserNUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*DAContext).e = _m
		}
		localctx.(*DAContext).nodo = &Nodo{}
		strVar := (func() string {
			if localctx.(*DAContext).GetE() == nil {
				return ""
			} else {
				return localctx.(*DAContext).GetE().GetText()
			}
		}())
		intVar, err := strconv.ParseFloat(strVar, 8)
		if err != nil {
			gen("error en conversion")
		}
		localctx.(*DAContext).nodo.dir = intVar
		localctx.(*DAContext).nodo.tipo = intVar

	case GramaticaParserLA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)

			var _m = p.Match(GramaticaParserLA)

			localctx.(*DAContext).e = _m
		}

		localctx.(*DAContext).nodo = &Nodo{}
		localctx.(*DAContext).nodo.dir = 10.00
		localctx.(*DAContext).nodo.tipo = 10.00

	case GramaticaParserLB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)

			var _m = p.Match(GramaticaParserLB)

			localctx.(*DAContext).e = _m
		}

		localctx.(*DAContext).nodo = &Nodo{}
		localctx.(*DAContext).nodo.dir = 11.00
		localctx.(*DAContext).nodo.tipo = 11.00

	case GramaticaParserLC:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)

			var _m = p.Match(GramaticaParserLC)

			localctx.(*DAContext).e = _m
		}

		localctx.(*DAContext).nodo = &Nodo{}
		localctx.(*DAContext).nodo.dir = 12.00
		localctx.(*DAContext).nodo.tipo = 12.00

	case GramaticaParserLD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)

			var _m = p.Match(GramaticaParserLD)

			localctx.(*DAContext).e = _m
		}

		localctx.(*DAContext).nodo = &Nodo{}
		localctx.(*DAContext).nodo.dir = 13.00
		localctx.(*DAContext).nodo.tipo = 13.00

	case GramaticaParserLE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(45)

			var _m = p.Match(GramaticaParserLE)

			localctx.(*DAContext).e = _m
		}

		localctx.(*DAContext).nodo = &Nodo{}
		localctx.(*DAContext).nodo.dir = 14.00
		localctx.(*DAContext).nodo.tipo = 14.00

	case GramaticaParserLF:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(47)

			var _m = p.Match(GramaticaParserLF)

			localctx.(*DAContext).e = _m
		}

		localctx.(*DAContext).nodo = &Nodo{}
		localctx.(*DAContext).nodo.dir = 15.00
		localctx.(*DAContext).nodo.tipo = 15.00

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ListaAContext = nil
		if localctx != nil {
			t = localctx.(*ListaAContext)
		}
		return p.ListaA_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) ListaA_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
