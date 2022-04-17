# Old Norwegian Dictionary Builder

Build "Dictionary of the Old Norwegian Language" into easier-to-use data formats.

Available formats:
- JSON
- DSL
- XML

Hardcoded packages/modules/classes for programming languages:
- Go package
- TypeScript module
- Rust module
- Python class

The hardcoded language packages exist to benchmark hardcoded entries against reading them from JSON/similar. The results vary based on language. For example, Rust and Go are much faster hardcoded. In case of Rust it hundreds of ms vs hundreds of µs.

TypeScript is slightly faster hardcoded, but Python on the other hand is much slower if hardcoded.

### Usage

Main package exposes `ToJson`, `ToDsl`, `ToXml`, `ToGoPackage`, `ToRustModule` and `ToTypeScriptModule` functions, which respecticely generate output files in /build/ directory. Running `main` function generates all outputs.

### About "Dictionary of the Old Norwegian Language"

_"Ordbog over det gamle norske Sprog"_ dictionary was published in late 1800s by Johan Fritzner. Its is the largest Old Norse to Norwegian dictionary, containing over 40 000 word definitions. While the original dictionary is called dictionary of "old norwegian", it is practically a dictionary of western Old Norse. Technically "Old Norwegian" would be a later stage in the language.
