package snake

import (
	"strings"

	"github.com/liuchong/econf/internal/bytes"
)

func ToSnake(s string, delimiter uint8, uppercase bool) string {
	s = strings.TrimSpace(s)
	l := len(s)

	newS := strings.Builder{}
	// allocate some more space for growth
	newS.Grow(l + 2)

	prevIsUpper := false
	prevIsLower := false
	prevIsNumber := false

	for i, v := range []byte(s) {
		vIsUpper := bytes.IsUppercase(v)
		vIsLower := bytes.IsLowercase(v)
		vIsNumber := bytes.IsNumber(v)

		if v == ' ' || v == '_' || v == '-' || v == '.' {
			newS.WriteByte(delimiter)
		} else {
			if vIsLower && uppercase {
				v += 'A'
				v -= 'a'
			} else if vIsUpper && !uppercase {
				v += 'a'
				v -= 'A'
			}

			if i > 0 {
				if vIsUpper {
					if prevIsLower || prevIsNumber {
						newS.WriteByte(delimiter)
					} else if nextIsLower := (i+1 < l && bytes.IsLowercase(s[i+1])); nextIsLower {
						// JSONRpc -> json_rpc
						// JsonRPC -> json_rpc
						newS.WriteByte(delimiter)
					}
				} else if vIsNumber {
					if prevIsUpper || prevIsLower {
						newS.WriteByte(delimiter)
					}
				} else if vIsLower {
					if prevIsNumber {
						newS.WriteByte(delimiter)
					}
				}

			}

			newS.WriteByte(v)

		}

		prevIsUpper = vIsUpper
		prevIsLower = vIsLower
		prevIsNumber = vIsNumber
	}

	return newS.String()
}
