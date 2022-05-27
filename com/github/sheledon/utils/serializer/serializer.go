package serializer

import (
	"encoding/json"
	"fmt"
)

type Serializer interface {
	Serialize(interface{})([]byte,error)
	Deserialize([]byte,interface{}) error
}

func NewDefaultSerializer() *JSONSerializer{
	return &JSONSerializer{}
}
type JSONSerializer struct {}

func (serializer *JSONSerializer) Serialize(v interface{}) (bytes []byte,err error){
	bytes, err = json.Marshal(v)
	if err != nil {
		err = fmt.Errorf("serialize fail, err: %v",err)
	}
	return
}
func (serializer *JSONSerializer) Deserialize(bytes []byte,v interface{}) (err error){
	err = json.Unmarshal(bytes, v)
	if err != nil {
		err = fmt.Errorf("deserialize fail, err : %v",err)
	}
	return
}

