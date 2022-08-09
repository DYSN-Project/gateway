package response

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewTokens(accessToken, refreshToken string) *Tokens {
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
