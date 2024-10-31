package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type sessionCache struct {
    rdb *redis.Client
}

func NewSessionCache(rdb *redis.Client) *sessionCache {
    return &sessionCache{
		rdb: rdb,
	}
}

// Stores the refresh token associated with a user
func (s *sessionCache) StoreRefreshToken(ctx context.Context, userId uint64, token string) error {
    // Store user ID to token mapping
    if err := s.rdb.Set(ctx, s.refreshTokenKey(userId), token, 24*time.Hour).Err(); err != nil {
        return err
    }
    // Store token to user ID mapping
    return s.rdb.Set(ctx, s.userIdKey(token), userId, 24*time.Hour).Err()
}

// Retrieves the refresh token for a user
func (s *sessionCache) GetRefreshToken(ctx context.Context, userId uint64) (string, error) {
    return s.rdb.Get(ctx, s.refreshTokenKey(userId)).Result()
}

// Deletes the refresh token for a user
func (s *sessionCache) DeleteRefreshToken(ctx context.Context, userId uint64) error {
    token, err := s.GetRefreshToken(ctx, userId)
    if err == nil {
        // Delete token-to-user ID mapping as well
        _ = s.rdb.Del(ctx, s.userIdKey(token)).Err()
    }
    return s.rdb.Del(ctx, s.refreshTokenKey(userId)).Err()
}

// Retrieves the user ID from a given token
func (s *sessionCache) GetUserIdFromRefreshToken(ctx context.Context, token string) (uint64, error) {
    userIdStr, err := s.rdb.Get(ctx, s.userIdKey(token)).Result()
    if err != nil {
        return 0, err
    }
    userId, err := strconv.ParseUint(userIdStr, 10, 64)
    if err != nil {
        return 0, err
    }
    return userId, nil
}

// Key for storing refresh token by user ID
func (s *sessionCache) refreshTokenKey(userId uint64) string {
    return "refresh_token_with_user_id:" + strconv.FormatUint(userId, 10)
}

// Key for storing user ID by token
func (s *sessionCache) userIdKey(token string) string {
    return "user_id_from_token:" + token
}







// func (s *sessionCache) StoreRefreshToken(ctx context.Context, userId uint64, token string) error {
//     return s.rdb.Set(ctx, s.refreshTokenKey(userId), token, 24*time.Hour).Err()
// }

// func (s *sessionCache) GetRefreshToken(ctx context.Context, userId uint64) (string, error) {
//     return s.rdb.Get(ctx, s.refreshTokenKey(userId)).Result()
// }



// func (s *sessionCache) DeleteRefreshToken(ctx context.Context, userId uint64) error {
//     return s.rdb.Del(ctx, s.refreshTokenKey(userId)).Err()
// }

// func (s *sessionCache) refreshTokenKey(userId uint64) string {
//     return "refresh_token:" + strconv.FormatUint(userId, 10)
// }
