// Copyright © 2019 The Things Industries B.V.

// Package tenant contains context handling and middleware for tenancy.
package tenant

import "time"

// Config represents tenancy configuration.
type Config struct {
	DefaultID string        `name:"default-id" description:"Default tenant ID"`
	CacheTTL  time.Duration `name:"ttl" description:"TTL of cached tenant configurations"`
}
