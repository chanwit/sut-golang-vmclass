package main

import (
    "fmt"
    . "gvm"
    "os"
)

func readSize(f *os.File) {
    state, _ := f.Stat()
    fmt.Printf("size %d bytes\n", state.Size())
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
    // decoder{file: f, bo: binary.BigEndian, cf: cf}
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

    cf := new(ClassFile)

    if len(os.Args) == 1 {
        fmt.Println("please input fileName !!!")
    } else {
        fileName  := os.Args[1]
        fileClass := fileName + ".class"
        fmt.Printf("  ClassFile: \"%s\"; ", fileClass)
        readFile(fileClass, cf)
        ca := FindMethod(ACC_PUBLIC | ACC_STATIC, "main([Ljava/lang/String;)V", cf)
        Interpret(ca, cf)
    }

}
