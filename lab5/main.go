package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
)

func main() {
    version := os.Getenv("VERSION")
    if version == "" {
        version = "dev"
    }

    hostname, _ := os.Hostname()

    addrs, _ := net.InterfaceAddrs()
    var ip string
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
            ip = ipnet.IP.String()
            break
        }
    }

    htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="pl">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>docker_lab5</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					color: #333;
					padding: 20px;
				}
				.container {
					background: #fff;
					border-radius: 8px;
					box-shadow: 0 2px 4px rgba(0,0,0,0.1);
					padding: 20px;
					max-width: 600px;
					margin: auto;
				}
				h1 {
					color: #0066cc;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Informacje o serwerze</h1>
				<ul>
					<li><strong>Adres IP serwera:</strong> %s</li>
					<li><strong>Nazwa serwera (hostname):</strong> %s</li>
					<li><strong>Wersja aplikacji:</strong> %s</li>
				</ul>
			</div>
		</body>
		</html>
`, ip, hostname, version)

    err := ioutil.WriteFile("index.html", []byte(htmlContent), 0644)
    if err != nil {
        fmt.Printf("Błąd zapisu pliku index.html: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Plik index.html został wygenerowany.")
}

