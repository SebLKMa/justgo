package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type FoodDelivery struct {
	PostCode string `json:"postcode"`
	Recipe   string `json:"recipe"`
	Delivery string `json:"delivery"`
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

func doMain() {
	start := time.Now()

	fileName := "hf.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not obtain stat, handle error: %v", err.Error())
	}

	// maps recipe to count
	var recipesMap map[string]int
	recipesMap = make(map[string]int)

	// maps postcode to count
	var postcodesMap map[string]int
	postcodesMap = make(map[string]int)

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	log.Println("Starting Main loop")
	i := 0
	d.Token() // start token
	countByPostCodeTimeRange := 0
	for d.More() {
		fd := &FoodDelivery{}
		d.Decode(fd)
		//fmt.Printf("%v \n", rd)
		i++
		//fmt.Printf("%d %v ", i, rd)
		recipesMap[fd.Recipe]++
		postcodesMap[fd.PostCode]++

		// 1AM - ... takes 4secs
		//if postcodeInTimeRange(fd, "10161", "8AM", "11PM") {
		//	countByPostCodeTimeRange++
		//fmt.Println("%v\n", fd)
		//}

	}
	d.Token() // end token
	elapsed := time.Since(start)
	fmt.Println("")
	fmt.Printf("To parse the file took [%v]\n", elapsed)
	fmt.Println("")
	fmt.Printf("Total of [%d] object created.\n", i)
	fmt.Printf("The [%s] is %s long\n", fileName, FileSize(fi.Size()))

	// 1. No. of recipes
	log.Printf("1. Number of recipes [%d]\n", len(recipesMap))

	start = time.Now()
	totalCheck := 0
	// 2. sort keys of recipesMap
	log.Println("2. Sorted keys of recipesMap")
	recipeKeys := make([]string, 0, len(recipesMap))
	for k, _ := range recipesMap {
		recipeKeys = append(recipeKeys, k)
	}
	sort.Strings(recipeKeys)
	index := 0
	for _, key := range recipeKeys {
		index++
		val := recipesMap[key]
		fmt.Println(index, key, val)
		totalCheck = totalCheck + val
	}
	fmt.Printf("Verification Total: [%d] [%d]\n", i, totalCheck)
	elapsed = time.Since(start)
	fmt.Printf("2. sort keys of recipesMap took [%v]\n", elapsed)
	fmt.Println("")

	start = time.Now()
	// 3. Postcode with most deliveries
	log.Printf("Number of postcodes [%d]\n", len(postcodesMap))
	log.Println("3. Postcode with most deliveries")

	type kv struct {
		Key   string // postcode
		Value int    // count
	}

	var sortedSlice []kv // slice of postcode=>count to be sorted by count
	for k, v := range postcodesMap {
		sortedSlice = append(sortedSlice, kv{k, v})
	}

	// sort the slice in descending order
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	log.Printf("Postcode with most deliveries [%s] [%d]\n", sortedSlice[0].Key, sortedSlice[0].Value)
	for _, kv := range sortedSlice {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
	elapsed = time.Since(start)
	fmt.Printf("3. Postcode with most deliveries took [%v]\n", elapsed)
	fmt.Println("")

	start = time.Now()
	// 4. Search recipes containing word(s)
	log.Println("4. Search recipes containing word(s)")
	words := []string{"Chicken", "Pork"}

	matchedRecipes := getMatchedKeys(recipeKeys, words)
	for _, matched := range matchedRecipes {
		fmt.Printf("[%s]\n", matched)
	}
	elapsed = time.Since(start)
	fmt.Printf("4. Search recipes containing word(s) took [%v]\n", elapsed)
	fmt.Println("")

	// 5. No. of deliveries by a postcode in time range
	fmt.Printf("Postcode Time Range Count: %d\n", countByPostCodeTimeRange)
}

func convertTo24(s string) int {
	format := s[len(s)-2:]
	strHour := s[:len(s)-2]
	hour, err := strconv.Atoi(strHour) // to integer
	if err != nil {
		fmt.Println(err)
	}
	if format == "AM" && hour == 12 { // 12 midnight = 0
		return 0
	}
	if format == "PM" && hour == 12 { // 12 noon = 12
		return 12
	}
	if format == "PM" { // if after 12pm, add 12 to hour
		hour = hour + 12
	}

	return hour
}

var re = regexp.MustCompile(`(\d+)(AM)|(\d+)(PM)`)

func doTime() {
	str := "Monday 11AM - 11PM"
	fmt.Println(str)

	match := re.FindAllString(str, -1)
	for _, element := range match {
		fmt.Printf("%s == %d\n", element, convertTo24(element))
	}
}

func postcodeInTimeRangeTest() {
	fd := FoodDelivery{PostCode: "10161", Recipe: "abc", Delivery: "Saturday 10AM - 6PM"}
	if postcodeInTimeRange(&fd, "10161", "10AM", "6PM") {
		fmt.Printf("%v is in range\n", fd)
	} else {
		fmt.Printf("%v is NOT in range\n", fd)
	}

	if postcodeInTimeRange(&fd, "10161", "11AM", "5PM") {
		fmt.Printf("%v is in range\n", fd)
	} else {
		fmt.Printf("%v is NOT in range\n", fd)
	}

	if postcodeInTimeRange(&fd, "10161", "10AM", "7PM") {
		fmt.Printf("%v is in range\n", fd)
	} else {
		fmt.Printf("%v is NOT in range\n", fd)
	}

	if postcodeInTimeRange(&fd, "10161", "9AM", "5PM") {
		fmt.Printf("%v is in range\n", fd)
	} else {
		fmt.Printf("%v is NOT in range\n", fd)
	}
}

func postcodeInTimeRange(fd *FoodDelivery, postcode string, startHrStr string, endHrStr string) bool {
	if postcode != fd.PostCode {
		return false
	}

	fdTimeStr := fd.Delivery
	match := re.FindAllString(fdTimeStr, -1)
	if len(match) != 2 {
		return false
	}

	startHr := convertTo24(startHrStr)
	endHr := convertTo24(endHrStr)
	fdStartHr := convertTo24(match[0])
	fdEndHr := convertTo24(match[1])

	if startHr >= fdStartHr && endHr <= fdEndHr {
		return true
	}

	return false
}

func getKeys(logs []string, searchStr string) []string {
	var keys []string
	pattern := regexp.MustCompile(searchStr)

	seen := make(map[string]bool)
	for _, log := range logs {
		result := pattern.FindAllStringSubmatch(log, -1)
		if len(result) != 2 {
			continue
		}
		for _, item := range result {
			key := item[1]
			if _, ok := seen[key]; !ok {
				keys = append(keys, key)
				seen[key] = true
			}
		}
	}
	return keys
}

func getMatchedKeys(logs []string, words []string) []string {
	searchExpr := strings.Join(words, "|")
	var keys []string
	re := regexp.MustCompile(searchExpr) //(`Chicken|Pork`)

	seen := make(map[string]bool)
	for _, log := range logs {
		result := re.FindString(log)
		if result == "" {
			continue
		}
		key := log
		if _, ok := seen[key]; !ok {
			keys = append(keys, key)
			seen[key] = true
		}

	}
	return keys
}

func main() {
	//doTime()
	//postcodeInTimeRangeTest()
	doMain()
}
