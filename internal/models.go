package internal

type OAuth2Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type User struct {
	UserName   string `json:"username"`
	Picture    string `json:"picture"`
	TotalFiles int    `json:"totalFiles"`
}

type GoogleAccountInfo struct {
	*User
	Email         string `json:"email"`
	ID            string `json:"id"`
	VerifiedEmail bool   `json:"verified_email"`
}

type DataFile struct {
	Name              string `json:"name"`
	Url               string `json:"url"`
	ContentType       string `json:"type"`
	AutoDeletedAt     int64  `json:"autoDeletedAt"`
	PrivateUrlExpires int    `json:"privateUrlExpires"`
	UploadedAt        int64  `json:"uploadedAt"`
	UpdatedAt         int64  `json:"updatedAt"`
	Size              int64  `json:"size"`
	IsPublic          bool   `json:"isPublic"`
}

type GuestToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"` // in seconds
}

type GOAuth2Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scopes       string `json:"scope"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
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
