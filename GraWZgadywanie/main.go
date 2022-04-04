package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sort"
	"strings"
	"os"
)
func main(){
	var a int 
	var b string
	var i int
	min := 1
	max := 100
	
	type Rekord struct{
		imie string
		wynik int
	}
	HallOfFame := make([]Rekord,0)
	
	Loop:
		for ;;{
			a = rand.Intn(max - min) + min
			fmt.Println("Witaj!")
			fmt.Println("Podaj liczbę z zakresu od",min,"do",max)
			proby := 0
			for a != i {
				fmt.Scanf("%s",&b)
				if b == "koniec"{
					fmt.Println("Żegnaj")
					break
				}
				i, _ = strconv.Atoi(b)
				if (a > i) {
					fmt.Println("Za mała")
				}
				if (a < i) {
					fmt.Println("Za duża")
				}
				proby += 1
				if (a == i) {
					fmt.Println("Gratulacje udało ci się w",proby,"prób!!!")
					var record Rekord
					var imie string
					fmt.Println("Podaj swoje imie: ")
					fmt.Scanf("%s\n",&imie)
					record.imie = imie
					record.wynik = proby
					HallOfFame = append(HallOfFame,record)
					fmt.Println("Czy chcesz zagrać jeszcze raz? [T/N]")
					var ans string
					fmt.Scanf("%s\n",&ans)
					if strings.ToUpper(ans) == "T"{
						break
					} else {
						break Loop
					}
				}
				
			}
		}
	sort.Slice(HallOfFame, func(i, j int) bool {
		return HallOfFame[i].wynik < HallOfFame[j].wynik
	})
	file, _ := os.Create("data.txt")
	file.Close()

	for x := range HallOfFame{
		fmt.Println(HallOfFame[x].imie,HallOfFame[x].wynik)
	}
	
}