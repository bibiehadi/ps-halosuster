package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ValidateNIP(nip int, isNurse bool) bool {
	var reg string = ""
	if isNurse {
		reg = `^303(1|2)(200[0-9]{1}|201[0-9]{1}|202[0-4]{1})(0[1-9]{1}|1[0-2]{1})([0-9]{3,5})$`
	} else {
		reg = `^615(1|2)(200[0-9]{1}|201[0-9]{1}|202[0-4]{1})(0[1-9]{1}|1[0-2]{1})([0-9]{3,5})$`
	}
	match, err := regexp.MatchString(reg, strconv.Itoa(nip))
	if err != nil {
		fmt.Fprintf(os.Stdout, "error validating phone number: %w", []any{err}...)
		return false
	}
	if !match {
		fmt.Printf("invalid phone number format (must start with + and valid international code)")
		return false
	}
	return true
}

func ValidateUrl(url string) bool {
	match, err := regexp.MatchString(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`, url)
	if err != nil {
		fmt.Printf("error validating phone number: %w", err)
		return false
	}
	if !match {
		fmt.Printf("invalid phone number format (must start with + and valid international code)")
		return false
	}
	return true
}
