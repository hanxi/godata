package hello

import (
	//"log"
	//"io/ioutil"
	//pb "github.com/hanxi/godata/go/example"
	//"google.golang.org/protobuf/proto"
)

func Hello() string {
    return "Hello, world."
}

/*
func Write() {
	pn := pb.NewPhoneNumber{}
	pn.SetNumber("123")
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(pn)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("a.out", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	log.Println("Write")
}

func Read() {
	// Read the existing address book.
	in, err := ioutil.ReadFile("a.out")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	pn := pb.NewPhoneNumber{}
	if err := proto.Unmarshal(in, pn); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	n := pn.GetNumber()
	log.Println("Read", n)
}
*/
