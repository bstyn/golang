package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	dni := [...]string{
		"Szeroki",
		"Wąski",
		"Łysy",
		"Gruby",
		"Zakolasty",
		"Epicki",
		"Nietypowy",
		"Robiący wrażenie",
		"Potężny",
		"Wyjątkowy",
		"Miły",
		"Niski",
		"Okrągły",
		"Zdziadziały",
		"Szybki",
		"Wolny",
		"Zdradliwy",
		"Różowy",
		"Waleczny",
		"Zawistny",
		"Cichy",
		"Martwy",
		"Nieumarły",
		"Angielski",
		"Brudny",
		"Pachnący",
		"Najgorszy",
		"Najlepszy",
		"Ulubiony",
		"Kolosalny"}

	imiona := map[string]string{
		"A": "Miodożer",
		"B": "Melon",
		"C": "Głaz",
		"D": "Tygrys",
		"E": "Niedźwiedź",
		"F": "Gigant",
		"G": "Wojownik",
		"H": "Ślimak",
		"I": "Maciuś",
		"J": "Złodziej",
		"K": "Rudzielec",
		"L": "Żyd",
		"M": "Zakolak",
		"N": "Dziad",
		"O": "Leniwiec",
		"P": "Mąż",
		"R": "Alkoholik",
		"S": "Artysta",
		"T": "Trębacz",
		"U": "Zawisza Czarny",
		"W": "Boombox",
		"V": "T-55",
		"X": "Ivan",
		"Y": "Barman",
		"Z": "Hydraulik",
		"Ź": "Krokodyl",
		"Ż": "Pies",
		"Ć": "Małysz",
		"Ś": "Pudzian",
		"Ó": "Oman"}

	nazwiska := map[string]string{
		"A": "Topiący się w jeziorze",
		"B": "Kąpiacy się w wannie",
		"C": "Idący do śpiulkolotu",
		"D": "Wyzywany przez goryla",
		"E": "Umiejący w regexa",
		"F": "w klubie",
		"G": "Na wykrzywionych wieżach",
		"H": "Spadający z drabiny",
		"I": "POG",
		"J": "Jadący na koniu",
		"K": "Uczący się Neo",
		"L": "Korzystający z Windowsa",
		"M": "z łysą glacą",
		"N": "kupujący w Stabucksie",
		"O": "Zamawiający kebaba",
		"P": "klepiący kotlety",
		"R": "Szukający zaczepki",
		"S": "Ukrywający się w sianie",
		"T": "Terroryzujący pierwszaków",
		"U": "Taplający się w błocie",
		"W": "który miał tu zjechać",
		"V": "Kradnący legitymacje",
		"X": "na emigracji",
		"Y": "na robocie w Niemczech",
		"Z": "Słuchający Edzia",
		"Ź": "Beka po kolce",
		"Ż": "Klaskający w samolocie",
		"Ć": "Skacząco 102 metry w Welingen",
		"Ś": "Co był drugi w kulach",
		"Ó": "Yemen"}

	dzien := flag.Int("dzien", 0, "Dzie urodzenia osoby")
	imie := flag.String("imie", "", "Imie osoby")
	nazwisko := flag.String("nazwisko", "", "Nazwisko osoby")
	flag.Parse()
	day := *dzien
	name := *imie
	lastname := *nazwisko
	if day == 0 {
		fmt.Print("Podaj dzien urodzenia: ")
		fmt.Scanf("%d\n", &day)
	}
	if name == "" {
		fmt.Print("Podaj imie: ")
		fmt.Scanf("%s\n", &name)
	}
	if lastname == "" {
		fmt.Print("Podaj nazwisko: ")
		fmt.Scanf("%s\n", &lastname)
	}
	name = strings.ToUpper(name[0:1])
	lastname = strings.ToUpper(lastname[0:1])
	fmt.Println(dni[day-1], imiona[name], nazwiska[lastname])
}
