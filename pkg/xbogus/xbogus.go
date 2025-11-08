package xbogus

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"
)

// XBogusService implements the X-Bogus algorithm for generating TikTok API parameters
type XBogusService struct {
	// Array is a lookup table for hex character conversion
	Array []*int
	// character is the base64-like encoding character set
	character string
	// uaKey is the user agent encryption key
	uaKey []byte
	// userAgent is the user agent string
	userAgent string
}

// NewXBogusService creates a new XBogusService instance
func NewXBogusService(userAgent string) *XBogusService {
	// Initialize Array lookup table (128 elements, matching Python implementation)
	// Most are None (nil), except for hex characters 0-9, A-F, a-f
	array := make([]*int, 128)

	// Pre-allocate values to avoid pointer issues
	values := make([]int, 22) // 10 digits + 6 uppercase + 6 lowercase

	// Set digits 0-9 (ASCII 48-57) to values 0-9
	for i := 0; i < 10; i++ {
		values[i] = i
		array[48+i] = &values[i]
	}

	// Set uppercase letters A-F (ASCII 65-70) to values 10-15
	for i := 0; i < 6; i++ {
		values[10+i] = 10 + i
		array[65+i] = &values[10+i]
	}

	// Set lowercase letters a-f (ASCII 97-102) to values 10-15
	for i := 0; i < 6; i++ {
		values[16+i] = 10 + i
		array[97+i] = &values[16+i]
	}

	// Set default user agent if not provided
	if userAgent == "" {
		userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0"
	}

	return &XBogusService{
		Array:     array,
		character: "Dkdpgh4ZKsQB80/Mfvw36XI1R25-WUAlEi7NLboqYTOPuzmFjJnryx9HVGcaStCe=",
		uaKey:     []byte{0x00, 0x01, 0x0c},
		userAgent: userAgent,
	}
}

// md5StrToArray converts an MD5 hex string to an array of integers
func (x *XBogusService) md5StrToArray(md5Str string) []int {
	if len(md5Str) > 32 {
		// If string is longer than 32 chars, return ord(char) for each char
		result := make([]int, len(md5Str))
		for i, char := range md5Str {
			result[i] = int(char)
		}
		return result
	}

	// Process pairs of hex characters
	result := make([]int, 0, len(md5Str)/2)
	for i := 0; i < len(md5Str); i += 2 {
		if i+1 < len(md5Str) {
			char1 := int(md5Str[i])
			char2 := int(md5Str[i+1])
			if char1 < len(x.Array) && char2 < len(x.Array) && x.Array[char1] != nil && x.Array[char2] != nil {
				val := (*x.Array[char1] << 4) | *x.Array[char2]
				result = append(result, val)
			}
		}
	}
	return result
}

// md5Encrypt encrypts URL path using multiple rounds of MD5 hashing
func (x *XBogusService) md5Encrypt(urlPath string) []int {
	// md5(md5_str_to_array(md5(url_path)))
	firstHash := x.md5(urlPath)
	firstArray := x.md5StrToArray(firstHash)
	secondHash := x.md5FromArray(firstArray)
	secondArray := x.md5StrToArray(secondHash)
	return secondArray
}

// md5 calculates MD5 hash of input data (string or array)
func (x *XBogusService) md5(input interface{}) string {
	var array []int

	switch v := input.(type) {
	case string:
		array = x.md5StrToArray(v)
	case []int:
		array = v
	default:
		panic("Invalid input type. Expected string or []int")
	}

	// Convert array to bytes
	bytes := make([]byte, len(array))
	for i, val := range array {
		bytes[i] = byte(val)
	}

	// Calculate MD5
	hash := md5.Sum(bytes)
	return fmt.Sprintf("%x", hash)
}

// md5FromArray calculates MD5 hash from an integer array
func (x *XBogusService) md5FromArray(arr []int) string {
	bytes := make([]byte, len(arr))
	for i, val := range arr {
		bytes[i] = byte(val)
	}
	hash := md5.Sum(bytes)
	return fmt.Sprintf("%x", hash)
}

// encodingConversion performs the first encoding conversion
func (x *XBogusService) encodingConversion(a, b, c, e, d, t, f, r, n, o, i, underscore, xVal, u, s, l, v, h, p int) string {
	y := []int{a}
	y = append(y, i)
	y = append(y, b, underscore, c, xVal, e, u, d, s, t, l, f, v, r, h, n, p, o)

	// Convert to bytes and decode as ISO-8859-1
	result := make([]byte, len(y))
	for idx, val := range y {
		result[idx] = byte(val)
	}
	return string(result)
}

// encodingConversion2 performs the second encoding conversion
// Returns chr(a) + chr(b) + c
func (x *XBogusService) encodingConversion2(a, b int, c string) string {
	return string([]byte{byte(a), byte(b)}) + c
}

// rc4Encrypt encrypts data using RC4 algorithm
func (x *XBogusService) rc4Encrypt(key []byte, data []byte) []byte {
	S := make([]int, 256)
	for i := 0; i < 256; i++ {
		S[i] = i
	}

	j := 0
	// Initialize S box
	for i := 0; i < 256; i++ {
		j = (j + S[i] + int(key[i%len(key)])) % 256
		S[i], S[j] = S[j], S[i]
	}

	// Generate ciphertext
	encryptedData := make([]byte, len(data))
	i, j := 0, 0
	for k := 0; k < len(data); k++ {
		i = (i + 1) % 256
		j = (j + S[i]) % 256
		S[i], S[j] = S[j], S[i]
		encryptedByte := data[k] ^ byte(S[(S[i]+S[j])%256])
		encryptedData[k] = encryptedByte
	}

	return encryptedData
}

// calculation performs bitwise operations and base64-like encoding
func (x *XBogusService) calculation(a1, a2, a3 int) string {
	x1 := (a1 & 255) << 16
	x2 := (a2 & 255) << 8
	x3 := x1 | x2 | a3

	result := string(x.character[(x3&16515072)>>18]) +
		string(x.character[(x3&258048)>>12]) +
		string(x.character[(x3&4032)>>6]) +
		string(x.character[x3&63])

	return result
}

// GetXBogus generates the X-Bogus parameter for the given URL path
// Returns: (params string, xbogus string, userAgent string)
func (x *XBogusService) GetXBogus(urlPath string) (string, string, string) {
	// array1 = md5_str_to_array(md5(base64(rc4_encrypt(ua_key, user_agent))))
	uaBytes := []byte(x.userAgent)
	rc4Encrypted := x.rc4Encrypt(x.uaKey, uaBytes)
	base64Encoded := base64.StdEncoding.EncodeToString(rc4Encrypted)
	md5Hash1 := x.md5(base64Encoded)
	array1 := x.md5StrToArray(md5Hash1)

	// array2 = md5_str_to_array(md5(md5_str_to_array("d41d8cd98f00b204e9800998ecf8427e")))
	emptyHashArray := x.md5StrToArray("d41d8cd98f00b204e9800998ecf8427e")
	md5Hash2 := x.md5FromArray(emptyHashArray)
	array2 := x.md5StrToArray(md5Hash2)

	// url_path_array = md5_encrypt(url_path)
	urlPathArray := x.md5Encrypt(urlPath)

	// timer = int(time.time())
	timer := int(time.Now().Unix())
	ct := 536919696

	// new_array = [64, 0.00390625, 1, 12, ...]
	// Note: 0.00390625 = 1/256, but in the XOR operation it's converted to int(0)
	newArray := []int{
		64, 0, 1, 12,
		urlPathArray[14], urlPathArray[15], array2[14], array2[15], array1[14], array1[15],
		(timer >> 24) & 255, (timer >> 16) & 255, (timer >> 8) & 255, timer & 255,
		(ct >> 24) & 255, (ct >> 16) & 255, (ct >> 8) & 255, ct & 255,
	}

	// XOR all elements
	xorResult := newArray[0]
	for i := 1; i < len(newArray); i++ {
		xorResult ^= newArray[i]
	}
	newArray = append(newArray, xorResult)

	// Split new_array into array3 and array4
	var array3, array4 []int
	idx := 0
	for idx < len(newArray) {
		array3 = append(array3, newArray[idx])
		if idx+1 < len(newArray) {
			array4 = append(array4, newArray[idx+1])
		}
		idx += 2
	}

	// merge_array = array3 + array4
	mergeArray := append(array3, array4...)

	// encoding_conversion(*merge_array)
	// Ensure we have exactly 19 elements
	if len(mergeArray) < 19 {
		// Pad with zeros if needed
		for len(mergeArray) < 19 {
			mergeArray = append(mergeArray, 0)
		}
	}
	garbledCode := x.encodingConversion2(
		2,
		255,
		string(x.rc4Encrypt(
			[]byte{0xff}, // "ÿ" in ISO-8859-1
			[]byte(x.encodingConversion(
				mergeArray[0], mergeArray[1], mergeArray[2], mergeArray[3], mergeArray[4],
				mergeArray[5], mergeArray[6], mergeArray[7], mergeArray[8], mergeArray[9],
				mergeArray[10], mergeArray[11], mergeArray[12], mergeArray[13], mergeArray[14],
				mergeArray[15], mergeArray[16], mergeArray[17], mergeArray[18],
			)),
		)),
	)

	// Calculate xb_ by processing garbled_code in chunks of 3
	var xb string
	idx = 0
	for idx < len(garbledCode) {
		if idx+2 < len(garbledCode) {
			xb += x.calculation(
				int(garbledCode[idx]),
				int(garbledCode[idx+1]),
				int(garbledCode[idx+2]),
			)
		}
		idx += 3
	}

	params := fmt.Sprintf("%s&X-Bogus=%s", urlPath, xb)
	return params, xb, x.userAgent
}
