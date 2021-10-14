package main

//based on https://github.com/komuw/goweb/blob/fa3829371ee9f8848a30a8e989cf592bb0fccbf7/main.go#L161
//learnings: always make sure you work with pointers to the server instead of copy by valuein order to reduce memory uage on every method addition.
import (
	"context"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type scoreServer struct {
	version string
	router 	*http.ServeMux //note that mux is more performant, but lets start with this. https://github.com/gorilla/mux
	logger *log.Logger //note; there are idfferent loggers that might make more sense, but let's start basic . https://github.com/uber-go/zap
}

func main() {
	err := run()
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%+v\n", err)
		if err != nil {
			log.Fatal("Error setting error")
		}
		os.Exit(1)
	}

}

// Make `myAPI` implement the http.Handler interface(https://golang.org/pkg/net/http/#Handler)
// use myAPI wherever you could use http.Handler(eg ListenAndServe)
func (s *scoreServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("Got request"+r.RequestURI)
	s.router.ServeHTTP(w, r)
}



func (s *scoreServer) addServerHeader(wrappedHandler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("server", "Powered by GoScore!")
		wrappedHandler(w,r)
	}
}


func run() error {
	localScoreServer := &scoreServer{
		version: "0.1",
		router: http.NewServeMux(),
		logger: log.New(os.Stdout, "logger- "+time.Now().String()+": ", log.Lshortfile),
	}
	ctx, cancel := context.WithCancel(context.Background())
	log.Println("Server started: running on port 3000")
	localScoreServer.Routes()
	serverPort := ":3000"
	address := fmt.Sprintf("127.0.0.1%s", serverPort)
	server := &http.Server{
		Addr: serverPort,

		// 1. https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production
		// 2. https://blog.cloudflare.com/exposing-go-on-the-internet/
		// 3. https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
		// 4. https://github.com/golang/go/issues/27375
		Handler:           http.TimeoutHandler(localScoreServer, 10*time.Second, "Custom Server timeout"),
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       120 * time.Second,

		BaseContext: func(net.Listener) context.Context { return ctx },
	}
	sigHandler(server, ctx, cancel)
	localScoreServer.logger.Printf("server listening at %s", address)
	return server.ListenAndServe()
}


func sigHandler(srv *http.Server, ctx context.Context, cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, unix.SIGTERM, unix.SIGINT, unix.SIGQUIT, unix.SIGHUP)
	go func() {
		<-sigs
		cancel()
		_ = srv.Shutdown(ctx)
	}()
}

