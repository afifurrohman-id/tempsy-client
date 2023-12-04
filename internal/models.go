package internal

type OAuth2Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type User struct {
	UserName   string `json:"username"`
	TotalFiles int    `json:"totalFiles"`
	Picture    string `json:"picture"`
}

type GoogleAccountInfo struct {
	*User
	Email         string `json:"email"`
	ID            string `json:"id"`
	VerifiedEmail bool   `json:"verified_email"`
}

type DataFile struct {
	Name              string `json:"name"`
	AutoDeletedAt     int64  `json:"autoDeletedAt"`     // milliseconds
	PrivateUrlExpires int    `json:"privateUrlExpires"` // seconds
	IsPublic          bool   `json:"isPublic"`
	UploadedAt        int64  `json:"uploadedAt"` // milliseconds
	UpdatedAt         int64  `json:"updatedAt"`  // milliseconds
	Url               string `json:"url"`
	Size              int64  `json:"size"` // byte count
	ContentType       string `json:"type"`
}

type GuestToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"` // in seconds
}

type GOAuth2Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"` // in seconds
	TokenType    string `json:"token_type"`
	Scopes       string `json:"scope"` // separated by space
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type GOAuth2Config struct {
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	CallbackURL  string   `json:"callbackUrl"`
	Scopes       []string `json:"scopes"`
}

type ApiError struct {
	Type        string `json:"errorType"`
	Description string `json:"errorDescription"`
}
