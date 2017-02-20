/*
	Created On: 7 Feb 2017
	Author: Nilav

	Structure the JSON in defined order
 */


package utility

import (
	"bytes"
	"encoding/json"
)

type KeyVal struct {
	Key interface{}
	Val interface{}
}

// Define an ordered map
type JsonOrder []KeyVal

// Implement the json.Marshaler interface
func (omap JsonOrder) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, kv := range omap {
		if i != 0 {
			buf.WriteString(",")
		}
		// marshal key
		key, err := json.Marshal(kv.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		// marshal value
		val, err := json.Marshal(kv.Val)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}