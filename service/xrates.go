package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
	"zumata/xrates"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	BaseURL = "http://openexchangerates.org/api/latest.json"
)

type Server struct {
	url       string
	timestamp time.Time
	rates     map[string]float64
}

func NewServer(appID string) *Server {
	s := &Server{url: BaseURL + "?app_id=" + appID}
	if err := s.populate(); err != nil {
		log.Fatal("Error in creating server. %v", err)
	}
	return s
}

func (s *Server) Get(ctx context.Context, cur *xrates.Currency) (*xrates.Rate, error) {
	if err := s.update(); err != nil {
		return nil, err
	}

	r, ok := s.rates[cur.Currency]
	if !ok {
		return nil, errors.New("Currency not recognized.")
	}

	return &xrates.Rate{r}, nil
}

func (s *Server) All(ctx context.Context, cur *xrates.Currencies) (*xrates.Rates, error) {
	if err := s.update(); err != nil {
		return nil, err
	}

	return &xrates.Rates{s.rates}, nil
}

func (s *Server) populate() error {
	log.Println("Querying OpenExchange")
	res, err := http.Get(s.url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var r result
	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}

	s.timestamp = time.Unix(r.Timestamp, 0)
	s.rates = r.Rates

	log.Println("Result cached.")
	return nil
}

func (s *Server) update() error {
	now := time.Now()
	dur := now.Sub(s.timestamp)

	if dur.Hours() < 1 {
		log.Println("No update needed.")
		return nil
	}

	log.Println("Upate needed.")
	return s.populate()
}

type result struct {
	Timestamp int64
	Rates     map[string]float64
}

func main() {
	id := flag.String("app_id", "", "App ID given by OpenExchange")
	port := flag.Int("port", 50800, "Port for this server to serve")
	flag.Parse()
	if *id == "" {
		log.Fatal("App ID cannot be empty.")
	}

	s := NewServer(*id)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on %d: %v", *port, err)
	}

	gs := grpc.NewServer()
	xrates.RegisterXRatesServer(gs, s)
	gs.Serve(lis)
}
