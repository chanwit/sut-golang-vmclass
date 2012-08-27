package java_lang

import . "gvm"
import "os"
//import "io"
//import "fmt"
import "strconv"

func Init(ct map[string]*Class) {
    java_lang_System := NewClass()
    java_lang_System.StaticFields["out"] = &Object{ClassName: "java/io/PrintStream", Native: os.Stdout}

    java_lang_StringBuilder := NewClass()
    //java_lang_StringBuilder.StaticFields["<init>()V"] = &Object{ClassName: "java/lang/StringBuilder", Native: ""}
    java_lang_StringBuilder.Methods["<init>()V"] = &init__v{native: true}
    java_lang_StringBuilder.Methods["append(Ljava/lang/String;)Ljava/lang/StringBuilder;"] = &append__Ljava_lang_String__Ljava_lang_StringBuilder{native: true}
    java_lang_StringBuilder.Methods["append(I)Ljava/lang/StringBuilder;"] = &append__I__Ljava_lang_StringBuilder{native: true}
    java_lang_StringBuilder.Methods["toString()Ljava/lang/String;"] = &toString__V__Ljava_lang_String{native: true}

    ct["java/lang/System"] = java_lang_System
    ct["java/lang/StringBuilder"] = java_lang_StringBuilder
}

//**************************************************

type append__Ljava_lang_String__Ljava_lang_StringBuilder struct {
    native bool
}

func (m *append__Ljava_lang_String__Ljava_lang_StringBuilder) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
    recv.Native = recv.Native.(string) + args[0].Native.(string)
    return false, recv
}

func (m *append__Ljava_lang_String__Ljava_lang_StringBuilder) GetArgCount() int { return 1 }

//**************************************************

type append__I__Ljava_lang_StringBuilder struct {
    native bool
}

func (m *append__I__Ljava_lang_StringBuilder) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
    recv.Native = recv.Native.(string) + strconv.Itoa(args[0].Native.(int))
    return false, recv
}

func (m *append__I__Ljava_lang_StringBuilder) GetArgCount() int { return 1 }

//**************************************************

type toString__V__Ljava_lang_String struct {
    native bool
}

func (m *toString__V__Ljava_lang_String) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
    obj := &Object{ClassName: "java/lang/String", Native: recv.Native.(string)}
    return false, obj
}

func (m *toString__V__Ljava_lang_String) GetArgCount() int { return 0 }

//**************************************************

type init__v struct {
    native bool
}

func (m *init__v) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
    recv.Native = ""
    return true, nil
}

func (m *init__v) GetArgCount() int { return 0 }