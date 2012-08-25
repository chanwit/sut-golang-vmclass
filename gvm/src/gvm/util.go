package gvm

func flagsOn(value uint16, flag uint16) bool {
    return (value & flag) == flag
}

// findMethod ACC_STATIC | ACC_PUBLIC main ([Ljava/lang/String;)V
func FindMethod(flags uint16, signature string, cf *ClassFile) (ca code_attribute) {
    //_debugf("\nFind method %s:\n", signature)
    for i := uint16(0); i < cf.method_count; i++ {
        method    := cf.methods[i]
        nameIndex := cf.constant_pool[method.name_index]
        descIndex := cf.constant_pool[method.descriptor_index]

        methodSignature := string(nameIndex.info[2:]) + string(descIndex.info[2:])
        //_debugf("\nChecking signature %s:\n", methodSignature)
        accessFlags := method.access_flags

        if methodSignature == signature && flagsOn(accessFlags, flags) {
            //_debugf("\nFound signature %s:\n", methodSignature)
            for j := uint16(0); j < method.attributes_count; j++ {
                attribute := method.attributes[j]
                niMain := cf.constant_pool[attribute.attribute_name_index]
                if string(niMain.info[2:]) == "Code" {
                    return attribute
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
