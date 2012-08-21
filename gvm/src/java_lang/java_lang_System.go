package java_lang

import . "gvm"
import "os"

func Init(ct map[string]*Class) {
    java_lang_System := NewClass()
    java_lang_System.StaticFields["out"] = &Object{ClassName: "java/io/PrintStream", Native: os.Stdout}

    ct["java/lang/System"] = java_lang_System
}