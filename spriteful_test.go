package main

import (
	"os"
	"testing"

	"github.com/emicklei/go-restful"
)

var (
	validRoutes = []string{
		"/api/v1/boot/{mac-addr}",
		"/api/v1/static/{resource:*}",
	}
	validMac    = "00:00:00:00:00:00"
	invalidMac  = "00:00:00:00:00:01"
	testFile    = "/tmp/test"
	invalidFile = "someweirdfile"
)

func TestRouteRegister(t *testing.T) {
	s := &Spriteful{}
	c := restful.NewContainer()
	s.register(c)
	services := c.RegisteredWebServices()
	if serviceCount := len(services); serviceCount != 1 {
		t.Errorf("only one service is expected, services: %d", serviceCount)
	}
	service := services[0]
	routes := service.Routes()
	if routeCount := len(routes); routeCount != 2 {
		t.Errorf("only two routes are expected. routes: %d", routeCount)
	}

	for _, route := range routes {
		found := false
		for _, validRoute := range validRoutes {
			if validRoute == route.Path {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("route %s is not a valid route", route)
		}
	}
}

func TestFindServer(t *testing.T) {
	s := &Spriteful{
		Servers: []Server{
			{
				MacAddress: validMac,
			},
		},
	}
	if _, err := s.findServerConfig(validMac); err != nil {
		t.Errorf("%s config should be found, but it's not", validMac)
	}
	if _, err := s.findServerConfig(invalidMac); err == nil {
		t.Errorf("%s config should not be found, but it is", invalidMac)
	}
}

func TestFindResource(t *testing.T) {
	s := &Spriteful{}
	os.Create(testFile)
	if _, err := s.findResource(testFile); err != nil {
		t.Errorf("%s should be found, but it's not", testFile)
	}
	os.Remove(invalidFile)
	if _, err := s.findResource(invalidFile); err == nil {
		t.Errorf("%s should not be found, but it is", invalidFile)
	}
}
