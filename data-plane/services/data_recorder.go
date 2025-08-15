package services

import (
	"fmt"
	"time"

	"github.com/oracle/nosql-go-sdk/nosqldb"
	"github.com/oracle/nosql-go-sdk/nosqldb/auth/iam"
	"github.com/oracle/nosql-go-sdk/nosqldb/types"
)

func PushDataToNoSql(activeRecordStatus *chan struct{}) {

	ticker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-*activeRecordStatus:
			fmt.Println("Finish recording!")
			return
		case <-ticker.C:
			client, err := createClient()
			if err != nil {
				fmt.Println("Error creating client:", err)
				return
			}
			putRequest := createPutRequest()

			_, err = client.Put(putRequest)
			if err != nil {
				fmt.Println("Error inserting row:", err)
				return
			}
		}
	}
}

func createPutRequest() *nosqldb.PutRequest {
	sensors := GetSensorStore()
	row := map[string]interface{}{
		"tmstamp":   time.Now().Format("2006-01-02T15:04:05.000"),
		"speedKPH":  float32(sensors["gps-speed-1"].SpeedKPH),
		"latitude":  sensors["gps-position-1"].Latitude,
		"longitude": sensors["gps-position-1"].Longitude,
	}

	putReq := nosqldb.PutRequest{
		TableName: "trackday_test0",
		Value:     types.NewMapValue(row),
	}

	return &putReq
}

func createClient() (*nosqldb.Client, error) {
	provider, err := iam.NewSignatureProviderFromFile("/home/astro/.oci/config", "", "", "")

	if err != nil {
		fmt.Println("Could not create config")
		return nil, err
	}

	cfg := nosqldb.Config{
		Region:                "af-johannesburg-1",
		AuthorizationProvider: provider,
	}

	client, err := nosqldb.NewClient(cfg)
	if err != nil {
		fmt.Println("Could not create client")
		return nil, err
	}
	return client, nil
}
