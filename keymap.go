package main

import (
	"syscall"
)

// https://www.kernel.org/doc/Documentation/input/input.txt
type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

type HidKey struct {
	Name   string
	Code   byte
	Mod    byte
	Active bool
}

func (e HidKey) GetName() string {
	return e.Name
}

func (e HidKey) IsNormalKey() bool {
	return e.Mod == 0
}

// Forked from
// https://github.com/ramius345/pizero-keyboard-proxy/blob/256f0924bef71e249e6e6a3b62b26e4a3d403e2e/evthid/evthid.go
func BuildEventToHidMap() map[uint16]*HidKey {
	m := make(map[uint16]*HidKey)
	m[0] = &HidKey{"KEY_RESERVED", 0, 0, false}
	m[1] = &HidKey{"KEY_ESC", 0x29, 0, false}
	m[2] = &HidKey{"KEY_1", 0x1e, 0, false}
	m[3] = &HidKey{"KEY_2", 0x1f, 0, false}
	m[4] = &HidKey{"KEY_3", 0x20, 0, false}
	m[5] = &HidKey{"KEY_4", 0x21, 0, false}
	m[6] = &HidKey{"KEY_5", 0x22, 0, false}
	m[7] = &HidKey{"KEY_6", 0x23, 0, false}
	m[8] = &HidKey{"KEY_7", 0x24, 0, false}
	m[9] = &HidKey{"KEY_8", 0x25, 0, false}
	m[10] = &HidKey{"KEY_9", 0x26, 0, false}
	m[11] = &HidKey{"KEY_0", 0x27, 0, false}
	m[12] = &HidKey{"KEY_MINUS", 0x2d, 0, false}
	m[13] = &HidKey{"KEY_EQUAL", 0x2e, 0, false}
	m[14] = &HidKey{"KEY_BACKSPACE", 0x2a, 0, false}
	m[15] = &HidKey{"KEY_TAB", 0x2b, 0, false}
	m[16] = &HidKey{"KEY_Q", 0x14, 0, false}
	m[17] = &HidKey{"KEY_W", 0x1a, 0, false}
	m[18] = &HidKey{"KEY_E", 0x08, 0, false}
	m[19] = &HidKey{"KEY_R", 0x15, 0, false}
	m[20] = &HidKey{"KEY_T", 0x17, 0, false}
	m[21] = &HidKey{"KEY_Y", 0x1c, 0, false}
	m[22] = &HidKey{"KEY_U", 0x18, 0, false}
	m[23] = &HidKey{"KEY_I", 0x0c, 0, false}
	m[24] = &HidKey{"KEY_O", 0x12, 0, false}
	m[25] = &HidKey{"KEY_P", 0x13, 0, false}
	m[26] = &HidKey{"KEY_LEFTBRACE", 0x2f, 0, false}
	m[27] = &HidKey{"KEY_RIGHTBRACE", 0x30, 0, false}
	m[28] = &HidKey{"KEY_ENTER", 0x28, 0, false}
	m[29] = &HidKey{"KEY_LEFTCTRL", 0, 0x01, false}
	m[30] = &HidKey{"KEY_A", 0x04, 0, false}
	m[31] = &HidKey{"KEY_S", 0x16, 0, false}
	m[32] = &HidKey{"KEY_D", 0x07, 0, false}
	m[33] = &HidKey{"KEY_F", 0x09, 0, false}
	m[34] = &HidKey{"KEY_G", 0x0a, 0, false}
	m[35] = &HidKey{"KEY_H", 0x0b, 0, false}
	m[36] = &HidKey{"KEY_J", 0x0d, 0, false}
	m[37] = &HidKey{"KEY_K", 0x0e, 0, false}
	m[38] = &HidKey{"KEY_L", 0x0f, 0, false}
	m[39] = &HidKey{"KEY_SEMICOLON", 0x33, 0, false}
	m[40] = &HidKey{"KEY_APOSTROPHE", 0x34, 0, false}
	m[41] = &HidKey{"KEY_GRAVE", 0x35, 0, false}
	m[42] = &HidKey{"KEY_LEFTSHIFT", 0, 0x02, false}
	m[43] = &HidKey{"KEY_BACKSLASH", 0x31, 0, false}
	m[44] = &HidKey{"KEY_Z", 0x1d, 0, false}
	m[45] = &HidKey{"KEY_X", 0x1b, 0, false}
	m[46] = &HidKey{"KEY_C", 0x06, 0, false}
	m[47] = &HidKey{"KEY_V", 0x19, 0, false}
	m[48] = &HidKey{"KEY_B", 0x05, 0, false}
	m[49] = &HidKey{"KEY_N", 0x11, 0, false}
	m[50] = &HidKey{"KEY_M", 0x10, 0, false}
	m[51] = &HidKey{"KEY_COMMA", 0x36, 0, false}
	m[52] = &HidKey{"KEY_DOT", 0x37, 0, false}
	m[53] = &HidKey{"KEY_SLASH", 0x38, 0, false}
	m[54] = &HidKey{"KEY_RIGHTSHIFT", 0, 0x20, false}
	m[55] = &HidKey{"KEY_KPASTERISK", 0x55, 0, false}
	m[56] = &HidKey{"KEY_LEFTALT", 0, 0x04, false}
	m[57] = &HidKey{"KEY_SPACE", 0x2c, 0, false}
	m[58] = &HidKey{"KEY_CAPSLOCK", 0x39, 0, false}
	m[59] = &HidKey{"KEY_F1", 0x3a, 0, false}
	m[60] = &HidKey{"KEY_F2", 0x3b, 0, false}
	m[61] = &HidKey{"KEY_F3", 0x3c, 0, false}
	m[62] = &HidKey{"KEY_F4", 0x3d, 0, false}
	m[63] = &HidKey{"KEY_F5", 0x3e, 0, false}
	m[64] = &HidKey{"KEY_F6", 0x3f, 0, false}
	m[65] = &HidKey{"KEY_F7", 0x40, 0, false}
	m[66] = &HidKey{"KEY_F8", 0x41, 0, false}
	m[67] = &HidKey{"KEY_F9", 0x42, 0, false}
	m[68] = &HidKey{"KEY_F10", 0x43, 0, false}
	m[69] = &HidKey{"KEY_NUMLOCK", 0x53, 0, false}
	m[70] = &HidKey{"KEY_SCROLLLOCK", 0x54, 0, false}
	m[71] = &HidKey{"KEY_KP7", 0x5f, 0, false}
	m[72] = &HidKey{"KEY_KP8", 0x60, 0, false}
	m[73] = &HidKey{"KEY_KP9", 0x61, 0, false}
	m[74] = &HidKey{"KEY_KPMINUS", 0x56, 0, false}
	m[75] = &HidKey{"KEY_KP4", 0x5c, 0, false}
	m[76] = &HidKey{"KEY_KP5", 0x5d, 0, false}
	m[77] = &HidKey{"KEY_KP6", 0x5e, 0, false}
	m[78] = &HidKey{"KEY_KPPLUS", 0x57, 0, false}

	m[87] = &HidKey{"KEY_F11", 0x44, 0, false}
	m[88] = &HidKey{"KEY_F12", 0x45, 0, false}

	m[97] = &HidKey{"KEY_RIGHTCTRL", 0, 0x10, false}

	m[100] = &HidKey{"KEY_RIGHTALT", 0, 0x40, false}

	m[102] = &HidKey{"KEY_HOME", 0x4a, 0, false}
	m[103] = &HidKey{"KEY_UP", 0x52, 0, false}
	m[104] = &HidKey{"KEY_PAGEUP", 0x4b, 0, false}
	m[105] = &HidKey{"KEY_LEFT", 0x50, 0, false}
	m[106] = &HidKey{"KEY_RIGHT", 0x4f, 0, false}
	m[107] = &HidKey{"KEY_END", 0x4d, 0, false}
	m[108] = &HidKey{"KEY_DOWN", 0x51, 0, false}
	m[109] = &HidKey{"KEY_PAGEDOWN", 0x4e, 0, false}
	m[110] = &HidKey{"KEY_INSERT", 0x49, 0, false}
	m[111] = &HidKey{"KEY_DELETE", 0x4c, 0, false}

	m[125] = &HidKey{"KEY_LEFTMETA", 0, 0x08, false}
	m[126] = &HidKey{"KEY_RIGHTMETA", 0, 0x80, false}
	m[127] = &HidKey{"KEY_COMPOSE", 0x65, 0, false}

	return m
}
