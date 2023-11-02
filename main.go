package main

import (
	"Paginations/dbConfig"
	"Paginations/apiHandlers"
	"fmt"
	"log"
	"net"
	"errors"
	"github.com/gofiber/fiber/v2"
	//"github.com/joho/godotenv"
)

func init() {
	//err := godotenv.Load("utils/.env")
	//if err != nil {
	//	log.Fatal("Cannot load environment file")
	//}
}

func main() {
	fmt.Println("Starting Application")

	app := fiber.New(fiber.Config{
		AppName: "Paginations",
	})

	// Connect To Database
	dbConfig.ConnectToMongoDB()

	//Remove Pre-Generated Outs
	dbConfig.RemoveGeneratedOuts()

   

	// Define the API routes
	apiHandlers.Router(app)

	// Start the server
	log.Fatal(app.Listen(":8888"))
}


func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network")
}
