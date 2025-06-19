package command

import (
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/usecases"
)

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
var commands = map[string]func(storage *infra.Storage, args []string) []byte{
	"help":               getHelp,
	"listAttribuitions":  usecases.GetAttribuitions,
	"listTypes":          usecases.GetTypes,
	"listLicences":       usecases.GetLicences,
	"addType":            usecases.AddType,
	"addLicence":         usecases.AddLicence,
	"updateType":         usecases.UpdateType,
	"deleteType":         usecases.DeleteType,
	"updateLicence":      usecases.UpdateLicence,
	"deleteLicence":      usecases.DeleteLicence,
	"addAttribuition":    usecases.AddAttribuition,
	"updateAttribuition": usecases.UpdateAttribuition,
	"deleteAttribuition": usecases.DeleteAttribuition,
}

func ParseCommand(storage *infra.Storage, args []string) string {
	if len(args) < 3 {
		return string(getHelp(nil, nil))
	}

	caller, has := commands[args[2]]
	if !has {
		caller = getHelp
	}

	return string(caller(storage, args))
}

func getHelp(_ *infra.Storage, _ []string) []byte {
	return []byte(staticHelp)
}

const staticHelp = `-> Attributions
attribuitions-amd64-linux ~/mygames/attributions.sqlite listAttribuitions
attribuitions-amd64-linux ~/mygames/attributions.sqlite listAttribuitions {"text":"<search>", "order": "ASC"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite addAttribuition {"name":"Test","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite updateAttribuition {"_id":1,"name":"_Test","filename":"_file","type":"_One","author":"_Ze","link":"_http://none","licence":"Beerware","type":"Plugin"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite deleteAttribuition {"_id":1}

-> Types
attribuitions-amd64-linux ~/mygames/attributions.sqlite listTypes
attribuitions-amd64-linux ~/mygames/attributions.sqlite addType {"name": "Font"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite updateType {"_id":1, "name": "FontNew"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite deleteType {"_id":1}

-> Licenses
attribuitions-amd64-linux ~/mygames/attributions.sqlite listLicences
attribuitions-amd64-linux ~/mygames/attributions.sqlite addLicence {"name": "Insaneware", "link": "https://example.com/license"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite updateLicence {"_id":1, "name": "Insaneware2", "link": "https://example.com/licenses"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite deleteLicence {"_id":1}
`
