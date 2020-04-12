package game

import (
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/markbates/pkger"
	"gopkg.in/yaml.v2"
)

var (
	languageMap   = map[string]string{
		"english": "words_en.yml",
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
func GetRandomWords(lobby *Lobby) (string, string) {
	rand.Seed(time.Now().Unix())
	wordsAlreadyUsed := lobby.alreadyUsedWords

	//We attempt to find a random word for a hundred times, afterwards we just use any.
	//randomnessAttempts := 0
	//var word string
	category := lobby.Categories[rand.Int()%len(lobby.Categories)]
	wordList := lobby.Words[category]

	randomnessAttempts := 0
	word := ""

	OuterLoop:
	for {
		word = wordList[rand.Int()%len(wordList)]
		for _, usedWord := range wordsAlreadyUsed {
			if usedWord == word {
				if randomnessAttempts == 100 {
					break OuterLoop
				}

				randomnessAttempts++
				continue OuterLoop
			}
		}
		break
	}

	return word, category
}
