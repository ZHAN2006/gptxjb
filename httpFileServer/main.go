package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var dir string = "."
var port int = 8081

func main() {
	// 使用 flag 包来处理命令行参数
	flag.IntVar(&port, "p", 8081, "specify the port to listen")
	flag.StringVar(&dir, "d", ".", "specify the directory to serve")
	showHelp := flag.Bool("h", false, "show help")

	flag.Parse()

	if *showHelp {
		printUsage()
		return
	}

	printServerInfo()
	startServer()
}

func processIP(ip string) {
	// You can add validation for the IP address if needed
	// For now, just set it as the specified IP
}

func processDir(d string) {
	stat, err := os.Stat(d)
	if err == nil && stat.IsDir() {
		dir = d
	} else {
		fmt.Println("Invalid directory specified.")
		os.Exit(1)
	}
}

func printServerInfo() {
	fmt.Printf("listen port is %d\n", port)
	fmt.Printf("| loc address : http://127.0.0.1:%d\n", port)
	fmt.Printf("| lan address : http://%s:%d\n", getIntranetIP(), port)
	fmt.Printf("| wan address : http://%s:%d\n", getWanIP(), port)
	fmt.Printf("| run path    : %s\n", dir)
	fmt.Println("----------------------------------------------")
}

func startServer() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getIntranetIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func getWanIP() string {
	// This function needs to be implemented based on how you obtain the WAN IP
	// You can use an external service or any other method to get the WAN IP
	return ""
}

func printUsage() {
	fmt.Println("httpFileServer用法:")
	fmt.Println("  -h                显示帮助")
	fmt.Println("  -p int            指定要侦听的端口（默认值 8081）")
	fmt.Println("  -d string         指定要服务的目录（默认值 .）")
}
