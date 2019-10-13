package helpers

import (
	"math/rand" // There is no need to specify seed. We do actually care about reproducible randomness.
	"net"
)

var devices = map[int]string{
	0:  "Unknown",
	1:  "iPhone 6",
	2:  "iPhone 7",
	3:  "iPhone 8",
	4:  "iPhone X",
	5:  "iPhone XS",
	6:  "Samsung Galaxy S6",
	7:  "Samsung Galaxy S7",
	8:  "Samsung Galaxy S8",
	9:  "Samsung Galaxy S9",
	10: "Samsung Galaxy S10",
}

type country struct {
	name string
	cities []string
}

var countries = map[int]country{
	0: country{"Germany", []string{"Berlin", "Frankfurt", "Hamburg", "Munich"}},
	1: country{"France", []string{"Paris", "Marseille", "Toulouse", "Lyon"}},
	2: country{"Italy", []string{"Rome", "Milan"}},
	3: country{"Russia", []string{"Moscow", "Saint Petersburg"}},
	4: country{"Poland", []string{"Krakow", "Warszawa", "Szczebrzeszyn"}},
	5: country{"Czech", []string{"Prague", "Brno"}},
}

var events = map[int]string{
	0: "BUY_BTC",
	1: "BUY_ETH",
	2: "SELL_BTC",
	3: "SELL_ETH",
}

// Device returns name of Mobile Device
func Device() string {
	return devices[randInt(0, len(devices))]
}

// Location consists of Country + City
func Location() (string, string) {
	country := countries[randInt(0, len(countries))]
	return country.name, country.cities[randInt(0, len(country.cities))]
}

// Event returns what action did user take.
func Event() string {
	return events[randInt(0, len(events))]
}

// IPV6Address returns string representation of IP address.
func IPV6Address() string {
	var ip net.IP
	for i := 0; i < net.IPv6len; i++ {
		number := uint8(randInt(0, 255))
		ip = append(ip, number)
	}
	return ip.String()
}

func randInt(min int, max int) int {
	return min + rand.Intn(max)
}
