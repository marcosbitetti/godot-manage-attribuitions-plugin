package main

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/domain"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/interfaces/command"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/usecases"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	t.Run("should handle no arguments", func(t *testing.T) {
		os.Args = []string{"app"}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "no command provided")
	})

	t.Run("should get help message", func(t *testing.T) {
		os.Args = []string{"app", "help"}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "-> Attributions")
	})

	t.Run("should handle no existing database, then create and populate it", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "listTypes"}
		_ = fakeMain()
		stat, err := os.Stat(databasePath)
		assert.NoError(t, err)
		assert.Equal(t, "nonexistent.db", stat.Name())
		assert.Greater(t, stat.Size(), int64(0))
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		var dataTypes _ResponseType
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataTypes))
		assert.Greater(t, len(dataTypes.Data), 3)
		assertHasType(t, "3D Model", dataTypes.Data)
		assertHasType(t, "Music", dataTypes.Data)

		os.Args = []string{"app", databasePath, "listLicences"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")

		var dataLicences _ResponseLicence
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataLicences))
		assert.Greater(t, len(dataLicences.Data), 3)
		assertHasLicence(t, "CC0 1.0 Universal (CC0 1.0) - Public Domain Dedication", dataLicences.Data)
		assertHasLicence(t, "MIT", dataLicences.Data)
	})

	t.Run("should add new Type", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		expectedType := "API"
		os.Args = []string{"app", databasePath, "addType", `{"name": "` + expectedType + `"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listTypes"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataTypes _ResponseType
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataTypes))
		assertHasType(t, expectedType, dataTypes.Data)
	})

	t.Run("should update Type", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		expectedType := "Changed"
		os.Args = []string{"app", databasePath, "updateType", `{"_id":1, "name": "` + expectedType + `"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listTypes"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataTypes _ResponseType
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataTypes))
		assertHasType(t, expectedType, dataTypes.Data)
	})

	t.Run("should delete Type", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "listTypes"}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataTypes _ResponseType
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataTypes))
		count := len(dataTypes.Data)

		os.Args = []string{"app", databasePath, "deleteType", `{"_id":1}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success", jsonRaw)

		os.Args = []string{"app", databasePath, "listTypes"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataTypes))
		finalCount := len(dataTypes.Data)
		assert.Equal(t, count-1, finalCount)
	})

	t.Run("should add new Licence", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		expectedLicence := "FajuteCodeLicence"
		expectedLicenceLink := "https://example.com"
		os.Args = []string{"app", databasePath, "addLicence", `{"name": "` + expectedLicence + `","link": "` + expectedLicenceLink + `"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listLicences"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataLicences _ResponseLicence
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataLicences))
		assertHasLicence(t, expectedLicence, dataLicences.Data)
	})

	t.Run("should update Licence", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		expectedLicence := "Changed"
		expectedLicenceLink := "https://example.com"
		os.Args = []string{"app", databasePath, "updateLicence", `{"_id":1, "name": "` + expectedLicence + `","link": "` + expectedLicenceLink + `"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listLicences"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataLicences _ResponseLicence
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataLicences))
		assertHasLicence(t, expectedLicence, dataLicences.Data)
	})

	t.Run("should delete Licence", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "listLicences"}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataLicences _ResponseType
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataLicences))
		count := len(dataLicences.Data)

		os.Args = []string{"app", databasePath, "deleteLicence", `{"_id":1}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success", jsonRaw)

		os.Args = []string{"app", databasePath, "listLicences"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataLicences))
		finalCount := len(dataLicences.Data)
		assert.Equal(t, count-1, finalCount)
	})

	t.Run("should list attribuitions after create db file", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "listAttribuitions"}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataAttribuitions _ResponseAttribuition
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		count := len(dataAttribuitions.Data)
		assert.Equal(t, count, 0)
	})

	t.Run("should add attribuition", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"Test","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listAttribuitions"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataAttribuitions _ResponseAttribuition
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		count := len(dataAttribuitions.Data)
		assert.Equal(t, count, 1)
	})

	t.Run("should retorne attribuitions in reverse order", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"A","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")
		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"B","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listAttribuitions", `{"order":"DESC"}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataAttribuitions _ResponseAttribuition
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		count := len(dataAttribuitions.Data)
		assert.Equal(t, count, 2)
		assert.Equal(t, dataAttribuitions.Data[0].Name, "B")
		assert.Equal(t, dataAttribuitions.Data[1].Name, "A")
	})

	t.Run("should retorne attribuitions filtered", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"Abcde","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")
		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"Fghij","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listAttribuitions", `{"text":"cde","order":"ASC"}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataAttribuitions _ResponseAttribuition
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		count := len(dataAttribuitions.Data)
		assert.Equal(t, count, 1)
		assert.Equal(t, dataAttribuitions.Data[0].Name, "Abcde")
	})

	t.Run("should update attribuition", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"Test","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "updateAttribuition",
			`{"_id":1,"name":"_Test","filename":"_file","type":"_One","author":"_Ze","link":"_http://none","licence":"Beerware","type":"Plugin"}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listAttribuitions"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataAttribuitions _ResponseAttribuition
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		assert.Equal(t, dataAttribuitions.Data[0].Name, "_Test")
		assert.Equal(t, dataAttribuitions.Data[0].Licence, "Beerware")
	})

	t.Run("should delete attribuition", func(t *testing.T) {
		tempDir := t.TempDir()
		databasePath := tempDir + "/nonexistent.db"

		os.Args = []string{"app", databasePath, "addAttribuition",
			`{"name":"Test","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}`}
		jsonRaw := fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listAttribuitions"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		var dataAttribuitions _ResponseAttribuition
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		assert.Equal(t, 1, len(dataAttribuitions.Data))

		os.Args = []string{"app", databasePath, "deleteAttribuition", `{"_id":1}`}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")

		os.Args = []string{"app", databasePath, "listAttribuitions"}
		jsonRaw = fakeMain()
		assert.Contains(t, jsonRaw, "success")
		assert.NoError(t, json.Unmarshal([]byte(jsonRaw), &dataAttribuitions))
		assert.Equal(t, 0, len(dataAttribuitions.Data))
	})
}

func fakeMain() string {
	argCount := len(os.Args)
	if argCount == 1 {
		return string(usecases.FormatJSON(nil, errors.New("no command provided")))
	}

	path, err := infra.ParseDatabasePath(os.Args)
	if err != nil {
		return string(usecases.FormatJSON(nil, err))
	}

	storage, err := infra.NewStorage(path)
	if err != nil {
		return string(usecases.FormatJSON(nil, err))
	}
	defer storage.CloseDatabase()

	return string(command.ParseCommand(storage, os.Args))
}

type _ResponseType struct {
	Status  string        `json:"status"`
	Message *string       `json:"message,omitempty"`
	Data    []domain.Type `json:"data"`
}

type _ResponseLicence struct {
	Status  string           `json:"status"`
	Message *string          `json:"message,omitempty"`
	Data    []domain.Licence `json:"data"`
}

type _ResponseAttribuition struct {
	Status  string                `json:"status"`
	Message *string               `json:"message,omitempty"`
	Data    []domain.Attribuition `json:"data"`
}

func assertHasType(t *testing.T, name string, list []domain.Type) {
	for _, type_ := range list {
		if type_.Name == name {
			assert.True(t, true)
			return
		}
	}
	assert.True(t, false)
}

func assertHasLicence(t *testing.T, name string, list []domain.Licence) {
	for _, licence := range list {
		if licence.Name == name {
			assert.True(t, true)
			return
		}
	}
	assert.True(t, false)
}
