# h2m

A tool that converts HTML (primarily from blogs) into Markdown.
The concept is to input a URL into the CLI tool, which then saves the Markdown file to the ~/Documents folder.
I know there are (better) tools out there but this is just a hobby project to learn golang.

## TODO
- [ ] fix bug parsing <article> content </article> -- endless loop
## DONE
- [x] install an LSP for markdown. I'm getting randomg error messaged in README.md... (nvim releated)
- [x] finish parsing first html element correctly
- [x] rename vars in lexer_test.go. atm confusing naming
## Steps
- Get html content
- setup lexer
- define html tags to scan (tokens)

## NOTES


### 15/08/2025
I wish I had some more time to code, but with my baby of 4 months it's just not possible. It is interesting tough to code in these small slots of 15 minutes to 1 hour in the evening. The most dangerous thing is getting to excited and staying up to late, which happens a lot :) 
I am pretty excited about golang so for. I was surprised as how fast you can become semi productive in Go. I am sure my knowledge of C helps a lot.

I notice that I like working with languages that have little keywords. The less keywords to more readable the code is. I don't have a lot of experience but I have seen some Java code that was more complex than needed to be. When I see rust I get tired. I feel Rust is code that can get overengineerd very quickly.

I also like go's explicit error handling. I Java you can throw exceptions and you don't see where they end up, unles you follow the path upstream. No hidden control flow. That's also the reason why I'm not a fan of abstract classes. You have to jump back and forth between the impl class and abstract class to have a good view of what is going on, plus is the abstract method overriden.

| Language | Keywords |
| -------- | -------- |
| **C**    | 32       |
| **Go**   | 25       |
| **Java** | 53–68    |
| **Rust** | \~53     |

### 12/08/2025
Difficulties of deciding where to put the responsabiity of making a token. Should readIdentifer return a token or just advance over the relevatn characters and return? I think it should return a Token. Altough I haven't written it in this way so far.
The more you nest readChar, the harder it is to have an overview of where in the code I am reading chars
### 02/08/2025
Started writing the lexer. I miss writing with an LSP, but the good thing is that I seem to remember the syntax better, and I see the use of writing short concise names (which I am not doing at the moment). Now that I don't have an LSP, I am dependend on the error messages of the compiler, which shows how important error messages are. I feel like the error messages of the Rust compiler are better, but atm I don't like writing Rust. 
Learned about writing and running tests in Go, packages, ...

### References
[What is an idiomatic way of representing enums in Go?](https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go) \
[Learning Go, and the `type` keyword is incredibly powerful and makes code more readable](https://www.reddit.com/r/golang/comments/1at369q/learning_go_and_the_type_keyword_is_incredibly/?share_id=o6o5D80TE1fQ8s9dnDslU&utm_content=2&utm_medium=android_app&utm_name=androidcss&utm_source=share&utm_term=2) \
[moving from classes to golang](https://www.reddit.com/r/golang/comments/1ebcp85/moving_from_classes_to_golang/) \
[labeld loops](https://stackoverflow.com/questions/46792159/labels-break-vs-continue-vs-goto)

Type definitions (not type aliases) zijn heel krachtig. Kan je bijvoorbeeld van een int een type maken specifiek voor je domein. Je type erft geen methods van het bovenliggen type en je kan je eigen methodes voor dat type specifieren. Verder kan je niet impliciet van int naar je type converten. Je moet __expliciet casten__. Verder type system info opzoeken en mee experimenteren. Waarschijnlijk ook meer lightweight dan maken van een class in Java. Per file kan je meerdere types definieren.

### Mapping html to markdown

| HTML Tag                   | Markdown Equivalent         |
|----------------------------|-----------------------------|
| `<h1>`                     | `#`                         |
| `<h2>`                     | `##`                        |
| `<h3>`                     | `###`                       |
| `<h4>`                     | `####`                      |
| `<h5>`                     | `#####`                     |
| `<h6>`                     | `######`                    |
| `<p>`                      | (blank line)                |
| `<br>`                     | `  ` or `\`                |
| `<strong>`, `<b>`          | `**` or `__`                |
| `<em>`, `<i>`              | `*` or `_`                  |
| `<u>`                      | *(none)*                    |
| `<del>`, `<s>`, `<strike>` | `~~`                        |
| `<a>`                      | `[]()`                      |
| `<img>`                    | `![]()`                     |
| `<ul>`                     | `-` or `*`                  |
| `<ol>`                     | `1.`                        |
| `<li>`                     | `-`, `*`, or `1.`           |
| `<code>`                   | `` ` ``                     |
| `<pre><code>`              | ```` ``` ````               |
| `<pre>`                    | (indented or fenced)        |
| `<blockquote>`             | `>`                         |
| `<hr>`                     | `---` or `***`              |
| `<table>`                  | `|`                         |
| `<tr>`                     | `|`                         |
| `<th>`                     | `|` + `---`                 |
| `<td>`                     | `|`                         |
| `<div>`, `<span>`          | *(none)*                    |
| `<iframe>`, `<script>`     | *(none)*                    |

