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

	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "impossible d'écrire le fichier", os.Args[2])
		os.Exit(1)
	}
}

// convertisseurHexEtBin gère (hex) et (bin) : le mot qui précède le marqueur
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

// capitaliser met la première lettre d'un mot en majuscule et le reste en minuscule.
// Elle passe par les runes pour ne pas couper une lettre accentuée.
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

// collerPonctuation colle la ponctuation (. , ! ? : ;) au mot précédent.
func collerPonctuation(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		reste := strings.TrimLeft(mot, ".,!?:;")
		debut := mot[:len(mot)-len(reste)]

		if debut != "" && len(resultat) > 0 {
			dernier := len(resultat) - 1
			resultat[dernier] = resultat[dernier] + debut

			if reste != "" {
				resultat = append(resultat, reste)
			}
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}

// Article transforme "a" en "an" devant une voyelle ou un h.
func Article(mots []string) []string {
	for i := 0; i < len(mots)-1; i++ {
		if mots[i] == "a" || mots[i] == "A" {
			lettres := []rune(strings.ToLower(mots[i+1]))

			if strings.Contains("aeiouh", string(lettres[0])) {
				mots[i] = mots[i] + "n"
			}
		}
	}
	return mots
}


// collerGuillemets colle chaque paire de ' aux mots qui les entourent.
func collerGuillemets(mots []string) []string {
	var resultat []string
	ouvert := false    
	enAttente := false 

	for _, mot := range mots {
		if mot == "'" {
			if !ouvert {
				ouvert = true
				enAttente = true
			} else {
				ouvert = false
				if enAttente {
					resultat = append(resultat, "''")
					enAttente = false
				} else {
					dernier := len(resultat) - 1
					resultat[dernier] = resultat[dernier] + "'"
				}
			}
		} else if enAttente {
				resultat = append(resultat, "'"+mot)
			enAttente = false
		} else {
			resultat = append(resultat, mot)
		}
	}
	if enAttente {
		resultat = append(resultat, "'")
	}

	return resultat
}

// transformer applique toutes les règles au texte, ligne par ligne, pour garder les retours à la ligne du fichier d'origine.
func transformer(texte string) string {
	
	lignes := strings.Split(texte, "\n")

	
	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")

		mots := strings.Fields(ligne)

		mots = convertisseurHexEtBin(mots)
		mots = transformationDuMarqueur(mots)
		mots = collerPonctuation(mots)
		mots = collerGuillemets(mots)
		mots = Article(mots)
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
