package pg_core_validators

import "strings"

func ValidateCreatorUserHandle(handle string) string {
	newHandle := strings.ReplaceAll(handle, " ", "-")

	return newHandle
}
