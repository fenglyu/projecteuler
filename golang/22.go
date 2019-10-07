package main

import (
	"bufio"
	"flag"
	"fmt"
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

		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[1 : i-1], nil
			}
		}
		// If we're at EOF, we have a final, non ',' terminated words "ALONSO".
		// Return it.
		if atEOF {
			return len(data), data[1 : len(data)-1], nil
		}
		// Request more data
		return 0, nil, nil
	}

	scanner.Split(onComma)

	var words []*Word
	words = make([]*Word, 6000, 6001)

	i := 0
	for scanner.Scan() {
		w := Word(scanner.Text())
		words[i] = &w
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	d := Dictionary{
		words: words,
		count: i,
	}

	sort.Sort(ByAlphabetical{d})
	//	fmt.Println(d)
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

	fmt.Println(sum)
}
