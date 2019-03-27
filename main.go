/**
 * @author Jose Nidhin
 */
package main

import (
	"encoding/json"
	"fmt"

	"github.com/josnidhin/go-rest-api-example/config"
)

func main() {
	config := config.Load()
	configJson, _ := json.Marshal(config)
	fmt.Println(string(configJson))
}
