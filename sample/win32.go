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
	cSWSHOW       = 5
	cSWUSEDEFAULT = 0x80000000
)

const (
	cWSMAXIMIZEBOX = 0x00010000
	cWSMINIMIZEBOX = 0x00020000
	cWSTHICKFRAME  = 0x00040000
	cWSSYSMENU     = 0x00080000
	cWSCAPTION     = 0x000C0000

	cWSOVERAPPENDWINDOW = 0x00CF0000
)
