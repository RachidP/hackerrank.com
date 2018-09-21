package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"runtime/trace"

	"strconv"
	"strings"
)

//global variables
var nextNodeID int

var priorityCount int

type node struct {
	//child of the current node
	idChild     []rune
	child       []*node
	isEndOfWord bool // isLeaf is true if the node represents

	fail *node //fail node

	id     int //id node (good for view and debug)
	output []string
}

// return new trie node (initialized to nil)
func getNode() *node {
	pNode := node{id: nextNodeID, idChild: make([]rune, 0), child: make([]*node, 0), output: make([]string, 0)}
	//	pNode.child = make(map[rune]*node)
	//	pNode.output = make(map[string]struct{})
	//pNode.output = make([]string, 0)
	nextNodeID++
	return &pNode

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

	//n.output = append(n.output, key)
	n.output = appendUnique(n.output, key)

	n.isEndOfWord = true

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

func appendUnique(a []string, x string) []string {
	for _, y := range a {
		if x == y {
			return a
		}
	}
	return append(a, x)
}

func (n *node) failsOnRoot() {

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

	n.bfsBuildFailure(pq)
}

// BFS to build failure
func (n *node) bfsBuildFailure(pq PriorityQueue) {

	for pq.Len() > 0 {
		//pop out a node
		item := heap.Pop(&pq).(*Item)

		cur := item.value //the node just poped out get our current node

		// for each children of the current node
		for id, ch := range cur.child {

			//get wich current node is failing
			failNode := cur.fail

			//keep failing till add ch to the prefix of pattern
			failNode = failNode.move(cur.idChild[id])

			//set the current node fail to the failNode
			//cur.child[id].fail = failNode
			ch.fail = failNode

			//add the output in the failNode to the output of the current node
			for _, v := range failNode.output {

				//cur.child[key].output = append(cur.child[key].output, v)
				ch.output = appendUnique(ch.output, v)

			}

			// add the node to the quee
			priorityCount++
			item = &Item{
				value:    ch,
				priority: priorityCount,
			}
			heap.Push(&pq, item)
		}

	}

}

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

// Returns true if key presents in trie, else false
func (n *node) search(key string) bool {

	var indice int
	for _, v := range key {

		if indice = n.getNodeID(v); indice < 0 {
			return false
		}

		n = n.child[indice]

	}
	if (n != nil) && n.isEndOfWord {
		//fmt.Println(n.charls)
		return true
	}
	return false
}

func (n *node) printTree() {

	if n.id == 0 {
		//fmt.Println("node:", n.id, "output :", n.output, "prefix :", n.fail)
		fmt.Println("node:", n.id, "output :", n.output)

	} else {

		//		fmt.Println("node:", n.id, "output :", n.output, "prefix :", n.fail.id)
		fmt.Println("node:", n.id, "output :", n.output)

	}
	fmt.Printf("child: ")
	for i, v := range n.idChild {
		fmt.Printf("%v:%v  ", n.child[i].id, string(v))
	}
	fmt.Printf("\n")
	for _, v := range n.child {
		v.printTree()
	}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var cputrace = flag.String("cputrace", "", "write cpu trace to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *cputrace != "" {
		f, err := os.Create(*cputrace)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)
		defer trace.Stop()

	}

	var t int
	var in string
	var healthiest, unhealthiest uint64
	healthiest = 0
	unhealthiest = math.MaxUint64
	genes := make(map[string][]int)

	f, err := os.Open("test33.txt")
	check(err)
	reader := bufio.NewReader(f)
	//reader := bufio.NewReader(os.Stdin)

	//skip first line
	_, _ = reader.ReadString('\n')
	//get patterns
	in, _ = reader.ReadString('\n')
	//in = strings.TrimSuffix(in, "\r\n")
	patterns := strings.Fields(in)
	//get the values of  health of geneas
	in, _ = reader.ReadString('\n')
	healthS := strings.Fields(in)          //health as a string
	health := make([]uint64, len(healthS)) //health as a integer

	//buil the root of the tree and the tree
	root := getNode()
	for i, v := range patterns {
		//build a map with. key: patterns(genes), value: a slice of indices of that genes
		genes[v] = append(genes[v], i)

		//convert the health value from string to int
		genHealth, _ := strconv.ParseUint(healthS[i], 10, 64)
		health[i] = genHealth

		//insert the pattern in the tree
		root.insert(v)

	}

	root.failsOnRootVec()

	//root.failsOnRoot()

	//root.printTree()
	//get number of test
	in, _ = reader.ReadString('\n')

	in = strings.TrimSuffix(in, "\r\n")
	//in = strings.TrimSuffix(in, "\r\n")// in the site take away \r

	t, _ = strconv.Atoi(in)
	//fmt.Println("number of test: ", t)

	for t > 0 {
		in, _ = reader.ReadString('\n')
		in = strings.TrimSuffix(in, "\r\n")
		strands := strings.Fields(in)
		//fmt.Printf("strands : \n %v \n", strands)
		start, _ := strconv.Atoi(strands[0])
		//fmt.Printf("start : \n %v \n", start)
		end, _ := strconv.Atoi(strands[1])
		//fmt.Printf("end : \n %v \n", end)
		text := strands[2]
		// fmt.Println("start", start)
		//fmt.Println("strands",strands)
		//fmt.Printf("genes : \n %v \n\n\n", genes)

		totHealth := root.matchingPatterns(text, genes, health, start, end)
		//fmt.Println("total health =", totHealth)
		if totHealth > healthiest {
			healthiest = totHealth
		}
		if totHealth < unhealthiest {
			unhealthiest = totHealth
		}
		t--
	}

	//fmt.Println("unhealthiest: ", unhealthiest, "healthiest", healthiest)
	fmt.Println(unhealthiest, healthiest)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
