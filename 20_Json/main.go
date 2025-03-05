package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)


var data =`
{
	"user": "John",
	"type": "deposit",
	"amount": 10.0
}
`

// Request is a bank transaction
type Request struct {
	Login string `json:"user"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
}

func main(){
	rdr := bytes.NewBufferString(data) // Simulate a file/socket

	// Decode request
	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}


	fmt.Printf("Decoded Request (got): %+v\n", req)

	// create a new request
	preBalance := 100.0
	response := map[string]interface{}{
		"ok": true,
		"message": "Deposit completed",
		"Balance": preBalance + req.Amount,
		"code":  200,
		
	}
	
	// Encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(response); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}

	fmt.Printf("Response: %+v\n", response)
}