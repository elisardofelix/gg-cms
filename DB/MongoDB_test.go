package DB

import (
	"context"
	"testing"
)

func TestConnectionToMongoDB(t *testing.T) {
	client, err := GetMongoDBClient()
	if err != nil {
		t.Fail()
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fail()
	}
}