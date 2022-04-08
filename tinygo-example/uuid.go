package main

import (
	"github.com/google/uuid"
)

func main() {
}

//export myUuid
func myUuid() string {
	u := uuid.New()
	return u.String()
}
