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

type RuleGraph struct {
	adjList     map[int][]int
	inDegree    map[int]int
	updateOrder []int
}

func NewRuleGraph(rules [][]int) *RuleGraph {
	graph := &RuleGraph{
		adjList:  make(map[int][]int),
		inDegree: make(map[int]int),
	}
	graph.Populate(rules)
	return graph
}

func (g *RuleGraph) Populate(rules [][]int) {
	for _, rule := range rules {
		g.AddEdge(rule[0], rule[1])
	}
}

func (g *RuleGraph) AddNode(node int) {
	if _, exists := g.adjList[node]; !exists {
		g.adjList[node] = make([]int, 0)
	}
}

func (g *RuleGraph) AddEdge(from int, to int) {
	g.AddNode(from)
	g.AddNode(to)

	if !slices.Contains(g.adjList[from], to) || from == to {
		g.adjList[from] = append(g.adjList[from], to)
		g.inDegree[to]++
	}

	if _, exists := g.inDegree[from]; !exists {
		g.inDegree[from] = 0
	}
}
func (g *RuleGraph) SubGraph(nodes []int) *RuleGraph {
	subgraph := &RuleGraph{
		adjList:     make(map[int][]int),
		inDegree:    make(map[int]int),
		updateOrder: nodes,
	}

	for _, node := range nodes {
		if edges, exists := g.adjList[node]; exists {
			for _, edgeNode := range edges {
				if slices.Contains(nodes, edgeNode) {
					subgraph.AddEdge(node, edgeNode)
				}
			}
		}
	}
	return subgraph
}

func (g *RuleGraph) isValidUpdateOrder() bool {
	if g.updateOrder == nil {
		fmt.Println("No update order provided for this graph")
		return false
	}

	for i := 0; i < len(g.updateOrder)-2; i++ {
		if !slices.Contains(g.adjList[g.updateOrder[i]], g.updateOrder[i+1]) {
			return false
		}
	}

	return true
}

func (g *RuleGraph) Print() {
	for key, value := range g.adjList {
		fmt.Printf("(In: %d, Out: %d)   %d -> %v\n", g.inDegree[key], len(value), key, value)
	}

}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)
	rules := getPageOrderingRules(content)
	updates := getUpdates(content)

	graph := NewRuleGraph(rules)

	var result int
	for _, update := range updates {
		subgraph := graph.SubGraph(update)
		if subgraph.isValidUpdateOrder() {
			result += subgraph.updateOrder[(len(subgraph.updateOrder)-1)/2]
		}
	}

	fmt.Println(result)

}
