// Code generated by go generate. DO NOT EDIT.
//
// Run the following command to generate this file:
// 		go run ./cmd/authelia-gen code csp
//

package server

const (
	placeholderCSPNonce = "${NONCE}"
	tmplCSPDefault      = "default-src 'self'; frame-src 'none'; object-src 'none'; style-src 'self' 'nonce-%s'; frame-ancestors 'none'; base-uri 'self'; require-trusted-types-for 'script'"
	tmplCSPDevelopment  = "default-src 'self' 'unsafe-eval'; frame-src 'none'; object-src 'none'; style-src 'self' 'nonce-%s'; frame-ancestors 'none'; base-uri 'self'; require-trusted-types-for 'script'"
)
