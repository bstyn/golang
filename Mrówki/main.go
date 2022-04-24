package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mrowka struct {
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

func ustawLiscie(arr [][]string, n int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		x := rand.Intn(len(arr)-2) + 1
		y := rand.Intn(len(arr[0])-2) + 1
		arr[x][y] = "L"
	}
}

func stworzMrowisko(height int, width int) [][]string {
	a := make([][]string, height)
	for i := range a {
		a[i] = make([]string, width)
		for x := range a[i] {
			if x == 0 || x == len(a[i])-1 {
				a[i][x] = "|"
			} else {
				a[i][x] = " "
			}

		}
		if i == 0 || i == len(a)-1 {
			for x := range a[i] {
				a[i][x] = "-"
			}
		}
	}
	return a
}

func stworzMrowki(mrowisko [][]string, ilosc int) []Mrowka {
	mrowki := make([]Mrowka, 10)
	for i := range mrowki {
		x := rand.Intn(len(mrowisko)-2) + 1
		y := rand.Intn(len(mrowisko[0])-2) + 1
		mrowki[i].set(x, y)
	}
	for i := range mrowki {
		mrowisko[mrowki[i].x][mrowki[i].y] = "M"
	}
	return mrowki
}

func moveMrowki(mrowisko [][]string, mrowki []Mrowka) {
	for i := range mrowki {
		mrowisko[mrowki[i].x][mrowki[i].y] = " "
		mrowki[i].move(rand.Intn(3)-1, rand.Intn(3)-1)
		mrowisko[mrowki[i].x][mrowki[i].y] = "M"
	}
}
func main() {

	mrowisko := stworzMrowisko(15, 30)
	ustawLiscie(mrowisko, 20)
	mrowki := stworzMrowki(mrowisko, 10)

	for i := 0; i < 5; i++ {
		fmt.Printf("\x1b[2J")
		for i := range mrowisko {
			for x := range mrowisko[i] {
				fmt.Print(mrowisko[i][x])
			}
			fmt.Println()
		}
		time.Sleep(1 * time.Second)
		moveMrowki(mrowisko, mrowki)
	}

}
