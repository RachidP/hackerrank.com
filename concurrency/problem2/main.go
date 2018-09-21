package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

type eats struct {
	foods   string
	morsels int
}

var wg sync.WaitGroup
var traceFlag = flag.String("trace", "", "execution trace file, e.g., ./trace.out")

func main() {
	flag.Parse()
	if *traceFlag != "" {
		traceFile, err := os.Create(*traceFlag)
		if err != nil {
			panic(err)
		}
		trace.Start(traceFile)
		defer func() {
			trace.Stop()
			if err := traceFile.Close(); err != nil {
				log.Panic(err)
			}
		}()
	}

	var foods = []string{"Chorizo", "Chopittos", "PiemontosDePadron", "Croquetas", "PatatasBravas"}
	//var foods = []string{"Chorizo", "Chopittos", "Croquetas", "PatatasBravas"}
	//var morsels = []int{rand.Intn(6) + 5, rand.Intn(6) + 5, rand.Intn(6) + 5, rand.Intn(6) + 5, rand.Intn(6) + 5}
	//	var morsels = []int{2, 2, 2, 2, 2}
	var morsels = []int{3, 12, 14, 25, 10}

	e := []eats{}
	for i, v := range foods {
		f1 := eats{foods: v, morsels: morsels[i]}
		e = append(e, f1)

	}
	fmt.Println(e)

	rand.Seed(time.Now().UTC().UnixNano())

	//var morsels = []int{1, 2, 4, 3, 1}
	names := []string{"Alice", "Bob", "Dave", "Charlie"}

	fmt.Println("Bon appétit!")

	ch := make(chan eats, len(e))
	//n := len(e.foods)

	for _, name := range names {
		wg.Add(1)
		go eat(name, ch)

	}
	for _, v := range e {
		ch <- v
	}

	wg.Wait()
	fmt.Println("That was delicious!")
}
func eat(name string, ch chan eats) {
	//	second := time.Second
	tmp := <-ch

	//fmt.Println("r= ", r)
	//fmt.Printf("%s is eating %s: portions= %d \n", name, tmp.foods[r], tmp.morsels[r])
	fmt.Printf("%s is enjoying some %s\n", name, tmp.foods)
	tmp.morsels--
	if tmp.morsels <= 0 {
		wg.Done()
		return
	}

	ch <- tmp
	eat(name, ch)
	//go eat(name, ch)

}
