<<<<<<< HEAD
<<<<<<< HEAD
// An Item is something we manage in a priority queue.
package main

import (
	"container/heap"
)

func (pq *graph) Len() int { return len(*pq) }

func (pq graph) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority <= pq[j].priority
}

func (pq graph) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j

}

func (pq *graph) Push(x interface{}) {

	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *graph) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority of the node  in the queue.
func (pq *graph) update(v *node) {

	heap.Fix(pq, v.index)
}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
// An Item is something we manage in a priority queue.
package main

import (
	"container/heap"
)

func (pq *graph) Len() int { return len(*pq) }

func (pq graph) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority <= pq[j].priority
}

func (pq graph) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j

}

func (pq *graph) Push(x interface{}) {

	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *graph) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority of the node  in the queue.
func (pq *graph) update(v *node) {

	heap.Fix(pq, v.index)
}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
