package define

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"time"
)

const (
	ProjectName = "db_proj"
)

var (
	ProjectPath = getCurrentFilepath()
)

const (
	defaultPort       = 10000
	defaultUseSwagger = false
	defaultDebugMode  = false
)

var (
	Port       = defaultPort
	UseSwagger = defaultUseSwagger
	DebugMode  = defaultDebugMode
)

// Perm const
const (
	NormalPerm = 1 << 0
	AdminPerm  = 1 << 1
)

// Redis ctx
var DefaultRedisContext = context.Background()

// JWT
var (
	JWTPrivateToken, _ = rsa.GenerateKey(rand.Reader, 2048)
)

const (
	Issuer     = "weasleychen"
	ExpireTime = 2 * time.Hour
)
