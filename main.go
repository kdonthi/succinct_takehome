package main

import (
	"awesomeProject/succinct"
	"fmt"
)

func main() {
	//Example 1: f(x) = x^2 + x + 5

	n := succinct.NewNodeBuilder()

	x := n.Var()
	xSq := n.Mul(x, x)
	five := n.Const(5)
	xSquaredPlusFive := n.Add(xSq, five)
	total := n.Add(xSquaredPlusFive, x)

	n.FillNodes(map[int]float64{
		x.GetId(): 2,
	})
	n.AssertEq(total, n.Val(11))
	fmt.Println(n.Verify(total))

	//Example 2: f(a) = (a+1) / 8
	//
	//function f(a):
	//   b = a + 1
	//   c = b / 8
	//   return c

	//n2 := succinct.NewNodeBuilder()
	//h2 := n2.Hint()
	//
	//a := n2.Var()
	//one := n2.Val(1)
	//b := n2.Add(a, one)
	//
	//c := h2.Build(h2.Div(h2.Val(b), h2.Const(8))) // get a node through the n2!
	//eight := n2.Const(8)
	//cTimesEight := n2.Mul(c, eight)
	//
	//n2.AssertEq(cTimesEight, b)
	//n2.AssertEq(c, n2.Const(0.25))
	//
	//n2.FillNodes(map[int]float64{
	//	a.GetId(): 1,
	//})
	//fmt.Println(n2.Verify(b))
	//fmt.Println(c.GetVal(), a.GetVal(), b.GetVal())

	//Example 3: f(x) = sqrt(x+7)
	//
	//Assume that x+7 is a perfect square (so x = 2 or 9, etc.).}

	//n3 := succinct.NewNodeBuilder()
	//h3 := n3.Hint()
	//x := n3.Var()
	//seven := n3.Const(7)
	//
	//xPlusSeven := n3.Add(x, seven)
	//y := h3.Build(h3.Sqrt(h3.Val(xPlusSeven)))
	//ySquared := n3.Mul(y, y)
	//n3.AssertEq(xPlusSeven, ySquared)
	//
	//n3.FillNodes(map[int]float64{
	//	x.GetId(): 2,
	//})
	//fmt.Println(xPlusSeven)
	//n3.AssertEq(y, n3.Val(3))
	//fmt.Println(n3.Verify(xPlusSeven))
	//fmt.Println(x.GetVal(), y.GetVal(), xPlusSeven.GetVal(), ySquared.GetVal())
}
