package main

import (
	"fmt"
	"github.com/followme1987/protoc/addressbook"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

//read and write Person in the Address Book.
func main() {
	ab := createAddressBook()
	writeToFile("addressBook.bin",ab)
	readFile("addressBook.bin")
}

func writeToFile(fileName string, ab proto.Message) {
	out, err := proto.Marshal(ab)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = ioutil.WriteFile(fileName, out, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createAddressBook() *addressbook.AddressBook {
	return &addressbook.AddressBook{
		People: []*addressbook.Person{
			{
				Name:  "Ganymede",
				Id:    1,
				Email: "123@test.com",
				Phones: []*addressbook.Person_PhoneNumber{
					{
						Number: "098765",
						Type:   0,
					},
				},
			},
		},
	}
}

func readFile(fileName string) {
	out, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	ab := addressbook.AddressBook{}

	err = proto.Unmarshal(out, &ab)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(ab)
}
