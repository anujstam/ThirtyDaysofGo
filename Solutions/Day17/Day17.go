package main

import(
  "errors"
  "fmt"
  "math"
)

type calculator struct {
	n int
	p int
}

func(c *calculator) power() (int,error) {
	if c.n > 0 && c.p > 0 {
		return int(math.Pow(float64(c.n),float64(c.p))),nil
	} else {
    	return 0, errors.New("n and p should be non-negative")
  	}
}

func NewCalc() *calculator {
	return &calculator{}
}

func main() {
	a := NewCalc()
  	var i int
  	var j int
  	fmt.Scanf("%d\n",&i)
  	fmt.Scanf("%d\n",&j)
  	a.n = i
	a.p = j
	pow,err := a.power()
  	if err!=nil{
    		fmt.Println(err)
    		return
  	}
  	fmt.Println(pow)
}
