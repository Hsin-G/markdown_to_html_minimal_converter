# Markdown to HTML Converter (Go)

A simple Go program that converts a Markdown (`.md` or `.mk`) file into an HTML file using [Goldmark](https://github.com/yuin/goldmark).

## Features
- Converts `.md` and `.mk` files to HTML
- Costumizable HTML output file with css
- Simple CLI usage

## Installation
```bash
git clone https://github.com/Hsin-G/markdown_to_html_minimal_converter.git
cd markdown_to_html_minimal_converter
go mod tidy
go build .
```
## Usage
`./<executable> ./example/example.md`

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
