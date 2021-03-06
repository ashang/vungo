package openrtbutil_test

import (
	"fmt"
	"net"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbutil"
	"golang.org/x/net/context"
)

func ExampleClient() {
	ctx := context.Background()

	br := &openrtb.BidRequest{
		Id: "1234",
		Impressions: []*openrtb.Impression{&openrtb.Impression{
			Id:            "imp-1234",
			Video:         &openrtb.Video{},
			BidFloorPrice: 4,
		}},
	}

	endpoint := "http://127.0.0.1:8080/requestBid"

	req, err := openrtbutil.NewRequest(ctx, br, endpoint, nil)
	if err != nil {
		panic(err)
	}

	// Setting some custom HTTP header.
	req.Http().Header.Set("X-Auction-Platform", "Vungle Exchange")

	resp, err := openrtbutil.DefaultClient.Do(req)

	if err != nil {
		switch e := err.(type) {
		case net.Error:
			if e.Timeout() {
				fmt.Println("Resposne timeout: ", e)
			} else {
				fmt.Println("Other unexpected network error: ", e)
			}
		case openrtbutil.NoBidError:
			fmt.Println("No bid: ", e)
			if e.Err() != nil {
				fmt.Println("Underlying error: ", e.Err())
			} else if e.Response() != nil {
				fmt.Println("From response body: ", e.Response())
			} else {
				fmt.Println("Summarized no bid reason: ", e.Reason())
			}
		default:
			fmt.Println("Unexpected error: ", e)
		}
	} else {
		fmt.Println("Got response: ", resp)
	}
}
