package main

import (
    "fmt"
    "os"
    . "gvm"

    "java_io"
    "java_lang"
)

func readSize(f *os.File) {
    // state, _ := f.Stat()
    // _debugf("size %d bytes\n", state.Size())
}

func readFile(fileClass string, cf *ClassFile) {
    f, err := os.Open(fileClass)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
    defer f.Close()
    readSize(f)

    r := NewClassReader(f, cf)
    r.ReadMagic()
    r.ReadVersion()
    r.ReadConstantPool()
    r.ReadFlag()
    r.ReadClass()
    r.ReadInterface()
    r.ReadField()
    r.ReadMethod()
    //d.readAttribute()
}

func main() {

    LoggingLevel = INFO

    cf := new(ClassFile)

    if len(os.Args) == 1 {
        fmt.Println("please input fileName !!!")
    } else {
        fileName  := os.Args[1]
        fileClass := fileName + ".class"
        // _debugf("  ClassFile: \"%s\"; ", fileClass)
        readFile(fileClass, cf)

        java_lang.Init(ClassTable)
        java_io.Init(ClassTable)

        ca := FindMethod(ACC_PUBLIC | ACC_STATIC, "main([Ljava/lang/String;)V", cf)
        Interpret(ca, cf.ConstantPool())
    }

}
