package define

import "context"

// export const part
const (
	ProjectName = "db_proj"
)

// not-export const part
const (
	defaultPort       = 80
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
