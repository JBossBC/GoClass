package Dao

import "os"

func init() {
	err := os.MkdirAll("C:\\ProgramData\\goClass", 0744)
	if err != nil {
		panic(err)
	}
}
