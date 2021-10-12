package memutil

func IsPressed(key uint32) bool {
	ret, _, _ := pGetAsyncKeyState.Call(uintptr(key))
	return ret != 0
}