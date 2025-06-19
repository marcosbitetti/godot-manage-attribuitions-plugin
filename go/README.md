# Godot Manage Attributions Plugin

A command-line tool to manage attributions for Godot projects.

## Why Manage Attributions?

Proper attribution management is crucial for any software project, especially games. When you use third-party assets, code, or tools, giving credit where credit is due isn't just good practice - it's often a legal requirement. This tool helps you maintain clear records of:

- Asset origins and creators
- License compliance 
- Usage rights and restrictions
- Required attributions for commercial/non-commercial use

By tracking attributions systematically, you:
- Show respect for other creators' work
- Build trust with the creative community
- Ensure legal compliance
- Make it easier to distribute your project
- Maintain clear documentation for future reference


## Commands

### General
- `help` - Display help information

### Attributions
- `listAttribuitions` - List all attributions
- `addAttribuition` - Add a new attribution
- `updateAttribuition` - Update an existing attribution
- `deleteAttribuition` - Delete an attribution

### Types
- `listTypes` - List all types
- `addType` - Add a new type
- `updateType` - Update an existing type
- `deleteType` - Delete a type

### Licenses
- `listLicences` - List all licenses
- `addLicence` - Add a new license
- `updateLicence` - Update an existing license
- `deleteLicence` - Delete a license

## Usage

The general command structure is:

```bash
attribuitions-amd64-linux <database-path> <json-command>
```

### Examples
#### Usage
```bash
attribuitions-amd64-linux help
```

#### Types
```bash
attribuitions-amd64-linux ~/mygames/attributions.sqlite listTypes
attribuitions-amd64-linux ~/mygames/attributions.sqlite addType {"name": "Font"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite updateType {"_id":1, "name": "FontNew"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite deleteType {"_id":1}
```

#### Licenses
```bash
attribuitions-amd64-linux ~/mygames/attributions.sqlite listLicences
attribuitions-amd64-linux ~/mygames/attributions.sqlite addLicence {"name": "Insaneware", "link": "https://example.com/license"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite updateLicence {"_id":1, "name": "Insaneware2", "link": "https://example.com/licenses"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite deleteLicence {"_id":1}
```

#### Attributions
```bash
attribuitions-amd64-linux ~/mygames/attributions.sqlite listAttribuitions
attribuitions-amd64-linux ~/mygames/attributions.sqlite listAttribuitions {"text":"<search>", "order": "ASC"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite addAttribuition {"name":"Test","filename":"file","type":"One","author":"Ze","link":"http://none","licence":"MIT","type":"Music"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite updateAttribuition {"_id":1,"name":"_Test","filename":"_file","type":"_One","author":"_Ze","link":"_http://none","licence":"Beerware","type":"Plugin"}
attribuitions-amd64-linux ~/mygames/attributions.sqlite deleteAttribuition {"_id":1}
```
