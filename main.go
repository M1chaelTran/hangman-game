package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var dictionary = []string{
	"conversation",
	"photography",
	"presentation",
	"orientation",
	"fluctuation",
	"deprivation",
	"stimulation",
	"development",
	"firefighter",
	"satisfaction",
	"requirement",
	"contraction",
	"administration",
	"circulation",
	"cooperation",
	"legislation",
	"supplementary",
	"progressive",
	"contribution",
	"instruction",
	"negotiation",
	"shareholder",
	"agriculture",
	"superintendent",
	"constituency",
	"integration",
	"astonishing",
	"spontaneous",
	"cooperative",
	"discriminate",
	"salesperson",
	"expectation",
	"participate",
	"grandmother",
	"responsible",
	"jurisdiction",
	"introduction",
	"registration",
	"continental",
	"nonremittal",
	"comfortable",
	"demonstration",
	"temperature",
	"legislature",
	"comprehensive",
	"publication",
	"intervention",
	"commemorate",
	"interference",
}

func main() {
	/*
		Print game state
		* print word you're guessing
		* print hangman state
		derive a word we have to guess
		read user input
		* validate it (eg. only letters)
		determine if it's correct
		* if correct, update guessed letters
		* if incorrect, update hangman state
		if word is guessed, print you win
		if hangman is dead, print you lose
	*/

	targetWord := getRandomWord()
	guessedLetters := getInitGuessedLetters(targetWord)
	enteredLetter := guessedLetters
	hangmanState := 0
	printGameState(targetWord, guessedLetters, hangmanState, enteredLetter)

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		input := readInput()
		letter := rune(input[0])
		if isEnterredLetter(enteredLetter, letter) {
			fmt.Println("You already entered this letter")
			continue
		}
		enteredLetter[letter] = true

		if len(input) != 1 {
			fmt.Println("Please enter a single letter")
			continue
		}

		if isCorrectGuess(input, targetWord) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
		printGameState(targetWord, guessedLetters, hangmanState, enteredLetter)

	}

	if isWordGuessed(targetWord, guessedLetters) {
		printStatusArt("win")
	} else {
		printStatusArt("lose")
	}

}

func getRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return dictionary[rand.Intn(len(dictionary))]
}

func getInitGuessedLetters(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[rune(targetWord[0])] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true
	return guessedLetters
}

func printGameState(targetWord string, guessedLetters map[rune]bool, state int, enteredLetters map[rune]bool) {
	fmt.Println()
	printHangmanState(state)
	fmt.Println()
	printGuessedLetters(targetWord, guessedLetters)
	fmt.Println()
	fmt.Println()
	printEnteredLetters(enteredLetters)
	fmt.Println()
	fmt.Println()
}

func printGuessedLetters(targetWord string, guessedLetters map[rune]bool) {
	fmt.Print("   ")
	for _, letter := range targetWord {
		if letter == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(letter)] {
			fmt.Print(strings.ToUpper(string(letter)))
		} else {
			fmt.Printf("_")
		}
		fmt.Print(" ")
	}
}

func printHangmanState(state int) {
	data, err := ioutil.ReadFile("states/" + strconv.Itoa(state))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func readInput() string {
	var input string
	fmt.Print("Enter a letter: ")
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

func isCorrectGuess(guess string, targetWord string) bool {
	return strings.Contains(targetWord, guess)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, letter := range targetWord {
		if !guessedLetters[letter] {
			return false
		}
	}
	return true
}

func isHangmanDead(state int) bool {
	return state >= 9
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, state int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanDead(state)
}

func printStatusArt(state string) {
	data, err := ioutil.ReadFile("status/" + state)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func printEnteredLetters(letter map[rune]bool) {
	println("Entered letters: " + strings.Join(getEnteredLetters(letter), ", "))
}

func getEnteredLetters(enteredLetters map[rune]bool) []string {
	letters := []string{}
	for letter, guessed := range enteredLetters {
		if guessed {
			letters = append(letters, string(letter))
		}
	}
	return letters
}

func isEnterredLetter(enteredLetters map[rune]bool, letter rune) bool {
	return enteredLetters[letter]
}
