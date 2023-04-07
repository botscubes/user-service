package tokenstorage

type TokenStorage interface {
	SaveToken(token string, lifeTime int) error
	DeleteToken(token string) error
	CheckToken(token string) bool
}
