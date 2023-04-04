package mongo

import (
	"os"
)

func getMongoDSN() string {
	if os.Getenv("MONGO_DSN_TEST") == "" {
		// return "mongodb://localhost:27017"
		return "mongodb+srv://sol:H3Xw6542Mx7D08JW@mongo-sol-35718b1e.mongo.ondigitalocean.com/test?replicaSet=mongo-sol&tls=true&authSource=admin"
	}
	return os.Getenv("MONGO_DSN")
}
