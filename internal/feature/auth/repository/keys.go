package repository

import "fmt"

// refreshTokenKey generates a Redis key for refresh tokens for a given userID.
type refreshTokenKey string

func newRefreshTokenKey(userID string) refreshTokenKey {
	return refreshTokenKey(fmt.Sprintf("refresh_token:%s", userID))
}

func (k refreshTokenKey) String() string {
	return string(k)
}
