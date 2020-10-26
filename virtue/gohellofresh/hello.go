// main package containing the main program and its supporting functions
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ============================================================================
// Structs definition
// ============================================================================
// FoodDelivery is used during input
type FoodDelivery struct {
	PostCode string `json:"postcode"`
	Recipe   string `json:"recipe"`
	Delivery string `json:"delivery"`
}

// KeyValuePair defines a string key and integer value pair
type KeyValuePair struct {
	Key   string // postcode
	Value int    // count
}

// RecipeCount defines a recipe and its count, used by Output
type RecipeCount struct {
	Recipe string `json:"recipe"`
	Count  int    `json:"count"`
}

// PostcodeCount defines a postal code and its count, used by Output
type PostcodeCount struct {
	Postcode string `json:"postcode"`
	Count    int    `json:"delivery_count"`
}

// PostcodePerformanceTimerange defines the performance of a postal code in a specific time range
// Used by Output
type PostcodePerformanceTimerange struct {
	Postcode string `json:"postcode"`
	From     string `json:"from"`
	To       string `json:"to"`
	Count    int    `json:"delivery_count"`
}

// Output is implements the API output specification
type Output struct {
	UniqueRecipeCount       int                          `json:"unique_recipe_count"`
	CountPerRecipe          []RecipeCount                `json:"count_per_recipe"`
	BusiestPostcode         PostcodeCount                `json:"busiest_postcode"`
	CountPerPostcodeAndTime PostcodePerformanceTimerange `json:"count_per_postcode_and_time"`
	MatchByName             []string                     `json:"match_by_name"`
}

// ============================================================================
// Supporting functions
// ============================================================================

// ArrayFlags provides a list of strings for command line arguments
type ArrayFlags []string

// Default string value
func (i *ArrayFlags) String() string {
	return "string representation"
}

// Set implements flag.Value interface
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// convertTo24 a string with format nnAM or nnPM to 24hr, e.g. "3PM" will return 15.
// if error envountered, -1 is returned.
func convertTo24(s string) (int, error) {
	if len(s) < 3 {
		return -1, errors.New("can't work length less than 3")
	}
	format := s[len(s)-2:]             // last 2 chars
	strHour := s[:len(s)-2]            // first chars except the last 2
	hour, err := strconv.Atoi(strHour) // to integer
	if err != nil {
		return -1, err
	}
	if format == "AM" && hour == 12 { // 12 midnight = 0
		return 0, nil
	}
	if format == "PM" && hour == 12 { // 12 noon = 12
		return 12, nil
	}
	if format == "PM" { // if after 12pm, add 12 to hour
		hour = hour + 12
	}

	return hour, nil
}

// regex used by postcodeInTimeRange()
var re = regexp.MustCompile(`(\d+)(AM)|(\d+)(PM)`)

// postcodeInTimeRange determines if the postcode timerange is within the same timerange of FoodDelivery record
func postcodeInTimeRange(fd *FoodDelivery, postcode string, startHrStr string, endHrStr string) bool {
	if postcode == "" || postcode != fd.PostCode {
		return false
	}

	fdTimeStr := fd.Delivery
	match := re.FindAllString(fdTimeStr, -1)
	if len(match) != 2 {
		return false
	}

	startHr, err := convertTo24(startHrStr)
	if err != nil {
		return false
	}
	endHr, err := convertTo24(endHrStr)
	if err != nil {
		return false
	}
	fdStartHr, err := convertTo24(match[0])
	if err != nil {
		return false
	}
	fdEndHr, err := convertTo24(match[1])
	if err != nil {
		return false
	}

	if startHr >= fdStartHr && endHr <= fdEndHr {
		return true
	}

	return false
}

// getMatchedKeys resembles a "SQL LIKE" search, it returns the list of strings that matches the words
func getMatchedKeys(inputlist []string, words []string) (matchedList []string) {
	searchExpr := strings.Join(words, "|")
	re := regexp.MustCompile(searchExpr) // e.g. (`Chicken|Pork`)

	seen := make(map[string]bool)
	for _, inputStr := range inputlist {
		result := re.FindString(inputStr)
		if result == "" {
			continue
		}
		key := inputStr
		if _, ok := seen[key]; !ok { // check as we do not want duplicates from inputList
			matchedList = append(matchedList, key)
			seen[key] = true
		}

	}
	return
}

// getSortedRecipes returns sorted list of recipes
func getSortedRecipes(recipesMap map[string]int) []string {
	//  sort the recipe keys
	recipeKeys := make([]string, 0, len(recipesMap))
	for k, _ := range recipesMap {
		recipeKeys = append(recipeKeys, k)
	}
	sort.Strings(recipeKeys)

	return recipeKeys
}

// returns a list of recipes and their counts
func getRecipesCount(recipeKeys []string, recipesMap map[string]int) (results []RecipeCount) {
	for _, key := range recipeKeys {
		val := recipesMap[key]
		//fmt.Println(index, key, val)
		results = append(results, RecipeCount{Recipe: key, Count: val})
	}
	return
}

// returns the recipes that contain the provided words
func getMatchedRecipes(recipesKeys []string, words []string) (results []string) {
	if len(words) == 0 || len(recipesKeys) == 0 {
		return
	}
	matchedRecipes := getMatchedKeys(recipesKeys, words)
	for _, matched := range matchedRecipes {
		results = append(results, matched)
	}
	return
}

// returns a ranked list of postcode=>count to be sorted descending by count
func getRankedPostcodes(postcodesMap map[string]int) (results []KeyValuePair) {
	// transfer map keys to results list
	for k, v := range postcodesMap {
		results = append(results, KeyValuePair{k, v})
	}

	// sort the results list in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	return
}

// a placeholder to hold the top ranked postcode
var topPostCode KeyValuePair = KeyValuePair{}

// updateTopRankedPostcode updates topPostCode that is higher ranked by count
func updateTopRankedPostcode(postcode string, count int) {
	if count > topPostCode.Value {
		topPostCode.Key = postcode
		topPostCode.Value = count
	}
}

// incrementPostcodeInTimeRange increments the provided postCodePerformance if it is within the provided FoodDelivery timerange
func incrementPostcodeInTimeRange(fd *FoodDelivery, postCodePerformance *PostcodePerformanceTimerange) {
	if postcodeInTimeRange(fd, postCodePerformance.Postcode, postCodePerformance.From, postCodePerformance.To) {
		postCodePerformance.Count++
	}
}

// ============================================================================
// Main functions
// ============================================================================

// doOutput processes and writes to JSON file using the Output struct definition
func doOutput(outputFile string, recipesMap map[string]int, postcodesMap map[string]int, postCodePerformance *PostcodePerformanceTimerange, words []string) {
	f, _ := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666) // overwrite if exists
	defer f.Close()

	uniqueRecipesCount := len(recipesMap)
	log.Printf("1. Number of recipes [%d]\n", uniqueRecipesCount)
	// 1.
	recipesKeys := getSortedRecipes(recipesMap)

	// 2.
	recipesCounts := getRecipesCount(recipesKeys, recipesMap)

	// 3.
	//rankedPostcodes := getRankedPostcodes(postcodesMap)
	//topRankedPostcode := rankedPostcodes[0]
	topRankedPostcode := topPostCode

	// 4. postCodePerformance computed from doInput

	// 5. Search recipes containing word(s)
	matchedRecipes := getMatchedRecipes(recipesKeys, words)

	myOutput := Output{
		UniqueRecipeCount: uniqueRecipesCount,
		CountPerRecipe:    recipesCounts,
		BusiestPostcode: PostcodeCount{
			Postcode: topRankedPostcode.Key,
			Count:    topRankedPostcode.Value,
		},
		CountPerPostcodeAndTime: *postCodePerformance,
		MatchByName:             matchedRecipes,
	}

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ") // leave empty prefix for first arg, second arg is the indent spaces
	if err := encoder.Encode(&myOutput); err != nil {
		log.Println(err)
	}
}

func doInput(inputFile string, recipesMap map[string]int, postcodesMap map[string]int, postCodePerformance *PostcodePerformanceTimerange) {
	start := time.Now()

	fileName := inputFile
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	log.Println("Starting Main loop")
	i := 0
	d.Token() // start token
	for d.More() {
		fd := &FoodDelivery{}
		d.Decode(fd)
		i++
		recipesMap[fd.Recipe]++
		postcodesMap[fd.PostCode]++

		updateTopRankedPostcode(fd.PostCode, postcodesMap[fd.PostCode])

		incrementPostcodeInTimeRange(fd, postCodePerformance)

	}
	d.Token() // end token
	elapsed := time.Since(start)
	log.Println("")
	log.Printf("Parsing input file took [%v]\n", elapsed)
	fmt.Printf("Total of [%d] object created.\n", i)
}

// Sample Usage:
// go run hello.go --word_list=Chicken -word_list=Pork
// go run hello.go -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork
// go run hello.go -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork -word_list=Pasta
// go run hello.go -input_file=input.json -output_file=output.json -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork -word_list=Pasta
func main() {
	inputFilePtr := flag.String("input_file", "input.json", "The input JSON file to be process, defaults to input.json")
	outputFilePtr := flag.String("output_file", "output.json", "The output JSON file, defaults to output.json")
	postcodePtr := flag.String("query_postcode", "10161", "postcode to query, defaults to 10161")
	postcodeFromPtr := flag.String("query_FromHr", "9AM", "postcode to query from hour, defaults to 9AM")
	postcodeToPtr := flag.String("query_ToHr", "11PM", "postcode to query from hour, defaults to 11PM")

	var words ArrayFlags
	flag.Var(&words, "word_list", "query recipe containing these words")

	flag.Parse()

	inputFileArg := *inputFilePtr
	outputFileArg := *outputFilePtr
	postcodeArg := *postcodePtr
	fromHrArg := *postcodeFromPtr
	toHrArg := *postcodeToPtr

	fmt.Println("Your arguments -")
	fmt.Println("  inputFileArg :", inputFileArg)
	fmt.Println("  outputFileArg:", outputFileArg)
	fmt.Println("  postcodeArg  :", postcodeArg)
	fmt.Println("  fromHrArg    :", fromHrArg)
	fmt.Println("  toHrArg      :", toHrArg)
	fmt.Printf("  words        : %v\n", words)

	// maps recipes to counts for doInput to update
	var recipesMap map[string]int
	recipesMap = make(map[string]int)

	// maps postcodes to counts for doInput to update
	var postcodesMap map[string]int
	postcodesMap = make(map[string]int)

	// setting postCodePerformance for doInput to update
	postCodePerformance := &PostcodePerformanceTimerange{Postcode: postcodeArg, From: fromHrArg, To: toHrArg, Count: 0}

	doInput(inputFileArg, recipesMap, postcodesMap, postCodePerformance)

	// words provide Search recipes containing word(s)	e.g. []string{"Chicken", "Pork"}
	doOutput(outputFileArg, recipesMap, postcodesMap, postCodePerformance, words)

}
