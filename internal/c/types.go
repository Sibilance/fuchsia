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
	Type Type `yaml:",inline"`
}

type Declaration struct {
	Name       string
	IsStatic   bool
	IsExtern   bool
	IsVolatile bool
	Type       Type `yaml:",inline"`
}

type FunctionDefinition struct {
	Name       string
	IsStatic   bool
	IsInline   bool
	Type       FunctionType `yaml:",inline"`
	Statements []Statement
}

type Statement struct{}

type Type struct {
	NamedType      string
	PointerType    *PointerType
	ArrayType      *ArrayType
	SizedArrayType *SizedArrayType
	FunctionType   *FunctionType
}

type PointerType struct {
	TargetType Type `yaml:",inline"`
}

type ArrayType struct {
	ItemType Type `yaml:",inline"`
}

type SizedArrayType struct {
	Length   int
	ItemType Type `yaml:",inline"`
}

type FunctionType struct {
	Arguments  []FunctionArgument
	ReturnType Type `yaml:",inline"`
}

type FunctionArgument struct {
	Name string
	Type Type `yaml:",inline"`
}
