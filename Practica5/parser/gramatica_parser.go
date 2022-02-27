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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 38, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 5, 2, 9, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10, 3, 12, 3, 14, 3,
	36, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 4, 3, 2, 5, 6, 3, 2, 3, 4, 2, 40,
	2, 8, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 9, 5, 4, 3, 2, 7, 9, 7, 2, 2, 3,
	8, 6, 3, 2, 2, 2, 8, 7, 3, 2, 2, 2, 9, 3, 3, 2, 2, 2, 10, 11, 8, 3, 1,
	2, 11, 12, 7, 7, 2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7, 8, 2, 2, 14, 15,
	8, 3, 1, 2, 15, 21, 3, 2, 2, 2, 16, 17, 7, 9, 2, 2, 17, 21, 8, 3, 1, 2,
	18, 19, 7, 10, 2, 2, 19, 21, 8, 3, 1, 2, 20, 10, 3, 2, 2, 2, 20, 16, 3,
	2, 2, 2, 20, 18, 3, 2, 2, 2, 21, 34, 3, 2, 2, 2, 22, 23, 12, 7, 2, 2, 23,
	24, 9, 2, 2, 2, 24, 25, 5, 4, 3, 8, 25, 26, 8, 3, 1, 2, 26, 33, 3, 2, 2,
	2, 27, 28, 12, 6, 2, 2, 28, 29, 9, 3, 2, 2, 29, 30, 5, 4, 3, 7, 30, 31,
	8, 3, 1, 2, 31, 33, 3, 2, 2, 2, 32, 22, 3, 2, 2, 2, 32, 27, 3, 2, 2, 2,
	33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 5, 3, 2,
	2, 2, 36, 34, 3, 2, 2, 2, 6, 8, 20, 32, 34,
}
var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}
var symbolicNames = []string{
	"", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "NUM", "ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "exp2",
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
	dir  string
	tipo string
}

var desp int = 0
var tmp int = 1

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
	GramaticaParserNUM        = 7
	GramaticaParserID         = 8
	GramaticaParserWHITESPACE = 9
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start = 0
	GramaticaParserRULE_exp2  = 1
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

func (s *StartContext) Exp2() IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
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

	p.SetState(6)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserPARI, GramaticaParserNUM, GramaticaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(4)
			p.exp2(0)
		}

	case GramaticaParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(5)
			p.Match(GramaticaParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExp2Context is an interface to support dynamic dispatch.
type IExp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExp2Context

	// GetE returns the e rule contexts.
	GetE() IExp2Context

	// GetE2 returns the e2 rule contexts.
	GetE2() IExp2Context

	// SetE1 sets the e1 rule contexts.
	SetE1(IExp2Context)

	// SetE sets the e rule contexts.
	SetE(IExp2Context)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExp2Context)

	// GetDir returns the dir attribute.
	GetDir() string

	// GetNum returns the num attribute.
	GetNum() int

	// SetDir sets the dir attribute.
	SetDir(string)

	// SetNum sets the num attribute.
	SetNum(int)

	// IsExp2Context differentiates from other interfaces.
	IsExp2Context()
}

type Exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	num    int
	e1     IExp2Context
	e      IExp2Context
	n      antlr.Token
	op     antlr.Token
	e2     IExp2Context
}

func NewEmptyExp2Context() *Exp2Context {
	var p = new(Exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_exp2
	return p
}

func (*Exp2Context) IsExp2Context() {}

func NewExp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2Context {
	var p = new(Exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_exp2

	return p
}

func (s *Exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp2Context) GetN() antlr.Token { return s.n }

func (s *Exp2Context) GetOp() antlr.Token { return s.op }

func (s *Exp2Context) SetN(v antlr.Token) { s.n = v }

func (s *Exp2Context) SetOp(v antlr.Token) { s.op = v }

func (s *Exp2Context) GetE1() IExp2Context { return s.e1 }

func (s *Exp2Context) GetE() IExp2Context { return s.e }

func (s *Exp2Context) GetE2() IExp2Context { return s.e2 }

func (s *Exp2Context) SetE1(v IExp2Context) { s.e1 = v }

func (s *Exp2Context) SetE(v IExp2Context) { s.e = v }

func (s *Exp2Context) SetE2(v IExp2Context) { s.e2 = v }

func (s *Exp2Context) GetDir() string { return s.dir }

func (s *Exp2Context) GetNum() int { return s.num }

func (s *Exp2Context) SetDir(v string) { s.dir = v }

func (s *Exp2Context) SetNum(v int) { s.num = v }

func (s *Exp2Context) PARI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARI, 0)
}

func (s *Exp2Context) PARD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARD, 0)
}

func (s *Exp2Context) AllExp2() []IExp2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExp2Context)(nil)).Elem())
	var tst = make([]IExp2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExp2Context)
		}
	}

	return tst
}

func (s *Exp2Context) Exp2(i int) IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
}

func (s *Exp2Context) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *Exp2Context) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Exp2Context) POR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPOR, 0)
}

func (s *Exp2Context) DIV() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIV, 0)
}

func (s *Exp2Context) MAS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAS, 0)
}

func (s *Exp2Context) MEN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMEN, 0)
}

func (s *Exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExp2(s)
	}
}

func (s *Exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExp2(s)
	}
}

func (p *GramaticaParser) Exp2() (localctx IExp2Context) {
	return p.exp2(0)
}

func (p *GramaticaParser) exp2(_p int) (localctx IExp2Context) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, GramaticaParserRULE_exp2, _p)
	var _la int

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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserPARI:
		{
			p.SetState(9)
			p.Match(GramaticaParserPARI)
		}
		{
			p.SetState(10)

			var _x = p.exp2(0)

			localctx.(*Exp2Context).e = _x
		}
		{
			p.SetState(11)
			p.Match(GramaticaParserPARD)
		}

		localctx.(*Exp2Context).num = localctx.(*Exp2Context).GetE().GetNum()
		localctx.(*Exp2Context).dir = localctx.(*Exp2Context).GetE().GetDir()

	case GramaticaParserNUM:
		{
			p.SetState(14)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*Exp2Context).n = _m
		}
		localctx.(*Exp2Context).num = 0
		localctx.(*Exp2Context).dir = (func() string {
			if localctx.(*Exp2Context).GetN() == nil {
				return ""
			} else {
				return localctx.(*Exp2Context).GetN().GetText()
			}
		}())

	case GramaticaParserID:
		{
			p.SetState(16)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Exp2Context).n = _m
		}
		localctx.(*Exp2Context).num = 0
		localctx.(*Exp2Context).dir = (func() string {
			if localctx.(*Exp2Context).GetN() == nil {
				return ""
			} else {
				return localctx.(*Exp2Context).GetN().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				localctx.(*Exp2Context).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp2)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(21)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp2Context).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserPOR || _la == GramaticaParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp2Context).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(22)

					var _x = p.exp2(6)

					localctx.(*Exp2Context).e2 = _x
				}

				tmp = tmp - localctx.(*Exp2Context).GetE1().GetNum() - localctx.(*Exp2Context).GetE2().GetNum()
				localctx.(*Exp2Context).num = 1
				localctx.(*Exp2Context).dir = "t" + strconv.Itoa(tmp)
				gen(localctx.(*Exp2Context).dir + "=" + localctx.(*Exp2Context).GetE1().GetDir() + (func() string {
					if localctx.(*Exp2Context).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp2Context).GetOp().GetText()
					}
				}()) + localctx.(*Exp2Context).GetE2().GetDir())
				tmp = tmp + 1

			case 2:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				localctx.(*Exp2Context).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp2)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(26)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp2Context).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserMAS || _la == GramaticaParserMEN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp2Context).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(27)

					var _x = p.exp2(5)

					localctx.(*Exp2Context).e2 = _x
				}

				tmp = tmp - localctx.(*Exp2Context).GetE1().GetNum() - localctx.(*Exp2Context).GetE2().GetNum()
				localctx.(*Exp2Context).num = 1
				localctx.(*Exp2Context).dir = "t" + strconv.Itoa(tmp)
				gen(localctx.(*Exp2Context).dir + "=" + localctx.(*Exp2Context).GetE1().GetDir() + (func() string {
					if localctx.(*Exp2Context).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp2Context).GetOp().GetText()
					}
				}()) + localctx.(*Exp2Context).GetE2().GetDir())
				tmp = tmp + 1

			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *Exp2Context = nil
		if localctx != nil {
			t = localctx.(*Exp2Context)
		}
		return p.Exp2_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Exp2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
