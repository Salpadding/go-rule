package rule

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var commentPattern = regexp.MustCompile(`/\*.*?\*/`)

func eval(node *astNode, e *env) (Ruler, error) {
	switch node.token.tokenType {
	case andToken:
		if node.next == nil {
			return nil, errors.New("and requires a expression")
		}
		rulers := make([]Ruler, 0)
		for n := node.next; n != nil; n = n.next {
			newRuler, err := eval(n, e)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return AndRuler(rulers), nil
	case orToken:
		if node.next == nil {
			return nil, errors.New("or requires a expression")
		}
		rulers := make([]Ruler, 0)
		for n := node.next; n != nil; n = n.next {
			newRuler, err := eval(n, e)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return OrRuler(rulers), nil
	case atLeastToken:
		if node.next == nil || node.next.next == nil {
			return nil, errors.New("atLeast requires a symbol and rule expressions")
		}
		rulers := make([]Ruler, 0)
		n, err := strconv.Atoi(node.next.token.buf)
		if err != nil {
			return nil, err
		}
		for n := node.next.next; n != nil; n = n.next {
			newRuler, err := eval(n, e)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return NewAtLeastNRuler(rulers, n), nil
	case atMostToken:
		if node.next == nil || node.next.next == nil {
			return nil, errors.New("atMost requires a symbol and rule expressions")
		}
		rulers := make([]Ruler, 0)
		n, err := strconv.Atoi(node.next.token.buf)
		if err != nil {
			return nil, err
		}
		for n := node.next.next; n != nil; n = n.next {
			newRuler, err := eval(n, e)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return NewAtMostNRuler(rulers, n), nil
	// 括号
	case leftParenthesisToken:
		return eval(node.child, e)
	// 否运算
	case notToken:
		newRuler, err := eval(node.next, e)
		if err != nil {
			return nil, err
		}
		return Negate(newRuler), nil
	case letToken:
		if node.next == nil || node.next.next == nil {
			return nil, errors.New("let requires a symbol and a rule expression")
		}
		newRuler, err := eval(node.next.next, e)
		if err != nil {
			return nil, err
		}
		e.setLocal(node.next.token.buf, newRuler)
		return newRuler, nil
	case exportToken:
		if node.next == nil || node.next.next == nil {
			return nil, errors.New("export requires a symbol and a rule expression")
		}
		newRuler, err := eval(node.next.next, e)
		if err != nil {
			return nil, err
		}
		e.setGlobal(node.next.token.buf, newRuler)
		return newRuler, nil
	case unitRuleToken:
		return parseUnitRule(node.token.buf)
	// symbol token
	default:
		return e.get(node.token.buf)
	}
}

type ExpectResult struct {
	RuleName  string `json:"rule_name"`
	Ok        bool   `json:"ok"`
	IsPending bool   `json:"isPending"`
}

type Evaluator struct {
	e *env
}

func (evaler *Evaluator) Eval(expressions string) (Ruler, error) {
	expressions = commentPattern.ReplaceAllString(expressions, "")
	var ruler Ruler
	var err error
	for _, expr := range strings.Split(expressions, ";") {
		if strings.TrimSpace(expr) == "" {
			continue
		}
		node, err := buildAST(tokenize(expr))
		if err != nil {
			return nil, errors.New("expression " + expr + " parse fail")
		}
		ruler, err = eval(node, evaler.e)
		if err != nil {
			return nil, err
		}
	}
	return ruler, err
}

func (evaler *Evaluator) Expect(c Context) map[string]*ExpectResult {
	results := make(map[string]*ExpectResult)
	for n, r := range evaler.e.globals {
		result, err := r.Rule(c)
		results[n] = &ExpectResult{
			RuleName:  n,
			Ok:        result,
			IsPending: err != nil,
		}
	}
	return results
}

func NewEvaluator() *Evaluator {
	return &Evaluator{
		e: newEnv(),
	}
}
