package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s) // Parse URL and check for errors
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // Get Schema

	fmt.Println(u.User)            // Get User, Username and Password
	fmt.Println(u.User.Username()) // Get Username
	p, _ := u.User.Password()      // Get Password
	fmt.Println(p)

	fmt.Println(u.Host)                        // Get Host, Hostname and Port
	host, port, _ := net.SplitHostPort(u.Host) // Split host to hostname and Port
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)     // Get Path
	fmt.Println(u.Fragment) // Get fragment

	fmt.Println(u.RawQuery)            // Get Key-Value pairs
	m, _ := url.ParseQuery(u.RawQuery) // Put into Map
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
