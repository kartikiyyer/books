// My version of implementing structs

package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Price  float64
	Volume int
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 12.1, 23, true}
	fmt.Println(t1)
	t1.Price = 33.3
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)

	t2 := Trade{}
	fmt.Println(t2)

	t3 := Trade{
		Symbol: "MSFT",
		Price:  12.12,
		Buy:    true,
		Volume: 4,
	}
	fmt.Printf("%+v", t3)
}
