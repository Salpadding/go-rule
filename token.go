package rule

import (
	"regexp"
	"strings"
)

// Tokens 表示 token 列表
type Tokens struct {
	tokens []string
}

// append 在末尾添加一个 token
func (tks *Tokens) append(token string) {
	tks.tokens = append(tks.tokens, token)
}

// shift 弹出第一个 token
func (tks *Tokens) shift() string {
	firstToken := tks.tokens[0]
	tks.tokens = tks.tokens[1:]
	return firstToken
}

// peak 查看第一个 token
func (tks *Tokens) peak() string {
	return tks.tokens[0]
}

// length 返回长度
func (tks *Tokens) length() int {
	return len(tks.tokens)
}

// tokenize 通过词法分析将 lisp 代码转为 token 列表，用于语法分析
func tokenize(program string) *Tokens {
	replaced := strings.Replace(program, "(", " ( ", -1)
	replaced = strings.Replace(replaced, ")", " ) ", -1)
	// 把换行符和制表符号替换成空格
	replaced = strings.Replace(replaced, "\n", " ", -1)
	replaced = strings.Replace(replaced, "\t", " ", -1)
	regx := regexp.MustCompile("[\\s]+")
	tokens := new(Tokens)
	tokens.tokens = regx.Split(strings.Trim(replaced, " "), -1)
	return tokens
}
