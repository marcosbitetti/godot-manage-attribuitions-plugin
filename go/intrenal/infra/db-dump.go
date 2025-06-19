package infra

import (
	"context"

	"github.com/pkg/errors"
)

func createBaseTable(ctx context.Context, storage *Storage) error {
	var err error

	// types
	_, err = storage.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS types (
			_id 	INTEGER PRIMARY KEY NOT NULL,
			name	TEXT
		);
	`)
	if err != nil {
		return errors.Wrap(err, "error creating table types")
	}

	// licences
	_, err = storage.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS licences (
			_id 	INTEGER PRIMARY KEY NOT NULL,
			name	TEXT DEFAULT "Attribution 4.0 International (CC BY 4.0)",
			link	TEXT DEFAULT "https://creativecommons.org/licenses/by/4.0/"
		);
	`)
	if err != nil {
		return errors.Wrap(err, "error creating table licences")
	}

	// credits
	_, err = storage.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS credits (
			_id 		INTEGER PRIMARY KEY NOT NULL,
			name		TEXT,
			filename	TEXT,
			type_id		INTEGER NOT NULL DEFAULT 1,
			author 		TEXT,
			link 		TEXT,
			licence_id 	INTEGER NOT NULL DEFAULT 1,
			FOREIGN KEY (type_id)
				REFERENCES types (_id)
					ON DELETE CASCADE
					ON UPDATE NO ACTION,
			FOREIGN KEY (licence_id)
				REFERENCES licences (_id)
					ON DELETE CASCADE
					ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		return errors.Wrap(err, "error creating table credits")
	}
	return nil
}

func dumpFirstTypes(storage *Storage) {
	types := []string{
		"3D Model",
		"Music",
		"Plugin",
		"Project",
		"Sound Effect",
		"Texture",
		"Shader",
		"Photo",
		"Dubbing/Narration",
		"Font",
		"Code Snippet",
	}
	for _, value := range types {
		storage.AddType(value)
	}
}

func dumpFirstLicences(storage *Storage) {
	licences := [][]string{
		{"Attribution 4.0 International (CC BY 4.0)", "https://creativecommons.org/licenses/by/4.0/"},
		{"Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)", "https://creativecommons.org/licenses/by-sa/4.0/"},
		{"Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)", "https://creativecommons.org/licenses/by-nc/4.0/"},
		{"Attribution-NonCommercial-ShareAlike 4.0 International (CC BY-NC-SA 4.0)", "https://creativecommons.org/licenses/by-nc-sa/4.0/"},
		{"Attribution-NoDerivatives 4.0 International (CC BY-ND 4.0)", "https://creativecommons.org/licenses/by-nd/4.0/"},
		{"Attribution-NonCommercial-NoDerivatives 4.0 International (CC BY-NC-ND 4.0)", "https://creativecommons.org/licenses/by-nc-nd/4.0/"},
		{"CC0 1.0 Universal (CC0 1.0) - Public Domain Dedication", "https://creativecommons.org/publicdomain/zero/1.0/"},
		{"MIT", "https://opensource.org/license/mit/"},
		{"GNU General Public Licence", "https://www.gnu.org/licenses/gpl-3.0.html"},
		{"Attribution-NonCommercial-ShareAlike 3.0 Unported (CC BY-NC-SA 3.0)", "https://creativecommons.org/licenses/by-nc-sa/3.0/"},
		{"Attribution-NonCommercial-NoDerivs 3.0 Unported (CC BY-NC-ND 3.0)", "https://creativecommons.org/licenses/by-nc-nd/3.0/"},
		{"Attribution-ShareAlike 3.0 Unported (CC BY-SA 3.0)", "https://creativecommons.org/licenses/by-sa/3.0/"},
		{"Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)", "https://creativecommons.org/licenses/by-nd/3.0/"},
		{"Attribution 3.0 Unported (CC BY 3.0)", "https://creativecommons.org/licenses/by/3.0/"},
		{"GNU Lesser General Public License (LGPL)", "https://www.gnu.org/licenses/lgpl-3.0.html"},
		{"Apache License 2.0", "https://www.apache.org/licenses/LICENSE-2.0"},
		{"Mozilla Public License 2.0", "https://www.mozilla.org/en-US/MPL/2.0/"},
		{"Beerware", "https://fedoraproject.org/wiki/Licensing/Beerware"},
		{"Royalty Free", "https://en.wikipedia.org/wiki/Royalty-free"},
		{"Open Font License (OFL)", "https://openfontlicense.org/"},
		{"OGA-BY 3.0 (Open Game Art)", "https://static.opengameart.org/OGA-BY-3.0.txt"},
		{"Free Standard (Sketchfab)", "https://www.youtube.com/watch?v=M2bKt1oZsi4"},
	}

	for _, licence := range licences {
		storage.AddLicence(licence[0], licence[1])
	}
}
