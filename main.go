package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Uygulamaya hoşgeldiniz, hava durumunu öğrenmek istediğiniz şehri giriniz:")

	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	if !checkInCSV("tr2.csv", city) {
		fmt.Println("Yanlış girdi")
		return
	}

	fmt.Println("Şimdi hava durumunu öğrenmek istediğiniz ilçeyi giriniz:")
	district, _ := reader.ReadString('\n')
	district = strings.TrimSpace(district)

	if !checkInCSV("tr2.csv", district) {
		fmt.Println("Yanlış girdi")
		return
	}

	fmt.Printf("https://openweathermap.org/find?q=%s\n", district)

}

func checkInCSV(filename, value string) bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Dosya açılamadı: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		for _, item := range record {
			if strings.EqualFold(item, value) {
				return true
			}
		}
	}
	return false
}
