package game

import (
	"io/ioutil"
	"log"
	"math/rand"
	//"strings"
	"time"
	"fmt"
	//"log"

	"github.com/markbates/pkger"
	"gopkg.in/yaml.v2"
	//"github.com/smallfish/simpleyaml"
)

var (
	languageMap   = map[string]string{
		"english": "words_en.yml",
		"italian": "words_it",
		"german":  "words_de",
	}
)

func readWordList(chosenLanguage string) ([]string, map[string][]string, error) {
	langFileName := languageMap[chosenLanguage]
	wordListFile, pkgerError := pkger.Open("/resources/words/" + langFileName)
	if pkgerError != nil {
		panic(pkgerError)
	}
	defer wordListFile.Close()

	data, err := ioutil.ReadAll(wordListFile)
	if err != nil {
		return nil, nil, err
	}

	// Parse word list to a dict
	m := make(map[string][]string)
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	// Compute categories
	categories := make([]string, len(m))
	i := 0
	for k := range m {
		categories[i] = k
		i++
	}

	return categories, m, nil
}

// GetRandomWords gets 3 random words for the passed Lobby. The words will be
// chosen from the custom words and the default dictionary, depending on the
// settings specified by the Lobby-Owner.
func GetRandomWords(lobby *Lobby) []string {
	rand.Seed(time.Now().Unix())
	wordsNotToPick := lobby.alreadyUsedWords
	word, category := getUnusedRandomWord(lobby, wordsNotToPick)

	return []string{ word, category }
}

func getUnusedRandomWord(lobby *Lobby, wordsAlreadyUsed []string) (string, string) {
	//We attempt to find a random word for a hundred times, afterwards we just use any.
	//randomnessAttempts := 0
	//var word string
//OUTER_LOOP:
//	for {
//		word = lobby.Words[rand.Int()%len(lobby.Words)]
//		for _, usedWord := range wordsAlreadyUsed {
//			if usedWord == word {
//				if randomnessAttempts == 100 {
//					break OUTER_LOOP
//				}
//
//				randomnessAttempts++
//				continue OUTER_LOOP
//			}
//		}
//		break
//	}

	return "word", "category"
}
