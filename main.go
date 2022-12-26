package main

import (
	"fmt"
	"gspr/helper"
	"log"

	proto "google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello World")

	p1 := &helper.Person{
		Name: "Michael",
		Age:  30,
		SocialFollowers: &helper.SocialFollowers{
			Youtube: 150,
			Twitter: 200,
		},
	}

	data, err := proto.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Println(data)

	newP1 := &helper.Person{}
	err = proto.Unmarshal(data, newP1)
	if err != nil {
		log.Println("unmarshalling error: ", err)
	}

	fmt.Printf("%v\n", newP1.String())
	fmt.Printf("%v\n", newP1.GetAge())
	fmt.Printf("%v\n", newP1.GetSocialFollowers())
	fmt.Printf("%v\n", newP1.SocialFollowers.GetTwitter())
}
