package main

import "fmt"
import "os"
import "strings"
import "java_lang"
import "java_io"
import . "gvm"

func readSize(f *os.File) {
    state, _ := f.Stat()
    fmt.Println("size", state.Size(), "bytes")
}

func readFile(fileName string, cf *ClassFile) {
    f, err := os.Open(fileName)
    if err != nil {
        fmt.Println("This file cannot open.")
        os.Exit(1)
    }
    defer f.Close()
    fmt.Print("Classfile ", fileName, "; ")
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

func getFile(fileName string, cf *ClassFile) {
    f, err := os.Open(fileName)
    if err != nil {
        fmt.Println("This file cannot open.")
        os.Exit(1)
    }
    defer f.Close()
    fmt.Print("Classfile ", fileName, "; ")
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

	if len(os.Args) == 3 {
		cf        := new(ClassFile)
		option    := os.Args[1]
		fileName  := os.Args[2]

		if !strings.HasSuffix(fileName, ".class") {
			fileName = fileName + ".class"
		}

		if option == "-c" {
			getFile(fileName, cf)
		}else if option == "-verbose" {
			readFile(fileName, cf)
		}else{
			fmt.Println("-option")
			fmt.Println("    -c")
			fmt.Println("    -verbose")
			os.Exit(0)
		}

        java_lang.Init(ClassTable)
        java_io.Init(ClassTable)

		ca := FindMethod("main", cf)
		Interpret(ca, cf)

	}else{
		fmt.Println("./gvm.go -option -fileName")
		os.Exit(0)
	}

}