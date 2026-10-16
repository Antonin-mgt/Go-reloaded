package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "il faut une entrée et une sortie")
		fmt.Fprintln(os.Stderr, "go run . <input_file> <output_file>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}


	resultat := transformer(string(data))

	// On écrit le résultat dans le fichier de sortie.
	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "impossible d'écrire le fichier", os.Args[2])
		os.Exit(1)
	}
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

// func mettreEnMajuscules(mots []string) []string {
// 	var resultat []string

// 	for _, mot := range mots {
// 		if mot == "(up)" && len(resultat) == 0 {
// 			continue
// 		}

// 		if mot == "(up)" {
// 			dernier := len(resultat) - 1
// 			resultat[dernier] = strings.ToUpper(resultat[dernier])
// 		} else {
// 			resultat = append(resultat, mot)
// 		}
// 	}

// 	return resultat
// }

func mettreEnMinuscules(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if mot == "(low)" && len(resultat) == 0 {
			continue
		}

		if mot == "(low)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToLower(resultat[dernier])
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}
func mettrePremiereLettreEnMajuscule(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if mot == "(cap)" && len(resultat) == 0 {
			continue
		}

		if mot == "(cap)" {
			dernier := len(resultat) - 1
			lettres := []rune(strings.ToLower(resultat[dernier]))
			lettres[0] = unicode.ToUpper(lettres[0])
			resultat[dernier] = string(lettres)
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
func transformer(texte string) string {
	// On coupe le texte à chaque retour à la ligne : on obtient une liste de lignes.
	lignes := strings.Split(texte, "\n")

	// On passe sur chaque ligne : i est son numéro (en commençant à 0), ligne est son contenu.
	for i, ligne := range lignes {
		
		ligne = strings.TrimRight(ligne, "\r")

	
		mots := strings.Fields(ligne)

		mots = convertirNombres(mots)
		mots = mettreEnMajuscules(mots)
		mots = mettreEnMinuscules(mots)
		mots = mettrePremiereLettreEnMajuscule(mots)

		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
