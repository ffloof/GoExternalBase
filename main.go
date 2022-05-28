package main

import (
	"GoExternalBase/memutil"
	"fmt"
)

func main() {
	fmt.Println("Hi")
	process := memutil.AttatchProcess("SomeExecutable.exe")
	module := memutil.GetModuleInfo("SomeExecutable.exe", "SomeDllOrExe.dll")

	//Multi level pointer in a 32 bit app
	ptr := process.MultiLevelPtr32([]uint64{uint64(module.DwBase) + 0xDEADBEEF, 0xF0, 0x0, 0xCC})

	//Offset 0xF4 from base ptr
	offset_ptr := ptr + 0xF4

	//Read a float32 value from the pointer
	readfloat := process.ReadF32(offset_ptr)
	fmt.Println(readfloat)

	//Read unsigned int 64 from the pointer
	readuint := process.ReadU64(offset_ptr)
	fmt.Println(readuint)

	//Write unsigned int 64 to memory
	memutil.WriteVal(process, offset_ptr, 20)

	//Write a series of bytes or a string
	memutil.WriteVal(process, offset_ptr, "I like turtles")
}
