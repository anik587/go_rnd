package main

import (
	"fmt"
	"sync"
)

var score = []int{
	0,
}

var wgrp = &sync.WaitGroup{}
var mutx = &sync.Mutex{}
var ranges int = 100

func main() {
	wgrp.Add(ranges)
	for i := 0; i < ranges; i++ {
		fmt.Printf("From %d \n", i)
		go func(wgr *sync.WaitGroup, m *sync.Mutex, i int) {
			fmt.Printf("From inside %d \n", i)
			mutx.Lock()
			score = append(score, i)
			mutx.Unlock()
			wgrp.Done()
		}(wgrp, mutx, i)

	}

	// go func(wgr *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("From 2")
	// 	mutx.Lock()
	// 	score = append(score, 2)
	// 	mutx.Unlock()
	// 	wgrp.Done()
	// }(wgrp, mutx)

	// go func(wgr *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("From 3")
	// 	mutx.Lock()
	// 	score = append(score, 3)
	// 	mutx.Unlock()
	// 	wgrp.Done()
	// }(wgrp, mutx)

	// go func(wrg *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("From 4")
	// 	mutx.Lock()
	// 	score = append(score, 4)
	// 	mutx.Unlock()
	// 	wgrp.Done()
	// }(wgrp, mutx)

	wgrp.Wait()
	fmt.Println(score)
}
