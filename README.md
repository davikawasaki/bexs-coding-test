<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-90%25-brightgreen.svg?longCache=true&style=flat)</a>

# Bexs Bank Coding Test - Trip Route Search

Coding test developed for Bexs position as Senior Backend, with the goal of building a CLI & REST service for querying the cheapest trip route from origin to destination (details can be found [here](https://github.com/davikawasaki/bexs-coding-test/blob/master/ABOUT.md).

Requirements for the REST API:

- :heavy_check_mark: Search for best price route between two points
- :heavy_check_mark: Register new routes, persisting them into a *input-routes.csv* file

Requirements for the CLI:

- :heavy_check_mark: Start CLI with the *input-routes.csv* file as the initial argument
- :heavy_check_mark: Console interface with FROM-TO format input
- :heavy_check_mark: Console interface input result should output the full path and the total price for the best route

This solution was developed on Linux Ubuntu 16.0.4 using Go as the main language, net/http framework to server as HTTP server, gotest for colorful tests and gopherbadger to export a coverage badge. More details can be found in the [*go.mod* file](https://github.com/davikawasaki/bexs-coding-test/blob/master/go.mod).

## Table of Contents

<!--ts-->
  * [Usage](#usage)
    * [Command-line interface](#command-line-interface)
    * [REST interface](#rest-interface)
    * [Testing](#testing)
  * [Files and packages structures](#files-and-packages-structures)
  * [Design decisions](#design-decisions)
    * [Brief algorithm explanation](#brief-algorithm-explanation)
    * [Decisions over Go](#decisions-over-go)
    * [HTTP server](#http-server)
    * [CLI](#http-server)
    * [Improvements](#improvements)
  * [REST endpoints](#rest-endpoints)
<!--te-->

## Usage

- Make sure you have Go 1.14+ installed:

```bash
go version  # go version go1.14.4 linux/amd64
```

Libraries and dependencies that aren't native to Go will be installed automatically when running the project.

### Command-line interface

- To run the command-line interface (CLI), run the following command to start reading the [data/input-routes.csv](https://github.com/davikawasaki/bexs-coding-test/blob/master/data/input-routes.csv) file:

```bash
go run ./cli/input-routes.csv
```

- Afterwards, you may start the search typing the routes with the format ORIGINCODE-DESTINATIONCODE. Examples can be seen below, with success and wrong results:

```bash
=========================================================================================
starting trip-route program. if you wish to exit anytime, just type exit or press ctrl+C.
=========================================================================================

please enter the route: GRU-BRC
best route: GRU - BRC > $10
please enter the route: GRU-SCL
best route: GRU - BRC - SCL > $15
please enter the route: GRU-ORL
best route: GRU - BRC - ORL > $16
please enter the route: SCL-ORL
best route: SCL - ORL > $20
please enter the route: BRC-ORL
best route: BRC - ORL > $6
please enter the route: BRC-GRU
path not found for BRC->GRU. try again
please enter the route: BRCC
invalid input BRCC. format must be XXX-XXX, try again
please enter the route: BRCC-GRUU
origin BRCC not found. try inserting an existent origin or add a route with this origin.
please enter the route: GRUU-BRCC
origin GRUU not found. try inserting an existent origin or add a route with this origin.
please enter the route: 
invalid input . format must be XXX-XXX, try again
please enter the route: _-_
origin _ not found. try inserting an existent origin or add a route with this origin.
please enter the route: exit

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$
```

### REST interface

- To run the REST HTTP server, run the following command to start listening on port 8080 and reading the [data/input-routes.csv](https://github.com/davikawasaki/bexs-coding-test/blob/master/data/input-routes.csv) file:

```bash
go run ./rest
```

- Afterwards, try running the search in another terminal using cURL or any other HTTP-based application with the query params as ?from=ORIGINCODE&to=DESTINATIONCODE. Examples can be seen below on the first terminal:

```bash
[INFO] Starting trip route server on route 8080...

[ERROR] '22-06-2020 15:10:13' | Path not found for 'BRC->GRU'. Try again with a proper path.
[ERROR] 22-06-2020 15:10:24 | Origin 'BRCD' not found. Try inserting an existent origin or add a route with this origin.
[ERROR] 22-06-2020 15:10:28 | Destination 'GRUD' not found. Try inserting an existent destination or add a route with this destination.
[ERROR] 22-06-2020 15:10:36 | Incorrect or missing params passed.
[ERROR] 22-06-2020 15:10:41 | Incorrect or missing params passed.
[ERROR] 22-06-2020 15:11:56 | Connection 'BLL->SXF' already registered.
```

and in the second terminal:

```bash
kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XGET 'localhost:8080/route?from=GRU&to=BRC'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 22 Jun 2020 18:09:56 GMT
Content-Length: 55

{"From":"GRU","To":"BRC","Path":"GRU - BRC","Price":10}

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XGET 'localhost:8080/route?from=BRC&to=GRU'
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 22 Jun 2020 18:10:13 GMT
Content-Length: 61

Path not found for 'BRC->GRU'. Try again with a proper path.

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XGET 'localhost:8080/route?from=BRCD&to=GRU'
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 22 Jun 2020 18:10:24 GMT
Content-Length: 91

Origin 'BRCD' not found. Try inserting an existent origin or add a route with this origin.

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XGET 'localhost:8080/route?from=BRC&to=GRUD'
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 22 Jun 2020 18:10:28 GMT
Content-Length: 106

Destination 'GRUD' not found. Try inserting an existent destination or add a route with this destination.

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XGET 'localhost:8080/route?from=BRC&tdo=GRUD'
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 22 Jun 2020 18:10:36 GMT
Content-Length: 36

Incorrect or missing params passed.

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XGET 'localhost:8080/route?fdrom=BRC&to=GRUD'
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 22 Jun 2020 18:10:41 GMT
Content-Length: 36

Incorrect or missing params passed.

curl -i -XPOST "localhost:8080/route" -H "Content-Type: application/json" --data '{"from": "BLL", "to": "SXF", "price": 10}'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 22 Jun 2020 18:11:54 GMT
Content-Length: 0

kawasaki@kawasaki:~/Git/tests/bexs-coding-test$ curl -i -XPOST "localhost:8080/route" -H "Content-Type: application/json" --data '{"from": "BLL", "to": "SXF", "price": 10}'
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 22 Jun 2020 18:11:56 GMT
Content-Length: 42

Connection 'BLL->SXF' already registered.
```

### Testing

For this project, packages were splitted between services and the main business logic. Coverage can be seen at the [coverage folder](https://github.com/davikawasaki/bexs-coding-test/tree/master/coverage), as well as in the coverage badge in the beginning of this file.

To run all tests at once, run the following command (to run go test you may need to verify its respective bin location, which is **normally on Linux at $HOME/go/bin**):

```bash
go test ./... -v
$HOME/go/bin/gotest ./... -v
```

To run the utils tests (i.e. string and file parsers) tests, run the following command (to run go test you may need to verify its respective bin location, which is **normally on Linux at $HOME/go/bin**):

```bash
go test ./services/utils -v
$HOME/go/bin/gotest ./services/utils -v
```

To run the csv parser tests, run the following command (to run go test you may need to verify its respective bin location, which is **normally on Linux at $HOME/go/bin**):

```bash
go test ./services/csv -v
$HOME/go/bin/gotest ./services/csv -v
```

To run the domain tests, run the following command (to run go test you may need to verify its respective bin location, which is **normally on Linux at $HOME/go/bin**):

```bash
go test ./domain -v
$HOME/go/bin/gotest ./domain -v
```

## REST endpoints

- GET /: Check the status of the API. Returns a header with 200 status if server is up

Example:

```bash
curl -i -GET "localhost:8080"
```

- GET /route: Search for the best route as long as the mandatory query params are passed (e.g. localhost:8080/route?from=GRU&to=BRC):
  - from (mandatory): origin code (e.g. GRU, BRC)
  - to (mandatory): destination code (e.g. GRU, BRC)
  - Returns a JSON object with the origin, destination, full path and price (e.g. {"From":"GRU","To":"BRC","Path":"GRU - BRC","Price":10}). If an error happened, the respective HTTP status will be returned

Example:

```bash
curl -i -GET "localhost:8080/route?from=GRU&to=BRC"
```

- POST /route: Insert a new route passing the following arguments in the JSON object:
  - from (mandatory, string): origin code (e.g. GRU, BRC)
  - to (mandatory, string): destination code (e.g. GRU, BRC)
  - price (mandatory, int): path price (e.g. 20, 5)
  - Returns a header with 200 status if it was a success. If an error happened, the respective HTTP status will be returned

Example:

```bash
curl -i -XPOST "localhost:8080/route" -H "Content-Type: application/json" --data '{"from": "BLL", "to": "SXF", "price": 10}'
```

## Files and packages structures
    .
    ├── cli                     # Command-line interface package
    ├── coverage                # Coverage test txt and html files
    ├── data                    # Data source for the CLI and REST HTTP server
    ├── domain                  # Package for the business logic
    ├── rest                    # REST HTTP server interface package
    ├── services                # Package for tools and parser utilities (i.e. csv, string, file)
    ├── ABOUT.md
    ├── go.mod
    ├── LICENSE
    └── README.md

## Design decisions

### Brief algorithm explanation

This test appeared as a traveling salesman problem, which focus mainly on getting the shortest/cheapest distance from one point to another. Even though this was a system requirement, which would be solved fairly well with the [Dijkstra algorithm](https://brilliant.org/wiki/dijkstras-short-path-finder/), tracking the whole path and getting only the main path between A and B made the problem a little bit complex, since keeping track of path was a must.

Therefore, the algorithm solution uses recursion, always searching the cheapest path sniffing inside the inner nodes and identifying if it's a possible destination or if the inner path should be discarded (similar to the Depth First Search on tree searches). Throughout this process, the paths are accumulated in a array, while the total price is increased accordingly with the cheapest route price.

In order to split the responsibilities, the main route logic was separated into a package, exposed as an interface to the command-line interface and the REST HTTP server. At the same rate, other responsibilities were split into service packages, such as a CSV parser to read/write the input routes file, as well as string/file parsers that combines methods into utilities (e.g. trim and upper a string for the origin/destination code comparison).

### Decision over Go

The decision over Go instead of a main-used language in my stack such as Javascript was due to the following points:

- Less bloat: instead of importing multiple dependencies, Go can leverage almost everything with the native libs and leaving a source code that is less boilerplate-ish

- More flexibility: instead of having to crack my head between multiple interfaces, super hirearchical structures or even messy code, Go can find a common balance between object-oriented and functional paradigms through its compositions, structs and error handling

- Proximity to low-level coding: this can perceive that "objects" in Go are more lightweight due to the lack of high hierarchy, with direct indications such as pointers

- Cross-platform: working with Java and .NET, building to different OS systems and environments were always a hassle. Even with Javascript (i.e. web or even Electron) was somewhat problematic. Always a fan of code reusal to avoid rework (DRY principle), Go provided me this opportunity, besides keeping simple to even testing and starting up a HTTP server

- Oriented for DevOps: since my career is turning towards DevOps and Data Engineering, having a flexible tool such as Go that can be a huge difference in this area is astoundingly great :)

### HTTP server

For the HTTP server, just the net/http package was sufficient. Routes were splitted into two different functions (ApiStatus and ServeHTTP), where the ServeHTTP that servers the /route path has a wrapper composition around the *domain.Routes* struct. This allows a manipulation directly to the routes struct class from the domain package, allowing the server to properly do searches and insert new routes to the *input-routes.csv* file.

### CLI

For the command-line interface, no secret surrounds the implementation. A simple verification is being checked for the arguments, checking if the *input-routes.csv* file is being passed correctly, and then doing the continuous searches as long as the user doesn't type **exit** or hits **Ctrl+C** signal to interrupt the process.

### Improvements

Some points are still missing in the implementation due to lack of knowledge or time:

- Implement some integration tests for the CLI and REST server

- Handle proper status for incorrect or not-implemented routes (i.e. throw error 501 or 405)

- Split some utility functions from REST package (e.g. getTimeNowFormatted)

- Improve code readability in some points, especially in the REST package and the domain package (i.e. split function accountability for better code handling)

- Allow circular route paths (i.e. starting and returing to point A)

- Implement file lock usage when HTTP server or CLI are up

- Implement CI/CD using Github Actions (i.e. releasing versions and modules)
