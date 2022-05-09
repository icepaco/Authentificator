package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var usrName = flag.String("u", "noName", "Your user name")
	var userPswd = flag.String("p", "insecure", "Your user password")
	flag.Parse()
	if result, err := bcrypt.GenerateFromPassword([]byte(*userPswd), 12); err == nil {
		var msg strings.Builder
		msg.WriteString("user name:")
		msg.WriteString(*usrName)
		msg.WriteString(" password:")
		msg.Write(result)
		if err2 := writeStringToFile(msg.String()); err2 != nil {
			fmt.Println(err2.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func writeStringToFile(message string) error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	// close the file with defer
	defer f.Close()

	//write directly into file
	f.Write([]byte(message))
	err = os.Chmod("test.txt", 0700)
	if err != nil {
		return err
	}
	return err
}
