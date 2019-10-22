package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"runtime/pprof"
)

//	maps := make(map[pair]int)
/*
	primes := make([]int64, 560)
	var i int64
	count := 0
	for i = 0; i < 4001; i++ {
		p := big.NewInt(i)
		if p.ProbablyPrime(20) {
			//fmt.Println(p)
			primes[count] = i
			count++
		}
	}
	//fmt.Println(count)
	fmt.Println(primes)
*/
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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
	var amax, bmax, imax int64 = 0, 0, 0
	var a, b int64
	for a = -1000; a < 1001; a++ {
		for b = -1000; b < 1001; b++ {
			var i int64 = 0

			for {
				n := i*i + a*i + b
				p := big.NewInt(n)
				if p.ProbablyPrime(3) {
					if i > imax {
						amax = a
						bmax = b
						imax = i
					}
					i++
				} else {
					break
				}
			}
		}
	}
	fmt.Println(amax, bmax, imax)
	fmt.Println(amax * bmax)
}
