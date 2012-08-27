package java_lang

import . "gvm"
import "os"
import . "strconv"
type init____V struct{
    native bool
}
func (m *init____V) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
      recv.Native=""
    return false, recv
}
func (m *init____V) GetArgCount() int { return 1 }
//=======================================================================================================
type append__I__Ljava_lang_StringBuilder struct{
    native bool
}
func (m *append__I__Ljava_lang_StringBuilder) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
      recv.Native=recv.Native.(string)+Itoa((args[0].Native).(int))
    return false, recv
}
func (m *append__I__Ljava_lang_StringBuilder) GetArgCount() int { return 1 }
//=======================================================================================================
type append__Ljava_lang_String__Ljava_lang_StringBuilder struct {
    native bool
}
func (m *append__Ljava_lang_String__Ljava_lang_StringBuilder) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
        recv.Native=recv.Native.(string)+(args[0].Native).(string) //
    return false, recv
}
func (m *append__Ljava_lang_String__Ljava_lang_StringBuilder) GetArgCount() int { return 1 }
//=======================================================================================================
type toString____Ljava_lang_String struct {
    native bool
}
func (m *toString____Ljava_lang_String) Invoke(recv *Object, args []*Object) (void bool, ret interface{}) {
    ret = &Object{ClassName:"java/lang/String", Native:recv.Native.(string)}
    return false,ret
}
func (m *toString____Ljava_lang_String) GetArgCount() int { return 0 } //return 0 by Sarawut.
//=======================================================================================================
func Init(ct map[string]*Class) {
    java_lang_System := NewClass()
    java_lang_System.StaticFields["out"] = &Object{ClassName: "java/io/PrintStream", Native: os.Stdout}

    java_lang_StringBuilder := NewClass()
    java_lang_StringBuilder.Methods["<init>()V"]=&init____V{native:true}
    //java_lang_StringBuilder.StaticFields["<init>()V"] = &Object{Native: ""}
    java_lang_StringBuilder.Methods["append(Ljava/lang/String;)Ljava/lang/StringBuilder;"] = &append__Ljava_lang_String__Ljava_lang_StringBuilder{native:true}
    java_lang_StringBuilder.Methods["append(I)Ljava/lang/StringBuilder;"] = &append__I__Ljava_lang_StringBuilder{native:true}
    java_lang_StringBuilder.Methods["toString()Ljava/lang/String;"] = &toString____Ljava_lang_String{native:true}
    ct["java/lang/System"] = java_lang_System
    ct["java/lang/StringBuilder"] = java_lang_StringBuilder
}