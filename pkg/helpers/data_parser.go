package helpers

import "encoding/json"

func DataParser[T1 any, T2 any](src T1, dst T2) {

	byteData, err := json.Marshal(src)

	if err != nil {
		return
	}

	json.Unmarshal(byteData, dst)
}
