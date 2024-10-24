package _89

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	client, err := Login("15315496321", "mogen123")
	fmt.Println(client, err)
}
