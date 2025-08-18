package services

import (
	"fmt"
	"os"
	"time"

	"github.com/oracle/nosql-go-sdk/nosqldb"
	"github.com/oracle/nosql-go-sdk/nosqldb/auth/iam"
	"github.com/oracle/nosql-go-sdk/nosqldb/types"
)

func PushDataToNoSql(activeRecordStatus *chan struct{}) {

	ticker := time.NewTicker(time.Second * 20)
	client, err := createClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	for {
		select {
		case <-*activeRecordStatus:
			fmt.Println("Finish recording!")
			return
		case <-ticker.C:
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
	var provider nosqldb.AuthorizationProvider
	var err error
	config := os.Getenv("OCI_LOCAL_BUILD")

	if config != "" {
		provider, err = iam.NewSignatureProviderFromFile(config, "", "", "")
		if err != nil {
			fmt.Println("failed to create local principal auth provider: %w", err)
			return nil, err
		}
	} else {
		provider, err = iam.NewSignatureProviderWithResourcePrincipal("")
		if err != nil {
			fmt.Println("failed to create local principal auth provider: %w", err)
			return nil, err
		}
	}

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
