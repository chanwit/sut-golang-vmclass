package gvm

type ClassFile struct {
    magic                   uint32
    minor_version           uint16
    major_version           uint16
    constant_pool_count     uint16
    constant_pool           []cp_info
    access_flags            uint16
    this_class              uint16
    super_class             uint16
    interfaces_count        uint16
    interfaces              []uint16
    fields_count            uint16
    fields                  []field_info
    method_count            uint16
    methods                 []method_info
    attributes_count        uint16
    attributes              []attribute_info
}

type cp_info struct {
    tag  uint8
    info []uint8
}

type CONSTANT_Class_info struct {
    tag        uint8
    name_index uint16
}

type CONSTANT_Fieldref_info struct {
    tag                 uint8
    class_index         uint16
    name_and_type_index uint16
}

type CONSTANT_Methodref_info struct {
    tag                 uint8
    class_index         uint16
    name_and_type_index uint16
}

type CONSTANT_InterfaceMethodref_info struct {
    tag                 uint8
    class_index         uint16
    name_and_type_index uint16
}

type CONSTANT_String_info struct {
    tag          uint8
    string_index uint16
}

type CONSTANT_Integer_info struct {
    tag   uint8
    bytes uint32
}

type CONSTANT_Float_info struct {
    tag   uint8
    bytes uint32
}

type CONSTANT_Long_info struct {
    tag        uint8
    high_bytes uint32
    low_bytes  uint32
}

type CONSTANT_Double_info struct {
    tag        uint8
    high_bytes uint32
    low_bytes  uint32
}

type CONSTANT_NameAndType_info struct {
    tag              uint8
    name_index       uint16
    descriptor_index uint16
}

type CONSTANT_Utf8_info struct {
    tag    uint8
    length uint16
    bytes  []uint8
}

type CONSTANT_MethodHandle_info struct {
    tag             uint8
    reference_kind  uint8
    reference_index uint16
}

type CONSTANT_MethodType_info struct {
    tag              uint8
    descriptor_index uint16
}

type CONSTANT_InvokeDynamic_info struct {
    tag                         uint8
    bootstrap_method_attr_index uint16
    name_and_type_index         uint16
}

type field_info struct {
    access_flags     uint16
    name_index       uint16
    descriptor_index uint16
    attributes_count uint16
    attributes       []attribute_info
}

type method_info struct {
    access_flags     uint16
    name_index       uint16
    descriptor_index uint16
    attributes_count uint16
    attributes       []code_attribute
}

type attribute_info struct {
    attribute_name_index uint16
    attribute_length     uint32
    info                 []uint8
}

type code_attribute struct {
    attribute_name_index   uint16
    attribute_length       uint32
    max_stack              uint16
    max_locals             uint16
    code_length            uint32
    code                   []uint8
    exception_table_length uint16
    exception              []exception_table
    attributes_count       uint16
    line_number_table_att  []LineNumberTable_attribute
}

type LineNumberTable_attribute struct {
    attribute_name_index     uint16
    attribute_length         uint32
    line_number_table_length uint16
    line_number_tables       []line_number_table
}

type line_number_table struct {
    start_pc    uint16
    line_number uint16
}

type exception_table struct {
    start_pc   uint16
    end_pc     uint16
    handler_pc uint16
    catch_type uint16
}