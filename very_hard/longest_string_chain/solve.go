package main

import "sort"

type Chain struct {
	NextString     string
	MaxChainLength int
}

func LongestStringChain(strings []string) []string {
	stringChains := map[string]*Chain{}
	for _, str := range strings {
		stringChains[str] = &Chain{NextString: "", MaxChainLength: 1}
	}

	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})
	sortedStrings := strings

	for _, str := range sortedStrings {
		findLongestStringChain(str, stringChains)
	}
	return buildLongestStringChain(strings, stringChains)
}

func findLongestStringChain(str string, stringChains map[string]*Chain) {
	for i := range str {
		smallerString := getSmallerString(str, i)
		if _, found := stringChains[smallerString]; !found {
			continue
		}
		tryUpdateLongestStringChain(str, smallerString, stringChains)
	}
}

func getSmallerString(str string, index int) string {
	return str[:index] + str[index+1:]
}

func tryUpdateLongestStringChain(currentString, smallerString string, stringChains map[string]*Chain) {
	smallerStringChainLength := stringChains[smallerString].MaxChainLength
	currentStringChainLength := stringChains[currentString].MaxChainLength

	if smallerStringChainLength+1 > currentStringChainLength {
		stringChains[currentString].MaxChainLength = smallerStringChainLength + 1
		stringChains[currentString].NextString = smallerString
	}
}

func buildLongestStringChain(strings []string, stringChains map[string]*Chain) []string {
	maxChainLength := 0
	chainStartingString := ""
	for _, str := range strings {
		if stringChains[str].MaxChainLength > maxChainLength {
			maxChainLength = stringChains[str].MaxChainLength
			chainStartingString = str
		}
	}

	ourLongestStringChain := []string{}
	currentString := chainStartingString
	for currentString != "" {
		ourLongestStringChain = append(ourLongestStringChain, currentString)
		currentString = stringChains[currentString].NextString
	}
	if len(ourLongestStringChain) == 1 {
		return []string{}
	}
	return ourLongestStringChain
}
