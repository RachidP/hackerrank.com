package main

import (
	"bufio"

	"fmt"
	"math"
	"os"

	"strconv"
	"strings"
)

type node struct {
	idChild []rune   // id child of the node
	child   []*node  //childs of the current node
	fail    *node    //fail node
	output  []string // save the pattern in the node

}

// return new trie node (initialized to nil)
func getNode() *node {

	return &node{idChild: make([]rune, 0), child: make([]*node, 0), output: make([]string, 0)}
}

// If not present, inserts key into trie
// If the key is prefix of trie node, just marks leaf node
func (n *node) insert(key string) {

	var tmpNode *node

	for _, letter := range key {
		//the key is in the tree
		if childIndex := n.getNodeID(letter); childIndex >= 0 {
			tmpNode = n.child[childIndex]
		} else { // If not present, inserts key into trie
			//make new node
			tmpNode = getNode()
			n.child = append(n.child, tmpNode)

			//add id node to idNode
			n.idChild = append(n.idChild, letter)
		}

		n = tmpNode

	}
	//add the string to the output
	n.output = appendUnique(n.output, key)
}

// get the Indice of the letter in the node
// -1 if don't found
func (n *node) getNodeID(letter rune) int {

	for i, ch := range n.idChild {

		if ch == letter {
			return i
		}

	}
	return -1
}

//append a string to slice only if doesn't already their
func appendUnique(a []string, x string) []string {
	for _, y := range a {
		if x == y {
			return a
		}
	}
	return append(a, x)
}

//Make all the children of the Root fail to the root
func (n *node) failsOnRootVec() {

	sq := make([]*node, 0, 5)
	for _, v := range n.child {
		sq = append(sq, v)
		//the child now field to the root
		v.fail = n
	}

	n.bfsBuildFailureVec(sq)

}

// BFS to build failure
func (n *node) bfsBuildFailureVec(pq []*node) {
	numberNode := len(pq)

	for numberNode > 0 {
		//get a node from the queue
		cur := pq[0]
		pq = pq[1:]

		numberNode--

		for id, ch := range cur.child {

			//get wich current node is failing
			failNode := cur.fail

			//keep failing till add ch to the prefix of pattern
			failNode = failNode.move(cur.idChild[id])

			//set the childe node fail to the failNode
			ch.fail = failNode

			//add the output in the failNode to the output of the current node
			for _, v := range failNode.output {

				ch.output = appendUnique(ch.output, v)
			}

			// add the node to the quee
			//pq = append(pq, cur.child[key])
			pq = append(pq, ch)

			numberNode++

		}

	}
	//delete the slice
	pq = nil

}

//looking for the node that failing with the character char
func (n *node) move(letter rune) *node {

	index := n.getNodeID(letter)

	for ; index < 0; index = n.getNodeID(letter) {
		if n.fail == nil {
			return n
		}
		n = n.fail

	}

	return n.child[index]

}

//calculate the tot health for each
func (n *node) matchingPatterns(text string, genes map[string][]int, health []uint64, start, end int) uint64 {

	var totHealth uint64

	for _, v := range text {

		n = n.move(v)

		//range over the output
		for _, k := range n.output {

			// if the character is present in the pattern
			if indGenes, exist := genes[k]; exist {

				// get all the genes ID that have the charact
				for _, genesID := range indGenes {

					if genesID >= start && genesID <= end {

						totHealth += health[genesID]

					}
				}

			}

			//	tmp := i - len(output) + 1
			//	m[output] = append(m[output], tmp)

		}
	}

	return totHealth
}

func main() {

	var t int
	var in string
	var healthiest, unhealthiest uint64
	unhealthiest = math.MaxUint64
	genes := make(map[string][]int)

	//Read from file
	f, err := os.Open("test33.txt")
	check(err)
	reader := bufio.NewReader(f)

	//use this only in the site
	//reader := bufio.NewReader(os.Stdin)

	//skip first line
	_, _ = reader.ReadString('\n')

	//get patterns
	in, _ = reader.ReadString('\n')
	patterns := strings.Fields(in)

	//get the values of  health of geneas
	in, _ = reader.ReadString('\n')
	healthS := strings.Fields(in)          //health as a string
	health := make([]uint64, len(healthS)) //health as a integer

	//buil the root of the tree and the tree
	root := getNode()

	// range over all patterns
	for i, v := range patterns {

		//build a map with. key: patterns(genes), value: a slice of indices of that genes
		genes[v] = append(genes[v], i)

		//convert the health value from string to int
		genHealth, _ := strconv.ParseUint(healthS[i], 10, 64)
		health[i] = genHealth

		//insert the pattern in the tree
		root.insert(v)
	}
	//build the failor tree
	root.failsOnRootVec()

	in, _ = reader.ReadString('\n')

	in = strings.TrimSuffix(in, "\r\n")
	//in = strings.TrimSuffix(in, "\r\n")// in the site take away \r
	t, _ = strconv.Atoi(in)

	for t > 0 {
		in, _ = reader.ReadString('\n')
		in = strings.TrimSuffix(in, "\r\n")
		strands := strings.Fields(in)

		start, _ := strconv.Atoi(strands[0])

		end, _ := strconv.Atoi(strands[1])

		text := strands[2]
		//calculate the health value
		totHealth := root.matchingPatterns(text, genes, health, start, end)

		if totHealth > healthiest {
			healthiest = totHealth
		}
		if totHealth < unhealthiest {
			unhealthiest = totHealth
		}

		t--
	}

	fmt.Println(unhealthiest, healthiest)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
