package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func getPageOrderingRules(content string) [][]int {
	pageOrderingRulePattern := `(\d+)\|(\d+)`
	pageOrderingRuleRe := regexp.MustCompile(pageOrderingRulePattern)
	pageOrderingRuleMatches := pageOrderingRuleRe.FindAllStringSubmatch(content, -1)

	var pageOrderingRules [][]int
	for _, match := range pageOrderingRuleMatches {
		var pageOrderingRule []int
		for _, number := range match[1:] {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			pageOrderingRule = append(pageOrderingRule, num)
		}
		pageOrderingRules = append(pageOrderingRules, pageOrderingRule)
	}

	return pageOrderingRules
}

func getUpdates(content string) [][]int {
	numberPattern := `\d+`
	updatePattern := `\d+(?:,\d+)+`
	numberRe := regexp.MustCompile(numberPattern)
	updateRe := regexp.MustCompile(updatePattern)
	updateMatches := updateRe.FindAllStringSubmatch(content, -1)

	var updates [][]int
	for _, match := range updateMatches {
		numbers := numberRe.FindAllString(match[0], -1)
		var update []int
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return updates
}

func generateRuleMap(rules [][]int) map[int][]int {
	var ruleMap = make(map[int][]int)
	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}
	return ruleMap
}

func validateElement(before []int, number int, ruleMap map[int][]int) bool {
	mustBeAfter := ruleMap[number]

	for _, el := range before {
		if slices.Contains(mustBeAfter, el) {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)

	rules := getPageOrderingRules(content)
	updates := getUpdates(content)
	ruleMap := generateRuleMap(rules)

	var result int

Outerloop:
	for _, update := range updates {
		for j, number := range update {
			if !validateElement(update[:j], number, ruleMap) {
				continue Outerloop
			}
		}
		result += update[(len(update)-1)/2]
	}

	fmt.Println(result)

}
