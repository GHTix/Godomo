package overkiz

type OverkizConfig struct {
	BaseUrl            string `mapstructure:"base_url"`
	UserName           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	OAuthLoginEndpoint string `mapstructure:"oauth_login_endpoint"`
	OAuthClientId      string `mapstructure:"oauth_client_id"`
	OAuthClientSecret  string `mapstructure:"oauth_client_secret"`
}
