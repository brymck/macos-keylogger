package keyboard

import "errors"

var ErrInvalidKeyCode = errors.New("invalid key code")

var keyCodeKeyMapping = []string{
	"a",          // 0
	"s",          // 1
	"d",          // 2
	"f",          // 3
	"h",          // 4
	"g",          // 5
	"z",          // 6
	"x",          // 7
	"c",          // 8
	"v",          // 9
	"",           // 10
	"b",          // 11
	"q",          // 12
	"w",          // 13
	"e",          // 14
	"r",          // 15
	"y",          // 16
	"t",          // 17
	"1",          // 18
	"2",          // 19
	"3",          // 20
	"4",          // 21
	"6",          // 22
	"5",          // 23
	"=",          // 24
	"9",          // 25
	"7",          // 26
	"-",          // 27
	"8",          // 28
	"0",          // 29
	"]",          // 30
	"o",          // 31
	"u",          // 32
	"[",          // 33
	"i",          // 34
	"p",          // 35
	"CR",         // 36
	"l",          // 37
	"j",          // 38
	"'",          // 39
	"k",          // 40
	";",          // 41
	"\\",         // 42
	",",          // 43
	"/",          // 44
	"n",          // 45
	"m",          // 46
	".",          // 47
	"Tab",        // 48
	"Space",      // 49
	"~",          // 50
	"BS",         // 51
	"",           // 52
	"Esc",        // 53
	"rightsuper", // 54
	"leftsuper",  // 55
	"leftshift",  // 56
	"capslock",   // 57
	"leftalt",    // 58
	"leftctrl",   // 59
	"rightshift", // 60
	"rightalt",   // 61
	"rightctrl",  // 62
	"",           // 63
	"F17",        // 64
	"",           // 65
	"",           // 66
	"",           // 67
	"",           // 68
	"",           // 69
	"",           // 70
	"",           // 71
	"",           // 72
	"",           // 73
	"",           // 74
	"",           // 75
	"",           // 76
	"",           // 77
	"",           // 78
	"F18",        // 79
	"F19",        // 80
	"",           // 81
	"",           // 82
	"",           // 83
	"",           // 84
	"",           // 85
	"",           // 86
	"",           // 87
	"",           // 88
	"",           // 89
	"F20",        // 90
	"",           // 91
	"",           // 92
	"",           // 93
	"",           // 94
	"",           // 95
	"F5",         // 96
	"F6",         // 97
	"F7",         // 98
	"F3",         // 99
	"F8",         // 100
	"F9",         // 101
	"",           // 102
	"F11",        // 103
	"",           // 104
	"F13",        // 105
	"F16",        // 106
	"F14",        // 107
	"",           // 108
	"F10",        // 109
	"",           // 110
	"F12",        // 111
	"",           // 112
	"F15",        // 113
	"",           // 114
	"Home",       // 115
	"PageUp",     // 116
	"Del",        // 117
	"F4",         // 118
	"End",        // 119
	"F2",         // 120
	"PageDown",   // 121
	"F1",         // 122
	"Left",       // 123
	"Right",      // 124
	"Down",       // 125
	"Up",         // 126
}

func ConvertKeyCode(keyCode int) (string, error) {
	if keyCode < 0 || keyCode >= len(keyCodeKeyMapping) {
		return "", ErrInvalidKeyCode
	}

	k := keyCodeKeyMapping[keyCode]
	if k == "" {
		return "", ErrInvalidKeyCode
	}

	return k, nil
}
