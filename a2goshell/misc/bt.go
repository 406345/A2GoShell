package misc

import (
	"fmt"
	"strings"
)

func ReadAria2BtTable(raw string) map[string]string {

	ret := make(map[string]string)
	lines := strings.Split(raw, "\n")
	state := 0
	for _, v := range lines {

		switch state {
		case 0: // ===
			{
				if strings.Index(v, "=") == 0 {
					state = 1
					continue
				}
			}
		case 1: // filename
			{
				if v == "" {
					continue
				}

				segs := strings.Split(v, "|")
				index := strings.Trim(segs[0], " ")
				path := segs[1]
				ret[index] = path
				fmt.Print("#", index, path)
				state = 2
			}
		case 2: // size
			{
				segs := strings.Split(v, "|")

				fmt.Print("size", segs[1])
				fmt.Println()

				state = 3
			}
		case 3: // -
			{
				state = 1
				continue
			}
		default:
			{

			}
		}

	}

	return ret
}
