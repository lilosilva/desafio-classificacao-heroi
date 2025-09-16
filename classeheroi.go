package main

import "fmt"

func main() {
	var numero int32
	var nome string
	var nivel string

	fmt.Print("Informe o nome do herói: ")
	fmt.Scan(&nome)

	fmt.Print("Informe a experiência (XP): ")
	fmt.Scan(&numero)

	switch {
	case numero < 1000:
		nivel = "Ferro"
	case numero > 1000 && numero <= 2000:
		nivel = "Bronze"
	case numero > 2000 && numero <= 5000:
		nivel = "Prata"
	case numero > 5000 && numero <= 7000:
		nivel = "Ouro"
	case numero > 7000 && numero <= 8000:
		nivel = "Platina"
	case numero > 8000 && numero <= 9000:
		nivel = "Ascendente"
	case numero > 9000 && numero <= 10000:
		nivel = "Imortal"
	case numero >= 10001:
		nivel = "Radiante"
	}

	fmt.Printf("O Herói de nome %s está no nível de %s\n", nome, nivel)
}
