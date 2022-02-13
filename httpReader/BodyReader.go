package httpReader

import (
	"strings"
)

// removeCharacter will remove (""),({),(}) from byteSlice
func removeCharacters(byteSlice []byte) []byte {
	var newByteSlice []byte = []byte{}
	for _, byteMember := range byteSlice {
		if byteMember != 34 && byteMember != 123 && byteMember != 125 {
			newByteSlice = append(newByteSlice, byteMember)
		}
	}
	return newByteSlice
}

// BodyReader will recieve a byte slice, delete its curly braces
// convert that byte slice into a string, split that string by (,)
// to then loop through that slice,then it will split that same sub string by (:) and
// and add the members of that sub slice to bodyMap
func BodyReader(body []byte) map[string]string {
	//delete curly braces from byte string
	var bodyMap map[string]string = make(map[string]string)
	body = removeCharacters(body)
	var stringBody string = string(body)
	for _, str := range strings.Split(stringBody, ",") {
		var subStr []string = strings.Split(str, ":")
		bodyMap[subStr[0]] = subStr[1]
	}
	return bodyMap
}
