# h2m

Hobby project to learn golang. The goals is convert the blog `What color is your function` from html to markdown.
Maybe I'll generalize it, so it works with all sorts of blogs.

## TODO

- [ ] I haven't used interfaces at all and don't remember how they work in go. Are they similar to java interfaces? look into it and see if i have a usecase for them here.
- [ ] output markdown
- [ ] create todos based on my comment of 21/08/2025
- [ ] If I want to be able to get rid of the bloat inside html tags. I must store the beginning of the tag but also of the content. Or mayby I don't need to store the conent, and storing the start and endo of the tags is just enough, and I just replace the complete tag with the equivalent markdown content.
- [ ] after parsing the header start generating markdown based on the token stream
- [ ] checkout Go release 1.25 and the note about [faster slices](https://go.dev/doc/go1.25#faster-slices), [other article](https://bitstack.substack.com/p/go-125-compiler-update-stronger-alignment),[dreams of code video](https://www.youtube.com/watch?v=8fcjcoXXMVQ)

#### optimize

- [ ] use channels : see one of my notes
- [ ] use trie in lexer (see todo)

### References

[What is an idiomatic way of representing enums in Go?](https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go) \
[Learning Go, and the `type` keyword is incredibly powerful and makes code more readable](https://www.reddit.com/r/golang/comments/1at369q/learning_go_and_the_type_keyword_is_incredibly/?share_id=o6o5D80TE1fQ8s9dnDslU&utm_content=2&utm_medium=android_app&utm_name=androidcss&utm_source=share&utm_term=2) \
[moving from classes to golang](https://www.reddit.com/r/golang/comments/1ebcp85/moving_from_classes_to_golang/) \
[labeld loops](https://stackoverflow.com/questions/46792159/labels-break-vs-continue-vs-goto)
[mutable-strings-in-golang](https://medium.com/kokster/mutable-strings-in-golang-298d422d01bc) \
[strings-in-golang](https://www.geeksforgeeks.org/go-language/strings-in-golang/) \
[stack-overflow-how-to-check-value-in-map-go](https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go) \
[strings package of go](https://pkg.go.dev/strings?utm_source=chatgpt.com#HasPrefix) \
[TRIE in crafting interpretres chapter 16](https://craftinginterpreters.com/scanning-on-demand.html#tries-and-state-machines) \
[type keyword in go](https://stackoverflow.com/questions/53689968/what-exactly-does-the-type-keyword-do-in-go)\
[go by example slices](https://gobyexample.com/slices)\
[Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)\
[Go slices : usage and internals](https://go.dev/blog/slices-intro)\
[regex is fun](https://pkg.go.dev/regexp)\
[How to remove redundant spaces/whitespace from a string in Golang?](https://stackoverflow.com/questions/37290693/how-to-remove-redundant-spaces-whitespace-from-a-string-in-golang?utm_source=chatgpt.com) \

Type definitions (not type aliases) zijn heel krachtig. Kan je bijvoorbeeld van een int een type maken specifiek voor je domein. Je type erft geen methods van het bovenliggen type en je kan je eigen methodes voor dat type specifieren. Verder kan je niet impliciet van int naar je type converten. Je moet **expliciet casten**. Verder type system info opzoeken en mee experimenteren. Waarschijnlijk ook meer lightweight dan maken van een class in Java. Per file kan je meerdere types definieren.

## DONE

- [x] converter test outputs UNKNOWN token?
- [x] read about type keyword, slices and arrays in go. Because I forgot the details of the go reference I read before. allready added the references.
- [x] use trie like data structure in lexer (see todo)
- [x] fix anchor tag token not recognized -> dummy it's not in convertTokenToString
- [x] Get the content out of the h1 tag
- [x] anchor token, content token, anchor closing token. is enough to lex an anchor tag
- [x] recognized bug, document and use go strings HasPrefix package
- [x] rename peekCurrent, currPos, peekNext(). naming is confusing
- [x] fix makeContentTokenTest. After making htmlToken the '>' doesn't get consumed and is still in l.char. Consume last char after making the token sounds better, so the nextToken is properly set for the next run! breakpoint in consumeTag()
- [x] default option in lexer should not be readChar(), but should be makeContentToken().
- [x] fix bug lexing:`<article><header> content </header> </article>`
- [x] fix bug parsing `<article>` content `</article>` -- endless loop
- [x] install an LSP for markdown. I'm getting randomg error messaged in README.md... (nvim releated)
- [x] finish parsing first html element correctly
- [x] rename vars in lexer_test.go. atm confusing naming

## CANCELLED TASKS

- [x] rip out l.char and just use currPos() and just use currPos. See log 08/09/2025.
- [x] fix regressionTest

## Steps

- Get html content
- setup lexer
- define html tags to scan (tokens)

### Mapping html to markdown

| HTML Tag                   | Markdown Equivalent  |
| -------------------------- | -------------------- | ------- |
| `<h1>`                     | `#`                  |
| `<h2>`                     | `##`                 |
| `<h3>`                     | `###`                |
| `<h4>`                     | `####`               |
| `<h5>`                     | `#####`              |
| `<h6>`                     | `######`             |
| `<p>`                      | (blank line)         |
| `<br>`                     | `  ` or `\`          |
| `<strong>`, `<b>`          | `**` or `__`         |
| `<em>`, `<i>`              | `*` or `_`           |
| `<u>`                      | _(none)_             |
| `<del>`, `<s>`, `<strike>` | `~~`                 |
| `<a>`                      | `[]()`               |
| `<img>`                    | `![]()`              |
| `<ul>`                     | `-` or `*`           |
| `<ol>`                     | `1.`                 |
| `<li>`                     | `-`, `*`, or `1.`    |
| `<code>`                   | `` ` ``              |
| `<pre><code>`              | ` ``` `              |
| `<pre>`                    | (indented or fenced) |
| `<blockquote>`             | `>`                  |
| `<hr>`                     | `---` or `***`       |
| `<table>`                  | `                    | `       |
| `<tr>`                     | `                    | `       |
| `<th>`                     | `                    | `+`---` |
| `<td>`                     | `                    | `       |
| `<div>`, `<span>`          | _(none)_             |
| `<iframe>`, `<script>`     | _(none)_             |
