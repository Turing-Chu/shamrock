// Author: Turing Zhu
// Date: 2021/8/23 7:58 PM
// File: json.go

package shamrock

import (
	"encoding/json"
	"io/ioutil"
)

func UnmarshalFile(filePath string, customTypeVariable interface{}) (interface{}, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &customTypeVariable)
	return customTypeVariable, err
}
