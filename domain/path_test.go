package domain

import (
	"strconv"
	"strings"
	"testing"
	"trip-route/services/utils"
)

type TestDataItem struct {
	from  string
	to    string
	path  string
	price uint64
	data  [][]string
}

func MockData3Airports() (Routes, []TestDataItem) {
	data := [][]string{
		{"BRC", "SCL", "5"},
		{"BRC", "ORL", "6"},
		{"SCL", "ORL", "20"},
	}

	dataTestItems := []TestDataItem{
		{"BRC", "SCL", "BRC - SCL", 5, data},
		{"BRC", "ORL", "BRC - ORL", 6, data},
		{"SCL", "ORL", "SCL - ORL", 20, data},
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
		{"GRU", "BRC", "GRU - BRC", 10, data},
		{"GRU", "SCL", "GRU - BRC - SCL", 15, data},
		{"GRU", "ORL", "GRU - BRC - ORL", 16, data},
		{"GRU", "CDG", "GRU - BRC - ORL - CDG", 21, data},
		{"BRC", "SCL", "BRC - SCL", 5, data},
		{"BRC", "ORL", "BRC - ORL", 6, data},
		{"BRC", "CDG", "BRC - ORL - CDG", 11, data},
		{"SCL", "ORL", "SCL - ORL", 20, data},
		{"SCL", "CDG", "SCL - ORL - CDG", 25, data},
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
		{"GRU", "BRC", "GRU - BRC", 10, data},
		{"GRU", "SCL", "GRU - BRC - SCL", 15, data},
		{"GRU", "ORL", "GRU - BRC - SCL - ORL", 35, data},
		{"GRU", "CDG", "GRU - BRC - SCL - ORL - CDG", 40, data},
		{"GRU", "CDG", "GRU - BRC - SCL - ORL - CDG", 40, data},
		{"GRU", "CPH", "GRU - BRC - SCL - ORL - CPH", 235, data},
		{"GRU", "SXF", "GRU - BRC - SCL - ORL - CPH - SXF", 240, data},
		{"GRU", "FRA", "GRU - BRC - SCL - ORL - CPH - FRA", 246, data},
		{"GRU", "BLL", "GRU - BRC - SCL - ORL - CPH - BLL", 314, data},
		{"GRU", "TXL", "GRU - BRC - SCL - ORL - CPH - BLL - TXL", 389, data},
		{"BRC", "SCL", "BRC - SCL", 5, data},
		{"BRC", "ORL", "BRC - SCL - ORL", 25, data},
		{"BRC", "CDG", "BRC - SCL - ORL - CDG", 30, data},
		{"BRC", "CPH", "BRC - SCL - ORL - CPH", 225, data},
		{"BRC", "SXF", "BRC - SCL - ORL - CPH - SXF", 230, data},
		{"BRC", "FRA", "BRC - SCL - ORL - CPH - FRA", 236, data},
		{"BRC", "BLL", "BRC - SCL - ORL - CPH - BLL", 304, data},
		{"BRC", "TXL", "BRC - SCL - ORL - CPH - BLL - TXL", 379, data},
		{"SCL", "ORL", "SCL - ORL", 20, data},
		{"SCL", "CDG", "SCL - ORL - CDG", 25, data},
		{"SCL", "CPH", "SCL - ORL - CPH", 220, data},
		{"SCL", "SXF", "SCL - ORL - CPH - SXF", 225, data},
		{"SCL", "FRA", "SCL - ORL - CPH - FRA", 231, data},
		{"SCL", "BLL", "SCL - ORL - CPH - BLL", 299, data},
		{"SCL", "TXL", "SCL - ORL - CPH - BLL - TXL", 374, data},
		{"ORL", "CDG", "ORL - CDG", 5, data},
		{"ORL", "CPH", "ORL - CPH", 200, data},
		{"ORL", "SXF", "ORL - CPH - SXF", 205, data},
		{"ORL", "FRA", "ORL - CPH - FRA", 211, data},
		{"ORL", "BLL", "ORL - CPH - BLL", 279, data},
		{"ORL", "TXL", "ORL - CPH - BLL - TXL", 354, data},
		{"CPH", "SXF", "CPH - SXF", 5, data},
		{"CPH", "FRA", "CPH - FRA", 11, data},
		{"CPH", "BLL", "CPH - BLL", 79, data},
		{"CPH", "TXL", "CPH - BLL - TXL", 154, data},
		{"BLL", "TXL", "BLL - TXL", 75, data},
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

func TestConnectionsAdded(t *testing.T) {
	routes5, _ := MockData5Airports()

	if len(routes5.connections) != 7 {
		t.Errorf("AddConnection() FAILED, expected 7 connections, got %d", len(routes5.connections))
	} else {
		t.Logf("AddConnection() PASSED, expected 7 connections, got %d", len(routes5.connections))
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

		hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, []string{}, 0)

		if len(routes.airports) != 3 {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected 3 aiports, got %d", item.from, item.to, len(routes.airports))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected path '%v' but got '%v', expected price %d but got %d, expected lastDestination to be true but got %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSED, expected path '%v' and got '%v', expected price %d and got %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}

	routes, dataTestItems = MockData5Airports()

	for _, item := range dataTestItems {

		hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, []string{}, 0)

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

		hasLastDestination, path, price := routes.BestPriceRoute(&Airport{item.from}, &Airport{item.to}, []string{}, 0)

		if len(routes.airports) != 10 {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected 10 aiports, got %d", item.from, item.to, len(routes.airports))
		} else if item.price != price || strings.Join(path, " - ") != item.path || hasLastDestination != true {
			t.Errorf("BestPriceRoute() [%v->%v] FAILED, expected path '%v' but got '%v', expected price %d but got %d, expected lastDestination to be true but got %v", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price, hasLastDestination)
		} else {
			t.Logf("BestPriceRoute() [%v->%v] PASSED, expected path '%v' and got '%v', expected price %d and got %d", item.from, item.to, item.path, strings.Join(path, " - "), item.price, price)
		}
	}
}
