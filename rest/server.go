package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	domain "trip-route/domain"
	csvparser "trip-route/services/csv"
	"trip-route/services/utils"
)

type SearchResponse struct {
	From  string
	To    string
	Path  string
	Price uint64
}

type Connection struct {
	From  string
	To    string
	Price uint64
}

type RoutesWrapper struct {
	routes domain.Routes
	data   [][]string
}

func main() {
	routesWrapper := RoutesWrapper{domain.Routes{}, nil}
	loadDataFromCSV(&routesWrapper)

	fmt.Println("[INFO] Starting trip route server on route 8080...")
	http.HandleFunc("/", ApiStatus)
	http.Handle("/route", &routesWrapper)
	http.ListenAndServe(":8080", nil)
}

func (routesWrapper *RoutesWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Search best route
		// Extract all values array of the query param
		fromParams := r.URL.Query()["from"]
		toParams := r.URL.Query()["to"]

		if len(fromParams) == 1 && len(toParams) == 1 {
			fromParams[0] = utils.TrimAndUpper(fromParams[0])
			toParams[0] = utils.TrimAndUpper(toParams[0])
			fromAirport := routesWrapper.routes.FindAirportByCode(fromParams[0])
			toAirport := routesWrapper.routes.FindAirportByCode(toParams[0])

			if fromAirport == nil {
				errorMsg := "Origin '" + fromParams[0] + "' not found. Try inserting an existent origin or add a route with this origin."
				fmt.Printf("\n[ERROR] %v | %v", getTimeNowFormatted(), errorMsg)
				http.Error(w, errorMsg, http.StatusBadRequest)
				return
			} else if toAirport == nil {
				errorMsg := "Destination '" + toParams[0] + "' not found. Try inserting an existent destination or add a route with this destination."
				fmt.Printf("\n[ERROR] %v | %v",
					getTimeNowFormatted(), errorMsg)
				http.Error(w, errorMsg, http.StatusBadRequest)
				return
			} else {
				_, _, path, price := routesWrapper.routes.BestPriceRoute(fromAirport, toAirport, fromAirport, []string{}, 0)
				pathString := strings.Join(path, " - ")

				if len(path) == 1 && price == 0 {
					errorMsg := "Path not found for '" + fromParams[0] + "->" + toParams[0] + "'. Try again with a proper path."
					fmt.Printf("\n[ERROR] '%v' | %v",
						getTimeNowFormatted(), errorMsg)
					http.Error(w, errorMsg, http.StatusBadRequest)
					return
				} else {
					js, err := json.Marshal(SearchResponse{fromParams[0], toParams[0], pathString, price})
					if err != nil {
						fmt.Printf("\n[ERROR] %v | %v",
							getTimeNowFormatted(), err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Header().Set("Content-Type", "application/json")
					w.Write(js)
				}
			}
		} else {
			errorMsg := "Incorrect or missing params passed."
			fmt.Printf("\n[ERROR] %v | %v",
				getTimeNowFormatted(), errorMsg)
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

	case http.MethodPost:
		// Insert new route
		var conn Connection

		err := json.NewDecoder(r.Body).Decode(&conn)
		if err != nil {
			fmt.Printf("\n[ERROR] %v | %v",
				getTimeNowFormatted(), err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if connection exists
		from := domain.NewAirport(utils.TrimAndUpper(conn.From))
		to := domain.NewAirport(utils.TrimAndUpper(conn.To))
		if routesWrapper.routes.HasConnection(from, to) {
			errorMsg := "Connection '" + conn.From + "->" + conn.To + "' already registered."
			fmt.Printf("\n[ERROR] %v | %v",
				getTimeNowFormatted(), errorMsg)
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

		// Add to memory new object
		routesWrapper.routes.AddConnection(from, to, conn.Price)
		routesWrapper.data = append(routesWrapper.data, []string{conn.From, conn.To, strconv.FormatUint(conn.Price, 10)})

		// Add to CSV file
		defaultCsvFilePath := "./data/input-routes.csv"
		err, newData := csvparser.CreateWrite(defaultCsvFilePath, routesWrapper.data)
		if err != nil {
			fmt.Printf("[ERROR] %v | Not possible to write the csv file due to error: '%v'\n", getTimeNowFormatted(), err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Updates the memory data with the return
		routesWrapper.data = newData

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	default:
		errorMsg := "Method not allowed."
		fmt.Printf("\n[ERROR] %v | %v",
			getTimeNowFormatted(), errorMsg)
		http.Error(w, errorMsg, http.StatusMethodNotAllowed)
		return
	}
}

func ApiStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Trip route server")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func loadDataFromCSV(routesWrapper *RoutesWrapper) {
	defaultCsvFilePath := "./data/input-routes.csv"

	// Validate file data read
	err, data := csvparser.Read(defaultCsvFilePath)
	if err != nil {
		fmt.Printf("[ERROR] %v | Not possible to read the csv file due to error: '%v'\n", getTimeNowFormatted(), err)
		os.Exit(1)
	}

	// Store temporarily on memory
	routesWrapper.data = data

	for _, connection := range data {
		if !routesWrapper.routes.HasConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1])) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routesWrapper.routes.AddConnection(domain.NewAirport(connection[0]), domain.NewAirport(connection[1]), price)
		}
	}
}

func getTimeNowFormatted() string {
	datetime := time.Now()
	return datetime.Format("02-01-2006 15:04:05")
}
