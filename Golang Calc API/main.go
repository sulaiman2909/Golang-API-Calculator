package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

type CalculatorRequest struct {
	Operation  string  `json:"operation"`
	First_num  float64 `json:"first_num"`
	Second_num float64 `json:"second_num"`
}

type CalculatorResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func main() {
	http.HandleFunc("/calculator", calculationHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func calculationHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	var request CalculatorRequest

	if error := dec.Decode(&request); error != nil {
		http.Error(w, "Failed to Decode Request", http.StatusBadRequest)
		return
	}

	var response CalculatorResponse
	switch request.Operation {
	case "+":
		response.Result = request.First_num + request.Second_num
	case "-":
		response.Result = request.First_num - request.Second_num
	case "*":
		response.Result = request.First_num * request.Second_num
	case "/":
		if request.Second_num == 0 {
			response.Error = "Cannot be divided by zero"
		} else {
			response.Result = request.First_num / request.Second_num
		}
	case "%":
		if request.Second_num == 0 {
			response.Error = "Cannot be Zero"
		} else {
			response.Result = math.Mod(request.First_num, request.Second_num)
		}

	default:
		response.Error = fmt.Sprintf("Invalid Operation: %s", request.Operation)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		log.Printf("Unable to encode")
	}
}
