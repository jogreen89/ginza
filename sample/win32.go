package main

import "syscall"

var (
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

func getModuleHandle() (syscall.Handle, error) {
	ret, _, err := pGetModuleHandleW.Call(uintptr(0))
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	pCreateWindowExW  = user32.NewProc("CreateWindowExW")
	pDefWindowW       = user32.NewProc("DefWindowProcW")
	pDestroyWindow    = user32.NewProc("DestroyWindow")
	pDispatchMessageW = user32.NewProc("DispatchMessageW")
	pGetMessageW      = user32.NewProc("GetMessageW")
	pLoadCursor       = user32.NewProc("LoadCursorW")
	pPostQuitMessage  = user32.NewProc("PostQuitMessage")
	pRegisterClassExw = user32.NewProc("RegisterClassExW")
	pTranslateMessage = user32.NewProc("TranslateMessage")
)

const (
	cSW_SHOW        = 5
	cSW_USE_DEFAULT = 0x80000000
)

const (
	cWS_MAXIMIZE_BOX = 0x00010000
	cWS_MINIMIZE_BOX = 0x00020000
	cWS_THICKFRAME   = 0x00040000
	cWS_SYSMENU      = 0x00080000
	cWS_CAPTION      = 0x000C0000

	cWS_OVERAPPENDWINDOW = 0x00CF0000
)
