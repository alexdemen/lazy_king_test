package main

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Name  string
	Nodes []*Node
}

func NewNode(name string) *Node {
	return &Node{Name: name, Nodes: make([]*Node, 0)}
}

func main() {
	pollResults := []string{
		"служанка Аня",
		"управляющий Семен Семеныч: крестьянин Федя, доярка Нюра",
		"дворянин Кузькин: управляющий Семен Семеныч, жена Кузькина, экономка Лидия Федоровна",
		"экономка Лидия Федоровна: дворник Гена, служанка Аня",
		"доярка Нюра",
		"кот Василий: человеческая особь Катя",
		"дворник Гена: посыльный Тошка",
		"киллер Гена",
		"зажиточный холоп: крестьянка Таня",
		"секретарь короля: зажиточный холоп, шпион Т",
		"шпион Т: кучер Д",
		"посыльный Тошка: кот Василий",
		"аристократ Клаус",
		"просветленный Антон",
	}

	king := NewNode("Король")
	printReport(king, pollResults)
}

func printReport(root *Node, pollResults []string) {
	buildTree(root, pollResults)
	printTree(root)
}

func printTree(root *Node) {
	printNode(root, 0)
}

func printNode(node *Node, level int) {
	fmt.Println(strings.Repeat("\t", level) + node.Name)
	sort.SliceStable(node.Nodes, func(i, j int) bool {
		return node.Nodes[i].Name < node.Nodes[j].Name
	})
	for _, child := range node.Nodes {
		printNode(child, level+1)
	}
}

func buildTree(root *Node, results []string) {
	for _, val := range results {
		nodeName, slaves := parseString(val)

		existNode := findNode(root, nodeName, false)
		parent := root
		if existNode == nil {
			parent = NewNode(nodeName)
			addNode(root, parent)
		} else {
			parent = existNode
		}

		for _, slave := range slaves {
			slave = strings.Trim(slave, " ")
			existNode = findNode(root, slave, true)
			if existNode == nil {
				addNode(parent, NewNode(slave))
			} else {
				addNode(parent, existNode)
			}
		}
	}
}

func parseString(str string) (nodeName string, slaves []string) {
	splitted := strings.Split(str, ":")
	nodeName = splitted[0]
	if len(splitted) > 1 {
		slaves = strings.Split(splitted[1], ",")
		for _, val := range slaves {
			val = strings.Trim(val, " ")
		}
	}
	return
}

func addNode(parent *Node, slave *Node) {
	parent.Nodes = append(parent.Nodes, slave)
}

func findNode(parent *Node, childNodeName string, removing bool) *Node {
	var resultNode *Node
	for idx, val := range parent.Nodes {
		if val.Name == childNodeName {
			resultNode = val
		} else {
			resultNode = findNode(val, childNodeName, removing)
		}

		if resultNode != nil {
			if removing && parent.Nodes[idx].Name == childNodeName {
				parent.Nodes = append(parent.Nodes[:idx], parent.Nodes[idx+1:]...)
			}
			return resultNode
		}
	}

	return nil
}
