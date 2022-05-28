package memutil

import (
	"reflect"
	"unsafe"
)

//TODO: rewrite like all of this when generics are released

/*
func (process Process) readMemory(address uint64, size uint64) []byte {
	value := make([]byte, size)
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(size))
	return value
} */

func (process Process) ReadF32(address uint64) float32 {
	var value float32
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}
func (process Process) ReadF64(address uint64) float64 {
	var value float64
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}

func (process Process) ReadI8(address uint64) int8 {
	var value int8
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}
func (process Process) ReadI32(address uint64) int32 {
	var value int32
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}
func (process Process) ReadI64(address uint64) int64 {
	var value int64
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}

func (process Process) ReadU8(address uint64) uint8 {
	var value uint8
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}
func (process Process) ReadU32(address uint64) uint32 {
	var value uint32
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}
func (process Process) ReadU64(address uint64) uint64 {
	var value uint64
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value
}

func (process Process) MultiLevelPtr32(offsets []uint64) uint64 {
	var pointer uint64 = 0
	for _, val := range offsets {
		pointer += val
		pointer = uint64(process.ReadU32(pointer))
		if pointer == 0 {
			return 0
		}
	}
	return pointer
}
func (process Process) MultiLevelPtr64(offsets []uint64) uint64 {
	var pointer uint64 = 0
	for _, val := range offsets {
		pointer += val
		pointer = process.ReadU64(pointer)
		if pointer == 0 {
			return 0
		}
	}
	return pointer
}

//TODO: implement ReadStr()
/*
func (process Process) ReadBytes(address uint64, amount uint64) []byte {
	return process.readMemory(address, amount)
} */

func WriteVal[V any](process Process, address uint64, value V) {
	str, isString := any(value).(string)
	if isString {
		process.writeStr(address, str)
		return
	}

	//For primitive types
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}

func (process Process) writeStr(address uint64, value string) {
	var oldprot uint32
	actualPointer := ((*reflect.SliceHeader)(unsafe.Pointer(&value))).Data //Gets pointer to first element not slice struct
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(len(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(actualPointer)), uintptr(len(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(len(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}

/*
In theory this should work

func (process Process) writeSlice(address uint64, value []any) {
	length := len(value)
	if length == 0 {
		return
	}

	var oldprot uint32
	actualPointer := ((*reflect.SliceHeader)(unsafe.Pointer(&value))).Data //Gets pointer to first element not slice struct

	trueSize := int(unsafe.Sizeof(value[0])) * length

	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(trueSize), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(actualPointer)), uintptr(trueSize), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(trueSize), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
*/
