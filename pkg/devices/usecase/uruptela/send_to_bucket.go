package uruptela

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sandronister/socket-go/pkg/devices"
	"gitlab.com/gobrax-dev/gobrax-tool/cloud/factory"
	"gitlab.com/gobrax-dev/gobrax-tool/cloud/types"
)

func (u *UseRuptela) SendToBucket(device devices.IDevice) {
	cloud, _ := factory.GetCloudInstance()

	if cloud == nil {
		log.Println("Cloud is not initialized")
		panic("Cloud is not initialized")
	}

	if cloud != nil {
		bucket := cloud.GetBucket()

		dataFolder := time.Now().Format("2006-01-02")
		timestamp := time.Now().Format("2006-01-02")

		res, err := bucket.PutObject(&types.ObjectInput{
			Bucket: os.Getenv("BUCKET_NAME"),
			Key:    fmt.Sprintf("ruptela/%s/%s/%s.bin", device.GetImei(), dataFolder, timestamp),
			Body:   bytes.NewReader(device.GetBuffer()),
		})

		if err != nil {
			log.Println(err)
		}

		fmt.Println("****************************************")
		fmt.Println(os.Getenv("BUCKET_NAME"))
		fmt.Println(res, "--------------------------------------")

	}

}
