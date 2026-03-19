package sessionredis

import "encoding/hex"

func sessionKey(tokenHash []byte) string {
	return "sess:" + hex.EncodeToString(tokenHash)
}

func userSessionVersionKey(userID string) string {
	return "user_sess_ver:" + userID
}

func revokedSessionKey(tokenHash []byte) string {
	return "sess_revoked:" + hex.EncodeToString(tokenHash)
}
