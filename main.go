package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    dict := dictionary.New("dictionary.json")
 


    router := mux.NewRouter()

   
    router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
       
        word := r.FormValue("word")
        definition := r.FormValue("definition")
        err := dict.Add(word, definition)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Entry added: %s - %s\n", word, definition)
    }).Methods("POST")

    router.HandleFunc("/define/{word}", func(w http.ResponseWriter, r *http.Request) {
        
        vars := mux.Vars(r)
        word := vars["word"]
        entry, found, err := dict.Get(word)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        if !found {
            http.NotFound(w, r)
            return
        }
        fmt.Fprintf(w, "Définition de '%s': %s\n", word, entry.Definition)
    }).Methods("GET")

    router.HandleFunc("/remove/{word}", func(w http.ResponseWriter, r *http.Request) {
        
        vars := mux.Vars(r)
        word := vars["word"]
        err := dict.Remove(word)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Entry removed: %s\n", word)
    }).Methods("DELETE")

    
    http.Handle("/", router)
    fmt.Println("Server is running on :8080...")
    http.ListenAndServe(":8080", nil)
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