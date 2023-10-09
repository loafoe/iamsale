package service

type AuthConfig struct {
	Username string
	Password string
}

type IAMConfig struct {
	Region            string
	Environment       string
	OrgID             string
	ServiceID         string
	ServicePrivateKey string
}
