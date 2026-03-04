import (
	"encoding/base64"
	"net/http"
)

// Transform your API key for Basic Auth
apiKey := "<zk_fda8ebf7be4542e0a66ed77a7c5c0554>"
apiKeyTransformed := base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))

// Make the request
client := &http.Client{}
req, _ := http.NewRequest("GET", "https://api.zerion.io/v1/wallets/0x42b9df65b219b3dd36ff330a4dd8f327a6ada990/portfolio", nil)
req.Header.Set("Authorization", "Basic " + apiKeyTransformed)
req.Header.Set("accept", "application/json")

resp, _ := client.Do(req)
defer resp.Body.Close()