<<<<<<< HEAD
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
	rand.Seed(time.Now().UTC().UnixNano())
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

	wg.Add(25)

	pcs := make(chan int, 8)
	go func() {
		const work = 25
		count := 1
		for w := 1; w <= work; w++ {
			if count > 8 {
				count = 1
			}

			pcs <- count
			count++

		}
		close(pcs)
	}()
	for i := 1; i <= 25; i++ {
		go usePc(i, pcs)
	}

	wg.Wait()
	fmt.Println("The place is empty, let's close up and go to the beach!")
}

func usePc(user int, pcs chan int) {

	// if len(pcs) == cap(pcs) {
	// 	fmt.Printf("Tourist %d waiting for turn. \n", user)
	// }
	tmp := <-pcs
	fmt.Printf("Tourist %d is online on PC %v.\n", user, tmp)
	//getGetReady := rand.Intn(50) //rand number [0,60)
	getGetReady := 10 //rand number [0,60)
	start := time.Now()
	time.Sleep(time.Duration(getGetReady) * time.Second)
	timeGetReady := time.Since(start)
	fmt.Printf("Touris %d is done, having spent %v minutes online.\n", user, timeGetReady)

	wg.Done()

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
=======
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
	rand.Seed(time.Now().UTC().UnixNano())
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

	wg.Add(25)

	pcs := make(chan int, 8)
	go func() {
		const work = 25
		count := 1
		for w := 1; w <= work; w++ {
			if count > 8 {
				count = 1
			}

			pcs <- count
			count++

		}
		close(pcs)
	}()
	for i := 1; i <= 25; i++ {
		go usePc(i, pcs)
	}

	wg.Wait()
	fmt.Println("The place is empty, let's close up and go to the beach!")
}

func usePc(user int, pcs chan int) {

	// if len(pcs) == cap(pcs) {
	// 	fmt.Printf("Tourist %d waiting for turn. \n", user)
	// }
	tmp := <-pcs
	fmt.Printf("Tourist %d is online on PC %v.\n", user, tmp)
	//getGetReady := rand.Intn(50) //rand number [0,60)
	getGetReady := 10 //rand number [0,60)
	start := time.Now()
	time.Sleep(time.Duration(getGetReady) * time.Second)
	timeGetReady := time.Since(start)
	fmt.Printf("Touris %d is done, having spent %v minutes online.\n", user, timeGetReady)

	wg.Done()

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
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
