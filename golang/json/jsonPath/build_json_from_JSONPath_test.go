package jsonPath

import (
	"encoding/json"
	"os"

	"github.com/ohler55/ojg/jp"
)

func Example_build_json_from_JSONPath() {
	jsonPathValues := map[string]any{
		"store.book.category.path":    "reference",
		"store.book.author.age.color": "Nigel Rees",
		"store.book.title.size":       "Moby Dick",
		"store.bicycle.color":         "red",
		"store.bicycle.price":         19.95,
	}

	data := map[string]any{}
	for jsonPath, value := range jsonPathValues {
		jp.MustParseString(jsonPath).MustSet(data, value)
	}

	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(j)

	// Output:
	// {
	//  "store": {
	//    "bicycle": {
	//      "color": "red",
	//      "price": 19.95
	//    },
	//    "book": {
	//      "author": {
	//        "age": {
	//          "color": "Nigel Rees"
	//        }
	//      },
	//      "category": {
	//        "path": "reference"
	//      },
	//      "title": {
	//        "size": "Moby Dick"
	//      }
	//    }
	//  }
	// }
}
