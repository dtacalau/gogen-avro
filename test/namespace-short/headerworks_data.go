// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

// Common information related to the event which must be included in any clean event  
type HeaderworksData struct {

	
	// Unique identifier for the event used for de-duplication and tracing.
	
	
		Uuid *UnionNullDatatypeUUID
	

	
	// Fully qualified name of the host that generated the event that generated the data.
	
	
		Hostname *UnionNullString
	

	
	// Trace information not redundant with this object
	
	
		Trace *UnionNullHeaderworksTrace
	

}

func NewHeaderworksData() (*HeaderworksData) {
	return &HeaderworksData{}
}

func DeserializeHeaderworksData(r io.Reader) (*HeaderworksData, error) {
	t := NewHeaderworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeHeaderworksDataFromSchema(r io.Reader, schema string) (*HeaderworksData, error) {
	t := NewHeaderworksData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeHeaderworksData(r *HeaderworksData, w io.Writer) error {
	var err error
	
	err = writeUnionNullDatatypeUUID( r.Uuid, w)
	if err != nil {
		return err			
	}
	
	err = writeUnionNullString( r.Hostname, w)
	if err != nil {
		return err			
	}
	
	err = writeUnionNullHeaderworksTrace( r.Trace, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *HeaderworksData) Serialize(w io.Writer) error {
	return writeHeaderworksData(r, w)
}

func (r *HeaderworksData) Schema() string {
	return "{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"headerworks.Data\",\"type\":\"record\"}"
}

func (r *HeaderworksData) SchemaName() string {
	return "headerworks.Data"
}

func (_ *HeaderworksData) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetInt(v int32) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetLong(v int64) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetString(v string) { panic("Unsupported operation") }
func (_ *HeaderworksData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *HeaderworksData) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.Uuid = NewUnionNullDatatypeUUID()
	
		
		
			return r.Uuid
		
	
	case 1:
		
			r.Hostname = NewUnionNullString()
	
		
		
			return r.Hostname
		
	
	case 2:
		
			r.Trace = NewUnionNullHeaderworksTrace()
	
		
		
			return r.Trace
		
	
	}
	panic("Unknown field index")
}

func (r *HeaderworksData) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.Uuid = NewUnionNullDatatypeUUID()

		return
	
	
        
	case 1:
       	 	r.Hostname = NewUnionNullString()

		return
	
	
        
	case 2:
       	 	r.Trace = NewUnionNullHeaderworksTrace()

		return
	
	
	}
	panic("Unknown field index")
}

func (_ *HeaderworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *HeaderworksData) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *HeaderworksData) Finalize() { }
