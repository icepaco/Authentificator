package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// userData is a simple structure to automate and make user login and logout easy
type UserData struct {
	userID          [16]byte
	userName        string
	userPassword    string
	encodedPassword []byte
}

// initUser will initialize name and password as well as create a unique ID for the user. That ID will then be used instead of
// userName
func (user *UserData) initUser(userName string, userPass string) {
	user.userID = uuid.New()
	user.userName = userName
	user.userPassword = userPass
}

// writeStringToFile takes 3 parameters and writes the first 2 parameters in a file named after the 3d parameters.
// The 2 values (first 2 parameters) will be separated by a return carrier \n
func (user *UserData) writeStringToFile(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	// close the file with defer
	defer f.Close()
	// Build the message, separating the username and password by \n
	var msg strings.Builder
	msg.WriteString(user.userName)
	msg.WriteString("\n")
	msg.Write([]byte(user.encodedPassword))

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
func (user *UserData) readFile(fileName string) (err error) {
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
		userInfo := strings.Split(strVar, "\n")
		user.userName = userInfo[0]
		user.encodedPassword = []byte(userInfo[1])
		return
	}
}

// verifyIdentity will compare the encrypted password with the password given by the user
func (user *UserData) verifyIdentity(userName string) bool {
	if errCrypt := bcrypt.CompareHashAndPassword(user.encodedPassword, []byte(user.userPassword)); errCrypt == nil {
		if userName == user.userName {
			return true
		}
	} else {
		return false
	}
	return false
}
func main() {
	// Using flag package as a CLI to enter username + password
	var usrName = flag.String("u", "noName", "Your user name")
	var userPswd = flag.String("p", "incertain", "Your user password")
	flag.Parse()
	var myUser UserData
	myUser.initUser(*usrName, *userPswd)
	// First verify if the user is already registered
	// Given that this is on Front End side we are assuming a single user
	if err1 := myUser.readFile("test.txt"); err1 == nil {
		// User exists so we now verify that user and password matches
		if myUser.verifyIdentity(*usrName) {
			fmt.Println("Succesful login")
		} else {
			fmt.Println("User or password do not match")
		}
	} else {
		// Can't access user file data so we are creating a new user on file
		// Using bcrypt library to encode the password
		if result, err := bcrypt.GenerateFromPassword([]byte(myUser.userPassword), 12); err == nil {
			// Building a text message containing user and encrypted password
			myUser.encodedPassword = result
			if err2 := myUser.writeStringToFile("test.txt"); err2 != nil {
				fmt.Println(err2.Error())
			} else {
				fmt.Println("new user added")
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}
