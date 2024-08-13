package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {

	// fmt.Println("hello world")

	// fmt.Println(arraySign([]int{2, 1}))                    // 1
	// fmt.Println(arraySign([]int{-2, 1}))                   // -1
	// fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // 1

	// fmt.Println(isAnagram("anak", "kana"))       // true
	// fmt.Println(isAnagram("anak", "mana"))       // false
	// fmt.Println(isAnagram("anagram", "managra")) // true

	// fmt.Println(string(findTheDifference("abcd", "abcde"))) // 'e'
	// fmt.Println(string(findTheDifference("abcd", "abced"))) // 'e'
	// fmt.Println(string(findTheDifference("", "y")))         // 'y'

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	//tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if len(nums) >= 1 && len(nums) <= 1000 {
		negativeCount := 0
		checkZero := false

		// check ada gak zero
		for _, num := range nums {
			if num >= -100 && num <= 100 {
				if num == 0 {
					checkZero = true
					break
				} else if num < 0 {
					negativeCount++
				}

			} else {
				panic("error arraySign : The length of nums is within the valid range.")
			}
		}

		// jika ada zero return 0
		if checkZero {
			return 0
		}

		// kalau habis dibagi dua balikin 1 kalau gk -1
		if negativeCount%2 == 0 {
			return 1
		} else {
			return -1
		}

	} else {
		panic("error arraySign : The length of nums is within the valid range.")
	}
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	upperLimit := 5 * int(math.Pow(10.0, 4.0))
	if len(s) >= 1 && len(t) <= upperLimit {
		s := strings.ToLower(s)
		t := strings.ToLower(t)

		// check if length strings different or not
		if len(s) != len(t) {
			return false
		}

		// create make for both of strings
		countS := map[rune]int{}
		countT := map[rune]int{}

		//populate map
		for _, v := range s {
			countS[v]++
		}

		for _, v := range t {
			countT[v]++
		}

		//check if s contains any extra character is not in t
		for idx, count := range countS {
			if countT[idx] != count {
				return false
			}
		}

		//check if t contains any extra character is not in s
		for idx, count := range countT {
			if countS[idx] != count {
				return false
			}
		}

		return true

	} else {
		panic("erroy anagram : Strings are within the valid length range.")
	}
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if len(s) >= 0 && len(s) <= 1000 {
		s := strings.ToLower(s)
		t := strings.ToLower(t)
		// loop map  characters in s
		countMap := map[rune]int{}
		for _, v := range s {
			countMap[v]++
		}

		// loop t, if character is 0, added character
		for _, v := range t {
			countMap[v]--
			if countMap[v] < 0 {
				return byte(v)
			}
		}

		return 0

	} else {
		panic("erroy findTheDifference : Strings are within the valid length range.")
	}
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	bottomLimit := -int(math.Pow(-10.0, 6.0))
	upperLimit := int(math.Pow(10.0, 6.0))
	if len(arr) >= 2 && len(arr) <= 1000 {
		// Sort  array
		sort.Ints(arr)

		// Calculate the common difference
		countDiff := arr[1] - arr[0]

		// Check if the difference between all consecutive elements is the same
		for i := 2; i < len(arr); i++ {
			if arr[i] >= bottomLimit && arr[i] <= upperLimit {
				if arr[i]-arr[i-1] != countDiff {
					return false
				}
			} else {
				panic("erroy canMakeArithmeticProgression : Value array are within the valid length range.")
			}
		}

		return true

	} else {
		panic("erroy canMakeArithmeticProgression : Array are within the valid length range.")
	}
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	return nil
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	return nil
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
