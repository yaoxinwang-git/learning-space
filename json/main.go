package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	//jsonData := []byte(`
	//{
	//    "Name": "Standard",
	//    "Fruit": [
	//         "Apple",
	//        "Banana",
	//        "Orange"
	//    ],
	//    "ref": 999,
	//    "Created": "2018-04-09T23:00:00Z"
	//}`)
	jsonData := []byte(`[
  {
    "Name": "Eve",
    "Age": 6,
    "Parents": [
      "Alice",
      "Bob"
    ]
  },
  {
    "Name": "Standard",
    "Fruit": [
      "Apple",
      "Banana",
      "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
  }
]`)
	//var v interface{}
	v := make([]interface{}, 0)
	json.Unmarshal(jsonData, &v)
	//data := v.(map[string]interface{})

	//for k, v := range v {
	//switch v := v.(type) {
	//case string:
	//	fmt.Println(k, v, "(string)")
	//case float64:
	//	fmt.Println(k, v, "(float64)")
	//case []interface{}:
	//	fmt.Println(k, "(array):")
	//	for i, u := range v {
	//		fmt.Println("    ", i, u)
	//	}
	//default:
	//	fmt.Println(k, v, "(unknown)")
	//}
	//}
	for k, v := range v {
		fmt.Printf("%d正在解析\n", k)
		data := v.(map[string]interface{})
		for k, v := range data {
			switch v := v.(type) {
			case string:
				fmt.Println(k, v, "(string)")
			case float64:
				fmt.Println(k, v, "(float64)")
			case []interface{}:
				fmt.Println(k, "(array):")
				for i, u := range v {
					fmt.Println("    ", i, u)
				}
			default:
				fmt.Println(k, v, "(unknown)")
			}
		}
	}

}
