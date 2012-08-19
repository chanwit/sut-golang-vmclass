package main

import "fmt"
import "os"
import . "gvm"

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

    d := NewDecoder(f, cf)
    d.ReadMagic()
    d.ReadVersion()
    d.ReadConstantPool()
    d.ReadFlag()
    d.ReadClass()
    d.ReadInterface()
    d.ReadField()
    d.ReadMethod()
    //d.ReadAttribute()
}

func getFile(fileClass string, cf *ClassFile) {
    f, err := os.Open(fileClass)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
    defer f.Close()
    readSize(f)

    d := NewDecoder(f, cf)
    d.GetMagic()
    d.GetVersion()
    d.GetConstantPool()
    d.GetFlag()
    d.GetClass()
    d.GetInterface()
    d.GetField()
    d.GetMethod()
    //d.GetAttribute()
}

func main() {

    cf := new(ClassFile)

    if len(os.Args) == 1 {
        fmt.Println("please input fileName !!!")
    } else {

        opt := os.Args[1]
        fileClass := os.Args[2] + ".class"
        fmt.Printf("  ClassFile: \"%s\"; ", fileClass)

        if opt == "-verbose" {
            readFile(fileClass, cf)
        }else if opt == "-c" {
            getFile(fileClass, cf)
        }else{
            os.Exit(0)
        }

        ca := FindMethod("main", cf)
        Interpret(ca, cf)
    }

}
