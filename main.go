package main

import (
	"fmt"
	"os"
	"regexp"
	"reloaded/piscine"
)

// Check whether element in slice is a modification
func modCheck(s string) bool {
	if s == "(hex)" || s == "(bin)" || s == "(up)" || s == "(low)" || s == "(cap)" {
		return true
	}

	match, _ := regexp.MatchString("(up, ([0-9]+))", s)
	if match {
		return true
	}

	match, _ = regexp.MatchString("(low, ([0-9]+))", s)
	if match {
		return true
	}

	match, _ = regexp.MatchString("(cap, ([0-9]+))", s)

	return match
}

// String is initially split by whitespace - this fixes all instances of modifications that get caught by that split
func modCorrection(ss []string) []string {
	var corrected []string
	u, _ := regexp.Compile("(up, ([0-9]+))")
	l, _ := regexp.Compile("(low, ([0-9]+))")
	c, _ := regexp.Compile("(cap, ([0-9]+))")

	for i := 0; i < len(ss); i++ {
		if i < len(ss)-1 {
			if u.MatchString(ss[i]+" "+ss[i+1]) || l.MatchString(ss[i]+" "+ss[i+1]) || c.MatchString(ss[i]+" "+ss[i+1]) {
				corrected = append(corrected, ss[i]+" "+ss[i+1])
				i++
			} else {
				corrected = append(corrected, ss[i])
			}
		}

		if i == len(ss)-1 {
			if u.MatchString(ss[i-1]+" "+ss[i]) || l.MatchString(ss[i-1]+" "+ss[i]) || c.MatchString(ss[i-1]+" "+ss[i]) {
				continue
			} else {
				corrected = append(corrected, ss[i])
			}
		}
	}

	return corrected
}

// Convert number from hexidecimal to decimal
func hexConvert(n string) string {
	return piscine.ConvertBase(n, "0123456789ABCDEF", "0123456789")
}

// Convert number from binary to decimal
func binConvert(n string) string {
	return piscine.ConvertBase(n, "01", "0123456789")
}

// Seperate function to mod numbers (hex, bin) as they don't play nice with punctuation!
func modNumbers(words []string) []string {
	for i := 0; i < len(words); i++ {
		if modCheck(words[i]) {
			if words[i] == "(hex)" {
				words[i-1] = hexConvert(piscine.ToUpper(words[i-1]))
				words = piscine.DeleteElement(words, i)
				i--
			} else if words[i] == "(bin)" {
				words[i-1] = binConvert(words[i-1])
				words = piscine.DeleteElement(words, i)
				i--
			}
		}
	}

	return words
}

// Applying text modifications
func modChanges(words []string) []string {
	for i := 0; i < len(words); i++ {
		if modCheck(words[i]) {
			if len(words[i]) == 5 || words[i] == "(up)" {
				if i > 0 {
					switch words[i] {
					case "(up)":
						words[i-1] = piscine.ToUpper(words[i-1])
					case "(low)":
						words[i-1] = piscine.ToLower(words[i-1])
					case "(cap)":
						words[i-1] = piscine.Capitalize(words[i-1])
					}
				}

				words = piscine.DeleteElement(words, i)
				i--
			} else {
				num := piscine.TrimAtoi(words[i])

				r, _ := regexp.Compile("(up, ([0-9]+))")
				if r.MatchString(words[i]) {
					for j := 1; j < num+1; j++ {
						if i-j >= 0 {
							words[i-j] = piscine.ToUpper(words[i-j])
						}
					}
					words = piscine.DeleteElement(words, i)
					i--
					continue
				}

				r, _ = regexp.Compile("(low, ([0-9]+))")
				if r.MatchString(words[i]) {
					for j := 1; j < num+1; j++ {
						if i-j >= 0 {
							words[i-j] = piscine.ToLower(words[i-j])
						}
					}
					words = piscine.DeleteElement(words, i)
					i--
					continue
				}

				r, _ = regexp.Compile("(cap, ([0-9]+))")
				if r.MatchString(words[i]) {
					for j := 1; j < num+1; j++ {
						if i-j >= 0 {
							words[i-j] = piscine.Capitalize(words[i-j])
						}
					}
					words = piscine.DeleteElement(words, i)
					i--
					continue
				}
			}
		} else {
			continue
		}
	}

	return words
}

// Check if element of slice only consists of punctuation
func puncOnly(s string) bool {
	for _, v := range s {
		if string(v) != "." && string(v) != "," && string(v) != "!" && string(v) != "?" && string(v) != ":" && string(v) != ";" {
			return false
		}
	}

	return true
}

// If punctuation is at the start of a word, move it to the end of the previous word
func puncTails(words []string) []string {
	for i := 1; i < len(words); i++ {
		if !puncOnly(words[i]) {
			if puncOnly(string(words[i][0])) {
				words[i-1] = words[i-1] + string(words[i][0])
				words[i] = words[i][1:]
			}
		}
	}

	return words
}

// Shift standalone punctuations to the first previous 'word' found that isn't a modification
func puncFix(words []string) []string {
	words = puncTails(words)

	for i := 0; i < len(words); i++ {
		noMod := -1
		if puncOnly(words[i]) {
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					if !modCheck(words[j]) {
						noMod = j
						break
					}
				}
				if noMod >= 0 {
					words[noMod] = words[noMod] + words[i]
					words = piscine.DeleteElement(words, i)
				}
			}
		}
	}

	return words
}

// Separate function to catch apostrophes as could be appended to previous or next word
func apostropheCheck(words []string) []string {
	// Bool flip to determine whether apostrophe attaches to previous or next word
	flip := true
	for i := 0; i < len(words); i++ {
		noMod := -1
		if words[i] == "'" {
			if flip {
				if i != len(words)-1 {
					for j := i + 1; j < len(words); j++ {
						if !modCheck(words[j]) {
							noMod = j
							break
						}
					}

					if noMod < 0 {
						continue
					}

					words[noMod] = words[i] + words[noMod]
					words = piscine.DeleteElement(words, i)
				}

				flip = !flip
			} else {
				for j := i - 1; j >= 0; j-- {
					if !modCheck(words[j]) {
						noMod = j
						break
					}
				}
				if noMod >= 0 {
					words[noMod] = words[noMod] + words[i]
					words = piscine.DeleteElement(words, i)
				}

				flip = !flip
			}
		}
	}

	return words
}

// Changes 'a' to 'an' if the following word begins with a vowel or 'h'
func vowelGrammar(words []string) {
	for i, w := range words {
		if i < len(words)-1 {
			if w == "a" || w == "A" {
				next := piscine.ToLower(string(words[i+1][0]))
				if next == "a" || next == "e" || next == "i" || next == "o" || next == "u" || next == "h" {
					words[i] = words[i] + "n"
				}
			}
		}
	}
}

func main() {
	// Checking for correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of arguments. (Should be 2)")
		return
	}

	var finalString string

	// Read first argument
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create second argument
	out, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	words := piscine.SplitWhiteSpaces(string(f))

	// If slice is empty or only contained whitespace
	if len(words) < 1 {
		out.Close()
		return
	}

	words = modCorrection(words)
	words = modNumbers(words)
	words = puncFix(words)
	words = apostropheCheck(words)
	vowelGrammar(words)
	words = modChanges(words)

	// Creating a string from the final slice to be written to new file
	for i := 0; i < len(words); i++ {
		finalString = finalString + words[i] + " "
	}

	finalString = finalString[:len(finalString)-1]

	out.WriteString(finalString)
	out.Close()
}
