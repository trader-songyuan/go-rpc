package serializer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type person struct {
	Name   string
	Age    int
	Gender string
}
type teacher struct {
	person
	TeachType string
}

func TestSerializer(t *testing.T) {
	t.Run("jsonSerializer", func(t *testing.T) {
		p := person{Name: "ren", Age: 10, Gender: "man"}
		tt := teacher{
			TeachType: "math",
			person:    p,
		}
		serializer := NewDefaultSerializer()
		bytes, err := serializer.serialize(tt)
		if err != nil {
			t.Fatal(err)
		}
		var nt teacher
		json.Unmarshal(bytes, &nt)
		err = serializer.deserialize(bytes, &nt)
		if err != nil {
			t.Fatal(err)
		}
		if ! reflect.DeepEqual(nt,tt) {
			t.Fatal("default serializer test fail")
		}
	})
}
func TestJson(t *testing.T)  {
	p := person{Name: "ren", Age: 10, Gender: "man"}
	tt := teacher{
		TeachType: "math",
		person:    p,
	}
	js, _ := json.Marshal(tt)
	fmt.Printf("JSON format: %s \n", js)
	pp := teacher{}
	json.Unmarshal(js,&pp)
	if !reflect.DeepEqual(tt,pp) {
		t.Fatal("json serializer fail")
	}
}
