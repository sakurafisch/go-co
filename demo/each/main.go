package main

import (
	"fmt"
	"time"

	"github.com/magicdawn/go-co"
	"github.com/magicdawn/go-co/coutil"
)

func main() {
	var items = []interface{}{1, 2, 3, 4, 5}

	// new Task
	var t = coutil.Each(items, func(item interface{}, index int, items []interface{}) *co.Task {
		return co.Async(func() interface{} {
			var s = item.(int)
			time.Sleep(time.Second * time.Duration(s))
			fmt.Printf("[%d] = %d\n", index, item)
			return s
		})
	})

	fmt.Println(co.Await(t))
}
