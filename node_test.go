// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"reflect"
	"testing"
)

func TestDecodeNode(t *testing.T) {
	enc := []byte(`{
	"array": [1, 2,
	{"z": "a"}],
	"object": {
		"foo": "bar"
	},
	"int": 12,
	"str": "s"
	}`)
	expect := Node{
		Start: 1,
		End:   90,
		Value: map[string]Node{
			"array": {
				Start:    13,
				End:      30,
				KeyStart: 3,
				KeyEnd:   10,
				Value: []Node{
					{
						Start: 13,
						End:   13,
						Value: float64(1),
					},
					{
						Start: 16,
						End:   16,
						Value: float64(2),
					},
					{
						Start: 21,
						End:   29,
						Value: map[string]Node{
							"z": {
								Start:    26,
								End:      28,
								KeyStart: 21,
								KeyEnd:   24,
								Value:    "a",
							},
						},
					},
				},
			},
			"object": {
				Start:    45,
				End:      62,
				KeyStart: 34,
				KeyEnd:   42,
				Value: map[string]Node{
					"foo": {
						Start:    55,
						End:      59,
						KeyStart: 48,
						KeyEnd:   53,
						Value:    "bar",
					},
				},
			},
			"int": {
				Start:    73,
				End:      74,
				KeyStart: 66,
				KeyEnd:   71,
				Value:    float64(12),
			},
			"str": {
				Start:    85,
				End:      87,
				KeyStart: 78,
				KeyEnd:   83,
				Value:    "s",
			},
		},
	}

	var dec Node
	err := Unmarshal(enc, &dec)
	if err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !reflect.DeepEqual(expect, dec) {
		t.Fatalf("expected %+v, got %+v", expect, dec)
	}
}
