<<<<<<< HEAD
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

//benchmarking of the function FailsOnRootVec
func BenchmarkFailsOnRootVec(b *testing.B) {

	root := temp(b)

	for i := 0; i < b.N; i++ {
		root.failsOnRootVec()
	}

}

//benchmarking of the function FailsOnRoot
func BenchmarkFailsOnRoot(b *testing.B) {

	root := temp(b)

	for i := 0; i < b.N; i++ {
		root.failsOnRoot()
	}

}

//benchmarking of the function FailsOnRootVec and
//FailsOnRoot in one function
//pro: we use only one function to test all our functions
//contro: we have the sum of benchmarking of all functions
//so we can't understand the time for each function
// that simply gives a single and pretty meaningless measure, mixing all performances
func BenchmarkMerge(b *testing.B) {

	root := temp(b)
	merges := []func(){
		root.failsOnRootVec,
		root.failsOnRoot,
	}
	for _, merge := range merges {
		for i := 0; i < b.N; i++ {
			merge()
		}
	}
}

//benchmarking of the function FailsOnRootVec and
//FailsOnRoot in one function
// we can now see differet results as we wished
// this tecnic is also very good idea from a documentation point of view
//we can easily create subbenchmarks by calling testing.B.Run
func BenchmarkSubMerge(b *testing.B) {
	root := temp(b)

	merges := []struct {
		name string
		fun  func()
	}{
		{"AhoCorasick with Vec as a queue", root.failsOnRootVec},
		{"AhoCorasick with queue", root.failsOnRoot},
	}

	for _, merge := range merges {
		b.Run(merge.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				merge.fun()
			}
		})
	}
}

func temp(b *testing.B) *node {
	b.StopTimer()
	defer b.StartTimer()
	var in string

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
	return root

}
=======
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

//benchmarking of the function FailsOnRootVec
func BenchmarkFailsOnRootVec(b *testing.B) {

	root := temp(b)

	for i := 0; i < b.N; i++ {
		root.failsOnRootVec()
	}

}

//benchmarking of the function FailsOnRoot
func BenchmarkFailsOnRoot(b *testing.B) {

	root := temp(b)

	for i := 0; i < b.N; i++ {
		root.failsOnRoot()
	}

}

//benchmarking of the function FailsOnRootVec and
//FailsOnRoot in one function
//pro: we use only one function to test all our functions
//contro: we have the sum of benchmarking of all functions
//so we can't understand the time for each function
// that simply gives a single and pretty meaningless measure, mixing all performances
func BenchmarkMerge(b *testing.B) {

	root := temp(b)
	merges := []func(){
		root.failsOnRootVec,
		root.failsOnRoot,
	}
	for _, merge := range merges {
		for i := 0; i < b.N; i++ {
			merge()
		}
	}
}

//benchmarking of the function FailsOnRootVec and
//FailsOnRoot in one function
// we can now see differet results as we wished
// this tecnic is also very good idea from a documentation point of view
//we can easily create subbenchmarks by calling testing.B.Run
func BenchmarkSubMerge(b *testing.B) {
	root := temp(b)

	merges := []struct {
		name string
		fun  func()
	}{
		{"AhoCorasick with Vec as a queue", root.failsOnRootVec},
		{"AhoCorasick with queue", root.failsOnRoot},
	}

	for _, merge := range merges {
		b.Run(merge.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				merge.fun()
			}
		})
	}
}

func temp(b *testing.B) *node {
	b.StopTimer()
	defer b.StartTimer()
	var in string

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
	return root

}
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
