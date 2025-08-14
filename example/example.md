# Markdown to HTML Converter (Go)

A simple Go program that converts a Markdown (`.md` or `.mk`) file into an HTML file using [Goldmark](https://github.com/yuin/goldmark).

## Features
- Converts `.md` and `.mk` files to HTML
- Simple CLI usage

## Installation
```bash
git clone https://github.com/yourusername/markdown-to-html.git
cd markdown-to-html
go mod tidy
```
## Usage
`go run main.go ./example/example.md`

or
```bash
go build main.go
./<executable> main.go ./example/example.md
```
## Example
Input (`example.md`)
```markdown
# Hello
This is **Markdown**.
```
Output (`output.html`)
```html
<h1>Hello</h1>
<p>This is <strong>Markdown</strong>.</p>
```
