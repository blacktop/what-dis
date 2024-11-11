<p align="center">
  <a href="https://github.com/blacktop/what-dis"><img alt="what-dis Logo" src="https://raw.githubusercontent.com/blacktop/what-dis/main/docs/logo.png" height="512"  /></a>
  <h1 align="center">what-dis</h1>
  <h4><p align="center">Dumb image-to-text experiment</p></h4>
  <p align="center">
    <a href="https://github.com/blacktop/what-dis/actions" alt="Actions">
          <img src="https://github.com/blacktop/what-dis/actions/workflows/go.yml/badge.svg" /></a>
    <a href="https://github.com/blacktop/what-dis/releases/latest" alt="Downloads">
          <img src="https://img.shields.io/github/downloads/blacktop/what-dis/total.svg" /></a>
    <a href="https://github.com/blacktop/what-dis/releases" alt="GitHub Release">
          <img src="https://img.shields.io/github/release/blacktop/what-dis.svg" /></a>
    <a href="http://doge.mit-license.org" alt="LICENSE">
          <img src="https://img.shields.io/:license-mit-blue.svg" /></a>
</p>
<br>

## Why? 🤔

🤷

## Getting Started

### Requirements

- [ollama](https://ollama.com)
- `/usr/bin/say` *(might be **macOS** only?)*

### Install

```bash
go install github.com/blacktop/what-dis@latest
```

Or download the latest [release](https://github.com/blacktop/what-dis/releases/latest)

```bash
❯ what-dis

Describe an image

Usage:
  what-dis <IMAGE> [flags]

Flags:
  -h, --help      help for what-dis
  -t, --text      Output as text
  -V, --verbose   More detailed output
```

```bash
❯ what-dis docs/logo.png --text
```

> The image depicts a young boy and a robot standing in front of an abstract painting. The boy is wearing a white t-shirt and khaki shorts, while the robot has a sleek, futuristic design with glowing blue eyes and a round body. The background features a gold-framed picture on the wall, adding to the overall sense of wonder and curiosity emanating from the scene.

## License

MIT Copyright (c) 2024 **blacktop**