// Package migrations embeds SQL migration files so the API can migrate on boot.
package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
