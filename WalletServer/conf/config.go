// Package conf parse config to configuration
package conf

// Config contains configuration items for sdk
type Config struct {
	IsHTTP         bool
	ChainID        int64
	CAFile         string
	TLSCAContext   []byte
	Key            string
	TLSKeyContext  []byte
	Cert           string
	TLSCertContext []byte
	IsSMCrypto     bool
	PrivateKey     []byte
	GroupID        int
	NodeURL        string
}
