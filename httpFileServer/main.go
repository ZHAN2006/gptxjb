package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var dir = "."
var port = 8081

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

func printServerInfo() {
	fmt.Printf("listen port is %d\n", port)
	fmt.Printf("| loc address : http://127.0.0.1:%d\n", port)
	fmt.Printf("| lan address : http://%s:%d\n", getIntranetIP(), port)
	fmt.Printf("| wan address : http://%s:%d\n", getWanIP(), port)
	fmt.Printf("| run path    : %s\n", dir)
	fmt.Println("----------------------------------------------")
}

func startServer() {
	logFile, err := os.Create("server.log")
	if err != nil {
		log.Fatal("Error creating log file: ", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)

	logger := log.New(logFile, "", log.LstdFlags)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getIntranetIP() string {
	adders, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range adders {
		if aspnet, ok := address.(*net.IPNet); ok && !aspnet.IP.IsLoopback() {
			if aspnet.IP.To4() != nil {
				return aspnet.IP.String()
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
