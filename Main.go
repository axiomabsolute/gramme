package main

import (
	"encoding/json"
	"fmt"

	"github.com/axiomabsolute/gramme/annotations"
)

var testText string = `The limerick packs laughs anatomical
Into space that is quite economical.
But the good ones I've seen
So seldom are clean
And the clean ones so seldom are comical.`

func main() {
	batch := annotations.NewBatch(&testText)
	for _, annotation := range batch.All() {
		result, _ := json.Marshal(annotation)
		fmt.Println(string(result))
	}
}
