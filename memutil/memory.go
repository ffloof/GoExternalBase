package memutil

import (
	"fmt"
	"unsafe"
	"strings"
	"syscall"
)

func Hello(){
	fmt.Println("wow")
}



//TODO: add error checks and error returns
type Process struct {
	hProcess HANDLE
}

type Module struct {
	DwBase uintptr
	DwSize uint64
}

func GetModuleInfo(targetProcessName string, targetModuleName string) Module {
	pid := getPID(targetProcessName)
	ret, _, _ := pCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPMODULE | TH32CS_SNAPMODULE32), uintptr(pid))
	hModule := HANDLE(ret)

	if int(hModule) != INVALID_HANDLE_VALUE {
		var entry MODULEENTRY32W
		entry.dwSize = uint32(unsafe.Sizeof(entry))

		for true {
			ret, _, _ := pModule32NextW.Call(uintptr(hModule), uintptr(unsafe.Pointer(&entry)))

			if uint(ret) == 0 {
				//Also did not do the thing ;'(
				break
			}

			modulename := syscall.UTF16ToString(entry.szExePath[:])

			if strings.Contains(modulename, targetModuleName) {
				pCloseHandle.Call(uintptr(hModule))
				return Module { DwBase: uintptr(unsafe.Pointer(entry.modBaseAddr)), DwSize: uint64(entry.modBaseSize)}
			}
		}
	}

	return Module {
		DwBase: 0,
		DwSize: 0,
	}
}

func getPID(targetname string) uint32 {
	ret, _, _ := pCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	hProc := HANDLE(ret)

	var entry PROCESSENTRY32W
	entry.dwSize = uint32(unsafe.Sizeof(entry))

	for true {
		ret, _, _ := pProcess32NextW.Call(uintptr(hProc), uintptr(unsafe.Pointer(&entry)))

		if uint(ret) == 0 {
			//It did not do the thing :(
			break
		}

		processname := syscall.UTF16ToString(entry.szExeFile[:])

		if strings.Contains(processname, targetname) {
			pCloseHandle.Call(uintptr(hProc))
			return entry.th32ProcessID
		}
	}

	return 0
}

func AttatchProcess(targetname string) Process {
	pid := getPID(targetname)
	ret, _, _ := pOpenProcess.Call(PROCESS_ALL_ACCESS, uintptr(0), uintptr(pid))
	
	return Process{ hProcess: HANDLE(ret) }
}





