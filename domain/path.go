package domain

import (
	"errors"
	"sort"
)

// Airport as vertex
type Airport struct {
	code string
}

// Connection as edge
type Connection struct {
	from  *Airport
	to    *Airport
	price uint64
}

type PathConnectionPrice struct {
	airportCodePath    []string
	totalPrice         uint64
	hasLastDestination bool
}

// Routes as graph structure
type Routes struct {
	connections []*Connection
	airports    []*Airport
}

const Infinity = uint64(^uint64(0) >> 1)

func New() *Routes {
	return &Routes{
		airports:    []*Airport{},
		connections: []*Connection{},
	}
}

func NewAirport(code string) *Airport {
	return &Airport{code}
}

func (r *Routes) AddConnection(from *Airport, to *Airport, price uint64) {
	connection := &Connection{
		from:  from,
		to:    to,
		price: price,
	}

	r.connections = append(r.connections, connection)
	r.AddAirport(from)
	r.AddAirport(to)
}

func (r *Routes) GetAllConnections() []*Connection {
	return r.connections
}

func (r *Routes) GetAllAirports() []*Airport {
	return r.airports
}

func (r *Routes) HasConnection(from *Airport, to *Airport) bool {
	for _, c := range r.connections {
		if c.from.code == from.code && c.to.code == to.code {
			return true
		}
	}
	return false
}

func (r *Routes) AddAirport(airport *Airport) {
	var isAirportPresent bool
	for _, a := range r.airports {
		if a.code == airport.code {
			isAirportPresent = true
		}
	}
	if !isAirportPresent {
		r.airports = append(r.airports, airport)
	}
}

func (r *Routes) FindAirportByCode(code string) *Airport {
	var airport *Airport
	airport = nil
	for _, a := range r.airports {
		if a.code == code {
			airport = a
		}
	}
	return airport
}

func (r *Routes) GetConnectionsFromAirport(airport *Airport) (connections []*Connection) {
	for _, connection := range r.connections {
		if connection.from.code == airport.code {
			connections = append(connections, connection)
		}
	}
	return connections
}

func (r *Routes) BestPriceRoute(from *Airport, to *Airport, origin *Airport, accumulatedPath []string, accumulatedPrice uint64) (error, bool, []string, uint64) {
	if from.code == to.code && from.code == origin.code && to.code == origin.code {
		return errors.New("No circular path is allowed."), false, nil, 0
	}

	pathConnectionPrices := []*PathConnectionPrice{}

	// Get all connections from airport node
	connections := r.GetConnectionsFromAirport(from)

	if from.code == to.code && from.code != origin.code {
		// Reached destination, but it might be not the end of the path. We end the path here
		return nil, true, append(accumulatedPath, from.code), accumulatedPrice
	} else if len(connections) == 0 {
		// Reached last node point. Should not continue
		return nil, false, append(accumulatedPath, from.code), accumulatedPrice
	} else {
		// Loop through each connection to get their path price
		for _, conn := range connections {
			if origin.code != to.code && conn.to.code == origin.code {
				// Avoid circular loop
				continue
			}
			_, hasLastDestination, recursionPath, recursionPrice := r.BestPriceRoute(conn.to, to, origin, append(accumulatedPath, from.code), (accumulatedPrice + conn.price))

			// Append path only if it has the last node destination
			if hasLastDestination {
				pathConnectionPrices = append(pathConnectionPrices, &PathConnectionPrice{recursionPath, recursionPrice, hasLastDestination})
			}
		}

		// Loop path connection prices last returns and check
		// If it has, check if path price is cheaper than the others (sort)
		sort.Slice(pathConnectionPrices, func(i, j int) bool {
			return pathConnectionPrices[i].totalPrice < pathConnectionPrices[j].totalPrice
		})

		if len(pathConnectionPrices) == 0 {
			return nil, false, append(accumulatedPath, from.code), accumulatedPrice
		} else {
			return nil, pathConnectionPrices[0].hasLastDestination, pathConnectionPrices[0].airportCodePath, pathConnectionPrices[0].totalPrice
		}
	}
}
