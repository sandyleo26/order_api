package order

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

type DistanceService interface {
	Calculate(start, end Location) (int, error)
}

type GoogleDistanceService struct{}

func (*GoogleDistanceService) Calculate(start, end Location) (int, error) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_DISTANCE_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	req := &maps.DistanceMatrixRequest{
		Origins:      []string{strings.Join(start, ",")},
		Destinations: []string{strings.Join(end, ",")},
	}
	resp, err := c.DistanceMatrix(context.Background(), req)
	if err != nil {
		log.Fatalln("DistanceMatrix fatal error:", err.Error())
	}

	pretty.Println(resp)
	return resp.Rows[0].Elements[0].Distance.Meters, nil
}
