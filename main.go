package goutils

import "os"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
