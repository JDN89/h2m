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

### References
[What is an idiomatic way of representing enums in Go?](https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go)
[Learning Go, and the `type` keyword is incredibly powerful and makes code more readable](https://www.reddit.com/r/golang/comments/1at369q/learning_go_and_the_type_keyword_is_incredibly/?share_id=o6o5D80TE1fQ8s9dnDslU&utm_content=2&utm_medium=android_app&utm_name=androidcss&utm_source=share&utm_term=2)


### Mapping html to markdown
### HTML to Markdown Tag Mapping


| HTML Tag                   | Markdown Equivalent                          |
|----------------------------|----------------------------------------------------------------------------------------------------|
| `<h1>`                     | `# Heading`                                   Add `#` symbols equal to heading level               |
| `<h2>`                     | `## Heading`                                | 
| `<h3>`                     | `### Heading`                               | 
| `<h4>`                     | `#### Heading`                              | 
| `<h5>`                     | `##### Heading`                             | 
| `<h6>`                     | `###### Heading`                            | 
| `<p>`                      | *(blank line)*                              | 
| `<br>`                     | `  ` (2 spaces) or `\`                      | 
| `<strong>`, `<b>`          | `**bold**` or `__bold__`                | 
| `<em>`, `<i>`              | `*italic*` or `_italic_`                    | 
| `<u>`                      | *(no equivalent)*                           | 
| `<del>`, `<s>`, `<strike>` | `~~strikethrough~~`                         | 
| `<a href="url">text</a>`   | `[text](url)`                               | 
| `<img src="url" alt="">`   | `![alt](url)`                               | 
| `<ul>`                     | *(wraps list items)*                        | 
| `<ol>`                     | *(wraps list items)*                        | 
| `<li>`                     | `- item` or `* item` or `1. item`         | 
| `<code>`                   | `` `inline code` ``                          | Inline code                                          |
| `<pre><code>`              | \`\`\`code block\`\`\`                       | Fenced code block                                    |
| `<pre>`                    | Indent with 4 spaces or fenced code block    | Often used with `<code>`                             |
| `<blockquote>`             | `> quoted text`                              | Quote block                                          |
| `<hr>`                     | `---` or `***`                               | Horizontal rule                                      |
| `<table>`                  | Pipe-based table                             | See `<tr>`, `<th>`, `<td>`                           |
| `<tr>`                     | `| row | row |`                              | Table row                                            |
| `<th>`                     | `| Header |` with `| --- |` below            | Table header cell                                    |
| `<td>`                     | `| Cell |`                                   | Table data cell                                      |
| `<div>`, `<span>`          | *(no equivalent)*                            | Often ignored or passed as raw HTML                 |
| `<iframe>`, `<script>`     | *(no equivalent)*                            | Usually ignored or kept as raw HTML                  |

