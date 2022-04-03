package main

import (
	"flag"
	"fmt"
)

func main() {
	liczba1 := flag.Int("liczba1", 0, "pierwsza liczba do sumy")
	liczba2 := flag.Int("liczba2", 0, "druga liczba do sumy")
	flag.Parse()
	a := *liczba1
	b := *liczba2
	if a == 0 {
		fmt.Print("Podaj 1 liczbę: ")
		fmt.Scanf("%d\n", &a)
	}
	if b == 0 {
		fmt.Print("Podaj 2 liczbę: ")
		fmt.Scanf("%d\n", &b)
	}
	fmt.Println("Wynikiem równania ", a, " + ", b, " jest liczba ", a+b)
}
