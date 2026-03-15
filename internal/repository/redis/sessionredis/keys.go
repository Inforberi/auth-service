package sessionredis

import "encoding/hex"

func sessionKey(tokenHash []byte) string {
	return "sess:" + hex.EncodeToString(tokenHash)
}

// func userSessionVersionKey(userID string) string {
// 	return "user_sess_ver:" + userID
// }
