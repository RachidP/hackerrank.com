package main

//https://www.hackerrank.com/challenges/primsmstsub/problem
import (
	"container/heap"
	"fmt"
	"math"
)

type graph []*node

func main() {

	getGraph()

}

func (g graph) mstPrime(origin int64) {
	var total int64
	var v *node
	infinite := int64(math.MaxInt64)
	checkM := make(map[int64]*node)  //map check if a node is still in the queue
	path := make([]int64, 0, len(g)) //path from origin g[origin] until end

	//set all the vertex priority as infinite
	for i, u := range g {
		u.priority = infinite
		u.father = nil
		u.index = i
		checkM[u.id] = u
	}
	//g[origin].father = g[origin]
	g[origin].priority = 0
	checkM[g[origin].id] = g[origin]

	heap.Init(&g) //push the graph in a priority queue

	path = append(path, g[origin].id)

	sizeQ := g.Len()

	for sizeQ > 0 {

		u := heap.Pop(&g).(*node) //get the veertex in queue with the lowest priority

		delete(checkM, u.id) //delete node from the map

		sizeQ = g.Len() //update size of the queue

		if u.father != nil {
			total += u.priority
			//	fmt.Printf("%v ----> %v (%v)\n", string((u.father.id-1)+'a'), u.id, string((u.id-1)+'a'))
			path = append(path, u.id)
		}

		//for each vertex v ∈ G.adj[u] (all edges from the node u)
		adjU := u.next
		for adjU != nil {
			//if v∈Q
			if val, ok := checkM[adjU.id]; ok {
				v = val
				// if  w(u,v)<v.key
				if adjU.weight < v.priority {
					//update priority and father
					v.father = u
					v.priority = adjU.weight
					g.update(v)

				}
			}

			adjU = adjU.next
		}

	}
	fmt.Println(total)

}
