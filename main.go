package main

import (
	"fmt"

	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	loggo "github.com/moatorres/go/modules/logger"
	"github.com/moatorres/go/modules/utils"
)

// file server
func FileServerHandler() http.Handler {
	return http.FileServer(http.Dir("./static"))
}

// logger instance
var logger = loggo.New(loggo.LoggerOptions{
	Service: "infrago",
})

// health handler
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
	logger.Log(r.Method + " " + r.Host + r.URL.Path)
}

// graceful process
func main() {
	port := utils.GetEnvVar("PORT", "3000")
	fileServer := FileServerHandler()
	serverTimeout := 10

	http.Handle("/", fileServer)
	http.HandleFunc("/healthz", healthz)

	signalsChannel := make(chan os.Signal, 1)
	watchedSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(signalsChannel, watchedSignals...)

	go func() {
		logger.Log("Server is running at %s", port)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sig := <-signalsChannel

	logger.Warn("Caught signal '%v'", sig)

	logger.Log("Server will shut down in %s seconds", fmt.Sprint(serverTimeout))

	time.Sleep(time.Second * time.Duration(serverTimeout))

	logger.Log("Bye 👋")

	os.Exit(0)
}
