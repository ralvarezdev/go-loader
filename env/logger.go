package env

import (
	"log/slog"
)

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
//
// Parameters:
//
//   - logger: The slog.Logger instance to use for logging.
//   - variablesName: A variadic list of environment variable names that were loaded.
func EnvironmentVariableLoaded(logger *slog.Logger, variablesName ...string) {
	if logger != nil {
		logger.Debug(
			"environment variable loaded",
			slog.Any("variables", variablesName),
		)
	}
}
