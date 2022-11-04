package main

import "fmt"

const (
	Not   = '!'
	And   = '&'
	Or    = '|'
	Left  = '('
	Right = ')'
	True  = 't'
	False = 'f'
	Empty = byte(0)
)

type Opt struct {
	C   byte
	Val []bool
}

func (o *Opt) String() string {
	return fmt.Sprintf("%s:%v", string([]byte{o.C}), o.Val)
}

func (o *Opt) Result() bool {
	switch o.C {
	case Not:
		if len(o.Val) != 1 {
			return false
		}
		return !o.Val[0]
	case And:
		if len(o.Val) == 0 {
			return false
		}
		for _, b := range o.Val {
			if !b {
				return false
			}
		}
		return true
	case Or:
		if len(o.Val) == 0 {
			return false
		}
		for _, b := range o.Val {
			if b {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func NewOpt() Opt {
	return Opt{
		C: Empty,
	}
}

func parseBoolExpr(expression string) bool {

	var stack []Opt
	cur := NewOpt()
	for i := 0; i < len(expression); i++ {
		c := expression[i]
		switch c {
		case Not, And, Or:
			if cur.C != Empty {
				stack = append(stack, cur)
				cur = NewOpt()
			}
			cur.C = c
		case Left:
			// fmt.Println("Left", stack, cur)
		case Right:
			// fmt.Println("Right Before", stack, cur)
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				last.Val = append(last.Val, cur.Result())
				cur = last
			}
			// fmt.Println("Right After", stack, cur)
		case True:
			cur.Val = append(cur.Val, true)
		case False:
			cur.Val = append(cur.Val, false)
		default:
			continue
		}
	}

	if len(stack) == 0 {
		return cur.Result()
	}
	return stack[0].Result()
}
