package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"romanNumerals"
	"slices"
	"strings"
	"strconv"
)

func GetRomanNumeral(num int) string {
	var romanKeys = []int{}
	if romanNumerals.Numerals[num] != "" {
		return romanNumerals.Numerals[num]
	}
	for key, _ := range romanNumerals.Numerals {
		romanKeys = append(romanKeys, key)
	}
	slices.Sort(romanKeys)
	chars := []string{}
	numP := &num

	for {
		if *numP == 0 {
			break
		}
		if romanNumerals.Numerals[*numP] != "" {
			chars = append(chars, romanNumerals.Numerals[*numP])
			*numP -= *numP
		} else {
			for i := len(romanKeys) - 1; i >= 0; i-- {
				if romanKeys[i] < *numP {
					chars = append(chars, romanNumerals.Numerals[romanKeys[i]])
					*numP -= romanKeys[i]
					break
				}
			}
		}

	}
	return strings.Join(chars, "")
}

type data = map[string]int

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method){
		case "GET":
			returnStr := ""
			for key, val := range romanNumerals.Numerals {
				returnStr += fmt.Sprintf("%-10s--> %+5s \n", strconv.FormatInt(int64(key), 10), val)
			}
			w.Write([]byte(returnStr))
		case "POST":
			var requestData data
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&requestData)
			if err != nil{
				fmt.Println(err)
				http.Error(w, "invalid request", http.StatusBadRequest)
				break
			}
			rn := GetRomanNumeral(requestData["number"])
			w.Write([]byte(rn))
	default:
		http.Error(w, "route not found", http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8000", mux)
}
