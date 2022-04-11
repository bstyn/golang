package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mrowka struct{
	x int
	y int
}

func (m *Mrowka) set(x int, y int) {
	(*m).x = x
	(*m).y = y
}

func (m *Mrowka) move(x int, y int) {
	(*m).x = m.x + x
	(*m).y = m.y + y
}

func UstawLiscie(arr [][]string,n int ){
	rand.Seed(time.Now().UnixNano())
	for i := 0; i<n;i++ {
		x := rand.Intn(len(arr) - 0 ) + 0
		y := rand.Intn(len(arr[0]) - 0 ) + 0
		arr[x][y] = "L"
	}
}

func main(){

	a := make([][]string, 10)
	for i := range a {
		if i == 0 || i == 10{
			a[i] = make([]string, 10)
			for x := range a[i]{
				a[i][x] = "--"
			}	
		}
    	a[i] = make([]string, 10)
	}
	
	UstawLiscie(a,20)

	for i := range a {
		fmt.Println(a[i])
	}
}