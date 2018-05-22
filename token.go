package rule

import (
	"bufio"
	"bytes"
	"strings"
)

const (
	// 不是规则单元，不是关键词 就是数值或者合法的变量名
	symbolToken = iota
	// 规则单元用反单引号包起来 例如 `c0 <= 1` 就是一个规则单元
	unitRuleToken
	// let 定义变量 例如： (let rs1 `c0 <= 10`)
	letToken
	// 左括号 ()
	leftParenthesisToken
	// 右括号 )
	rightParenthesisToken
	// and 对规则进行与运算 (let rs2 (and rs1 `c2 >= 100`) )
	andToken
	// not 对规则进行否运算 (let rs2 (not rs1 ) )
	notToken
	// or 对规则进行与运算 (let rs6 (rs3 rs3 rs2) )
	orToken
	// atLeast (let rs1 (atLeast 2 rs2 `c3 <= 100`) )
	atLeastToken
	// atMost (let rs1 (atMost 2 rs2 `c3 <= 100`) )
	atMostToken
	// export
	exportToken
)

var keyWords = map[string]int{
	"and":     andToken,
	"not":     notToken,
	"or":      orToken,
	"atLeast": atLeastToken,
	"atMost":  atMostToken,
	"let":     letToken,
	"export":  exportToken,
}

type token struct {
	tokenType int
	buf       string
}

// tokens 表示 token 列表
type tokens struct {
	tokens []*token
}

// append 在末尾添加一个 token
func (tks *tokens) append(tk *token) {
	tks.tokens = append(tks.tokens, tk)
}

// shift 弹出第一个 token
func (tks *tokens) shift() *token {
	firstToken := tks.tokens[0]
	tks.tokens = tks.tokens[1:]
	return firstToken
}

// peak 查看第一个 token
func (tks *tokens) peak() *token {
	return tks.tokens[0]
}

// length 返回长度
func (tks *tokens) length() int {
	return len(tks.tokens)
}

type scanner struct {
	reader *bufio.Reader
	buf    *bytes.Buffer
}

func (s *scanner) clearBuf() {
	s.buf = bytes.NewBufferString("")
}

func isBlank(r rune) bool {
	return r == '\r' || r == '\n' || r == '\t' || r == ' '
}

// tokenize 词法分析
func (s *scanner) tokenize() *tokens {
	tks := new(tokens)
	tks.tokens = make([]*token, 0)
	r, _, err := s.reader.ReadRune()
	for err == nil {
		if isBlank(r) {
			r, _, err = s.reader.ReadRune()
			continue
		}
		switch r {
		case '\'':
			s.clearBuf()
			tks.append(s.tokenizeUnitRuleToken())
		case '(':
			tks.append(&token{
				tokenType: leftParenthesisToken,
			})
		case ')':
			tks.append(&token{
				tokenType: rightParenthesisToken,
			})
		default:
			s.clearBuf()
			s.buf.WriteRune(r)
			tks.append(s.tokenizeKeywordSymbolToken())
		}
		r, _, err = s.reader.ReadRune()
	}
	return tks
}

func (s *scanner) tokenizeUnitRuleToken() *token {
	r, _, err := s.reader.ReadRune()
	for r != '\'' && err == nil {
		s.buf.WriteRune(r)
		r, _, err = s.reader.ReadRune()
	}
	return &token{
		tokenType: unitRuleToken,
		buf:       strings.TrimSpace(s.buf.String()),
	}
}

func (s *scanner) tokenizeKeywordSymbolToken() *token {
	r, _, err := s.reader.ReadRune()
	for !isBlank(r) && err == nil {
		s.buf.WriteRune(r)
		r, _, err = s.reader.ReadRune()
	}
	buf := s.buf.String()
	tkType, ok := keyWords[buf]
	if ok {
		return &token{
			tokenType: tkType,
		}
	}
	return &token{
		tokenType: symbolToken,
		buf:       buf,
	}
}

func tokenize(input string) *tokens {
	inputReplaced := strings.Replace(input, "'", " ' ", -1)
	inputReplaced = strings.Replace(input, "(", " ) ", -1)
	inputReplaced = strings.Replace(input, ")", " ) ", -1)
	s := scanner{
		reader: bufio.NewReader(bytes.NewBufferString(inputReplaced)),
	}
	return s.tokenize()
}
