package uruptela

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sandronister/socket-go/pkg/devices"
	"gitlab.com/gobrax-dev/gobrax-tool/cloud/factory"
	"gitlab.com/gobrax-dev/gobrax-tool/cloud/types"
)

func (u *UseRuptela) getPayloadObject(device devices.IDevice) *types.ObjectInput {
	dataFolder := time.Now().Format("2006-01-02")
	name := uuid.New().String()

	return &types.ObjectInput{
		Bucket: os.Getenv("BUCKET_NAME"),
		Key:    fmt.Sprintf("ruptela/%s/%s/%s.bin", device.GetImei(), dataFolder, name),
		Body:   bytes.NewReader(device.GetBuffer()),
	}
}

func (u *UseRuptela) SendToBucket(device devices.IDevice) {
	cloud, _ := factory.GetCloudInstance()

	if cloud == nil {
		log.Println("Cloud is not initialized")
		panic("Cloud is not initialized")
	}

	if cloud != nil {
		bucket := cloud.GetBucket()

		payload := u.getPayloadObject(device)
		_, err := bucket.PutObject(payload)

		if err != nil {
			log.Println(err)
		}
	}
}
