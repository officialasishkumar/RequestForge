package models

type Auth struct {
    Type         string   `json:"type"` // "basic" or "oauth2"
    Username     string   `json:"username,omitempty"`
    Password     string   `json:"password,omitempty"`
    TokenURL     string   `json:"token_url,omitempty"`
    ClientID     string   `json:"client_id,omitempty"`
    ClientSecret string   `json:"client_secret,omitempty"`
    Scopes       []string `json:"scopes,omitempty"`
}

type Request struct {
    Method  string            `json:"method"`
    URL     string            `json:"url"`
    Headers map[string]string `json:"headers"`
    Body    interface{}       `json:"body"`
    Auth    *Auth             `json:"auth,omitempty"`
}
