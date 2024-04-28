package neo4j

import (
	"context"
	"fmt"
	"time"

	"github.com/9ssi7/music-recommender/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Connect(cnf config.Neo4j) neo4j.DriverWithContext {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		cnf.Uri,
		neo4j.BasicAuth(cnf.Username, cnf.Password, ""))
	if err != nil {
		panic(err)
	}
	verifyOrRetryConnectivity(driver, ctx)

	fmt.Println("connected to neo4j")
	return driver
}

func verifyOrRetryConnectivity(driver neo4j.DriverWithContext, ctx context.Context) {
	retryCount := 0
	for {
		err := driver.VerifyConnectivity(ctx)
		if err != nil {
			if retryCount > 3 {
				panic(err)
			}
			time.Sleep(5 * time.Second)
			fmt.Println("retrying to connect to neo4j")
			retryCount++
		} else {
			break
		}
	}
}
