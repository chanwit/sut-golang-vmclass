package java_io

import "io"
import "fmt"
import . "gvm"

type println__Ljava_lang_String__V struct {
    native bool
}

type println__I__V struct {
    native bool
}

func (m *println__Ljava_lang_String__V) Invoke(recv *Object, args []interface{}) (void bool, ret interface{}) {
    fmt.Fprintln(recv.Native.(io.Writer), args[0].(*Object).Native)
    return true, nil
}

func (m *println__Ljava_lang_String__V) GetArgCount() int { return 1 }

func (m *println__I__V) Invoke(recv *Object, args []interface{}) (void bool, ret interface{}) {
    fmt.Fprintln(recv.Native.(io.Writer), args[0])
    return true, nil
}

func (m *println__I__V) GetArgCount() int { return 1 }

func Init(ct map[string]*Class) {
    java_io_PrintStream := NewClass()
    java_io_PrintStream.Methods["println(Ljava/lang/String;)V"] = &println__Ljava_lang_String__V{native: true}
    java_io_PrintStream.Methods["println(I)V"] = &println__I__V{native: true}

    ct["java/io/PrintStream"] = java_io_PrintStream
}
