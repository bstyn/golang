package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Rekord struct {
	imie  string
	wynik int
	data  string
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
func trimLastRune(s string) string {
	return s[:len(s)-1]
}

func printLines(filePath string, values interface{}) error {
	f, err := os.Create(filePath)
	fmt.Println()
	if err != nil {
		return err
	}
	defer f.Close()
	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Slice {
		return errors.New("Not a slice")
	}
	for i := 0; i < rv.Len(); i++ {
		fmt.Fprintln(f, rv.Index(i).Interface())
	}
	return nil
}
func main() {
	var a int
	var b string
	var i int
	min := 1
	max := 100
	rand.Seed(time.Now().UnixNano())

	HallOfFame := make([]Rekord, 0)

	f, e := os.Open("data.txt")
	if e != nil {
		fmt.Println("Error is = ", e)
	}

	fmt.Println("Start")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		s := strings.Split(scanner.Text(), " ")
		wynik, _ := strconv.Atoi(s[1])
		var tempRekord Rekord
		tempRekord.imie, tempRekord.wynik, tempRekord.data = trimFirstRune(s[0]), wynik, trimLastRune(s[2])
		HallOfFame = append(HallOfFame, tempRekord)

	}
	fmt.Println("Witaj!")
Loop:
	for {
		a = rand.Intn(max-min) + min
		fmt.Println("Podaj liczbę z zakresu od", min, "do", max)
		proby := 0
		for a != i {
			fmt.Scanf("%s\n", &b)
			if b == "koniec" {
				fmt.Println("Żegnaj")
				break Loop
			}
			if b == "statystyki" {
				for x := range HallOfFame {
					fmt.Println(HallOfFame[x].imie, HallOfFame[x].wynik, HallOfFame[x].data)
				}
				break
			}
			i, _ = strconv.Atoi(b)
			if a > i {
				fmt.Println("Za mała")
			}
			if a < i {
				fmt.Println("Za duża")
			}
			proby += 1
			if a == i {
				fmt.Println("Gratulacje udało ci się w", proby, "prób!!!")
				var record Rekord
				var imie string
				fmt.Println("Podaj swoje imie: ")
				fmt.Scanf("%s\n", &imie)
				record.imie = imie
				record.wynik = proby
				record.data = time.Now().Format("2006-01-02")
				HallOfFame = append(HallOfFame, record)
				sort.Slice(HallOfFame, func(i, j int) bool {
					return HallOfFame[i].wynik < HallOfFame[j].wynik
				})
				fmt.Println("Czy chcesz zagrać jeszcze raz? [T/N]")
				var ans string
				fmt.Scanf("%s\n", &ans)
				if strings.ToUpper(ans) == "T" {
					break
				} else {
					break Loop
				}
			}

		}
	}
	printLines("data.txt", HallOfFame)
}
