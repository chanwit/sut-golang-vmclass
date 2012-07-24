package main

import ( "fmt"
		"os"
	//	"io"	
		"encoding/binary"
)
type classfile struct{
	magic 								uint32
	minor_version 							uint16
	major_version 							uint16
	constant_pool_count						uint16
	constant_pool							[]cp_info
	access_flags							uint16
	this_class							uint16
	super_class							uint16
	interfaces_count						uint16
	interfaces							[]uint16
	fields_count							uint16
	fields								[]field_info
	methods_count							uint16
	//methods							[]method_info
	attributes_count						uint16
	//attributes[attributes_count]					*attributes_info
}	

const (
    ACC_PUBLIC      = 0x0001
    ACC_FINAL       = 0x0010
    ACC_SUPER       = 0x0020
    ACC_INTERFACE   = 0x0200
    ACC_ABSTRACT    = 0x0400
)
const(		//ACC_PUBLIC 	= 0x0001 	
		ACC_PRIVATE 	= 0x0002 	
		ACC_PROTECTED 	= 0x0004 	
		ACC_STATIC 	= 0x0008 
		//ACC_FINAL 	= 0x0010 	
		ACC_VOLATILE 	= 0x0040 	 
		ACC_TRANSIENT 	= 0x0080 	
		ACC_SYNTHETIC 	= 0x1000 	
		ACC_ENUM 	= 0x4000
)
type cp_info struct{
	tag	uint8
	info 	[]uint8
}
type field_info struct {
    access_flags 	uint16
    name_index      	uint16
    descriptor_index	uint16
    attributes_count	uint16
    attributes		[]attribute_info 	//[attributes_count]
}
type attribute_info struct{
    attribute_name_index	uint16
    attribute_length		uint32
    info			[]uint8	//[attribute_length]
}

type CONSTANT_Class_info struct{
	tag		uint8
	name_index	uint16
}

type CONSTANT_Fieldref_info struct{
    tag 			uint8
    class_index			uint16
    name_and_type_index		uint16
}
type CONSTANT_Methodref_info struct{
    tag				uint8
    class_index			uint16
    name_and_type_index	uint16
}

type CONSTANT_InterfaceMethodref_info struct{
    tag				uint8
    class_index			uint16
    name_and_type_index		uint16
}

type CONSTANT_String_info struct{
    tag			uint8
    string_index	uint16
}

type CONSTANT_Integer_info struct{
    tag		uint8
    bytes	uint32
}

type CONSTANT_Float_info struct {
    tag		uint8
   	bytes	uint32
}

type CONSTANT_Long_info struct{
    tag			uint8
    high_bytes 	uint32
    low_bytes	uint32
}

type CONSTANT_Double_info struct{
    tag		uint8
    high_bytes	uint32
    low_bytes	uint32
}

type CONSTANT_NameAndType_info struct{
    tag			uint8
    name_index 		uint16
    descriptor_index	uint16
}

type CONSTANT_Utf8_info struct{
    tag			uint8
    length		uint16
    bytes 		[]uint8
}

type CONSTANT_MethodHandle_info struct{
    tag				uint8
    reference_kind		uint8
    reference_index		uint16
}
type CONSTANT_MethodType_info struct{
    tag			uint8
    descriptor_index 	uint16
}
type CONSTANT_InvokeDynamic_info struct{
    tag					uint8
    bootstrap_method_attr_index		uint16
    name_and_type_index			uint16
}


const (		CONSTANT_Class				=	7
		CONSTANT_Fieldref			=	9
		CONSTANT_Methodref			=	10
		CONSTANT_InterfaceMethodref		=	11
		CONSTANT_String				=	8
		CONSTANT_Integer			=	3
		CONSTANT_Float				=	4
		CONSTANT_Long				=	5
		CONSTANT_Double				=	6
		CONSTANT_NameAndType			=	12
		CONSTANT_Utf8				=	1
		CONSTANT_MethodHandle			=	15
		CONSTANT_MethodType			=	16
		CONSTANT_InvokeDynamic			=	18
)

func readMagic(file *os.File,cf classfile) {
    binary.Read(file, binary.BigEndian, &(cf.magic))
     fmt.Printf("Magic co*os.Filede = %x\n", cf.magic)
}

func readVersion(file *os.File,cf classfile) {
    binary.Read(file, binary.BigEndian, &(cf.minor_version))
    binary.Read(file, binary.BigEndian, &(cf.major_version))
    fmt.Printf("Version = %d.%d\n", cf.major_version,cf.minor_version)
}

func readConstantPool(file *os.File,cf *classfile) {
	binary.Read(file,binary.BigEndian,&(cf.constant_pool_count))
	fmt.Printf("cp count = %d\n", cf.constant_pool_count)

	cf.constant_pool = make([]cp_info, cf.constant_pool_count)
	for i:= uint16(1); i<cf.constant_pool_count ;i++{
		
		var tag uint8
		binary.Read(file,binary.BigEndian,&(tag))
			switch tag { 
				case CONSTANT_Class: fallthrough
				case CONSTANT_MethodType: fallthrough
				case CONSTANT_String:
						info := make([]byte, 2)
						binary.Read(file,binary.BigEndian,info)
						cf.constant_pool[i]=cp_info{tag:tag, info:info }
				case CONSTANT_Fieldref: fallthrough
				case CONSTANT_Methodref: fallthrough
				case CONSTANT_InterfaceMethodref: fallthrough
				case CONSTANT_Integer: fallthrough
				case CONSTANT_Float: fallthrough
				case CONSTANT_NameAndType:
						info := make([]byte, 4)
						binary.Read(file,binary.BigEndian,info)
						cf.constant_pool[i]=cp_info{tag:tag, info:info }
				case CONSTANT_Long:	fallthrough
				case CONSTANT_Double:
						info := make([]byte, 8)
						binary.Read(file,binary.BigEndian,info)
						cf.constant_pool[i]=cp_info{tag:tag, info:info }
				
				case CONSTANT_Utf8:
						var length uint16
						binary.Read(file,binary.BigEndian,&(length))
						info := make([]byte, length+2)

						binary.BigEndian.PutUint16(info[0:2], length)

						binary.Read(file,binary.BigEndian,info[2:])
						cf.constant_pool[i]=cp_info{tag:tag, info:info }
						fmt.Printf("%d %s\n",i,info[2:])


				case CONSTANT_MethodHandle:
						info := make([]byte, 3)
						binary.Read(file,binary.BigEndian,info)
						cf.constant_pool[i]=cp_info{tag:tag, info:info }
				
				case CONSTANT_InvokeDynamic:
						info := make([]byte, 4)
						binary.Read(file,binary.BigEndian,info)
						cf.constant_pool[i]=cp_info{tag:tag, info:info }
		}
	}
	
}
func readFlag(file *os.File, cf classfile) {
	binary.Read(file, binary.BigEndian, &(cf.access_flags))
	if cf.access_flags & ACC_PUBLIC == ACC_PUBLIC {
		fmt.Println("ACC_PUBLIC")
	}
	if cf.access_flags & ACC_ABSTRACT == ACC_ABSTRACT{
		fmt.Println("ACC_ABSTRACT")
	}
	if cf.access_flags & ACC_SUPER == ACC_SUPER{
		fmt.Println("ACC_SUPER")
	}

}
func readThis(file *os.File, cf classfile) {
	binary.Read(file,binary.BigEndian,&(cf.this_class))

	thisc := cf.constant_pool[cf.this_class]
	fmt.Println(string(cf.constant_pool[(binary.BigEndian.Uint16(thisc.info))].info[2:]))

	binary.Read(file,binary.BigEndian,&(cf.super_class))

	thiss := cf.constant_pool[cf.super_class]
	fmt.Println(string(cf.constant_pool[(binary.BigEndian.Uint16(thiss.info))].info[2:]))

}
func readInterface(file *os.File, cf *classfile) {
	binary.Read(file,binary.BigEndian,&(cf.interfaces_count))
	fmt.Println(cf.interfaces_count)
	
	
	cf.interfaces = make([]uint16 , cf.interfaces_count)
	

	for i:=uint16(0); i<cf.interfaces_count ;i++{
		binary.Read(file,binary.BigEndian,&(cf.interfaces[i]))
		inter := cf.constant_pool[cf.interfaces[i]]
		fmt.Printf("interface = ")
		fmt.Println(string(cf.constant_pool[(binary.BigEndian.Uint16(inter.info))].info[2:]))


	}
}

func readField(file *os.File, cf classfile) {
	binary.Read(file,binary.BigEndian,&(cf.fields_count))
	fmt.Println(cf.fields_count)
}


func checkSize(f *os.File) (){
	state,_:=f.Stat()
	fmt.Printf("size = %d bytes\n", state.Size())
}


func openfile(filename string,cf classfile) {
	f , err := os.Open(filename) // For read file.
		if err != nil {
			fmt.Printf("%v", err)	
		}
	defer f.Close()

	checkSize(f)
    readMagic(f,cf)
    readVersion(f,cf)
    readConstantPool(f, &cf)
    readFlag(f,cf)
    readThis(f,cf)
    readInterface(f, &cf)
    readField(f, cf)
}


func main() {
var (
		filename string	
	 	cf classfile
	)
	
		if len(os.Args)==1 {
			filename ="file not found"
		}else{
			filename=os.Args[1]
		} 
 openfile(filename,cf)
}



