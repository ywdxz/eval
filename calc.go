package eval

import (
	"fmt"
	"strconv"
)

type simpleLex struct {
	pos       int
	tokenList []*token
	ast       *ast
	err       error
}

// 这个接口必须实现，goyacc 是词法分析的接口
func (s *simpleLex) Lex(lval *yySymType) (tokID int) {

	w := s.Scan()
	if w == nil {
		return
	}

	var err error

	switch w.id {
	case NON_NEGATIVE_INTEGER:
		lval.value, err = strconv.ParseUint(w.raw, 10, 64)
		tokID = NUMBER
	default:
		tokID = int(w.id)
	}

	if err != nil {
		s.err = err
	}

	return
}

func (s *simpleLex) Scan() *token {

	if s.pos >= len(s.tokenList) {
		return nil
	}

	tok := s.tokenList[s.pos]

	s.pos++
	return tok
}

// 词法分析异常处理 goyacc 错误接口实现
func (s *simpleLex) Error(s1 string) {
	s.err = fmt.Errorf("%s", s1)
}

func Calc(raw string) (vale int64, err error) {

	tokenList, err := lexicalAnalysis(raw)
	if err != nil {
		return
	}

	lex := &simpleLex{
		tokenList: tokenList,
	}

	yyParse(lex)

	if lex.err != nil {
		err = lex.err
		return
	}

	return eval(lex.ast)
}
