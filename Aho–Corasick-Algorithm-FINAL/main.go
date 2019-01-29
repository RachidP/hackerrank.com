package main

import (
	"container/heap"
	"fmt"
)

//global variables
var nextNodeID int

var priorityCount int

type node struct {
	//child of the current node
	child map[rune]*node

	isEndOfWord bool // isLeaf is true if the node represents

	fail *node //fail node

	id int //id node (good for view and debug)

	output []string // if the current node determine a intere world in the dictionary
}

// return new trie node (initialized to nil)
func getNode() *node {
	pNode := node{id: nextNodeID}
	pNode.child = make(map[rune]*node)
	pNode.output = make([]string, 0)
	nextNodeID++
	return &pNode
}

// If not present, inserts key into trie
// If the key is prefix of trie node, just marks leaf node
func (n *node) insert(key string) {

	if key == "d" {
		println("Ciao")
	}

	for _, v := range key {

		// If not present, inserts key into trie
		if _, exist := n.child[v]; !exist {
			n.child[v] = getNode()
		}

		n = n.child[v]
	}

	n.output = append(n.output, key)

	n.isEndOfWord = true

}

func (n *node) fialsOnRoot() PriorityQueue {

	priorityCount = 0
	pq := make(PriorityQueue, 1)
	var item *Item

	for _, v := range n.child {

		//increment priority queue
		priorityCount++

		item = &Item{
			value:    v,
			priority: priorityCount,
		}

		//if we don't have inizialized the quee
		if priorityCount == 1 {
			pq[0] = item
			heap.Init(&pq)
		} else {
			// push  the new node in the queue.
			heap.Push(&pq, item)
		}

		//the child now field to the root
		v.fail = n

	}

	return pq
}

// BFS to build failure
func (n *node) bfsBuildFailure(pq PriorityQueue) {

	for pq.Len() > 0 {
		//pop out a node
		item := heap.Pop(&pq).(*Item)

		cur := item.value //the node just poped out get our current node

		// for each char appared in this node
		for key := range cur.child {

			//get wich current node is failing
			failNode := cur.fail

			//keep failing till add ch to the prefix of pattern
			failNode = failNode.move(key)

			//set the current node fail to the failNode
			cur.child[key].fail = failNode

			//add the output in the failNode to the output of the current node
			for _, v := range failNode.output {
				if v != "" {
					cur.child[key].output = append(cur.child[key].output, v)
				}

			}

			// add the node to the quee
			priorityCount++
			item = &Item{
				value:    cur.child[key],
				priority: priorityCount,
			}
			heap.Push(&pq, item)
		}

	}

}

//looking for the node that failing with the character char
func (n *node) move(index rune) *node {

	for _, exist := n.child[index]; !exist; _, exist = n.child[index] {
		if n.fail == nil {
			return n
		}

		n = n.fail

	}

	n = n.child[index]
	return n
}

// Returns true if key presents in trie, else false
func (n *node) search(key string) bool {

	var exist bool
	child := n
	for _, v := range key {

		// If not present, inserts key into trie
		if child, exist = n.child[v]; !exist {
			return false
		}

		n = child

	}
	if (n != nil) && n.isEndOfWord {
		//fmt.Println(n.charls)
		return true
	}
	return false
}

func (n *node) printTree() {

	if n.id == 0 {
		fmt.Println("node:", n.id, "output :", n.output)

	} else {

		fmt.Println("node:", n.id, "output :", n.output, "prefix :", n.fail.id)
	}

	//fmt.Println("node:", n.id, " childs: ", n.charls, "output :", n.output)

	for _, v := range n.child {
		//index := int(i - 'a')
		v.printTree()
	}

}

func (n *node) matchingPatterns(text string) {

	m := make(map[string][]int)

	for i, v := range text {
		n = n.move(v)
		for _, output := range n.output {
			if output != "" {
				tmp := i - len(output) + 1
				m[output] = append(m[output], tmp)
			}

		}
	}

	fmt.Println(m)
}

func main() {
	//patterns := []string{"he", "she", "his", "hers"}
	patterns := []string{"ab", "bc", "bab", "d", "abcde"}
	//patterns := []string{"a", "ab", "bc", "aab", "aac", "bd"}
	//output := []string{"Not present in trie", "Present in trie"}
	//tree := &tree{}
	root := getNode()

	for _, v := range patterns {
		root.insert(v)

	}
	fmt.Println("patterns salved in the tree")

	//root.search("he")

	pq := root.fialsOnRoot()
	root.bfsBuildFailure(pq)
	root.printTree()
	//text := "xbabcdex"
	text := "abckjfkdffkajfknvnfbf"
	root.matchingPatterns(text)

}

/*****************************************************************************/
