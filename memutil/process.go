package memutil

import (
	"unsafe"
	"reflect"
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

//TODO: maybe change 256 byte max hopefully theres a cleaner version of this with generics
//TODO: read str and bytes do the same thing just case string from []bytes next time
func (process Process) ReadStrFixed(address uint64, size int) string {
	var value [256]byte
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return string(value[:size])
}

func (process Process) ReadBytesFixed(address uint64, size int) []byte {
	var value [256]byte
	pReadProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	return value[:size]
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

func (process Process) WriteF32(address uint64, value float32) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
func (process Process) WriteF64(address uint64, value float64){
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}


func (process Process) WriteI8(address uint64, value int8) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
func (process Process) WriteI32(address uint64, value int32) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
func (process Process) WriteI64(address uint64, value int64) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}


func (process Process) WriteU8(address uint64, value uint8) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
func (process Process) WriteU32(address uint64, value uint32) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
func (process Process) WriteU64(address uint64, value uint64) {
	var oldprot uint32
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Sizeof(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Sizeof(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}

func (process Process) WriteStr(address uint64, value string) {
	var oldprot uint32
	actualPointer := ((*reflect.SliceHeader)(unsafe.Pointer(&value))).Data //Gets pointer to first element not slice struct
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(len(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(actualPointer)), uintptr(len(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(len(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
func (process Process) WriteBytes(address uint64, value []byte) {
	var oldprot uint32
	actualPointer := ((*reflect.SliceHeader)(unsafe.Pointer(&value))).Data //Gets pointer to first element not slice struct
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(len(value)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&oldprot)))
	pWriteProcessMemory.Call(uintptr(process.hProcess), uintptr(address), uintptr(unsafe.Pointer(actualPointer)), uintptr(len(value)), uintptr(0))
	pVirtualProtectEx.Call(uintptr(process.hProcess), uintptr(address), uintptr(len(value)), uintptr(oldprot), uintptr(unsafe.Pointer(&oldprot)))
}
