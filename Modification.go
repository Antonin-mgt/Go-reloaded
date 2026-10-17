package main

import ("strings"
		"strconv"
)
// transformationDuMarqueur gère les marqueurs (up), (low), (cap) et leurs versions avec un nombre (up, 2), (low, 2), (cap, 2).
func transformationDuMarqueur(mots []string) []string {
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

	if (mot == "(up)" || mot == "(low)" || mot == "(cap)") && len(resultat) == 0 {
			continue
		}

		if mot == "(up)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		
			} else if mot == "(low)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToLower(resultat[dernier])

			} else if mot == "(cap)" {
			dernier := len(resultat) - 1
			resultat[dernier] = capitaliser(resultat[dernier])

		
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
			} else if mot == "(low," && i+1 < len(mots) {
			texteNombre := strings.Trim(mots[i+1], ")")
			n, err := strconv.Atoi(texteNombre)
			if err == nil {
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				for j := debut; j < len(resultat); j++ {
					resultat[j] = strings.ToLower(resultat[j])
				}
				i = i + 1
				} else {
				resultat = append(resultat, mot)
			}
			
			} else if mot == "(cap," && i+1 < len(mots) {
			texteNombre := strings.Trim(mots[i+1], ")")
			n, err := strconv.Atoi(texteNombre)

			if err == nil {
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				for j := debut; j < len(resultat); j++ {
					resultat[j] = capitaliser(resultat[j]) 
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
