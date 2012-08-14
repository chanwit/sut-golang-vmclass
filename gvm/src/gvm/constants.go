package gvm

const (
    CONSTANT_Class              =   7
    CONSTANT_Fieldref           =   9
    CONSTANT_Methodref          =   10
    CONSTANT_InterfaceMethodref =   11
    CONSTANT_String             =   8
    CONSTANT_Integer            =   3
    CONSTANT_Float              =   4
    CONSTANT_Long               =   5
    CONSTANT_Double             =   6
    CONSTANT_NameAndType        =   12
    CONSTANT_Utf8               =   1
    CONSTANT_MethodHandle       =   15
    CONSTANT_MethodType         =   16
    CONSTANT_InvokeDynamic      =   18
)

const (
    ACC_PUBLIC    = 0x0001
    ACC_PRIVATE   = 0x0002
    ACC_PROTECTED = 0x0004
    ACC_STATIC    = 0x0008
    ACC_FINAL     = 0x0010
    ACC_SUPER     = 0x0020
    ACC_VOLATILE  = 0x0040
    ACC_TRANSIENT = 0x0080
    ACC_INTERFACE = 0x0200
    ACC_ABSTRACT  = 0x0400
    ACC_SYNTHETIC = 0x1000
    ACC_ENUM      = 0x4000
)