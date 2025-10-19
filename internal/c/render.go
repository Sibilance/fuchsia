package c

import "strings"

func (h Header) Render(b strings.Builder) {
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

func (p Program) Render(b strings.Builder) {
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

func (i Include) Render(b strings.Builder) {
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

func (s StructDefinition) Render(b strings.Builder) {
	b.WriteString("struct ")
	b.WriteString(s.Name)
	b.WriteString("{\n")
	for _, field := range s.Fields {
		field.Render(b)
	}
	b.WriteString("};\n")
}

func (d Declaration) Render(b strings.Builder) {
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

func (f FunctionDefinition) Render(b strings.Builder) {
	if f.IsStatic {
		b.WriteString("static ")
	}
	if f.IsInline {
		b.WriteString("inline ")
	}
	Type{FunctionType: &f.Type}.RenderDeclaration(b, f.Name)
	b.WriteString("{\n")
	for _, statement := range f.Statements {
		statement.Render(b)
	}
	b.WriteString("}\n")
}

func (f StructField) Render(b strings.Builder) {}

func (t Type) RenderDeclaration(b strings.Builder, name string) {

}

func (s Statement) Render(b strings.Builder) {}
