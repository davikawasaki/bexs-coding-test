package domain

import (
	"strconv"
	"strings"
	"testing"
	"trip-route/services/utils"
)

type TestDataItem struct {
	from     string
	to       string
	path     string
	price    uint64
	data     [][]string
	hasError bool
}

func MockData3AirportsCircularData() (Routes, []TestDataItem) {
	data := [][]string{
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"SCL", "ORL", "20"},
		{"SCL", "BRC", "5"},
	}

	dataTestItems := []TestDataItem{
		{"BRC", "BRC", "No circular path is allowed.", 0, nil, true},
		{"SCL", "SCL", "No circular path is allowed.", 0, nil, true},
		{"ORL", "ORL", "No circular path is allowed.", 0, nil, true},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func MockData3Airports() (Routes, []TestDataItem) {
	data := [][]string{
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"SCL", "ORL", "20"},
		{"SCL", "BRC", "5"},
	}

	dataTestItems := []TestDataItem{
		{"BRC", "SCL", "BRC - SCL", 5, data, false},
		{"BRC", "ORL", "BRC - ORL", 6, data, false},
		{"SCL", "ORL", "SCL - BRC - ORL", 11, data, false},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func MockData5Airports() (Routes, []TestDataItem) {
	data := [][]string{
		{"GRU", "BRC", "10"},
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"GRU", "CDG", "75"},
		{"GRU", "SCL", "20"},
		{"GRU", "ORL", "56"},
		{"ORL", "CDG", "5"},
		{"SCL", "ORL", "20"},
	}

	dataTestItems := []TestDataItem{
		{"GRU", "BRC", "GRU - BRC", 10, data, false},
		{"GRU", "SCL", "GRU - BRC - SCL", 15, data, false},
		{"GRU", "ORL", "GRU - BRC - ORL", 16, data, false},
		{"GRU", "CDG", "GRU - BRC - ORL - CDG", 21, data, false},
		{"BRC", "SCL", "BRC - SCL", 5, data, false},
		{"BRC", "ORL", "BRC - ORL", 6, data, false},
		{"BRC", "CDG", "BRC - ORL - CDG", 11, data, false},
		{"SCL", "ORL", "SCL - ORL", 20, data, false},
		{"SCL", "CDG", "SCL - ORL - CDG", 25, data, false},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func MockData10Airports() (Routes, []TestDataItem) {
	data := [][]string{
		{"GRU", "BRC", "10"},
		{"BRC", "SCL", "5"},
		{"GRU", "CDG", "75"},
		{"GRU", "SCL", "20"},
		{"GRU", "ORL", "56"},
		{"ORL", "CDG", "5"},
		{"SCL", "ORL", "20"},
		{"ORL", "CPH", "200"},
		{"CPH", "BLL", "79"},
		{"CPH", "FRA", "11"},
		{"CPH", "SXF", "5"},
		{"BLL", "TXL", "75"},
		{"GRU", "BLL", "955"},
	}

	dataTestItems := []TestDataItem{
		{"GRU", "BRC", "GRU - BRC", 10, data, false},
		{"GRU", "SCL", "GRU - BRC - SCL", 15, data, false},
		{"GRU", "ORL", "GRU - BRC - SCL - ORL", 35, data, false},
		{"GRU", "CDG", "GRU - BRC - SCL - ORL - CDG", 40, data, false},
		{"GRU", "CDG", "GRU - BRC - SCL - ORL - CDG", 40, data, false},
		{"GRU", "CPH", "GRU - BRC - SCL - ORL - CPH", 235, data, false},
		{"GRU", "SXF", "GRU - BRC - SCL - ORL - CPH - SXF", 240, data, false},
		{"GRU", "FRA", "GRU - BRC - SCL - ORL - CPH - FRA", 246, data, false},
		{"GRU", "BLL", "GRU - BRC - SCL - ORL - CPH - BLL", 314, data, false},
		{"GRU", "TXL", "GRU - BRC - SCL - ORL - CPH - BLL - TXL", 389, data, false},
		{"BRC", "SCL", "BRC - SCL", 5, data, false},
		{"BRC", "ORL", "BRC - SCL - ORL", 25, data, false},
		{"BRC", "CDG", "BRC - SCL - ORL - CDG", 30, data, false},
		{"BRC", "CPH", "BRC - SCL - ORL - CPH", 225, data, false},
		{"BRC", "SXF", "BRC - SCL - ORL - CPH - SXF", 230, data, false},
		{"BRC", "FRA", "BRC - SCL - ORL - CPH - FRA", 236, data, false},
		{"BRC", "BLL", "BRC - SCL - ORL - CPH - BLL", 304, data, false},
		{"BRC", "TXL", "BRC - SCL - ORL - CPH - BLL - TXL", 379, data, false},
		{"SCL", "ORL", "SCL - ORL", 20, data, false},
		{"SCL", "CDG", "SCL - ORL - CDG", 25, data, false},
		{"SCL", "CPH", "SCL - ORL - CPH", 220, data, false},
		{"SCL", "SXF", "SCL - ORL - CPH - SXF", 225, data, false},
		{"SCL", "FRA", "SCL - ORL - CPH - FRA", 231, data, false},
		{"SCL", "BLL", "SCL - ORL - CPH - BLL", 299, data, false},
		{"SCL", "TXL", "SCL - ORL - CPH - BLL - TXL", 374, data, false},
		{"ORL", "CDG", "ORL - CDG", 5, data, false},
		{"ORL", "CPH", "ORL - CPH", 200, data, false},
		{"ORL", "SXF", "ORL - CPH - SXF", 205, data, false},
		{"ORL", "FRA", "ORL - CPH - FRA", 211, data, false},
		{"ORL", "BLL", "ORL - CPH - BLL", 279, data, false},
		{"ORL", "TXL", "ORL - CPH - BLL - TXL", 354, data, false},
		{"CPH", "SXF", "CPH - SXF", 5, data, false},
		{"CPH", "FRA", "CPH - FRA", 11, data, false},
		{"CPH", "BLL", "CPH - BLL", 79, data, false},
		{"CPH", "TXL", "CPH - BLL - TXL", 154, data, false},
		{"BLL", "TXL", "BLL - TXL", 75, data, false},
	}

	routes := Routes{}
	for _, connection := range data {
		if !routes.HasConnection(&Airport{connection[0]}, &Airport{connection[1]}) {
			price, err := strconv.ParseUint(connection[2], 10, 64)
			if err != nil {
				price = 0
			}
			routes.AddConnection(&Airport{connection[0]}, &Airport{connection[1]}, price)
		}
	}

	return routes, dataTestItems
}

func TestNew(t *testing.T) {
	routes := Routes{}

	if len(routes.airports) == 0 && len(routes.connections) == 0 {
		t.Logf("New() PASSED, expected 0 connections and 0 airports and got %v connections and %v airports.",
			routes.connections, routes.airports)
	} else {
		t.Errorf("New() FAILED, expected 0 connections and 0 airports but got %v connections and %v airports.",
			routes.connections, routes.airports)
	}
}

func TestNewAirport(t *testing.T) {
	airportCode := "GRU"
	newAirport := NewAirport(airportCode)
	if newAirport.code != airportCode {
		t.Errorf("NewAirport() FAILED, expected code '%v' but got '%v'", airportCode, newAirport.code)
	} else {
		t.Logf("NewAirport() PASSED, expected code '%v' and got '%v'", airportCode, newAirport.code)
	}
}

func TestGetAllConnections(t *testing.T) {
	routes5, _ := MockData5Airports()

	connections := routes5.GetAllConnections()

	if len(connections) != 8 {
		t.Errorf("GetAllConnections() FAILED, expected 8 connections, got %d", len(connections))
	} else {
		t.Logf("GetAllConnections() PASSED, expected 8 connections, got %d", len(connections))
	}

	routes10, _ := MockData10Airports()

	connections = routes10.GetAllConnections()

	if len(connections) != 13 {
		t.Errorf("GetAllConnections() FAILED, expected 13 connections, got %d", len(connections))
	} else {
		t.Logf("GetAllConnections() PASSED, expected 13 connections, got %d", len(connections))
	}
}

func TestGetAllAirports(t *testing.T) {
	routes5, _ := MockData5Airports()

	airports := routes5.GetAllAirports()

	if len(airports) != 5 {
		t.Errorf("GetAllConnections() FAILED, expected 5 airports, got %d", len(airports))
	} else {
		t.Logf("GetAllConnections() PASSED, expected 5 airports, got %d", len(airports))
	}

	routes10, _ := MockData10Airports()

	airports = routes10.GetAllAirports()

	if len(airports) != 10 {
		t.Errorf("GetAllConnections() FAILED, expected 10 airports, got %d", len(airports))
	} else {
		t.Logf("GetAllConnections() PASSED, expected 10 airports, got %d", len(airports))
	}
}

func TestHasConnection(t *testing.T) {
	routes5, _ := MockData5Airports()

	hasConn5 := routes5.HasConnection(&Airport{"GRU"}, &Airport{"BRC"})

	if !hasConn5 {
		t.Errorf("HasConnection() FAILED, expected to have connection between GRU->BRC, got %v", hasConn5)
	} else {
		t.Logf("HasConnection() PASSED, expected to have connection between GRU->BRC, got %v", hasConn5)
	}

	routes10, _ := MockData10Airports()

	hasConn10 := routes10.HasConnection(&Airport{"GRU"}, &Airport{"BLL"})

	if !hasConn10 {
		t.Errorf("HasConnection() FAILED, expected to have connection between GRU->BLL, got %v", hasConn10)
	} else {
		t.Logf("HasConnection() PASSED, expected to have connection between GRU->BLL, got %v", hasConn10)
	}
}

func TestAirportsAdded(t *testing.T) {
	routes5, _ := MockData5Airports()

	if len(routes5.airports) != 5 {
		t.Errorf("AddAirport() FAILED, expected 5 airports, got %d", len(routes5.airports))
	} else {
		t.Logf("AddAirport() PASSED, expected 5 airports, got %d", len(routes5.airports))
	}

	routes10, _ := MockData10Airports()

	if len(routes10.airports) != 10 {
		t.Errorf("AddAirport() FAILED, expected 10 airports, got %d", len(routes10.airports))
	} else {
		t.Logf("AddAirport() PASSED, expected 10 airports, got %d", len(routes10.airports))
	}
}

func TestFindAirportByCode(t *testing.T) {
	routes5, _ := MockData5Airports()

	airport := routes5.FindAirportByCode("GRU")

	if airport == nil {
		t.Errorf("FindAirportByCode() FAILED, expected GRU airport object, got %v", airport.code)
	} else {
		t.Logf("FindAirportByCode() PASSED, expected GRU airport object, got %v", airport.code)
	}

	routes10, _ := MockData10Airports()

	airport = routes10.FindAirportByCode("CPH")

	if airport == nil {
		t.Errorf("FindAirportByCode() FAILED, expected CPH airport object, got %v", airport.code)
	} else {
		t.Logf("FindAirportByCode() PASSED, expected CPH airport object, got %v", airport.code)
	}
}

func TestConnectionsAdded(t *testing.T) {
	routes5, _ := MockData5Airports()

	if len(routes5.connections) != 8 {
		t.Errorf("AddConnection() FAILED, expected 8 connections, got %d", len(routes5.connections))
	} else {
		t.Logf("AddConnection() PASSED, expected 8 connections, got %d", len(routes5.connections))
	}

	routes13, _ := MockData10Airports()

	if len(routes13.connections) != 13 {
		t.Errorf("AddConnection() FAILED, expected 13 connections, got %d", len(routes13.connections))
	} else {
		t.Logf("AddConnection() PASSED, expected 13 connections, got %d", len(routes13.connections))
	}
}

func TestConnectionsFromAirport(t *testing.T) {
	routes, _ := MockData5Airports()
	airport := &Airport{"GRU"}
	expectedConnections := []string{"BRC", "SCL", "ORL", "CDG"}
	connections := routes.GetConnectionsFromAirport(airport)
	var convertedConnections []string

	for _, conn := range connections {
		convertedConnections = append(convertedConnections, conn.to.code)
	}

	if len(expectedConnections) != len(convertedConnections) {
		t.Errorf("GetConnectionsFromAirport() FAILED, expected %d connections, got %d", len(expectedConnections), len(convertedConnections))
	} else if !utils.CompareStringArrays(convertedConnections, expectedConnections) {
		t.Errorf("GetConnectionsFromAirport() FAILED, expected %v, got %v", expectedConnections, convertedConnections)
	} else {
		t.Logf("GetConnectionsFromAirport() PASSED, expected %v, got %v", expectedConnections, convertedConnections)
	}
}

func TestBestPriceRoute(t *testing.T) {
	routes, dataTestItems := MockData3Airports()

	for _, item := range dataTestItems {

		_, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if len(routes.airports) != 3 {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected 3 aiports, got %d", item.from, item.to, len(routes.airports))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected path '%v' but got '%v', expected price %d but got %d, expected lastDestination to be true but got %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSED, expected path '%v' and got '%v', expected price %d and got %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}

	routes, dataTestItems = MockData3AirportsCircularData()

	for _, item := range dataTestItems {

		err, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if err != nil {
			t.Logf("BestPriceRoute() [%v->%v] PASSED, expected error '%v', got error '%v', path '%v', hasLastDestination '%v', price '%d'", item.from, item.to, item.path, err, path, hasLastDestination, price)
		} else {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected error '%v', got no error, path '%v', hasLastDestination '%v', price '%d'", item.from, item.to, item.path, path, hasLastDestination, price)
		}
	}

	routes, dataTestItems = MockData5Airports()

	for _, item := range dataTestItems {

		_, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if len(routes.airports) != 5 {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected 5 aiports, got %d", item.from, item.to, len(routes.airports))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected path '%v' but got '%v', expected price %d but got %d, expected lastDestination to be true but got %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSED, expected path '%v' and got '%v', expected price %d and got %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}

	routes, dataTestItems = MockData10Airports()

	for _, item := range dataTestItems {

		_, hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, &Airport{item.from}, []string{}, 0)

		if len(routes.airports) != 10 {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected 10 aiports, got %d", item.from, item.to, len(routes.airports))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected path '%v' but got '%v', expected price %d but got %d, expected lastDestination to be true but got %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSED, expected path '%v' and got '%v', expected price %d and got %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}
}
