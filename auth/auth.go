package auth

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "RequestForge/models"
    "net/http"
    "net/url"
    "strings"
    "time"
)

func BasicAuthHeader(username, password string) string {
    auth := username + ":" + password
    return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetOAuthToken(auth *models.Auth) (string, error) {
    data := url.Values{}
    data.Set("grant_type", "client_credentials")
    data.Set("client_id", auth.ClientID)
    data.Set("client_secret", auth.ClientSecret)
    if len(auth.Scopes) > 0 {
        data.Set("scope", strings.Join(auth.Scopes, " "))
    }

    req, err := http.NewRequest("POST", auth.TokenURL, strings.NewReader(data.Encode()))
    if err != nil {
        return "", err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        bodyBytes, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("failed to get token: %s", string(bodyBytes))
    }

    var tokenResp struct {
        AccessToken string `json:"access_token"`
        TokenType   string `json:"token_type"`
        ExpiresIn   int    `json:"expires_in"`
    }

    err = json.NewDecoder(resp.Body).Decode(&tokenResp)
    if err != nil {
        return "", err
    }

    return tokenResp.AccessToken, nil
}
