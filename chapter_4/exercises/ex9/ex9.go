package main

import "fmt"

func main() {
	m := map[string][]string{ //don't use the {} twice
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["doc_oc"] = []string{`spiderman`, `robots`, `nuclear fusion`}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
