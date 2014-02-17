package elaboratecsvfile

import (
	//	"encoding/csv"
	"bufio"
	"bytes"
	"encoding/binary"
//	"fmt"
	"log/syslog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
	"domains"
//	"io"
)

const (
	unknownEndian = iota
	bigEndian
	littleEndian
)

func UTF16BytesToString(b []byte, o binary.ByteOrder) string {
	utf := make([]uint16, (len(b)+(2-1))/2)
	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}
	return string(utf16.Decode(utf))
}

func dropCREndian(data []byte, t1, t2 byte) []byte {
	if len(data) > 1 {
		if data[len(data)-2] == t1 && data[len(data)-1] == t2 {
			return data[0 : len(data)-2]
		}
	}
	return data
}
func dropCRBE(data []byte) []byte {
	return dropCREndian(data, '\x00', '\r')
}

func dropCRLE(data []byte) []byte {
	return dropCREndian(data, '\r', '\x00')
}

func dropCR(data []byte) ([]byte, int) {
	var endian = unknownEndian
	switch ld := len(data); {
	case ld != len(dropCRLE(data)):
		endian = littleEndian
	case ld != len(dropCRBE(data)):
		endian = bigEndian
	}
	return data, endian
}

func ScanUTF16LinesFunc(byteOrder binary.ByteOrder) (bufio.SplitFunc, func() binary.ByteOrder) {

	// Function closure variables
	var endian = unknownEndian
	switch byteOrder {
	case binary.BigEndian:
		endian = bigEndian
	case binary.LittleEndian:
		endian = littleEndian
	}
	const bom = 0xFEFF
	var checkBOM bool = endian == unknownEndian

	// Scanner split function
	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if checkBOM {
			checkBOM = false
			if len(data) > 1 {
				switch uint16(bom) {
				case uint16(data[0])<<8 | uint16(data[1]):
					endian = bigEndian
					return 2, nil, nil
				case uint16(data[1])<<8 | uint16(data[0]):
					endian = littleEndian
					return 2, nil, nil
				}
			}
		}

		// Scan for newline-terminated lines.
		i := 0
		for {
			j := bytes.IndexByte(data[i:], '\n')
			if j < 0 {
				break
			}
			i += j
			switch e := i % 2; e {
			case 1: // UTF-16BE
				if endian != littleEndian {
					if i > 1 {
						if data[i-1] == '\x00' {
							endian = bigEndian
							// We have a full newline-terminated line.
							return i + 1, dropCRBE(data[0 : i-1]), nil
						}
					}
				}
			case 0: // UTF-16LE
				if endian != bigEndian {
					if i+1 < len(data) {
						i++
						if data[i] == '\x00' {
							endian = littleEndian
							// We have a full newline-terminated line.
							return i + 1, dropCRLE(data[0 : i-1]), nil
						}
					}
				}
			}
			i++
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			// drop CR.
			advance = len(data)
			switch endian {
			case bigEndian:
				data = dropCRBE(data)
			case littleEndian:
				data = dropCRLE(data)
			default:
				data, endian = dropCR(data)
			}
			if endian == unknownEndian {
				if runtime.GOOS == "windows" {
					endian = littleEndian
				} else {
					endian = bigEndian
				}
			}
			return advance, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}

	// Endian byte order function
	orderFunc := func() (byteOrder binary.ByteOrder) {
		switch endian {
		case bigEndian:
			byteOrder = binary.BigEndian
		case littleEndian:
			byteOrder = binary.LittleEndian
		}
		return byteOrder
	}

	return splitFunc, orderFunc
}

func Elaborate(golog syslog.Writer, csvfile string) map[string]domains.GoogleAdsVal {

	var intmap map[string]domains.GoogleAdsVal

	file, err := os.Open(csvfile)
	if err != nil {

		golog.Crit(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	golog.Info(file.Name())

	rdr := bufio.NewReader(file)
	scanner := bufio.NewScanner(rdr)
	bo := binary.LittleEndian

	splitFunc, orderFunc := ScanUTF16LinesFunc(bo)

	scanner.Split(splitFunc)

	intmap = make(map[string]domains.GoogleAdsVal)

	for scanner.Scan() {
		b := scanner.Bytes()
		s := UTF16BytesToString(b, orderFunc())

		splitedstrarr := strings.Split(s, "\t")

		if !strings.HasSuffix(splitedstrarr[0], "Ad group") && !strings.HasSuffix(splitedstrarr[0], "Seed Keywords") {

//			fmt.Println(splitedstrarr[1], splitedstrarr[2], splitedstrarr[3], splitedstrarr[4])

			phrase := splitedstrarr[1]	
			var avgmonthlysearch int =0
			if splitedstrarr[3] != "" {
				avgmonthlysearch, err = strconv.Atoi(splitedstrarr[3])
				if err != nil {
					golog.Crit(err.Error() + "-->" + splitedstrarr[3])
					os.Exit(2)

				}
			}

			var competition float64 = 0
			if splitedstrarr[4] != "" {
				competition, err = strconv.ParseFloat(splitedstrarr[4], 64)
				if err != nil {
					golog.Crit(err.Error())
					os.Exit(2)

				}
			}
			var suggestedbid float64=0
			if splitedstrarr[5] != "" {
				suggestedbid, err = strconv.ParseFloat(splitedstrarr[5], 64)
				if err != nil {
					golog.Crit(err.Error())
					os.Exit(2)
				}
			}

			var googleAdsVal domains.GoogleAdsVal
			googleAdsVal.Avgmonthlysearch = avgmonthlysearch
			googleAdsVal.Competition = competition
			googleAdsVal.Suggestedbid = suggestedbid

			intmap[phrase] = googleAdsVal

		}

	}

	return intmap
}
