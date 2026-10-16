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

func convertisseurHexEtBin(mots []string) []string {
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


func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}


func collerPonctuation(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		// reste : le mot sans la ponctuation du début (",and" donne "and", "!!" donne "")
		reste := strings.TrimLeft(mot, ".,!?:;")
		// debut : la ponctuation retirée (",and" donne ",", "!!" donne "!!")
		debut := mot[:len(mot)-len(reste)]

		if debut != "" && len(resultat) > 0 {
			dernier := len(resultat) - 1
			resultat[dernier] = resultat[dernier] + debut

			// S'il reste un mot après la ponctuation, on l'ajoute comme mot normal.
			if reste != "" {
				resultat = append(resultat, reste)
			}
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

		mots = convertisseurHexEtBin(mots)
		mots = mettreEnMajuscules(mots)
		mots = collerPonctuation(mots)
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
