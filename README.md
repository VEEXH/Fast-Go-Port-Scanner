# Fast-Go-Port-Scanner
This is a simple command-line port scanner written in Go. It scans for open ports on a given IP address.

## Prerequisites
`Go` installed on your system.
Usage
Open your terminal and navigate to the directory where the `scan.go` file is located.

Run the following command to build and run the program: `go run scan.go <ip>`

Replace `<ip>` with the IP address you want to scan for open ports.

The program will scan for open ports on the specified IP address and display the results in the console.

## How it works
The program uses the `net` package in Go to establish a TCP connection with each port number on the specified IP address. If the connection is successful, it means that the port is open, and the program prints a message to the console indicating that the port is open.

To make the program more user-friendly, an animation is displayed in the console while the program is running. The animation shows a spinning line that indicates that the program is still working.

The program uses a `sync.WaitGroup` to ensure that all the goroutines finish executing before the program exits.

## Limitations
The program uses a simple algorithm to scan for open ports, which means that it may take a long time to complete if the IP address has a large number of ports. Additionally, the program only scans for open TCP ports, which means that it may miss other types of open ports. Finally, the program does not provide any information about the services running on the open ports.
