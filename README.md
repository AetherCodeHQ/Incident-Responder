# 🌐 Incident Responder

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.1.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Infrastructure tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`infrastructure` `devops` `cli` `golang`

---

## What is Incident-Responder?

**Incident-Responder** is an infrastructure tool for monitoring, inspecting, and managing systems and services.

## Features

- ✅ `severity()` — Severity
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Incident-Responder.git
cd Incident-Responder

# Build
go build -o incident-responder .

# Run
./incident-responder <log> [ack|escalate|resolve]
```

### Or directly with `go run`:
```bash
go run main.go <log> [ack|escalate|resolve]
```

## Usage

```bash
# Basic usage
./incident-responder <log> [ack|escalate|resolve]

# With flags
./incident-responder <log> [ack|escalate|resolve] value <log> [ack|escalate|resolve]
```

### Example Output

```
$ ./incident-responder <log> [ack|escalate|resolve]
<log> [ack|escalate|resolve]
incident acknowledged at %s\n
severity: %s\n
```

## Project Structure

```
Incident-Responder/
  main.go          # Entry point (62 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
