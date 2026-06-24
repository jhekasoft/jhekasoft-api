//go:build health || all
// +build health all

package modules

import "jhekasoft-api/modules/health"

func init() {
	m := health.NewModule()
	EnabledModules = append(EnabledModules, m)
}
