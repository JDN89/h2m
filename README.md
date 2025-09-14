# h2m

Hobby project to learn golang. The goals is convert the blog `What color is your function` from html to markdown.
Maybe I'll generalize it, so it works with all sorts of blogs.

## TODO

- [ ] recognized bug, document and use go strings HasPrefix package
- [ ] anchor token, content token, anchor closing token. is enough to lex an anchor tag
- [ ] fix anchor tag token not recognized -> dummy it's not in convertTokenToString
- [ ] Get the content out of the h1 tag
- [ ] parse a href
- [ ] create todos based on my comment of 21/08/2025
- [ ] If I want to be able to get rid of the bloat inside html tags. I must store the beginning of the tag but also of the content. Or mayby I don't need to store the conent, and storing the start and endo of the tags is just enough, and I just replace the complete tag with the equivalent markdown content.
- [ ] after parsing the header start generating markdown based on the token stream
- [ ] checkout Go release 1.25 and the note about [faster slices](https://go.dev/doc/go1.25#faster-slices), [other article](https://bitstack.substack.com/p/go-125-compiler-update-stronger-alignment),[dreams of code video](https://www.youtube.com/watch?v=8fcjcoXXMVQ)

## DONE

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

## LOG

### 14/09/2025

I had some time this afternoon, to program, but instead I spend the better part of an hour on tweaking my neovim lsp config... Got it working tough :D Somehow signature help didn't work (anymore?).
I found a bug in recoginzing html tags. I usualy look exact match for recognizing tags. but tags can have attributes, ... so exact match for a tag doesn't work. I found a strings.HasPrefix builtin method of go, and intead I'll create the relevant tokens if the tag starts with ....

I still have an issue with recognizing anchor tags, but is for tomorrow. Today I fixed my lsp. Which is also time well spend, seeing that I love using neovim.

### 11/09/2025

Using 2 booleans was making it very cumbersome to see in which state (you get it state) I am. I was using start and stop to indincate that I encountered the open_article and close_article tokens and that I could start generating and stop generating the tokens. I now converted it to a state machine which make it easier to have an overview of which state we are in. from there I als call the appropriat lexing functions, which is logic I pulled of the nextToken function. to make it more readable.
If your program swithces between states/modes just use enums instead of booleans. More descriptive.
Inside the lex functions I can use substates to parse complex tags. Tag_name tag_xxx. but this I probably won't need in my case.

### 08/09/2025

Current bug is mismathc between when I use currPos and lastConsumedChar. sometimes I switch on currPos and other times on lastConsumedChar, which causes bugs and mismatch between how I expect the lexer to work. Because In some cases I have to sonsume the currPos so it gets loaded in l.char because that is the position I am switching on. What if I rip out lastConsumedChar and just use currPos and currPos + 1? I think that will cause less confusion for me.

Try this week on fix branch

### 03/09/2025

I was going to make the default case in my lexer, makeContentToken and consider everyting
between tags as content. What with tabs, whitespace, enter,...
often tags are seperated by whitespace.
So consume as long as we don't encounter a char or a number? I don't think so
because you can have sentences that start with special characters.
Alternative, consume whitespace, tabs, newlines,... until we encounter something.

### 02/09/2025

Default option in lexer should not be readChar(), but should be makeContentToken. This is text that can be left as is by the transpiler. The exception is the text that follows the a href tag, because that is the link, name? but this can be easilty handled by the transpiler, but checking the nextToken(). Based on the token sequence create `[title](link)`

### 01/09/2025

Html can contain, classes, id's, styling, ... In consume tag, I only advance when I encounter a char (a-z) or a number. An easier solution is to consume every possible char until I encounter >.

I find like stretches of 5 minutes to 30 minutes a night. Having a newborn is demanding. 5 min sesj done!

### 31/08/2025

Issue with what color has your function article. The H1 has an `<a href ...>` wrapped inside of it and the href wraps the the header title. My first idea was, once you parse a tag, to ingnore all the inner tags, but this won't work for obvious reasons. But I can do this for h1 that I skip all the inner tags and just get the content out. I don't see a scenario where I need something besides the h1 title inside the h1 header.
I just realised this is a bad idea. I was going consume the href nested in the `<h1>` header, but this adds all sorts of complexity to my lexer. Like instead of just lexing the <a href.. > and returing an anchor tag token, I was goign to build an exception that just consumes the chars inside het `<h1>` tags until I hit a '>' of the anchor tag element. SO WHAT if my lexer consumes and produces an anchor tag token!!?? My lexer shouldn't care and just spit out tokens wether or not the position of the tokens is convenient or not. It's up to the markdown Generator to decide wheter a token should be discarded or not, based on the position in the token stream.
Learned. Don't add unecessary complexity, just lex and spit out tokens -- keep your lexer **dumb**. Then decide in the next steps what to do with the generated tokens.

This does mean that I must start spitting out content tokens, because otherwise I'd don't know where the content is situated in nested tags!

```html
<h1>
  <a
    href="/2015/02/01/what-color-is-your-function/"
    rel="bookmark"
    title="Permanent Link to What Color is Your Function?"
  >
    What Color is Your Function?
  </a>
</h1>
```

My original idea was to just see where tag it's start and ending position was of `<h1>` and `</h1>`, replace `<h1>` with # and remove `</h1>`, and leave all the rest untouched. writing this out i can just do the same with the nested <a tag> in case of `<h1>` i have to remove it, but this means that the content is still untouched. I'll just add the content tokes to be sure, because I'll probably encounter some edge cases. In the polish verison, I can add not the start and end pos of the content, but use the actual string, and in another case I can leave it out and do what I just described. Then I can see what is the most efficient.

### 21/08/2025

Just found out that in go String are immutable, meaning that if want to replace parts of the input based on the start and en position of a tag, then I'm creating a new string each time... My original Idea was to store start and end of the tags, and go over the token stream in reverse and replace the input from end to start. But this plan is now ruined... Or expensive, I think.

Scratch that -> byte buffers are mutable, so I can, instead of mutating the bye[] to string just pass the byte buffer to lexer and work and that and then pass it to the parser and start mutating the byte buffer in de markdown gen by working from back to front! yay

Why front to back, because in that case i don't have to recalculate the token positions in the text based on the allready replaced text. Write this out clearer, but. `<h1> some title </h1> ` would be start token.H1 token.pos = 0 token.end 4 -> after replacing the html i'd get # some title `</h1>` BUT the start and end position of `</h1>` now have moved up with 3 positions in the text. This would not happen when I manupulate the input in revere.

Parellel processing token stream? What if lex and store the stream in memomry before passing it to the markdown generator? In that case I can split upt the token stream in equal parts -> pass the to named go channels to process parallell and then join the input from the channels in logical order. Would this make my h2m converter faster??? experiment with this in the end

I can do the same for lexing. Do first pass to identify logical boundries in the markdown -> then pass the logical boundries to channels that lex. `<h1> <h2>, <div> <p>.`

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

[mutable-strings-in-golang](https://medium.com/kokster/mutable-strings-in-golang-298d422d01bc) \
[strings-in-golang](https://www.geeksforgeeks.org/go-language/strings-in-golang/) \
[stack-overflow-how-to-check-value-in-map-go](https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go)

[strings package of go](https://pkg.go.dev/strings?utm_source=chatgpt.com#HasPrefix)
Type definitions (not type aliases) zijn heel krachtig. Kan je bijvoorbeeld van een int een type maken specifiek voor je domein. Je type erft geen methods van het bovenliggen type en je kan je eigen methodes voor dat type specifieren. Verder kan je niet impliciet van int naar je type converten. Je moet **expliciet casten**. Verder type system info opzoeken en mee experimenteren. Waarschijnlijk ook meer lightweight dan maken van een class in Java. Per file kan je meerdere types definieren.

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
