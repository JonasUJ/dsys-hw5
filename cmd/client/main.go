package main

import (
	"bufio"
	"context"
	"fmt"
	"main/pkg/grpc/frontend"
	"main/pkg/netutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	client := frontend.NewFrontendClient(netutil.Dial("localhost", "50050"))

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("commands:\n'bid <amount>' to bid\n'result' to get the result\n'crashf' to crash a frontend\n'crashr' to crash a replica")
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		// Dispatch commands
		switch input[0] {
		case "bid":
			amt, err := strconv.Atoi(input[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			resp, err := client.Bid(context.Background(), &frontend.Amount{Amount: uint64(amt)})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp.Status)
			}
		case "result":
			resp, err := client.Result(context.Background(), &frontend.Void{})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp.Amount)
			}
		case "crashf":
			_, err := client.Crash(context.Background(), &frontend.Void{})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("crashed a frontend")
			}
		case "crashr":
			_, err := client.CrashReplica(context.Background(), &frontend.Void{})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("crashed a replica")
			}
		default:
			fmt.Printf("unknown command '%s'\n", input)
		}
	}

	if scanner.Err() != nil {
		panic("fail to read stdin")
	}
}
