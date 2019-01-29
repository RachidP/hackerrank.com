<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
	"fmt"
	"math"
)

type graph []*node

func main() {

	getGraph()

}

/*
book introduction to algorithms page 595
Given a graph G D .V; E/ and a distinguished source vertex s, breadth-ﬁrst
search systematically explores the edges of G to “discover” every vertex that is
reachable from s. It computes the distance (smallest number of edges) from s
to each reachable vertex. It also produces a “breadth-ﬁrst tree” with root s that
contains all reachable vertices. For any vertex v reachable from s, the simple path
in the breadth-ﬁrst tree from s to v corresponds to a “shortest path” from s to v
in G, that is, a path containing the smallest number of edges. The algorithm works
on both directed and undirected graphs.*/

func (g graph) bfs(s int64) {
	sorgent := g[s-1]

	totalNumNode := len(g)
	infinite := int64(math.MaxInt64)

	//for each vertex u ∈ G.V- {s}
	for _, v := range g {
		if v == sorgent {
			continue
		}

		v.color = 'w'
		v.distance = infinite
		v.father = nil
	}

	sorgent.color = 'g'
	sorgent.distance = 0

	//Queue
	sq := make([]*node, 0, totalNumNode)
	sq = append(sq, sorgent) //add the s to to queue
	numberNodeQ := 1         //number node in the queue

	for numberNodeQ > 0 { //until the queue is empty

		u := sq[0]  //get a node from queue
		sq = sq[1:] //move the queue to next element
		numberNodeQ--

		adjU := u.next
		for adjU != nil {
			v := g[adjU.id-1]

			if v.color == 'w' {
				v.color = 'g'
				v.distance = u.distance + adjU.distance
				v.father = u
				sq = append(sq, v)
				numberNodeQ++

			}

			adjU = adjU.next
		}

		u.color = 'b'
	}

}

//print the path from S to D in the graph if exist
func (g graph) printDistance(s int64) {
	infinite := int64(math.MaxInt64)
	for _, v := range g {
		if v.id == s {
			continue
		}
		if v.distance == infinite {
			fmt.Printf("%d ", -1)

		} else {

			fmt.Printf("%d ", v.distance)
		}

	}
	fmt.Printf("\n")

}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
package main

import (
	"fmt"
	"math"
)

type graph []*node

func main() {

	getGraph()

}

/*
book introduction to algorithms page 595
Given a graph G D .V; E/ and a distinguished source vertex s, breadth-ﬁrst
search systematically explores the edges of G to “discover” every vertex that is
reachable from s. It computes the distance (smallest number of edges) from s
to each reachable vertex. It also produces a “breadth-ﬁrst tree” with root s that
contains all reachable vertices. For any vertex v reachable from s, the simple path
in the breadth-ﬁrst tree from s to v corresponds to a “shortest path” from s to v
in G, that is, a path containing the smallest number of edges. The algorithm works
on both directed and undirected graphs.*/

func (g graph) bfs(s int64) {
	sorgent := g[s-1]

	totalNumNode := len(g)
	infinite := int64(math.MaxInt64)

	//for each vertex u ∈ G.V- {s}
	for _, v := range g {
		if v == sorgent {
			continue
		}

		v.color = 'w'
		v.distance = infinite
		v.father = nil
	}

	sorgent.color = 'g'
	sorgent.distance = 0

	//Queue
	sq := make([]*node, 0, totalNumNode)
	sq = append(sq, sorgent) //add the s to to queue
	numberNodeQ := 1         //number node in the queue

	for numberNodeQ > 0 { //until the queue is empty

		u := sq[0]  //get a node from queue
		sq = sq[1:] //move the queue to next element
		numberNodeQ--

		adjU := u.next
		for adjU != nil {
			v := g[adjU.id-1]

			if v.color == 'w' {
				v.color = 'g'
				v.distance = u.distance + adjU.distance
				v.father = u
				sq = append(sq, v)
				numberNodeQ++

			}

			adjU = adjU.next
		}

		u.color = 'b'
	}

}

//print the path from S to D in the graph if exist
func (g graph) printDistance(s int64) {
	infinite := int64(math.MaxInt64)
	for _, v := range g {
		if v.id == s {
			continue
		}
		if v.distance == infinite {
			fmt.Printf("%d ", -1)

		} else {

			fmt.Printf("%d ", v.distance)
		}

	}
	fmt.Printf("\n")

}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
