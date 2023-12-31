// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Foo struct {
	Label string `json:"label"`
}

const FooAvroCRC64Fingerprint = "\xb2k\x10~\xee\x8dߔ"

func NewFoo() Foo {
	r := Foo{}
	return r
}

func DeserializeFoo(r io.Reader) (Foo, error) {
	t := NewFoo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFooFromSchema(r io.Reader, schema string) (Foo, error) {
	t := NewFoo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFoo(r Foo, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Label, w)
	if err != nil {
		return err
	}
	return err
}

func (r Foo) Serialize(w io.Writer) error {
	return writeFoo(r, w)
}

func (r Foo) Schema() string {
	return "{\"fields\":[{\"name\":\"label\",\"type\":\"string\"}],\"name\":\"org.apache.avro.Foo\",\"type\":\"record\"}"
}

func (r Foo) SchemaName() string {
	return "org.apache.avro.Foo"
}

func (_ Foo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Foo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Foo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Foo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Foo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Foo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Foo) SetString(v string)   { panic("Unsupported operation") }
func (_ Foo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Foo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Label}

		return w

	}
	panic("Unknown field index")
}

func (r *Foo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Foo) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Foo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Foo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Foo) HintSize(int)                     { panic("Unsupported operation") }
func (_ Foo) Finalize()                        {}

func (_ Foo) AvroCRC64Fingerprint() []byte {
	return []byte(FooAvroCRC64Fingerprint)
}

func (r Foo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["label"], err = json.Marshal(r.Label)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Foo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["label"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Label); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for label")
	}
	return nil
}
