package main

type Cost [3]int

const (
	R = iota
	B
	G
)

func (c *Cost) SetR(v int) { (*c)[R] = v }
func (c *Cost) SetG(v int) { (*c)[G] = v }
func (c *Cost) SetB(v int) { (*c)[B] = v }

func (c Cost) GetR() int { return c[R] }
func (c Cost) GetG() int { return c[G] }
func (c Cost) GetB() int { return c[B] }

func minCost(costs [][]int) int {

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var pre Cost

	for _, cost := range costs {
		cur := pre
		// R看GB
		cur.SetR(min(pre.GetG(), pre.GetB()) + cost[R])
		// G看RB
		cur.SetG(min(pre.GetR(), pre.GetB()) + cost[G])
		// B看RG
		cur.SetB(min(pre.GetR(), pre.GetG()) + cost[B])
		pre = cur
	}

	var last = pre
	return min(min(last.GetR(), last.GetG()), last.GetB())
}
