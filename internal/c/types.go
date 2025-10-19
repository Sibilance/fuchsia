package c

type Include struct {
	IsSystem bool
	Path     string
}

type Header struct {
	Includes          []Include
	StructDefinitions []StructDefinition
	Declarations      []Declaration
}

type Program struct {
	Includes            []Include
	StructDefinitions   []StructDefinition
	Declarations        []Declaration
	FunctionDefinitions []FunctionDefinition
}

type StructDefinition struct {
	Name   string
	Fields []StructField
}

type StructField struct {
	Name string
	Type Type
}

type Declaration struct {
	Name       string
	IsStatic   bool
	IsExtern   bool
	IsVolatile bool
	Type       Type
}

type FunctionDefinition struct {
	Name       string
	IsStatic   bool
	IsInline   bool
	Type       FunctionType
	Statements []Statement
}

type Type struct {
	NamedType      *NamedType
	PointerType    *PointerType
	ArrayType      *ArrayType
	SizedArrayType *SizedArrayType
	FunctionType   *FunctionType
}

type NamedType struct {
	Name string
}

type PointerType struct {
	TargetType Type
}

type ArrayType struct {
	ItemType Type
}

type SizedArrayType struct {
	Length   int
	ItemType Type
}

type FunctionType struct {
	Arguments  []FunctionArgument
	ReturnType Type
}

type FunctionArgument struct {
	Name string
	Type Type
}

type Statement struct{}
