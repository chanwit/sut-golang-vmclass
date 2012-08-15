package gvm

import "io"
import "os"
import "fmt"

var classTable map[string]*class
var fieldTable map[string]*object

type class struct {
    staticFields map[string]*object
    methods      map[string]method
}

type object struct {
    class  string
    native interface{}
    fields map[string]*object
}

type method interface {
    invoke(recv *object, args []*object) (bool, interface{})
    getArgCount() int
}

type InterpretMethod struct {
}

func NewClass() *class {
    c := &class{}
    c.staticFields = make(map[string]*object)
    c.methods      = make(map[string]method)
    return c
}

type println__Ljava_lang_String__V struct {
    native bool
}

func (m *println__Ljava_lang_String__V) invoke(recv *object, args []*object) (void bool, ret interface{}) {
    fmt.Fprintln(recv.native.(io.Writer), args[0].native)
    return true, nil
}

func (m *println__Ljava_lang_String__V) getArgCount() int {
    return 1
}

func ClassTableInit() {
    classTable = make(map[string]*class)

    java_lang_System := NewClass()
    java_lang_System.staticFields["out"] = &object{class: "java/io/PrintStream", native: os.Stdout}

    java_io_PrintStream := NewClass()
    java_io_PrintStream.methods["println(Ljava/lang/String;)V"] =
        &println__Ljava_lang_String__V{native: true}

    classTable["java/lang/System"]    = java_lang_System
    classTable["java/io/PrintStream"] = java_io_PrintStream
}

func CT(name string) *class {
    return classTable[name]
}