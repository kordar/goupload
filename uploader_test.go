package goupload_test

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestT1(t *testing.T) {
	sourceURL := "aa/bb?123"
	id := []string{}
	if strings.HasPrefix(sourceURL, "http://") || strings.HasPrefix(sourceURL, "https://") {
		log.Println("sourceURL format is invalid.")
	}
	surl := strings.SplitN(sourceURL, "/", 2)
	log.Printf("=========%+v", surl)
	if len(surl) < 2 {
		log.Println(fmt.Sprintf("x-cos-copy-source format error: %s", sourceURL))
	}
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("%s/%s?versionId=%s", surl[0], encodeURIComponent(surl[1]), id[0])
	} else if len(id) == 0 {
		keyAndVer := strings.SplitN(surl[1], "?", 2)
		if len(keyAndVer) < 2 {
			u = fmt.Sprintf("%s/%s", surl[0], encodeURIComponent(surl[1], []byte{'/'}))
		} else {
			u = fmt.Sprintf("%v/%v?%v", surl[0], encodeURIComponent(keyAndVer[0], []byte{'/'}), encodeURIComponent(keyAndVer[1], []byte{'='}))
		}
	} else {
		log.Println("wrong params")
	}
	log.Println(u)
}

func encodeURIComponent(s string, excluded ...[]byte) string {
	var b bytes.Buffer
	written := 0

	for i, n := 0, len(s); i < n; i++ {
		c := s[i]

		switch c {
		case '-', '_', '.', '!', '~', '*', '\'', '(', ')':
			continue
		default:
			// Unreserved according to RFC 3986 sec 2.3
			if 'a' <= c && c <= 'z' {

				continue

			}
			if 'A' <= c && c <= 'Z' {

				continue

			}
			if '0' <= c && c <= '9' {

				continue
			}
			if len(excluded) > 0 {
				conti := false
				for _, ch := range excluded[0] {
					if ch == c {
						conti = true
						break
					}
				}
				if conti {
					continue
				}
			}
		}

		b.WriteString(s[written:i])
		fmt.Fprintf(&b, "%%%02X", c)
		written = i + 1
	}

	if written == 0 {
		return s
	}
	b.WriteString(s[written:])
	return b.String()
}
