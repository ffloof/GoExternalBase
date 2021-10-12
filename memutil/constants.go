package memutil

import (
	"syscall"
)

type (
	HANDLE  uintptr
	HMODULE HANDLE
)

const (
	INVALID_HANDLE_VALUE = int(-1)
	MAX_MODULE_NAME32    = 255
	MAX_PATH             = 260
)

const (
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010
	TH32CS_INHERIT      = 0x80000000
	TH32CS_SNAPALL      = TH32CS_SNAPHEAPLIST | TH32CS_SNAPMODULE | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD
)

const (
	PAGE_EXECUTE_READWRITE = 0x40

	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	SYNCHRONIZE              = 0x00100000

	PROCESS_ALL_ACCESS = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF
)

//Win32api structs
type MODULEENTRY32W struct {
	dwSize        uint32
	th32ModuleID  uint32
	th32ProcessID uint32
	glblcntUsage  uint32
	proccntUsage  uint32
	modBaseAddr   *uint8
	modBaseSize   uint32
	hMODULE       HMODULE
	szModule      [MAX_MODULE_NAME32 + 1]uint16
	szExePath     [MAX_PATH]uint16
}

type PROCESSENTRY32W struct {
	dwSize uint32
	cntUsage uint32
	th32ProcessID uint32
	th32DefaultHeapID uintptr //might 
	th32ModuleID uint32
	cntThreads uint32
	th32ParentProcessID uint32
	pcPriClassBase int32
	dwFlags uint32
	szExeFile [MAX_PATH]uint16
}

var (
	k32 = syscall.NewLazyDLL("kernel32.dll")
	u32 = syscall.NewLazyDLL("user32.dll")

	// Input
	pGetAsyncKeyState = u32.NewProc("GetAsyncKeyState")

	// Read / Write mem
	pReadProcessMemory  = k32.NewProc("ReadProcessMemory")
	pWriteProcessMemory = k32.NewProc("WriteProcessMemory")

	// Process enumeration
	pOpenProcess               = k32.NewProc("OpenProcess")
	pCreateToolhelp32Snapshot  = k32.NewProc("CreateToolhelp32Snapshot")
	pModule32FirstW            = k32.NewProc("Module32FirstW")
	pModule32NextW             = k32.NewProc("Module32NextW")
	pProcess32FirstW           = k32.NewProc("Process32FirstW")
	pProcess32NextW            = k32.NewProc("Process32NextW")
	pVirtualProtectEx          = k32.NewProc("VirtualProtectEx")

	// Other
	pCloseHandle = k32.NewProc("CloseHandle")
)