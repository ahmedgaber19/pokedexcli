# Pokedex CLI

A command-line interface (CLI) application for exploring Pokemon locations and catching Pokemon using the [PokeAPI](https://pokeapi.co/).

## Features

- 🗺️ Browse Pokemon location areas
- 🔍 Explore specific areas to find Pokemon
- ⚾ Catch Pokemon with a chance-based system
- 📊 Inspect caught Pokemon stats and attributes
- 📝 View your Pokedex collection
- 💾 Built-in caching for improved performance

## Installation

### Prerequisites

- Go 1.16 or higher

### Build from Source

```bash
git clone <repository-url>
cd pokedexcli
go build
```

## Usage

Run the application:

```bash
./pokedexcli
```

You'll be greeted with a REPL prompt:

```
Pokedex >
```

### Available Commands

| Command | Description | Example |
|---------|-------------|---------|
| `help` | Displays available commands | `help` |
| `map` | Displays the next 20 location areas | `map` |
| `mapb` | Displays the previous 20 location areas | `mapb` |
| `explore <location>` | Explores a specific location area and shows Pokemon found there | `explore canalave-city-area` |
| `catch <pokemon>` | Attempts to catch a Pokemon | `catch pikachu` |
| `inspect <pokemon>` | Shows detailed information about a caught Pokemon | `inspect pikachu` |
| `pokedex` | Lists all caught Pokemon | `pokedex` |
| `exit` | Exits the Pokedex | `exit` |

### Example Session

```bash
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays available commands
map: Displays location areas
mapb: Displays previous location areas
explore: Explore a specific location area
catch: Catch a specific Pokemon
inspect: Inspect a caught Pokemon
pokedex: List all caught Pokemon
exit: Exits the Pokedex

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
- tentacool
- tentacruel
- staryu
- magikarp
- gyarados
- wingull

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 455
Stats:
- hp: 40
- attack: 40
- defense: 35
- special-attack: 50
- special-defense: 100
- speed: 70
Types:
- water
- poison

Pokedex > pokedex
Your Pokedex:
 - tentacool

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Project Structure

```
pokedexcli/
├── main.go                          # Application entry point
├── repl/
│   └── repl.go                      # REPL (Read-Eval-Print Loop) logic
├── internal/
│   ├── pokeapi/
│   │   ├── types.go                 # API response type definitions
│   │   └── client.go                # HTTP client for PokeAPI
│   ├── pokecache/
│   │   ├── pokecache.go             # Caching implementation
│   │   └── pokecache_test.go        # Cache tests
│   ├── pokedex/
│   │   └── pokedex.go               # Application state/config
│   └── commands/
│       ├── commands.go              # Command registry
│       ├── help.go                  # Help and exit commands
│       ├── map.go                   # Map navigation commands
│       ├── explore.go               # Location exploration
│       ├── catch.go                 # Pokemon catching
│       └── inspect.go               # Pokemon inspection & listing
└── go.mod                           # Go module definition
```

## Architecture

### Separation of Concerns

- **main.go**: Minimal entry point that starts the REPL
- **repl**: Handles user input, command parsing, and execution loop
- **pokeapi**: Encapsulates all API interactions and response types
- **pokecache**: Implements a time-based cache to reduce API calls
- **pokedex**: Manages application state and configuration
- **commands**: Modular command handlers, one file per logical group

### Caching

The application uses an in-memory cache with a 5-second TTL to improve performance and reduce API calls. The cache automatically evicts expired entries.

### Catch Mechanics

When attempting to catch a Pokemon, there's a 50% chance of success based on a random number generator. This simulates the challenge of catching Pokemon in the games.

## Development

### Running Tests

```bash
go test ./...
```

### Running Specific Package Tests

```bash
go test ./internal/pokecache
```

## API

This project uses the [PokeAPI](https://pokeapi.co/) (v2), a free and open RESTful API for Pokemon data.

## License

This project is for educational purposes.

## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data
- The Pokemon Company for creating Pokemon
