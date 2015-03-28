package main

import (
	"flag"
	"log"
	"zumata/xrates"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	server := flag.String("server", "", "Server which host the service")
	flag.Parse()
	if *server == "" {
		log.Fatal("Server cannot be empty.")
	}

	var opts []grpc.DialOption
	conn, err := grpc.Dial(*server, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()

	client := xrates.NewXRatesClient(conn)

	log.Println("Getting one currency")
	for _, c := range []string{"MYR", "UNK"} {
		rate, err := client.Get(context.Background(), &xrates.Currency{c})
		if err != nil {
			log.Printf("Couldn't get currency. %v\n", err)
		} else {
			log.Printf("Rate for MYR: %f\n", rate.Rate)
		}
	}

	log.Println("Getting multiple currencies")
	rates, err := client.All(context.Background(), &xrates.Currencies{[]string{}})
	if err != nil {
		log.Printf("Couldn't get currency. %v\n", err)
	} else {
		for cur, rate := range rates.Rates {
			log.Printf("Rate for %s: %f\n", cur, rate)
		}
	}
}
