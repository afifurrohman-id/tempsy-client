package models

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
	MimeType          string `json:"mimeType"`
	AutoDeleteAt      int64  `json:"autoDeleteAt"`
	PrivateUrlExpires uint   `json:"privateUrlExpires"`
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

type Error struct {
	Kind        string `json:"kind"`
	Description string `json:"description"`
}

type ApiError struct {
	*Error `json:"apiError"`
}
