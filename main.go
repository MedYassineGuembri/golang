package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
)

func main() {
	dict := dictionary.New("dictionary.json")
	reader := bufio.NewReader(os.Stdin)

	actionAdd(dict, "pomme", "Un fruit comestible d'un pommier.", reader)
	actionAdd(dict, "ordinateur", "Une machine électronique qui traite les données.", reader)

	actionDefine(dict, "pomme", reader)



	actionList(dict)
}

func actionAdd(d *dictionary.Dictionary, word, definition string, reader *bufio.Reader) {
	d.Add(word, definition)
	
}

func actionDefine(d *dictionary.Dictionary, word string, reader *bufio.Reader) {
	entry, found, err := d.Get(word)
	if err != nil {
		fmt.Printf("Erreur lors de la recherche du mot '%s': %v\n", word, err)
		return
	}
	if !found {
		fmt.Printf("Le mot '%s' n'a pas été trouvé dans le dictionnaire.\n", word)
	} else {
		fmt.Printf("Définition de '%s': %s\n", word, entry.Definition)
	}
}

func actionRemove(d *dictionary.Dictionary, word string, reader *bufio.Reader) {
	d.Remove(word)
}

func actionList(d *dictionary.Dictionary) {
words, err := d.List()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération de la liste: %v\n", err)
		return
	}
	for _, word := range words {
		entry, found, _ := d.Get(word)
		if found {
			fmt.Printf("%v: %v\n", word, entry.Definition)
		}
	}
}