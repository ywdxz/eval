package eval

import (
	"errors"
	"fmt"
	"math"
)

var (
	AstIsEmptyErr error = errors.New("ast node is empty")
	DivisorIs0Err error = errors.New("divisor is 0")
)

type ast struct {
	tokenID   int
	leftNode  *ast
	rightNode *ast
	value     uint64
}

func genOpNode(tokenID int, l, r *ast) *ast {

	return &ast{
		leftNode:  l,
		rightNode: r,
		tokenID:   tokenID,
	}
}

func genValueNode(val uint64) *ast {

	return &ast{
		tokenID: NUMBER,
		value:   val,
	}
}

func printlnAST(cur *ast) {

	if cur == nil {
		return
	}

	if cur.tokenID != NUMBER {
		fmt.Printf("%p,%+v:%c\n", cur, cur, cur.tokenID)
	} else {
		fmt.Printf("%p,%+v: uint\n", cur, cur)
	}

	printlnAST(cur.leftNode)

	printlnAST(cur.rightNode)
}

func eval(cur *ast) (ret int64, err error) {

	if cur == nil {
		err = AstIsEmptyErr
		return
	}

	switch cur.tokenID {
	case NUMBER:
		ret = int64(cur.value)
	default:
		leftValue, er1 := eval(cur.leftNode)
		if er1 != nil {
			err = er1
			return
		}
		rightValue, er2 := eval(cur.rightNode)
		if er2 != nil {
			err = er2
			return
		}
		switch cur.tokenID {
		case '+':
			ret = leftValue + rightValue
		case '-':
			ret = leftValue - rightValue
		case '*':
			ret = leftValue * rightValue
		case '/':
			if rightValue == 0 {
				err = DivisorIs0Err
				return
			}
			ret = leftValue / rightValue
		case '^':
			ret = int64(math.Pow(float64(leftValue), float64(rightValue)))
		case '%':
			ret = leftValue % rightValue
		}
	}
	return
}
