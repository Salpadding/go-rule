package rule

import (
	"errors"
	"regexp"
	"strconv"
)

var ops = []string{"=", "!=", ">", ">=", "<", "<="}

var unitRulePattern = regexp.MustCompile("(.+?)(=|!=|>=|<=|<|>)(.+)")

func parseUnitRule(token string) (Ruler, error) {
	params := unitRulePattern.FindStringSubmatch(token)
	if len(params) != 4 {
		return nil, errors.New("parse error: invalid expression " + token)
	}
	index := -1
	for idx, s := range ops {
		if s == params[2] {
			index = idx
		}
	}
	if index == -1 {
		return nil, errors.New("parse error: invalid operator " + params[2])
	}
	val, err := strconv.ParseFloat(params[3], 64)
	if err != nil {
		return nil, errors.New("parse error: " + params[3] + " is not a valid number")
	}
	return NewComparatorRule(params[1], index, val)
}

func parse(node *ASTNode) (Ruler, error) {
	switch node.NodeType {
	case andNode:
		rulers := make([]Ruler, 0)
		for n := node.Next; n != nil; n = n.Next {
			newRuler, err := parse(n)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return AndRuler(rulers), nil
	case orNode:
		rulers := make([]Ruler, 0)
		for n := node.Next; n != nil; n = n.Next {
			newRuler, err := parse(n)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return OrRuler(rulers), nil
	case atLeastNode:
		rulers := make([]Ruler, 0)
		n, err := strconv.Atoi(node.Next.Token)
		if err != nil {
			return nil, err
		}
		for n := node.Next.Next; n != nil; n = n.Next {
			newRuler, err := parse(n)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return NewAtLeastNRuler(rulers, n), nil
	case atMostNode:
		rulers := make([]Ruler, 0)
		n, err := strconv.Atoi(node.Next.Token)
		if err != nil {
			return nil, err
		}
		for n := node.Next.Next; n != nil; n = n.Next {
			newRuler, err := parse(n)
			if err != nil {
				return nil, err
			}
			rulers = append(rulers, newRuler)
		}
		return NewAtMostNRuler(rulers, n), nil
	case parenthesisNode:
		return parse(node.Child)
	default:
		return parseUnitRule(node.Token)
	}
}

func ParseRule(s string) (Ruler, error) {
	return parse(buildAST(tokenize(s)))
}
