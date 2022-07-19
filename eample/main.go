package main

import (
	"fmt"
	"queue_interface"
)

func main() {
	t := queue_interface.New(10)

	err := t.Push(1)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Push("珍惜")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Push(1)
	if err != nil {
		fmt.Println(err)
	}
	/*err := t.Push("珍惜")
	if err != nil {
		fmt.Println(err)
	}*/
	/*for i := 0; i < 5; i++ {
		v, err := t.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v)
	}*/
	/*v, err := t.Top()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	t.Show()*/

}
