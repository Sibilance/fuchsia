package c

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderDeclaration(t *testing.T) {
	t.Run("int my_int", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{NamedType: &NamedType{
			Name: "int",
		}}
		myType.RenderDeclaration(&b, "my_int")
		assert.Equal(t, "int my_int", b.String())
	})
	t.Run("int *my_int_ptr", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{PointerType: &PointerType{
			TargetType: Type{NamedType: &NamedType{
				Name: "int",
			}},
		}}
		myType.RenderDeclaration(&b, "my_int_ptr")
		assert.Equal(t, "int *my_int_ptr", b.String())
	})
	t.Run("int **my_int_ptr_ptr", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{PointerType: &PointerType{
			TargetType: Type{PointerType: &PointerType{
				TargetType: Type{NamedType: &NamedType{
					Name: "int",
				}},
			}},
		}}
		myType.RenderDeclaration(&b, "my_int_ptr_ptr")
		assert.Equal(t, "int **my_int_ptr_ptr", b.String())
	})
	t.Run("char *my_cstring_array[]", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{ArrayType: &ArrayType{
			ItemType: Type{PointerType: &PointerType{
				TargetType: Type{NamedType: &NamedType{
					Name: "char",
				}},
			}},
		}}
		myType.RenderDeclaration(&b, "my_cstring_array")
		assert.Equal(t, "char *my_cstring_array[]", b.String())
	})
	t.Run("char *my_cstring_array[5]", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{SizedArrayType: &SizedArrayType{
			Length: 5,
			ItemType: Type{PointerType: &PointerType{
				TargetType: Type{NamedType: &NamedType{
					Name: "char",
				}},
			}},
		}}
		myType.RenderDeclaration(&b, "my_cstring_array")
		assert.Equal(t, "char *my_cstring_array[5]", b.String())
	})
	t.Run("int (*my_array_pointer)[]", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{PointerType: &PointerType{
			TargetType: Type{ArrayType: &ArrayType{
				ItemType: Type{NamedType: &NamedType{
					Name: "int",
				}},
			}},
		}}
		myType.RenderDeclaration(&b, "my_array_pointer")
		assert.Equal(t, "int (*my_array_pointer)[]", b.String())
	})
	t.Run("int (*my_array_pointer)[5]", func(t *testing.T) {
		var b strings.Builder
		myType := &Type{PointerType: &PointerType{
			TargetType: Type{SizedArrayType: &SizedArrayType{
				Length: 5,
				ItemType: Type{NamedType: &NamedType{
					Name: "int",
				}},
			}},
		}}
		myType.RenderDeclaration(&b, "my_array_pointer")
		assert.Equal(t, "int (*my_array_pointer)[5]", b.String())
	})
}
