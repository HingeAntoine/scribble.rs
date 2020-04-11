package game

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/markbates/pkger"
)

var (
	wordListCache = make(map[string][]string)
	languageMap   = map[string]string{
		"english": "words_en",
		"italian": "words_it",
		"german":  "words_de",
	}
)

func readWordList(chosenLanguage string) ([]string, error) {
	langFileName := languageMap[chosenLanguage]
	list, available := wordListCache[langFileName]
	if available {
		return list, nil
	}

	wordListFile, pkgerError := pkger.Open("/resources/words/" + langFileName)
	if pkgerError != nil {
		panic(pkgerError)
	}
	defer wordListFile.Close()

	data, err := ioutil.ReadAll(wordListFile)
	if err != nil {
		return nil, err
	}

	tempWords := strings.Split(string(data), "\n")
	var words []string
	for _, word := range tempWords {
		word = strings.TrimSpace(word)
		if strings.HasSuffix(word, "#i") {
			continue
		}

		lastIndexNumberSign := strings.LastIndex(word, "#")
		if lastIndexNumberSign == -1 {
			words = append(words, word)
		} else {
			words = append(words, word[:lastIndexNumberSign])
		}
	}

	wordListCache[langFileName] = words

	return words, nil
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
	randomnessAttempts := 0
	var word string
OUTER_LOOP:
	for {
		word = lobby.Words[rand.Int()%len(lobby.Words)]
		for _, usedWord := range wordsAlreadyUsed {
			if usedWord == word {
				if randomnessAttempts == 100 {
					break OUTER_LOOP
				}

				randomnessAttempts++
				continue OUTER_LOOP
			}
		}
		break
	}

	return word, "category"
}
