package main

//https://www.hackerrank.com/challenges/bfsshortreach/problem

import (
	"fmt"
	"os"
)

//false= directed graph, true = undirected graph
var undirected = false

type node struct {
	id       int64 //
	next     *node //
	color    rune  //color  of the node
	father   *node
	distance int64 // distance from the source to vertex u
}

// return new node (initialized to nil)
func getNode(v int64) *node {

	return &node{
		id: v,
	}
}

// return new node pointing to n
func setNeighbor(verNumber int64, n *node) *node {

	return &node{
		id:       verNumber,
		distance: 6,
		next:     n,
	}

}

// printGraph all the graph
func (g graph) printGraph() {
	for i, v := range g {
		fmt.Printf("%d (%v)", i+1, v.distance)
		//fmt.Printf("%d", i)
		v.next.printNeighbors()

	}
}

//printNeighbors all the neighbor of the current Node
func (n *node) printNeighbors() {
	tmp := n
	for tmp != nil {
		fmt.Printf("---->%d(%v)", tmp.id, tmp.distance)
		tmp = tmp.next
	}
	fmt.Println()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// build the graph
func getGraph() {
	var g graph
	var q int
	var n, m int
	var v1, v2, idOrigin int64
	f, err := os.Open("test1.txt")
	check(err)

	_, err = fmt.Fscanf(f, "%d\n", &q)
	check(err)

	for i := 0; i < q; i++ {
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
			_, err = fmt.Fscanf(f, "%d %d\n", &v1, &v2)
			check(err)
			g[v1-1].next = setNeighbor(v2, g[v1-1].next)
			g[v2-1].next = setNeighbor(v1, g[v2-1].next)
		}

		_, err = fmt.Fscanf(f, "%d\n", &idOrigin)
		check(err)
		// fmt.Println("\n original graph")
		// g.printGraph()
		g.bfs(idOrigin)
		// fmt.Println("\n new graph")
		// g.printGraph()
		// fmt.Println("\nshortest path")
		g.printDistance(idOrigin)

	}

}
