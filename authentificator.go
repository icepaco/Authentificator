package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Using flag package as a CLI to enter username + password
	var usrName = flag.String("u", "noName", "Your user name")
	var userPswd = flag.String("p", "nono", "Your user password")
	flag.Parse()
	// First verify if the user is already registered
	// Given that this is on Front End side we are assuming a single user
	if recordUser, err1 := readFile("test.txt"); err1 == nil {
		// User exists so we now verify that user and password matches
		if errCrypt := bcrypt.CompareHashAndPassword([]byte(recordUser[1]), []byte(*userPswd)); errCrypt == nil {
			if *usrName == recordUser[0] {
				fmt.Println("Succesful login")
			} else {
				fmt.Println("User or password do not match")
			}

		} else {
			fmt.Println("User or password do not match")
		}
	} else {
		// Can't access user file data so we are creating a new user on file
		// Using bcrypt library to encode the password
		if result, err := bcrypt.GenerateFromPassword([]byte(*userPswd), 12); err == nil {
			// Building a text message containing user and encrypted password
			if err2 := writeStringToFile(*usrName, result, "test.txt"); err2 != nil {
				fmt.Println(err2.Error())
			} else {
				fmt.Println("new user added")
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

// writeStringToFile takes 3 parameters and writes the first 2 parameters in a file named after the 3d parameters.
// The 2 values (first 2 parameters) will be separated by a return carrier \n
func writeStringToFile(userName string, encodedPswd []byte, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	// close the file with defer
	defer f.Close()
	// Build the message, separating the username and password by \n
	var msg strings.Builder
	msg.WriteString(userName)
	msg.WriteString("\n")
	msg.Write(encodedPswd)

	//write directly into file
	f.Write([]byte(msg.String()))
	err = os.Chmod(fileName, 0700)
	if err != nil {
		return err
	}
	return err
}

// readFile will read from a file named after it's parameter and return a slice of size 2 containing the user name and the user password
// on each line.  If it can't open the file or read the file it will return an error
func readFile(fileName string) (userData []string, err error) {
	// try to open the file
	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		err = errOpen
		return
	}
	defer file.Close()

	if content, fileErr := ioutil.ReadAll(file); err != nil {
		err = fileErr
		return
	} else {
		// read the 2 value as a string slice separated by \n
		strVar := string(content)
		userData = strings.Split(strVar, "\n")
		return
	}
}
