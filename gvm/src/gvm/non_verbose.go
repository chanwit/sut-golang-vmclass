package gvm

import "encoding/binary"

func (d *Decoder) GetMagic() {
    binary.Read(d.file, d.bo, &(d.cf.magic))
}

func (d *Decoder) GetVersion() {
    binary.Read(d.file, d.bo, &(d.cf.minor_version))
    binary.Read(d.file, d.bo, &(d.cf.major_version))
}

func (d *Decoder) GetConstantPool() {
    binary.Read(d.file, d.bo, &(d.cf.constant_pool_count))
    d.cf.constant_pool = make([]cp_info, d.cf.constant_pool_count)
    for i := uint16(1); i < d.cf.constant_pool_count; i++ {
        var tag uint8
        binary.Read(d.file, d.bo, &(tag))
        switch tag {
        case CONSTANT_Class:
            info := make([]byte, 2)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Fieldref:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Methodref:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_InterfaceMethodref:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_String:
            info := make([]byte, 2)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Integer:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Float:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Long:
            info := make([]byte, 8)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Double:
            info := make([]byte, 8)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_NameAndType:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Utf8:
            var length uint16
            binary.Read(d.file, d.bo, &(length))
            info := make([]byte, 2+length)
            binary.BigEndian.PutUint16(info[0:2], length)
            binary.Read(d.file, d.bo, info[2:])
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_MethodHandle:
            info := make([]byte, 3)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_MethodType:
            info := make([]byte, 2)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_InvokeDynamic:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        }
    }
}

func (d *Decoder) GetFlag() {
    binary.Read(d.file, d.bo, &(d.cf.access_flags))
    if d.cf.access_flags&ACC_PUBLIC == ACC_PUBLIC {

    }
    if d.cf.access_flags&ACC_PRIVATE == ACC_PRIVATE {

    }
    if d.cf.access_flags&ACC_PROTECTED == ACC_PROTECTED {

    }
    if d.cf.access_flags&ACC_STATIC == ACC_STATIC {

    }
    if d.cf.access_flags&ACC_FINAL == ACC_FINAL {

    }
    if d.cf.access_flags&ACC_SUPER == ACC_SUPER {

    }
    if d.cf.access_flags&ACC_VOLATILE == ACC_VOLATILE {

    }
    if d.cf.access_flags&ACC_TRANSIENT == ACC_TRANSIENT {

    }
    if d.cf.access_flags&ACC_INTERFACE == ACC_INTERFACE {

    }
    if d.cf.access_flags&ACC_ABSTRACT == ACC_ABSTRACT {

    }
    if d.cf.access_flags&ACC_SYNTHETIC == ACC_SYNTHETIC {

    }
    if d.cf.access_flags&ACC_ENUM == ACC_ENUM {

    }
}

func (d *Decoder) GetClass() {
    binary.Read(d.file, d.bo, &(d.cf.this_class))
    binary.Read(d.file, d.bo, &(d.cf.super_class))
}

func (d *Decoder) GetInterface() {
    binary.Read(d.file, d.bo, &(d.cf.interfaces_count))
    d.cf.interfaces = make([]uint16, d.cf.interfaces_count)
}

func (d *Decoder) GetField() {
    binary.Read(d.file, d.bo, &(d.cf.fields_count))
    d.cf.fields = make([]field_info, d.cf.fields_count)
    for i := uint16(0); i < d.cf.fields_count; i++ {
        var fi field_info
        binary.Read(d.file, d.bo, &fi.access_flags)
        binary.Read(d.file, d.bo, &fi.name_index)
        binary.Read(d.file, d.bo, &fi.descriptor_index)
        binary.Read(d.file, d.bo, &fi.attributes_count)
        d.cf.fields[i] = field_info{access_flags: fi.access_flags, name_index: fi.name_index, descriptor_index: fi.descriptor_index, attributes_count: fi.attributes_count}

        fi.attributes = make([]attribute_info, fi.attributes_count)
        for j := uint16(0); j < fi.attributes_count; j++ {
            var name_index uint16
            var length uint32
            binary.Read(d.file, d.bo, &name_index)
            binary.Read(d.file, d.bo, &length)
            info := make([]uint8, length)
            binary.Read(d.file, d.bo, &info)
            d.cf.fields[i].attributes[j] = attribute_info{attribute_name_index: name_index, attribute_length: length}
        }
    }
}

func (d *Decoder) GetMethod() {
    binary.Read(d.file, d.bo, &(d.cf.method_count))
    d.cf.methods = make([]method_info, d.cf.method_count)
    for i := uint16(0); i < d.cf.method_count; i++ {
        var mi method_info
        binary.Read(d.file, d.bo, &mi.access_flags)
        binary.Read(d.file, d.bo, &mi.name_index)
        binary.Read(d.file, d.bo, &mi.descriptor_index)
        binary.Read(d.file, d.bo, &mi.attributes_count)
        d.cf.methods[i] = method_info{access_flags: mi.access_flags, name_index: mi.name_index, descriptor_index: mi.descriptor_index, attributes_count: mi.attributes_count}

        d.cf.methods[i].attributes = make([]code_attribute, mi.attributes_count)
        for j := uint16(0); j < mi.attributes_count; j++ {
            var name_index uint16
            var length uint32
            binary.Read(d.file, d.bo, &name_index)
            binary.Read(d.file, d.bo, &length)
            info := make([]uint8, length)
            binary.Read(d.file, d.bo, &info)

            lookup := string(d.cf.constant_pool[name_index].info[2:])
            if lookup == "Code" {
                var ca code_attribute
                ca.attribute_name_index = name_index
                ca.attribute_length = length
                ca.max_stack = d.bo.Uint16(info[0:2])
                ca.max_locals = d.bo.Uint16(info[2:4])
                ca.code_length = d.bo.Uint32(info[4:8])
                ca.code = info[8 : 8+ca.code_length]
                d.cf.methods[i].attributes[j].code = make([]uint8, ca.code_length)
                d.cf.methods[i].attributes[j].code = ca.code

                ca.exception_table_length = d.bo.Uint16(info[8+ca.code_length : 10+ca.code_length])
                d.cf.methods[i].attributes[j].exception = make([]exception_table, ca.exception_table_length)
                for l := uint16(0); l < ca.exception_table_length; l++ {
                    var start_pc uint16
                    var end_pc uint16
                    var handler_pc uint16
                    var catch_type uint16
                    start_pc = d.bo.Uint16(info[10+ca.code_length : 12+ca.code_length])
                    end_pc = d.bo.Uint16(info[12+ca.code_length : 14+ca.code_length])
                    handler_pc = d.bo.Uint16(info[14+ca.code_length : 16+ca.code_length])
                    catch_type = d.bo.Uint16(info[16+ca.code_length : 18+ca.code_length])
                    d.cf.methods[i].attributes[j].exception[l] = exception_table{start_pc: start_pc, end_pc: end_pc, handler_pc: handler_pc, catch_type: catch_type}
                }
                index := uint16(ca.code_length) + ca.exception_table_length
                ca.attributes_count = d.bo.Uint16(info[index+10 : index+12])
                d.cf.methods[i].attributes[j].line_number_table_att = make([]LineNumberTable_attribute, ca.attributes_count)
                var lnt_a LineNumberTable_attribute
                for m := uint16(0); m < ca.attributes_count; m++ {
                    var name_index uint16
                    var length uint32
                    name_index = d.bo.Uint16(info[index+12 : index+14])
                    length = d.bo.Uint32(info[index+14 : index+18])
                    lnt_a.attribute_name_index = name_index
                    lnt_a.attribute_length = length
                    lnt_a.line_number_table_length = d.bo.Uint16(info[index+18 : index+20])
                    lnt_a.line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
                    d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
                    for o := uint16(0); o < lnt_a.line_number_table_length; o++ {
                        var start_pc uint16
                        var line_number uint16
                        start_pc = d.bo.Uint16(info[index+20+(o*4) : index+22+(o*4)])
                        line_number = d.bo.Uint16(info[index+22+(o*4) : index+24+(o*4)])

                        d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables[o] = line_number_table{start_pc: start_pc, line_number: line_number}

                    }

                    d.cf.methods[i].attributes[j].line_number_table_att[m] = LineNumberTable_attribute{attribute_name_index: lnt_a.attribute_name_index, attribute_length: lnt_a.attribute_length, line_number_table_length: lnt_a.line_number_table_length, line_number_tables: d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables}

                }

                d.cf.methods[i].attributes[j] = code_attribute{attribute_name_index: name_index, attribute_length: length, max_stack: ca.max_stack, max_locals: ca.max_locals, code_length: ca.code_length, code: d.cf.methods[i].attributes[j].code, exception_table_length: ca.exception_table_length, exception: d.cf.methods[i].attributes[j].exception, attributes_count: ca.attributes_count, line_number_table_att: d.cf.methods[i].attributes[j].line_number_table_att}

            }
        }
    }
}

func (d *Decoder) GetAttribute() {
    binary.Read(d.file, d.bo, &(d.cf.attributes_count))
    d.cf.attributes = make([]attribute_info, d.cf.attributes_count)
}