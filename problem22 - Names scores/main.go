package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type byIndex []string

func (a byIndex) Len() int           { return len(a) }
func (a byIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIndex) Less(i, j int) bool { return a[i] < a[j] }

func read(fileName string) ([]string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("File opening failed: %v", err)
	}

	names := []string{}
	for _, value := range strings.Split(string(file), ",") {
		names = append(names, strings.Trim(value, "\""))
	}

	return names, nil
}

func sumOfLetters(name string) int {
	sum := 0
	name = strings.ToUpper(name)
	for i := 0; i < len(name); i++ {
		sum += int(name[i]-'A') + 1
	}

	return sum
}

func main() {
	m, err := read("input.txt")

	if err != nil {
		panic(err)
	}

	sort.Sort(byIndex(m))

	sum := 0
	for i, value := range m {
		sum += sumOfLetters(value) * (i + 1)
	}

	fmt.Println(sum)
}
