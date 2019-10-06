package main

import (
	"bufio"
	"flag"
	"fmt"
	//	"strings"
	"log"
	"os"
	"sort"
)

/* sed -e 's/,/\n/g' -e 's/\"//g' p022_names.txt|sort > 22_sed.txt
➜ golang git:(master) ✗ diff  22.txt 22_sed.txt
146a147
> ALONSO

 why does the program miss string "ALONSO" ?
*/

type Word string

//func (d Word) Len() int { return len(d) }
//func (d Word) Swap(i, j int) {
//	b := []byte(d)
//	b[i], b[j] = b[j], b[i]
//}

//type Dictionary []*Word

type Dictionary struct {
	words []*Word
	count int
}

func (d Dictionary) Len() int { return d.count }

func (d Dictionary) Swap(i, j int) { d.words[i], d.words[j] = d.words[j], d.words[i] }

func (d Dictionary) String() string {
	str := "[Dictionary]\n"
	for i, v := range d.words {
		if i >= d.Len() {
			break
		}
		str += fmt.Sprintf("%s\n", *v)
	}
	return str
}

type ByAlphabetical struct{ Dictionary }

func (s ByAlphabetical) Less(i, j int) bool {
	a, b := []byte(*s.Dictionary.words[i]), []byte(*s.Dictionary.words[j])
	//	fmt.Println(*s.Dictionary.words[i], *s.Dictionary.words[j])
	//	fmt.Println(a, b)
	count := len(a)
	if len(b) < count {
		count = len(b)
	}

	idx := 0
	for {
		if idx >= count {
			break
		}

		if a[idx] == b[idx] {
			idx++
			continue
		} else if a[idx] < b[idx] {
			return true
		} else {
			return false
		}
	}

	if idx == len(a) {
		return true
	}

	return false
}

func main() {

	fileName := flag.String("file", "p022_names.txt", "File name")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[1 : i-1], nil
				//return i + 1, data[:i], nil
			}
		}

		//return 0, data, bufio.ErrFinalToken
		// fix me plz
		return
	}

	scanner.Split(onComma)

	var words []*Word
	words = make([]*Word, 6000, 6001)

	i := 0
	for scanner.Scan() {
		//		fmt.Printf("%s\n", scanner.Text())
		//words = append(words, Word{scanner.Text()})
		w := Word(scanner.Text())
		words[i] = &w
		i++
	}

	//	fmt.Println(words)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	d := Dictionary{
		words: words,
		count: i,
	}
	//	fmt.Println(d.Len())
	sort.Sort(ByAlphabetical{d})

	fmt.Println(d)

	sum := 0
	for i, v := range d.words {
		if i >= d.Len() {
			break
		}
		t := 0
		for _, c := range []byte(*v) {
			//			fmt.Println(c)
			t += int(c-'A') + 1
		}
		//	fmt.Println(*v, t, i+1, t*(i+1))
		sum += t * (i + 1)
	}

	//	fmt.Println(sum)
}
