package c

import (
	"fmt"
	"strings"
)

func (h *Header) Render(b *strings.Builder) {
	for _, include := range h.Includes {
		include.Render(b)
	}
	for _, structDefinition := range h.StructDefinitions {
		structDefinition.Render(b)
	}
	for _, declaration := range h.Declarations {
		declaration.Render(b)
	}
}

func (p *Program) Render(b *strings.Builder) {
	for _, include := range p.Includes {
		include.Render(b)
	}
	for _, structDefinition := range p.StructDefinitions {
		structDefinition.Render(b)
	}
	for _, declaration := range p.Declarations {
		declaration.Render(b)
	}
	for _, functionDefinition := range p.FunctionDefinitions {
		functionDefinition.Render(b)
	}
}

func (i *Include) Render(b *strings.Builder) {
	var start, end string
	if i.IsSystem {
		start = "<"
		end = ">\n"
	} else {
		start = `"`
		end = `"\n`
	}
	b.WriteString("#include ")
	b.WriteString(start)
	b.WriteString(i.Path)
	b.WriteString(end)

}

func (s *StructDefinition) Render(b *strings.Builder) {
	b.WriteString("struct ")
	b.WriteString(s.Name)
	b.WriteString("{\n")
	for _, field := range s.Fields {
		field.Render(b)
	}
	b.WriteString("};\n")
}

func (d *Declaration) Render(b *strings.Builder) {
	if d.IsStatic {
		b.WriteString("static ")
	}
	if d.IsExtern {
		b.WriteString("extern ")
	}
	if d.IsVolatile {
		b.WriteString("volatile ")
	}
	d.Type.RenderDeclaration(b, d.Name)
	b.WriteString(";\n")
}

func (f *FunctionDefinition) Render(b *strings.Builder) {
	if f.IsStatic {
		b.WriteString("static ")
	}
	if f.IsInline {
		b.WriteString("inline ")
	}
	(&Type{FunctionType: &f.Type}).RenderDeclaration(b, f.Name)
	b.WriteString("{\n")
	for _, statement := range f.Statements {
		statement.Render(b)
	}
	b.WriteString("}\n")
}

func (f *StructField) Render(b *strings.Builder) {}

func (t *Type) RenderDeclaration(b *strings.Builder, name string) {
	t.renderType(b, func(bool) {
		// Variable names are already atomic. No need to bind tightly.
		b.WriteString(name)
	})
}

func (t *Type) renderType(b *strings.Builder, renderTarget func(bindTightly bool)) {
	if t.NamedType != nil {
		t.NamedType.renderType(b, renderTarget)
	} else if t.PointerType != nil {
		t.PointerType.renderType(b, renderTarget)
	} else if t.ArrayType != nil {
		t.ArrayType.renderType(b, renderTarget)
	} else if t.SizedArrayType != nil {
		t.SizedArrayType.renderType(b, renderTarget)
	} else if t.FunctionType != nil {
		t.FunctionType.renderType(b, renderTarget)
	} else {
		b.WriteString("void ")
		renderTarget(false)
	}
}

func (t *NamedType) renderType(b *strings.Builder, renderTarget func(bool)) {
	b.WriteString(t.Name)
	b.WriteString(" ")
	renderTarget(false)
}

func (t *PointerType) renderType(b *strings.Builder, renderTarget func(bool)) {
	t.TargetType.renderType(b, func(bindTightly bool) {
		if bindTightly {
			b.WriteString("(")
		}
		b.WriteString("*")
		renderTarget(false)
		if bindTightly {
			b.WriteString(")")
		}
	})
}

func (t *ArrayType) renderType(b *strings.Builder, renderTarget func(bool)) {
	t.ItemType.renderType(b, func(bool) {
		// Array types bind more strongly than other types, so no need to parenthesize
		// this type. However, the target type must be bound tightly.
		renderTarget(true)
		b.WriteString("[]")
	})
}

func (t *SizedArrayType) renderType(b *strings.Builder, renderTarget func(bool)) {
	t.ItemType.renderType(b, func(bool) {
		// Array types bind more strongly than other types, so no need to parenthesize
		// this type. However, the target type must be bound tightly.
		renderTarget(true)
		b.WriteString(fmt.Sprintf("[%d]", t.Length))
	})
}

func (t *FunctionType) renderType(b *strings.Builder, renderTarget func(bool)) {
	t.ReturnType.renderType(b, func(bool) {
		// Function types bind more strongly than other types, so no need to parenthesize
		// this type. However, the target type must be bound tightly.
		renderTarget(true)
		b.WriteString("(")
		for i, argument := range t.Arguments {
			if i != 0 {
				b.WriteString(", ")
			}
			argument.Type.RenderDeclaration(b, argument.Name)
		}
		b.WriteString(")")
	})
}

func (s *Statement) Render(b *strings.Builder) {}
