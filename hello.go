package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Glass())

	if len(os.Args) < 2 {
		log.Fatal("ERRO: Dont forget the args")
	}

	// STEP 1
	str_from_args := os.Args[1]
	msf_from_args, err := greetings.Hello(str_from_args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msf_from_args)

	// STEP 2
	msg_from_str, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg_from_str)
}
