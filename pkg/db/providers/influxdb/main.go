package influxdb

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
)

var InFluxClient influxdb2.Client

func OpenInfluxDbConnection() influxdb2.Client {
	//bucket := "example-bucket"
	//org := "example-org"
	token := "fOakMj9F8PwL9DngYdcShEf8spJyLK-SmtrdKm6P53XbcEqTR1FocuKr3kzOV-9riQXfi6z54No9b6o6nOaRMA=="
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"

	InFluxClient = influxdb2.NewClient(url, token)
	isConnected, err := InFluxClient.Ping(context.Background())
	if err != nil {
		fmt.Print("error in influx connection ", err)
	}
	fmt.Println("isConnected ", isConnected)

	return InFluxClient
}
