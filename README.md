# Mux RouteRegistry

Solves the problem of easily registering routes and subroutes for different modules when using mux.

## Installation

```
go get -u github.com/cruzj6/muxRouteRegistry.go
```

## Usage
```
const BASE_ROUTE = '/someModule/'

// someModule.go
func RegisterRoutes() {

	baseRoute := routes.MakeRoute(BASE_ROUTE, "GET", baseHandler)
	dataRoute := routes.MakeRoute("/data/", "GET", dataHandler)

	baseRoute.AddSubRoute(dataRoute)
	routes.RegisterRoute(baseRoute)
}

// main.go
func main() {
  someModule.RegisterRoutes()

  routesHandler := routes.SetupRoutes(mux.NewRouter())

  srv := &http.Server{
		Handler: routesHandler,
		Addr: ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	 }

	srv.ListenAndServe()
}
```
