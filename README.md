# Old Norwegian Dictionary Builder

Build "Dictionary of the Old Norwegian Language" into easier-to-use data formats.

Available formats:
- JSON
- DSL
- XML

Hardcoded packages/modules/classes for programming languages:
- Go package
- TypeScript module
- Python class

The hardcoded language packages exist to benchmark hardcoded entries against reading them from JSON/similar. Seems to depend on language which approach is faster. For example, Go is much faster hardcoded, TypeScript seems a bit faster without file reading, whereas Python seems to be much slower when not read from a file.

### Usage

Main package exposes `ToJson`, `ToDsl`, `ToXml`, `ToGoPackage` and `ToTypeScriptModule` functions, which respecticely generate output files in /build/ directory. Running `main` function generates all outputs.

### About "Dictionary of the Old Norwegian Language"

_"Ordbog over det gamle norske Sprog"_ dictionary was published in late 1800s by Johan Fritzner. Its is the largest Old Norse to Norwegian dictionary, containing over 40 000 word definitions. While the original dictionary is called dictionary of "old norwegian", it is practically a dictionary of western Old Norse. Technically "Old Norwegian" would be a later stage in the language.
