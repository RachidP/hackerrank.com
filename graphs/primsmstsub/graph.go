<<<<<<< HEAD
<<<<<<< HEAD
package main

//https://www.youtube.com/watch?v=zVrPdF7f4-I
import (
	"fmt"
	"os"
)

//false= directed graph, true = undirected graph
var undirected = false
var time int

type node struct {
	id           int64 //id node
	next         *node // next neighbor
	weight       int64 // weight of the edge
	father       *node // father of the node
	priority     int64 // The priority of the item in the queue.
	isUndirected bool  //false= directed graph, true = undirected graph
	letter       string
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// return new node (initialized to nil)
func getNode(v int64) *node {

	return &node{
		id:     v,
		letter: string((v - 1) + 'a'),
	}
}

// setNeighbor make a node with pointing to n
func setNeighbor(verNumber int64, n *node, weight int64, isUndirected bool) *node {

	return &node{
		id:           verNumber,
		next:         n,
		weight:       weight,
		isUndirected: isUndirected,
		letter:       string((verNumber - 1) + 'a'),
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// build the graph
func getGraph() {
	var g graph

	var n, m int
	var v1, v2, weight, idOrigin int64
	f, err := os.Open("test1.txt")
	check(err)

	_, err = fmt.Fscanf(f, "%d %d\n", &n, &m)
	check(err)
	lenG := n
	g = make(graph, lenG)

	//make n nodes
	for j := 0; j < n; j++ {
		id := int64(j + 1) //id start from 1
		g[j] = getNode(id)
	}
	//make  m paths
	for j := 0; j < m; j++ {
		_, err = fmt.Fscanf(f, "%d %d %d\n", &v1, &v2, &weight)
		check(err)
		g[v1-1].next = setNeighbor(v2, g[v1-1].next, weight, undirected)
		g[v2-1].next = setNeighbor(v1, g[v2-1].next, weight, undirected)

	}

	_, err = fmt.Fscanf(f, "%d\n", &idOrigin)
	check(err)

	g.mstPrime(idOrigin - 1)

}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
package main

//https://www.youtube.com/watch?v=zVrPdF7f4-I
import (
	"fmt"
	"os"
)

//false= directed graph, true = undirected graph
var undirected = false
var time int

type node struct {
	id           int64 //id node
	next         *node // next neighbor
	weight       int64 // weight of the edge
	father       *node // father of the node
	priority     int64 // The priority of the item in the queue.
	isUndirected bool  //false= directed graph, true = undirected graph
	letter       string
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// return new node (initialized to nil)
func getNode(v int64) *node {

	return &node{
		id:     v,
		letter: string((v - 1) + 'a'),
	}
}

// setNeighbor make a node with pointing to n
func setNeighbor(verNumber int64, n *node, weight int64, isUndirected bool) *node {

	return &node{
		id:           verNumber,
		next:         n,
		weight:       weight,
		isUndirected: isUndirected,
		letter:       string((verNumber - 1) + 'a'),
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// build the graph
func getGraph() {
	var g graph

	var n, m int
	var v1, v2, weight, idOrigin int64
	f, err := os.Open("test1.txt")
	check(err)

	_, err = fmt.Fscanf(f, "%d %d\n", &n, &m)
	check(err)
	lenG := n
	g = make(graph, lenG)

	//make n nodes
	for j := 0; j < n; j++ {
		id := int64(j + 1) //id start from 1
		g[j] = getNode(id)
	}
	//make  m paths
	for j := 0; j < m; j++ {
		_, err = fmt.Fscanf(f, "%d %d %d\n", &v1, &v2, &weight)
		check(err)
		g[v1-1].next = setNeighbor(v2, g[v1-1].next, weight, undirected)
		g[v2-1].next = setNeighbor(v1, g[v2-1].next, weight, undirected)

	}

	_, err = fmt.Fscanf(f, "%d\n", &idOrigin)
	check(err)

	g.mstPrime(idOrigin - 1)

}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
