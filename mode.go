package memfs

import (
	"errors"
	"os"
	"strings"
)

func ParseFileMode(input string) (os.FileMode, error) {
	var mode os.FileMode

	if len(input) < 10 {
		return 0, errors.New("unable to parse file mode string too short")
	}
	input = strings.ToLower(input)
	switch input[0] {
	case '-':
	case 'd':
		mode |= os.ModeDir // d: is a directory
	case 'a':
		mode |= os.ModeAppend // a: append-only
	case 'l':
		mode |= os.ModeExclusive // l: exclusive use
	case 'T':
		mode |= os.ModeTemporary // T: temporary file; Plan 9 only
	case 'L':
		mode |= os.ModeSymlink // L: symbolic link
	case 'D':
		mode |= os.ModeDevice // D: device file
	case 'p':
		mode |= os.ModeNamedPipe // p: named pipe (FIFO)
	case 'S':
		mode |= os.ModeSocket // S: Unix domain socket
	case 'u':
		mode |= os.ModeSetuid // u: setuid
	case 'g':
		mode |= os.ModeSetgid // g: setgid
	case 'c':
		mode |= os.ModeCharDevice // c: Unix character device, when ModeDevice is set
	case 't':
		mode |= os.ModeSticky // t: sticky
	}

	if input[1] == 'r' {
		mode |= OS_USER_R
	}
	if input[2] == 'w' {
		mode |= OS_USER_W
	}
	if input[3] == 'x' {
		mode |= OS_USER_X
	}
	if input[4] == 'r' {
		mode |= OS_GROUP_R
	}
	if input[5] == 'w' {
		mode |= OS_GROUP_W
	}
	if input[6] == 'x' {
		mode |= OS_GROUP_X
	}
	if input[7] == 'r' {
		mode |= OS_OTH_R
	}
	if input[8] == 'w' {
		mode |= OS_OTH_W
	}
	if input[9] == 'x' {
		mode |= OS_OTH_X
	}

	return mode, nil
}
