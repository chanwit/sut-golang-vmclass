package gvm

var classTable map[string]*class
var fieldTable map[string]*object

type class struct {
    staticFields map[string]*object
    methods      map[string][]byte
}

type object struct {
    className string
    fields    map[string]*object
}

func ClassTableInit() {
    classTable = make(map[string]*class)

    java_lang_System := &class{}
    java_lang_System.staticFields = make(map[string]*object)
    java_lang_System.staticFields["out"] = &object{className: "java/io/PrintStream"}

    classTable["java/lang/System"] = java_lang_System
}

func CT(name string) *class {
    return classTable[name]
}