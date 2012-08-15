package gvm


var ClassTable = make(map[string]*Class)
// var fieldTable map[string]*Object

type Class struct {
    StaticFields map[string]*Object
    Methods      map[string]Method
}

type Object struct {
    ClassName string
    Native    interface{}
    Fields    map[string]*Object
}

type Method interface {
    Invoke(recv *Object, args []*Object) (bool, interface{})
    GetArgCount() int
}

func NewClass() *Class {
    c := new(Class)
    c.StaticFields = make(map[string]*Object)
    c.Methods      = make(map[string]Method )
    return c
}

func CT(name string) *Class {
    return ClassTable[name]
}