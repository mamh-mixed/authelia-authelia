package validator

var validKeys = []string{
	// Root Keys.
	"host",
	"port",
	"log_level",
	"log_format",
	"log_file_path",
	"default_redirection_url",
	"jwt_secret",
	"theme",
	"tls_key",
	"tls_cert",
	"certificates_directory",

	// Server Keys.
	"server.read_buffer_size",
	"server.write_buffer_size",
	"server.path",

	// TOTP Keys.
	"totp.issuer",
	"totp.period",
	"totp.skew",

	// Access Control Keys.
	"access_control.rules",
	"access_control.default_policy",
	"access_control.networks",

	// Session Keys.
	"session.name",
	"session.secret",
	"session.expiration",
	"session.inactivity",
	"session.remember_me_duration",
	"session.domain",

	// Redis Session Keys.
	"session.redis.host",
	"session.redis.port",
	"session.redis.password",
	"session.redis.database_index",

	// Local Storage Keys.
	"storage.local.path",

	// MySQL Storage Keys.
	"storage.mysql.host",
	"storage.mysql.port",
	"storage.mysql.database",
	"storage.mysql.username",
	"storage.mysql.password",

	// PostgreSQL Storage Keys.
	"storage.postgres.host",
	"storage.postgres.port",
	"storage.postgres.database",
	"storage.postgres.username",
	"storage.postgres.password",
	"storage.postgres.sslmode",

	// FileSystem Notifier Keys.
	"notifier.filesystem.filename",
	"notifier.disable_startup_check",

	// SMTP Notifier Keys.
	"notifier.smtp.username",
	"notifier.smtp.password",
	"notifier.smtp.host",
	"notifier.smtp.port",
	"notifier.smtp.identifier",
	"notifier.smtp.sender",
	"notifier.smtp.subject",
	"notifier.smtp.startup_check_address",
	"notifier.smtp.disable_require_tls",
	"notifier.smtp.trusted_cert", // TODO: Deprecated: Remove in 4.28.
	"notifier.smtp.disable_html_emails",
	"notifier.smtp.tls.minimum_version",
	"notifier.smtp.tls.skip_verify",
	"notifier.smtp.tls.server_name",
	"notifier.smtp.disable_verify_cert", // TODO: Deprecated: Remove in 4.28.

	// Regulation Keys.
	"regulation.max_retries",
	"regulation.find_time",
	"regulation.ban_time",

	// DUO API Keys.
	"duo_api.hostname",
	"duo_api.integration_key",
	"duo_api.secret_key",

	// Authentication Backend Keys.
	"authentication_backend.disable_reset_password",
	"authentication_backend.refresh_interval",

	// LDAP Authentication Backend Keys.
	"authentication_backend.ldap.implementation",
	"authentication_backend.ldap.url",
	"authentication_backend.ldap.base_dn",
	"authentication_backend.ldap.username_attribute",
	"authentication_backend.ldap.additional_users_dn",
	"authentication_backend.ldap.users_filter",
	"authentication_backend.ldap.additional_groups_dn",
	"authentication_backend.ldap.groups_filter",
	"authentication_backend.ldap.group_name_attribute",
	"authentication_backend.ldap.mail_attribute",
	"authentication_backend.ldap.display_name_attribute",
	"authentication_backend.ldap.user",
	"authentication_backend.ldap.password",
	"authentication_backend.ldap.start_tls",
	"authentication_backend.ldap.tls.minimum_version",
	"authentication_backend.ldap.tls.skip_verify",
	"authentication_backend.ldap.tls.server_name",
	"authentication_backend.ldap.skip_verify",         // TODO: Deprecated: Remove in 4.28.
	"authentication_backend.ldap.minimum_tls_version", // TODO: Deprecated: Remove in 4.28.

	// File Authentication Backend Keys.
	"authentication_backend.file.path",
	"authentication_backend.file.password.algorithm",
	"authentication_backend.file.password.iterations",
	"authentication_backend.file.password.key_length",
	"authentication_backend.file.password.salt_length",
	"authentication_backend.file.password.memory",
	"authentication_backend.file.password.parallelism",

	// OIDC Keys
	"openid_connect.clients.redirect_uris",
	"openid_connect.clients.client_id",
	"openid_connect.clients.client_secret",
	"openid_connect.clients.policy",
	"openid_connect.oauth2_hmac_secret",
	"openid_connect.oidc_issuer_private_key_path",

	// Secret Keys.
	"authelia.jwt_secret",
	"authelia.duo_api.secret_key",
	"authelia.session.secret",
	"authelia.authentication_backend.ldap.password",
	"authelia.notifier.smtp.password",
	"authelia.session.redis.password",
	"authelia.storage.mysql.password",
	"authelia.storage.postgres.password",
	"authelia.jwt_secret.file",
	"authelia.duo_api.secret_key.file",
	"authelia.session.secret.file",
	"authelia.authentication_backend.ldap.password.file",
	"authelia.notifier.smtp.password.file",
	"authelia.session.redis.password.file",
	"authelia.storage.mysql.password.file",
	"authelia.storage.postgres.password.file",
	"authelia.openid_connect.oauth2_hmac_secret.file",
}

var specificErrorKeys = map[string]string{
	"logs_file_path":   "config key replaced: logs_file is now log_file",
	"logs_level":       "config key replaced: logs_level is now log_level",
	"google_analytics": "config key removed: google_analytics - this functionality has been deprecated",
	"authentication_backend.file.password_options.algorithm":   "config key incorrect: authentication_backend.file.password_options should be authentication_backend.file.password",
	"authentication_backend.file.password_options.iterations":  "config key incorrect: authentication_backend.file.password_options should be authentication_backend.file.password",
	"authentication_backend.file.password_options.key_length":  "config key incorrect: authentication_backend.file.password_options should be authentication_backend.file.password",
	"authentication_backend.file.password_options.salt_length": "config key incorrect: authentication_backend.file.password_options should be authentication_backend.file.password",
	"authentication_backend.file.password_options.memory":      "config key incorrect: authentication_backend.file.password_options should be authentication_backend.file.password",
	"authentication_backend.file.password_options.parallelism": "config key incorrect: authentication_backend.file.password_options should be authentication_backend.file.password",
	"authentication_backend.file.password_hashing.algorithm":   "config key incorrect: authentication_backend.file.password_hashing should be authentication_backend.file.password",
	"authentication_backend.file.password_hashing.iterations":  "config key incorrect: authentication_backend.file.password_hashing should be authentication_backend.file.password",
	"authentication_backend.file.password_hashing.key_length":  "config key incorrect: authentication_backend.file.password_hashing should be authentication_backend.file.password",
	"authentication_backend.file.password_hashing.salt_length": "config key incorrect: authentication_backend.file.password_hashing should be authentication_backend.file.password",
	"authentication_backend.file.password_hashing.memory":      "config key incorrect: authentication_backend.file.password_hashing should be authentication_backend.file.password",
	"authentication_backend.file.password_hashing.parallelism": "config key incorrect: authentication_backend.file.password_hashing should be authentication_backend.file.password",
	"authentication_backend.file.hashing.algorithm":            "config key incorrect: authentication_backend.file.hashing should be authentication_backend.file.password",
	"authentication_backend.file.hashing.iterations":           "config key incorrect: authentication_backend.file.hashing should be authentication_backend.file.password",
	"authentication_backend.file.hashing.key_length":           "config key incorrect: authentication_backend.file.hashing should be authentication_backend.file.password",
	"authentication_backend.file.hashing.salt_length":          "config key incorrect: authentication_backend.file.hashing should be authentication_backend.file.password",
	"authentication_backend.file.hashing.memory":               "config key incorrect: authentication_backend.file.hashing should be authentication_backend.file.password",
	"authentication_backend.file.hashing.parallelism":          "config key incorrect: authentication_backend.file.hashing should be authentication_backend.file.password",
}

const denyPolicy = "deny"

const argon2id = "argon2id"
const sha512 = "sha512"

const schemeLDAP = "ldap"
const schemeLDAPS = "ldaps"

const testBadTimer = "-1"
const testInvalidPolicy = "invalid"
const testJWTSecret = "a_secret"
const testLDAPBaseDN = "base_dn"
const testLDAPPassword = "password"
const testLDAPURL = "ldap://ldap"
const testLDAPUser = "user"
const testModeDisabled = "disable"
const testTLSCert = "/tmp/cert.pem"
const testTLSKey = "/tmp/key.pem"
