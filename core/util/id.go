package util

import "fmt"

// GetID : gives the id for contact. its a combination of source id and number
func GetID(sourceID, number string) string {
	return fmt.Sprintf("%v::%v", sourceID, number)
}
