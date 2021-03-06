// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullHeaderworksTraceTypeEnum int
const (

	 UnionNullHeaderworksTraceTypeEnumNull UnionNullHeaderworksTraceTypeEnum = 0

	 UnionNullHeaderworksTraceTypeEnumHeaderworksTrace UnionNullHeaderworksTraceTypeEnum = 1

)

type UnionNullHeaderworksTrace struct {

	Null *types.NullVal

	HeaderworksTrace *HeaderworksTrace

	UnionType UnionNullHeaderworksTraceTypeEnum
}

func writeUnionNullHeaderworksTrace(r *UnionNullHeaderworksTrace, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullHeaderworksTraceTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullHeaderworksTraceTypeEnumHeaderworksTrace:
		return writeHeaderworksTrace(r.HeaderworksTrace, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksTrace")
}

func NewUnionNullHeaderworksTrace() *UnionNullHeaderworksTrace {
	return &UnionNullHeaderworksTrace{}
}

func (_ *UnionNullHeaderworksTrace) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullHeaderworksTrace) SetLong(v int64) { 
	r.UnionType = (UnionNullHeaderworksTraceTypeEnum)(v)
}
func (r *UnionNullHeaderworksTrace) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.HeaderworksTrace = NewHeaderworksTrace()
		
		
		return r.HeaderworksTrace
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullHeaderworksTrace) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) Finalize()  { }
