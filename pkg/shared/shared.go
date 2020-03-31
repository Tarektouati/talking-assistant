package shared

import (
	"fmt"
	"os"
)

//GetEnvWithError func returns error if the env is not found
func GetEnvWithError(env string) (string, error) {
	envValue, found := os.LookupEnv(env)
	if !found {
		return "", fmt.Errorf("%s env not found", env)
	}
	return envValue, nil
}
