//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerInfo(servers map[string]int) {
	statusCount := make(map[int]int)

	for _, status := range servers {
		statusCount[status]++
	}

	fmt.Println("")
	fmt.Println("Number of servers:", len(servers))
	fmt.Println("    Online", statusCount[Online])
	fmt.Println("    Offline", statusCount[Offline])
	fmt.Println("    Maintenance", statusCount[Maintenance])
	fmt.Println("    Retired", statusCount[Retired])
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)
	for _, name := range servers {
		serverStatus[name] = Online
	}

	//  - call display server info function
	printServerInfo(serverStatus)

	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired

	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline

	//  - call display server info function
	printServerInfo(serverStatus)

	//  - change server status of all servers to `Maintenance`
	for name, _ := range serverStatus {
		serverStatus[name] = Maintenance
	}

	//  - call display server info function
	printServerInfo(serverStatus)
}
