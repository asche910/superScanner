package main

import (
	"fmt"
	"testing"
)

func TestCheckFileExists(t *testing.T) {
	exists := CheckFileExists("output/subdomain_1618750953.txt")
	fmt.Println(exists)
}

func TestFileToList(t *testing.T) {
	list := FileToList("output/subdomain_1618751008.txt")
	fmt.Println(list)
}

func TestTempTest(t *testing.T) {
	arr := []string{"11", "22", "11", "33"}
	duplication := RemoveDuplication(arr)
	fmt.Println(duplication)

}
