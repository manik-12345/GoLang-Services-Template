package main

import (
	"fmt"
	"github.com/manik-12345/GoLang-Services-Template/pkg/API/HelloWorld"
	"github.com/manik-12345/GoLang-Services-Template/pkg/DataProvider"
	"github.com/manik-12345/GoLang-Services-Template/pkg/Listener/HelloListener"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)


func Get() string {
	return "Member Service"
}

func main() {
	setZeroLog()

	// instantiate the main router
	router := chi.NewRouter()

	// standard package values
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger) // modify to use stackdriver
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Instantiate the DataProvider and defer the close
	helloServiceDataProvider := DataProvider.HelloServiceDatabase{}
	helloServiceDataProvider.Configure()
	defer func() { _ = helloServiceDataProvider.HelloServiceDatabase.Close() }()


	//Setting the url path and directory of the sso routes
	helloWorldRoute := HelloWorld.ServiceTemplateComponent{DataProvider: &helloServiceDataProvider}
	router.Route("/hello/world", helloWorldRoute.Router)

	// activate the router to listen to port
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Info().Msgf("%s is starting to listen at: http://%s%s.  Public IP: %s", os.Getenv("ServiceModule"), GetHostIP(), port, GetPublicIP())
	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Error().Msgf("%s could not start listener.  Error: %s", os.Getenv("ServiceModule"), err.Error())
		panic(err)
	}
}

func GetHostIP() string {
	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		log.Error().Msgf("GetHostIP: could not get IP address.  " + err.Error())
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip := networkIp.IP.String()
			//fmt.Println("Resolved Host IP: " + ip)
			return ip
		}
	}
	return ""
}

func GetPublicIP() string {
	url := "https://api.ipify.org?format=text"
	// we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	//fmt.Printf("Getting IP address from  ipify ...\n")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", ip)
}

func setZeroLog() {
	// Set time format of logs
	var logLevel int

	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := strconv.Atoi(logLevelStr)
	if err != nil {
		logLevel = int(zerolog.WarnLevel)
	}
	/*
		panic (zerolog.PanicLevel, 5)
		fatal (zerolog.FatalLevel, 4)
		error (zerolog.ErrorLevel, 3)
		warn (zerolog.WarnLevel, 2)
		info (zerolog.InfoLevel, 1)
		debug (zerolog.DebugLevel, 0)
		trace (zerolog.TraceLevel, -1)
	*/
	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	// Setting time zone and time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	loc, _ := time.LoadLocation("America/Toronto")
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(loc)
	}
}
