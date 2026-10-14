package main

import (
	"fmt"
	"log"
	"os"
	"strconv" 
	"strings" 
)


func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "il faut une entrée et une sortie")
		fmt.Fprintln(os.Stderr, "go run main.go <input_file> <output_file>")
		os.Exit(1)
	}
	
	fmt.Println(os.Args)
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)

	fmt.Println(convertirNombres(strings.Fields(string(data))))
}

func convertirNombres(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if (mot == "(hex)" || mot == "(bin)") && len(resultat) == 0 {
			continue
		}

		if mot == "(hex)" {
			dernier := len(resultat) - 1
			valeur, err := strconv.ParseInt(resultat[dernier], 16, 64)
			if err == nil {
				resultat[dernier] = strconv.FormatInt(valeur, 10)
			}
		} else if mot == "(bin)" {
			dernier := len(resultat) - 1
			valeur, err := strconv.ParseInt(resultat[dernier], 2, 64)
			if err == nil {
				resultat[dernier] = strconv.FormatInt(valeur, 10)
			}
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}