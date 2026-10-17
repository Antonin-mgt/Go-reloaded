# Go-reloaded

Go-reloaded est un petit outil en ligne de commande, écrit en Go, qui complète, édite et corrige automatiquement un texte. Il lit un fichier texte, applique une série de règles (conversion de nombres, changement de casse, ponctuation, guillemets, article « a » / « an »), puis écrit le texte corrigé dans un second fichier.

Il n'utilise que la bibliothèque standard de Go.

## Installation

Il faut avoir Go installé, dans la version indiquée dans `go.mod` (1.27.1) ou une version plus récente. Pour le vérifier :

```bash
go version
```

Ensuite, récupérer le projet et se placer dans son dossier :

```bash
git clone <url-du-depot>
cd <nom-du-dossier>
```

Aucune autre installation n'est nécessaire.

## Utilisation

Le programme se lance **depuis le dossier du projet** (celui qui contient `go.mod`), avec deux arguments : le fichier à corriger, puis le fichier à créer.

```bash
go run . <fichier_entree> <fichier_sortie>
```

Si le fichier de sortie existe déjà, son contenu est remplacé.

Exemple :

```bash
$ echo "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." > exemple.txt
$ go run . exemple.txt sortie.txt
$ cat sortie.txt
Simply add 66 and 2 and you will see the result is 68.
```

## Les règles

| Règle | Effet | Exemple | Résultat |
|---|---|---|---|
| `(hex)` | Remplace le mot précédent (un nombre hexadécimal) par sa valeur décimale | `1E (hex) files were added` | `30 files were added` |
| `(bin)` | Remplace le mot précédent (un nombre binaire) par sa valeur décimale | `It has been 10 (bin) years` | `It has been 2 years` |
| `(up)` | Met le mot précédent en majuscules | `Ready, set, go (up) !` | `Ready, set, GO!` |
| `(low)` | Met le mot précédent en minuscules | `I should stop SHOUTING (low)` | `I should stop shouting` |
| `(cap)` | Met une majuscule au début du mot précédent (et le reste en minuscules) | `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| `(up, n)`, `(low, n)`, `(cap, n)` | Applique la règle aux `n` mots précédents | `This is so exciting (up, 2)` | `This is SO EXCITING` |
| Ponctuation `. , ! ? : ;` | Collée au mot précédent, séparée du mot suivant par un espace | `I was sitting over there ,and then BAMM !!` | `I was sitting over there, and then BAMM!!` |
| Groupes de ponctuation (`...`, `!?`) | Restent ensemble et sont collés au mot précédent | `I was thinking ... You were right` | `I was thinking... You were right` |
| Guillemets simples `'` | Les deux `'` sont collés au mot qu'ils entourent | `I am exactly how they describe me: ' awesome '` | `I am exactly how they describe me: 'awesome'` |
| Guillemets autour de plusieurs mots | Le premier `'` est collé au premier mot, le second au dernier | `As Elton John said: ' I am the most well-known homosexual in the world '` | `As Elton John said: 'I am the most well-known homosexual in the world'` |
| `a` devant une voyelle ou un `h` | `a` devient `an` | `There it was. A amazing rock!` | `There it was. An amazing rock!` |

### Ordre d'application

Les règles s'appliquent **ligne par ligne**, toujours dans cet ordre :

1. `(hex)` et `(bin)`
2. `(up)`, `(low)`, `(cap)` et leurs versions avec un nombre
3. la ponctuation
4. les guillemets simples
5. `a` devient `an`

Cet ordre compte : dans `Ready, set, go (up) !`, le marqueur est d'abord appliqué, puis la ponctuation est collée, ce qui donne `GO!`.

### Cas particuliers

- Un marqueur placé **avant tout mot** sur sa ligne est ignoré : `(up) bonjour` donne `bonjour`.
- Un marqueur inconnu est laissé tel quel : `un (truc) inconnu` reste `un (truc) inconnu`.
- Si le mot avant `(hex)` ou `(bin)` n'est pas un nombre valide, il reste inchangé et le marqueur disparaît : `zz (hex)` donne `zz`.
- Si `n` est plus grand que le nombre de mots disponibles, tous les mots de la ligne sont transformés : `un deux (cap, 5) trois` donne `Un Deux trois`.
- Les lignes vides sont conservées, et les espaces en trop sont réduits à un seul.

## Les messages d'erreur

Les erreurs sont affichées sur la sortie d'erreur, et le programme s'arrête avec le code 1.

| Situation | Message affiché |
|---|---|
| Nombre d'arguments incorrect (aucun, un seul, ou plus de deux) | `il faut une entrée et une sortie` suivi de `go run . <input_file> <output_file>` |
| Fichier d'entrée introuvable ou illisible | Le message de Go, précédé de la date et de l'heure, par exemple `2026/10/17 19:10:48 open absent.txt: no such file or directory` |
| Écriture impossible (par exemple, dossier de sortie inexistant) | `impossible d'écrire le fichier dossier_absent/sortie.txt` |

## Les tests

Le dossier contient deux fichiers de test :

- `sample.txt` : un texte d'entrée qui utilise plusieurs règles.
- `result.txt` : le résultat que le programme doit produire pour ce texte.

Pour lancer le test, on écrit le résultat dans un **autre** fichier (`sortie.txt`), puis on le compare à `result.txt`. Il ne faut pas écrire directement dans `result.txt`, sinon le résultat attendu serait écrasé.

```bash
go run . sample.txt sortie.txt
diff sortie.txt result.txt
```

Si les deux fichiers sont identiques, `diff` n'affiche rien. Sous PowerShell, la commande de comparaison est `fc.exe sortie.txt result.txt`. Dans VS Code, on peut aussi comparer les deux fichiers côte à côte.

Pour tester les messages d'erreur :

```bash
go run .                                    # aucun argument
go run . sample.txt                         # un seul argument
go run . absent.txt sortie.txt              # fichier d'entrée absent
go run . sample.txt dossier_absent/s.txt    # dossier de sortie absent
```

## Organisation du code

Arborescence du projet :

```text
Go-reloaded/
├── main.go
├── Modification.go
├── go.mod
├── sample.txt
├── result.txt
└── README.md
```

| Fichier | Contenu |
|---|---|
| `main.go` | Le point d'entrée et la plupart des règles |
| `Modification.go` | Les marqueurs de casse : `(up)`, `(low)`, `(cap)` et leurs versions avec un nombre |
| `go.mod` | Le nom du module et la version de Go |
| `sample.txt` | Texte d'entrée du test |
| `result.txt` | Résultat attendu pour `sample.txt` |
| `README.md` | Cette documentation |

Les fonctions, et ce qu'elles font :

| Fonction | Fichier | Rôle |
|---|---|---|
| `main` | `main.go` | Vérifie les arguments, lit le fichier d'entrée, écrit le fichier de sortie |
| `transformer` | `main.go` | Découpe le texte en lignes et applique toutes les règles, dans l'ordre, à chaque ligne |
| `convertisseurHexEtBin` | `main.go` | Gère `(hex)` et `(bin)` |
| `transformationDuMarqueur` | `Modification.go` | Gère `(up)`, `(low)`, `(cap)` et `(up, n)`, `(low, n)`, `(cap, n)` |
| `capitaliser` | `main.go` | Met la première lettre d'un mot en majuscule et le reste en minuscules, sans abîmer les lettres accentuées |
| `collerPonctuation` | `main.go` | Colle la ponctuation au mot précédent |
| `collerGuillemets` | `main.go` | Colle chaque paire de `'` aux mots qui les entourent |
| `Article` | `main.go` | Transforme `a` en `an` devant une voyelle ou un `h` |

## Limites

- `(up,2)`, écrit **sans espace** après la virgule, n'est pas reconnu et reste dans le texte. Il faut écrire `(up, 2)`.
- Un marqueur collé à de la ponctuation, comme `(up).`, n'est pas reconnu.
- La ponctuation collée **au milieu** d'un mot n'est pas séparée : `bonjour,monde` reste inchangé.
- Un guillemet simple n'est reconnu que s'il est **isolé**, entouré d'espaces. Un `'` collé à un autre caractère (comme `'"`) n'est pas traité. Les apostrophes à l'intérieur d'un mot (`j'ai`) ne sont pas touchées, ce qui est voulu.
- `a` devient `an` seulement si le mot suivant commence par `a`, `e`, `i`, `o`, `u` ou `h` **sans accent**. `a élan` reste inchangé. La règle suit le sujet à la lettre : `a university` devient `an university`.
- Un `a` collé à de la ponctuation (`a,`) n'est pas reconnu.
- Il n'y a **pas de vérification** quand le fichier de sortie porte le même nom que le fichier d'entrée : le fichier est alors remplacé par le texte corrigé.
- Les fins de ligne Windows (`\r\n`) sont converties en fins de ligne simples (`\n`).
- Les tests se font en comparant des fichiers à la main. Il n'y a pas encore de tests unitaires Go (`go test`).

## Auteur

Antonin Mingot