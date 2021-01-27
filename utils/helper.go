package utils

import (
	"time"
	"math/rand"
	"strings"
)

func CountDigit(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

func GenerateBookingCode() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 10

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func ExpiredDate() string {
	t := time.Now()
	return t.Add(time.Hour * 3).Format("2006-01-02 15:04:05")
}

func GenerateCustomerCode() string {
	t := time.Now()
	var dateNow string = t.Format("20060102")

	rand.Seed(t.UnixNano())
	chars := []rune("0123456789")
	length := 6

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return dateNow + b.String()
}

func GenerateActivationcode() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 6

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func DateNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}