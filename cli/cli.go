package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	domain "trip-route/domain"
	csvparser "trip-route/services/csv"
	utils "trip-route/services/utils"
)

func main() {
	// Validate args
	if len(os.Args) != 2 {
		fmt.Println("[ERROR] A CSV input file path is required!")
		os.Exit(1)
	} else {
		matched, err := regexp.MatchString(`^.*\.csv$`, os.Args[1])
		if err != nil || !matched {
			fmt.Printf("[ERROR] A CSV input file path must be inserted instead of %v\n", os.Args[1])
			os.Exit(1)
		}
	}

	// Validate file data read
	err, data := csvparser.Read(os.Args[1])
	if err != nil {
		fmt.Printf("[ERROR] Not possible to read the csv file due to error: '%v'\n", err)
		os.Exit(1)
	}

	// Load data to memory
	routes := domain.Routes{}
	for _, connection := range data {
		if !routes.HasConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1])) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1]), price)
		}
	}

	fmt.Println("=========================================================================================")
	fmt.Println("starting trip-route program. if you wish to exit anytime, just type exit or press ctrl+C.")
	fmt.Println("=========================================================================================")

	for {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("\nplease enter the route: ")
		sentence, err := buf.ReadBytes('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			splitRoutes := strings.Split(string(sentence), "-")
			if len(splitRoutes) == 2 {

				splitRoutes[0] = utils.TrimAndUpper(splitRoutes[0])
				splitRoutes[1] = utils.TrimAndUpper(splitRoutes[1])
				from := domain.NewAirport(splitRoutes[0])
				to := domain.NewAirport(splitRoutes[1])

				if routes.FindAirportByCode(splitRoutes[0]) == nil {
					fmt.Printf("origin %v not found. try inserting an existent origin or add a route with this origin.", splitRoutes[0])
				} else if routes.FindAirportByCode(splitRoutes[1]) == nil {
					fmt.Printf("destination %v not found. try inserting an existent destination or add a route with this destination.", splitRoutes[1])
				} else {
					_, path, price := routes.BestPriceRoute(from, to, []string{}, 0)
					pathString := strings.Join(path, " - ")

					if len(path) == 1 && price == 0 {
						fmt.Printf("path not found for %v->%v. try again", splitRoutes[0], splitRoutes[1])
					} else {
						fmt.Printf("best route: %v > $%v", pathString, price)
					}
				}

			} else {
				fmt.Printf("invalid input %v. format must be XXX-XXX, try again", strings.Trim(string(sentence), "\t \n"))
			}
		}
		if string(sentence) == "exit" {
			break
		}
	}
}
