package main

import ("strings"
		"strconv"
)

func mettreEnMajuscules(mots []string) []string {
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

	if mot == "(up)" && len(resultat) == 0 {
			continue
		}

		if mot == "(up)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		} else if mot == "(up," && i+1 < len(mots) {
			texteNombre := strings.Trim(mots[i+1], ")")
			n, err := strconv.Atoi(texteNombre)

			if err == nil {
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				for j := debut; j < len(resultat); j++ {
					resultat[j] = strings.ToUpper(resultat[j])
				}
				i = i + 1
				} else {
				resultat = append(resultat, mot)
			}
		} else {
			resultat = append(resultat, mot)
		}
		}
		return resultat
	}
