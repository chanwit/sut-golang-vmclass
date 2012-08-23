package gvm

import "fmt"

func FindMethod(name string, cf *ClassFile) (ca code_attribute) {
    fmt.Printf("\nFind method %s:\n", name)
    for i := uint16(0); i < cf.method_count; i++ {
        ni := cf.constant_pool[cf.methods[i].name_index]
        if string(ni.info[2:]) == name {
            for j := uint16(0); j < cf.methods[i].attributes_count; j++ {
                niMain := cf.constant_pool[cf.methods[i].attributes[j].attribute_name_index]
                if string(niMain.info[2:]) == "Code" {
                    return cf.methods[i].attributes[j]
                }
            }
        }
    }
    return
}

func u16(b []byte) uint16 {
    return uint16(b[1]) | uint16(b[0])<<8
}

func u32(b []byte) uint32 {
    return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func i32(b []byte) int {
    return int(uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24)
}