package httpclient

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "RequestForge/auth"
    "RequestForge/logger"
    "RequestForge/models"
    "net/http"
    "time"
)

const MaxRetries = 3

type Client struct {
    HTTPClient *http.Client
    Verbose    bool
}

func NewClient(timeout time.Duration, verbose bool) *Client {
    return &Client{
        HTTPClient: &http.Client{
            Timeout: timeout,
        },
        Verbose: verbose,
    }
}

func (c *Client) SendRequest(reqData models.Request) error {
    var err error
    var reqBody []byte
    if reqData.Body != nil {
        reqBody, err = json.Marshal(reqData.Body)
        if err != nil {
            fmt.Println("Error marshalling request body:", err)
            logger.Logger.Println("Error marshalling request body:", err)
            return err
        }
    }

    // Handle Authentication
    if reqData.Auth != nil {
        switch reqData.Auth.Type {
        case "basic":
            reqData.Headers["Authorization"] = auth.BasicAuthHeader(
                reqData.Auth.Username, reqData.Auth.Password)
        case "oauth2":
            token, err := auth.GetOAuthToken(reqData.Auth)
            if err != nil {
                fmt.Println("Error getting OAuth token:", err)
                logger.Logger.Println("Error getting OAuth token:", err)
                return err
            }
            reqData.Headers["Authorization"] = "Bearer " + token
        default:
            fmt.Println("Unsupported auth type:", reqData.Auth.Type)
            logger.Logger.Println("Unsupported auth type:", reqData.Auth.Type)
            return fmt.Errorf("unsupported auth type: %s", reqData.Auth.Type)
        }
    }

    // Implement retries
    var resp *http.Response
    for attempt := 1; attempt <= MaxRetries; attempt++ {
        req, err := http.NewRequest(reqData.Method, reqData.URL, bytes.NewBuffer(reqBody))
        if err != nil {
            fmt.Println("Error creating HTTP request:", err)
            logger.Logger.Println("Error creating HTTP request:", err)
            return err
        }

        // Set headers
        for key, value := range reqData.Headers {
            req.Header.Set(key, value)
        }

        resp, err = c.HTTPClient.Do(req)
        if err != nil {
            fmt.Printf("Attempt %d: Error sending HTTP request: %v\n", attempt, err)
            logger.Logger.Printf("Attempt %d: Error sending HTTP request: %v\n", attempt, err)
            time.Sleep(time.Duration(attempt) * time.Second) // Exponential back-off
            continue
        }
        break
    }

    if err != nil {
        fmt.Println("Failed to send HTTP request after retries:", err)
        logger.Logger.Println("Failed to send HTTP request after retries:", err)
        return err
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        logger.Logger.Println("Error reading response body:", err)
        return err
    }

    fmt.Printf("Response for %s %s (Status: %s):\n", reqData.Method, reqData.URL, resp.Status)
    if c.Verbose {
        fmt.Println(string(respBody))
    }
    return nil
}
