package usecases

import "github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"

func GetHelp(_ *infra.Storage, _ []string) []byte {
	return []byte(staticHelp)
}

const staticHelp = `Usage:

-> Attributions
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
