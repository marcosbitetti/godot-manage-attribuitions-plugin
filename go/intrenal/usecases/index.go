package usecases

import "github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"

// commands is a map that associates command names (as strings) to their corresponding handler functions.
// This approach offers several advantages over using a traditional switch-case statement:
//   - Internal Mechanism: Map String Search
//     In Go, map lookups for string keys use a hash-based mechanism. When a command name is
//     provided, the runtime computes a hash of the string and uses it to efficiently locate
//     the corresponding value in the map. This provides average-case constant time complexity
//     for lookups, making command dispatching both fast and scalable even as the number
//     of commands grows. The hash function and map bucket management are handled internally
//     by the Go runtime, ensuring optimal performance for string-keyed maps.
//   - Extensibility: New commands can be added or removed easily by updating the map, without modifying control flow logic.
//   - Maintainability: The mapping between command names and their handlers is centralized, making the code easier to read and maintain.
//   - Flexibility: The map can be iterated, inspected, or even modified at runtime if needed, enabling dynamic command registration.
//   - Decoupling: Command logic is decoupled from the command parsing logic, promoting separation of concerns and cleaner code organization.
var commands = map[string]func(storage *infra.Storage, args []string) []byte {
	"help":               GetHelp,
	"listAttribuitions":  GetAttribuitions,
	"listTypes":          GetTypes,
	"listLicences":       GetLicences,
	"addType":            AddType,
	"addLicence":         AddLicence,
	"updateType":         UpdateType,
	"deleteType":         DeleteType,
	"updateLicence":      UpdateLicence,
	"deleteLicence":      DeleteLicence,
	"addAttribuition":    AddAttribuition,
	"updateAttribuition": UpdateAttribuition,
	"deleteAttribuition": DeleteAttribuition,
}

func Commands() map[string]func(storage *infra.Storage, args []string) []byte {
	return commands
}
