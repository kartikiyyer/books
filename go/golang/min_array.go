package main
import "fmt"

func min_arrray(a []int) min int {
	if len(a) == 1 {
		return a[0];
	}

	for i=1;i<len(a)-1;i++ {
		if a[i] < a[i-1] && a[i] >  {

		}
	}
}

func main() {
	a := []int {1,2,3,4,5,6,7,8};
	b := []int {8,9,10,0,1,2,3,4};
	c := []int {10,11,12,13,14,17,3,4,5,6,7,8};
	fmt.Println(min_array(a));
	fmt.Println(min_array(b));
	fmt.Println(min_array(c));
}