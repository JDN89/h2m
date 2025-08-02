# h2m

A tool that converts HTML (primarily from blogs) into Markdown.
The concept is to input a URL into the CLI tool, which then saves the Markdown file to the ~/Documents folder.

## TODO
- [ ] install an LSP for markdown. I'm getting randomg error messaged in README.md... (nvim releated)
- [ ] finish parsing first html element correctly
- [ ] rename vars in lexer_test.go. atm confusing naming

## Steps
- Get html content
- setup lexer
- define html tags to scan (tokens)

## NOTES

### 02/08/2025
Started writing the lexer. I miss writing with an LSP, but the good thing is that I seem to remember the syntax better, and I see the use of writing short concise names (which I am not doing at the moment). Now that I don't have an LSP, I am dependend on the error messages of the compiler, which shows how important error messages are. I feel like the error messages of the Rust compiler are better, but atm I don't like writing Rust. 
Learned about writing and running tests in Go, packages, ...
