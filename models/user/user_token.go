package user

type TokenData struct {
    ID           int64  `json:"id"`
    Email        string `json:"email"`
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}
