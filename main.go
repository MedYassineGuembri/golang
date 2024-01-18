package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
)

func main() {
	dict := dictionary.New()
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
	entry, exists := d.Get(word)
	if !exists {
		fmt.Printf("Le mot '%s' n'existe pas.\n", word)
		return
	}
	fmt.Printf("La définition de '%s' est: %s\n", word, entry.String())
}

func actionRemove(d *dictionary.Dictionary, word string, reader *bufio.Reader) {
	d.Remove(word)
}

func actionList(d *dictionary.Dictionary) {
words := d.List()
	if len(words) == 0 {
		fmt.Println("Le dictionnaire est vide.")
		return
	}

	for _, word := range words {
		entry, exists := d.Get(word)
		if exists {
			fmt.Printf("%v: %v\n", word, entry.String())
		}
	}
}