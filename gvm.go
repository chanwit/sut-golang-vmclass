
package main

import ( "fmt"
		"os"
		"io"	
		"encoding/binary"
		//"bytes"
)
type classfile struct{
	magic 									uint32
	minor_version 							uint16
	major_version 							uint16
	constant_pool_count						uint16
	constant_pool							[]cp_info
	access_flags							uint16
	this_class								uint16
	super_class								uint16
	interfaces_count						uint16
	interfaces								[]uint16
	fields_count							uint16
	fields									[]field_info
	method_count							uint16
	method									[]method_info
	attributes_count						uint16
	attributes								[]attribute_info
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
	tag		uint8
	info 	[]uint8
}
type field_info struct{
    access_flags 		uint16
    name_index      	uint16
    descriptor_index	uint16
    attributes_count	uint16
    attributes			[]attribute_info 	//[attributes_count]
}
type constantValue_attribute struct{
    attribute_name_index	uint16
    attribute_length		uint32
    constantvalue_index		uint16
}
type method_info struct {
	access_flags		uint16
	name_index 			uint16
	descriptor_index 	uint16
	attributes_count 	uint16
	attributes 			[]attribute_info
}
type code_attribute struct{
    attribute_name_index	uint16
    attribute_length		uint32
    max_stack				uint16
    max_locals				uint16
    code_length				uint32
    code					[]uint8
 	exception_table_length 	uint16
    exception_tables		[]exception_table
    attributes_count 		uint16
    attributes 				[]attribute_info
}
type exception_table struct{
	start_pc	uint16
    end_pc		uint16
    handler_pc	uint16
    catch_type	uint16
}
type attribute_info struct{
    attribute_name_index	uint16
    attribute_length		uint32
    info					[]uint8	//[attribute_length]
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
type decoder struct{
	file 	io.Reader
	bo		binary.ByteOrder
	cf 		*classfile
}

func (d *decoder) readMagic() {

    binary.Read(d.file, d.bo, &(d.cf.magic))
    fmt.Printf("Magic code = %x\n", d.cf.magic)
}

func (d *decoder) readVersion() {
    binary.Read(d.file, d.bo, &(d.cf.minor_version))
    binary.Read(d.file, d.bo, &(d.cf.major_version))
    fmt.Printf("Version = %d.%d\n", d.cf.major_version,d.cf.minor_version)
}

func (d *decoder) readConstantPool() {
	binary.Read(d.file,d.bo,&(d.cf.constant_pool_count))
	fmt.Printf("cp count = %d\n", d.cf.constant_pool_count)

	d.cf.constant_pool = make([]cp_info, d.cf.constant_pool_count)
	for i:= uint16(1); i<d.cf.constant_pool_count ;i++{
		
		var tag uint8
		binary.Read(d.file,d.bo,&(tag))
			switch tag { 
			case CONSTANT_Class :fallthrough
			case CONSTANT_String :fallthrough
			case CONSTANT_MethodType :
				info := make([]byte, 2)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }

			case CONSTANT_MethodHandle :
				info := make([]byte, 3)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }

			case CONSTANT_Integer :fallthrough
			case CONSTANT_Float :fallthrough
			case CONSTANT_NameAndType :fallthrough
			case CONSTANT_Fieldref :fallthrough
			case CONSTANT_Methodref :fallthrough
			case CONSTANT_InterfaceMethodref :fallthrough
			case CONSTANT_InvokeDynamic :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }

			case CONSTANT_Long :fallthrough
			case CONSTANT_Double :
				info := make([]byte, 8)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
		
			case CONSTANT_Utf8 :
				var length uint16
				binary.Read(d.file, d.bo, &(length))
				info := make([]byte, 2 + length)
				binary.BigEndian.PutUint16(info[0:2], length)
				binary.Read(d.file, d.bo, info[2:])
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
				fmt.Printf("#%d\t %s\n",i,info[2:])
		}
	}
	
}
func (d *decoder) readFlag() {
	binary.Read(d.file, d.bo, &(d.cf.access_flags))
	if d.cf.access_flags & ACC_PUBLIC == ACC_PUBLIC {
		fmt.Print("ACC_PUBLIC ")
	}
	if d.cf.access_flags & ACC_ABSTRACT == ACC_ABSTRACT{
		fmt.Print("ACC_ABSTRACT ")
	}
	if d.cf.access_flags & ACC_SUPER == ACC_SUPER{
		fmt.Print("ACC_SUPER ")
	}

}
func (d *decoder) readThis() {
	binary.Read(d.file,d.bo,&(d.cf.this_class))

	thisc := d.cf.constant_pool[d.cf.this_class]
	fmt.Println("\nthis_class is "+string(d.cf.constant_pool[(d.bo.Uint16(thisc.info))].info[2:]))

	binary.Read(d.file,d.bo,&(d.cf.super_class))
	thiss := d.cf.constant_pool[d.cf.super_class]
	fmt.Println("super_class is "+string(d.cf.constant_pool[(d.bo.Uint16(thiss.info))].info[2:]))

}
func (d *decoder) readInterface() {
	binary.Read(d.file,d.bo,&(d.cf.interfaces_count))
	fmt.Println(d.cf.interfaces_count)
	
	
	d.cf.interfaces = make([]uint16 , d.cf.interfaces_count)
	

	for i:=uint16(0); i<d.cf.interfaces_count ;i++{
		binary.Read(d.file,d.bo,&(d.cf.interfaces[i]))
		inter := d.cf.constant_pool[d.cf.interfaces[i]]
		fmt.Printf("interface = ")
		fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(inter.info))].info[2:]))


	}
}

func (d *decoder) readField() {
		binary.Read(d.file, d.bo, &(d.cf.fields_count))
		fmt.Printf("field count : %d\n", d.cf.fields_count)
		d.cf.fields = make([]field_info, d.cf.fields_count)
		for i := uint16(0); i < d.cf.fields_count; i++ {
			var fi field_info
				binary.Read(d.file, d.bo, &fi.access_flags)
					
				binary.Read(d.file, d.bo, &fi.name_index)
				binary.Read(d.file, d.bo, &fi.descriptor_index)
				binary.Read(d.file, d.bo, &fi.attributes_count)
			d.cf.fields[i]=field_info{access_flags:fi.access_flags,name_index:fi.name_index,descriptor_index:fi.descriptor_index,attributes_count:fi.attributes_count}
			

			fi.attributes = make([]attribute_info, fi.attributes_count)
					for j := uint16(0); j < fi.attributes_count; j++ {
						var name_index uint16
						var length uint32
						binary.Read(d.file, d.bo, &name_index)
						binary.Read(d.file, d.bo, &length)
						info := make([]uint8, length)
						binary.Read(d.file, d.bo, &info)
					}
/*
			fi.attributes = make([]constantValue_attribute, fi.attributes_count)
			for k := uint16(0); k < fi.attributes_count; k++ {
				var name_index uint16
				var length uint32
				var constantvalue_index uint16
				binary.Read(d.file, d.bo, &name_index)
				binary.Read(d.file, d.bo, &length)
				binary.Read(d.file, d.bo, &constantvalue_index)
				fi.attributes[k]=constantValue_attribute{attribute_name_index:name_index, attribute_length:length,constantvalue_index:constantvalue_index}
			}
		}
		
		cp := d.cf.constant_pool
    	for i := uint16(0); i < d.cf.fields_count; i++ {
       		fi := d.cf.fields[i]
       		cp1 := cp[fi.name_index]
       		cp2 := cp[fi.descriptor_index]
       				if fi.access_flags & ACC_PUBLIC == ACC_PUBLIC {
						fmt.Print("public ")
					}
					if fi.access_flags & ACC_PRIVATE == ACC_PRIVATE{
						fmt.Print("private ")
					}
					if fi.access_flags & ACC_PROTECTED == ACC_PROTECTED{
						fmt.Print("protected ")
					}
       		fmt.Println( string(cp2.info[2:]), string(cp1.info[2:]))
       		
       	}*/
}
}
func (d *decoder) readMethod() {
	binary.Read(d.file, d.bo, &(d.cf.method_count))
	fmt.Printf("method count : %d\n", d.cf.method_count)
	d.cf.method = make([]method_info, d.cf.method_count)
	for i := uint16(0); i < d.cf.method_count; i++ {
		var mi method_info
		binary.Read(d.file, d.bo, &mi.access_flags)
		binary.Read(d.file, d.bo, &mi.name_index)
		binary.Read(d.file, d.bo, &mi.descriptor_index)
		binary.Read(d.file, d.bo, &mi.attributes_count)
		fmt.Println("method detail: ",mi.access_flags, mi.name_index, mi.descriptor_index, mi.attributes_count)
		d.cf.method[i]=method_info{access_flags:mi.access_flags, name_index:mi.name_index, descriptor_index:mi.descriptor_index, attributes_count:mi.attributes_count}			
		/*	metname := d.cf.constant_pool[mi.name_index]
			metdes := d.cf.constant_pool[mi.descriptor_index]
			if mi.access_flags & ACC_PUBLIC == ACC_PUBLIC {
						fmt.Print("public ")
					}
					if mi.access_flags & ACC_PRIVATE == ACC_PRIVATE{
						fmt.Print("private ")
					}
					if mi.access_flags & ACC_PROTECTED == ACC_PROTECTED{
						fmt.Print("protected ")
					}
			fmt.Println(string(metname.info[2:]),string(metdes.info[2:])) 
		*/
			mi.attributes = make([]attribute_info, mi.attributes_count)
					for j := uint16(0); j < mi.attributes_count; j++ {
						var name_index uint16
						var length uint32
						var ca code_attribute
						binary.Read(d.file, d.bo, &name_index)
						binary.Read(d.file, d.bo, &length)
							info := make([]uint8, length)
						binary.Read(d.file, d.bo, info)
					nameCheck:=d.cf.constant_pool[name_index]
					fmt.Println(string(nameCheck.info[2:]))
						if string(nameCheck.info[2:]) == "Code"{
							ca.attribute_name_index=name_index
							ca.attribute_length=length
							ca.max_stack=d.bo.Uint16(info[0:2])
							ca.max_locals=d.bo.Uint16(info[2:4])
							ca.code_length=d.bo.Uint32(info[4:8])

							ca.code=info[8:8+ca.code_length]
							fmt.Printf("%x\n",ca.code)



			/*				ca := make([]code_attribute, length)
				for k := uint16(0); k < ca.attributes_count; k++ {
					var (
						attribute_name_index	uint16
    					attribute_length		uint32
    					max_stack				uint16
    					max_locals				uint16
    					code_length				uint32
  				  		exception_table_length 	uint16
   				 		attributes_count 		uint16
   				 		//attributes 				[]attribute_info 
   				 	) 
					binary.Read(d.file, d.bo, &attribute_name_index)
					binary.Read(d.file, d.bo, &attribute_length)
					binary.Read(d.file, d.bo, &max_stack)
					binary.Read(d.file, d.bo, &max_locals)
					binary.Read(d.file, d.bo, &code_length)
						codes:=make([]uint8,code_length)
					binary.Read(d.file, d.bo, codes)
					binary.Read(d.file, d.bo, &exception_table_length)
						excep := make([]exception_table,exception_table_length)//array
						binary.Read(d.file, d.bo, excep)
					binary.Read(d.file, d.bo, &attributes_count)
						info := make([]attribute_info, attributes_count)
						binary.Read(d.file, d.bo, info)
					ca.attributes[k]=code_attribute{attribute_name_index:attribute_name_index,attribute_length:attribute_length,max_stack:max_stack,max_locals:max_locals,code_length:code_length,code:codes,exception_table_length:exception_table_length,exception_tables:excep,attributes_count:attributes_count,attributes:info}
					fmt.Println("Code Length: ",code_length)
					
					fmt.Printf("Code: %x\n",codes)

				} */
						}

					}	
				/*	
				mi.attributes = make([]code_attribute, mi.attributes_count)
				for j := uint16(0); j < mi.attributes_count; j++ {
					var (
						//ca 						code_attribute
						attribute_name_index	uint16
    					attribute_length		uint32
    					max_stack				uint16
    					max_locals				uint16
    					code_length				uint32
  				  		exception_table_length 	uint16
    					//exception_table
   				 		attributes_count 		uint16
   				 		//attributes 				[]attribute_info 
   				 	) 
					binary.Read(d.file, d.bo, &attribute_name_index)
					binary.Read(d.file, d.bo, &attribute_length)
					binary.Read(d.file, d.bo, &max_stack)
					binary.Read(d.file, d.bo, &max_locals)
					binary.Read(d.file, d.bo, &code_length)
						codes:=make([]uint8,code_length)
					binary.Read(d.file, d.bo, codes)
					binary.Read(d.file, d.bo, &exception_table_length)
						excep := make([]exception_table,exception_table_length)//array
						binary.Read(d.file, d.bo, excep)
					binary.Read(d.file, d.bo, &attributes_count)
						info := make([]attribute_info, attributes_count)
						binary.Read(d.file, d.bo, info)
					mi.attributes[j]=code_attribute{attribute_name_index:attribute_name_index,attribute_length:attribute_length,max_stack:max_stack,max_locals:max_locals,code_length:code_length,code:codes,exception_table_length:exception_table_length,exception_tables:excep,attributes_count:attributes_count,attributes:info}
					fmt.Println("Code Length: ",code_length)
					
					fmt.Printf("Code: %x\n",codes)

				} 	
				*/
				/* Nut Method
				mi.attributes = make([]attribute_info, mi.attributes_count)
					for j := uint16(0); j < mi.attributes_count; j++ {
						var name_index uint16
						var length uint32
						binary.Read(d.file, d.bo, &name_index)
						binary.Read(d.file, d.bo, &length)
						info := make([]uint8, length)
						binary.Read(d.file, d.bo, &info)
					}
				end of nut's method*/ 

			
	}	/*
		for k:=uint16(0) ;k<d.cf.method_count ;k++{
			mi:=d.cf.method[k]
			metname := d.cf.constant_pool[mi.name_index]
			metdes := d.cf.constant_pool[mi.descriptor_index]
			if mi.access_flags & ACC_PUBLIC == ACC_PUBLIC {
						fmt.Print("public ")
					}
					if mi.access_flags & ACC_PRIVATE == ACC_PRIVATE{
						fmt.Print("private ")
					}
					if mi.access_flags & ACC_PROTECTED == ACC_PROTECTED{
						fmt.Print("protected ")
					}
			fmt.Println(string(metname.info[2:]),string(metdes.info[2:])) 
       	}*/
}    	
func (d *decoder) readAttribute() {
		binary.Read(d.file, d.bo, &(d.cf.attributes_count))
		fmt.Printf("attribute count : %d\n", d.cf.attributes_count)
		d.cf.attributes = make([]attribute_info, d.cf.attributes_count)
		for i := uint16(0); i < d.cf.attributes_count; i++ {
			var name_index uint16
			var length uint32
			binary.Read(d.file, d.bo, &name_index)
			binary.Read(d.file, d.bo, &length)
				info := make([]uint8, length)
				binary.Read(d.file, d.bo, &info)
					d.cf.attributes[i] = attribute_info{ attribute_name_index:name_index, attribute_length:length, info:info }
					att := d.cf.constant_pool[name_index]
					fmt.Println(att, string(att.info[2:]))
		}
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
	d := decoder{file:f, bo:binary.BigEndian ,cf: &cf}

	checkSize(f)
    d.readMagic()
    d.readVersion()
    d.readConstantPool()
    d.readFlag()
    d.readThis()
    d.readInterface()
    d.readField()
    d.readMethod()
    d.readAttribute()
    
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

