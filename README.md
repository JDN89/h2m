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

## LOG -- quick notes, full of spelling mistakes and chaotic ramblings

### 02/10/2025

** NO POINTERS**

[value receiver vs pointer receiver](https://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver?utm_source=chatgpt.com) \
[using pointers vs copy in struct functions](https://stackoverflow.com/questions/22685062/using-pointers-vs-copy-in-struct-functions) \

`For example, you could define an interface that defined the Log() method and create a variable of this type. Then you could assign an instance of the Logger structure to that variable. You could also assign a pointer to an instance of the Logger structure to this variable. Both would work, because the Log() method is callable from both instances of the structure and pointers to instances. If the method took a pointer argument, then you would only be able to call it on pointers. It’s therefore good style in Go to only require methods to take a pointer when they modify the structure, or if the structure is so large that copying it on every method call would be prohibitive...`

- copying small structs is cheap
- no pointer juggling
- pointer can be usefull if I want to mutate the value

In the converter I am at the moment just using a for loop and indexing into the tokens one by one.
I think that it's better to work with next token and current token. I don't wat to start peeking the next token (in some cases) inside my for loop. So i'll add currToken and nextToken to my converter struct. and just keep looping as long as token.type is not EOF. maybe currToken and nextToken should be the actual token or a pointer to the actual tokens?
Now that I think about it, I think the same approach was used in `crafting your own `

[learned about the title attribute in an anchor tag](https://www.w3schools.com/html/html_links.asp) : The title attribute specifies extra information about an element. The information is most often shown as a tooltip text when the mouse moves over the element.
[also learned bout the rel attribute](https://www.w3schools.com/TAGS/att_a_rel.asp) The rel attribute specifies the relationship between the current document and the linked document. Only used when href attribute is present.

For my markdown I don't need those so I could just skip these attributes. I am also thinking of ignoring href wrapped inside headers. But for now I'll leave them in and I'll reconsider once the converter is finished.

### 30/09/2025

Fixed the lexer. For some reason, I thought the lexer allready worked correctly in the past. Looking at my log of 21/09/2025. zero deaths refactor? The fuck Jan - zeker moe? Anyway, fixed it this time - first attempt refactor (but not really)

### 29/09/2025

The first time in forever that I have underarm pain from using my mouse.

NOt much done today. The past days I was reading some regex reference guides and doing some tutorials. I managed to parse the href attributes, via regexp, and put them in the token.Attributes.
Fix the broken lexer later, Arm is to messed up.

### 23/09/2025

Hit a block with converting the anchor tag. Decided to expand the token and add a map of attributes that I just fill in in case of an open anchor tag. The issue is that I don't know the string libraries enough to lex the anchor tag in a good way. Buuuut, I saw somewhere that you can use regex to get substrings out of a string. The first anchor tag contains href, rel, title attributes and I think regex is the best way to parse that. downside is that I don't know regex. Upside, I always want to learn the basics of regex, so I've go the perfect excuse. I was starting to use the string library functions of go, but it seems a bit cumbersome to do waht I want to do.
Maybe I just do it, so I get used to it as well. and If i have time and motivation left I implement a regex version as well.

No I'll immediatly do the regex impl [regex is fun](https://pkg.go.dev/regexp)

### 21/09/2025

Fixed a bug in my lexer. Because l.char loads the last consumed character. currPos always points one character to far for what I want, so l.startPos = l.currPos always points one char to far. I feel this is unnecessary complexity. This has been bothering me since the start. I am doing stuff like +1 and -1 which doesn't make sense. let currPos just look at the currPos and don't work with consumed char,... create a function that just returns the char at currPos and another funtion that peeks to the next char and returns it in cased I need a functio like that.

Lol, first try refactoring on refactor-lexer-v2 and immediatly got it right. I think subcounciously I was thinking and solving this allready. It was also an easy refactor, but still. zero deaths. time to merge!

### 20/09/2025

LOOOL -- my whole setup was wrong. I thought I was going to maninpualte the existing slice, BUUUT I can just use the golang append function. WHICH means that I don't have to go over the tokens in reverse order and thus replace the input from back to front. I can just go over the tokens in chronological order and keep appending.

This is what happens if you code in 5 to 30 minute sessions over a period of 2 to 3 months. When did I start this repo?

I read the go spec 2 months ago. Now that I reread the slice and array part. I realize my whole plan didn't make sense. I even think I don't need the closing tag tokens... Or at least not for a simple blog article that doesn't use tags like li ul th ,....

Even my initial setup of going over it in reverse order is not good, because the performance would be bonkers if after each token conversion I move the input x chars over to replace the empty spots
It is interesting tough to find out if there will be performance gains if I just manipulate the existing input. The only question I have is how I willI deal with the null values. I can just move all the chars over x positions, with each html token replacement, but that will be an expinsive action. Better to mark the empty spaces with nill and do a sort that pushes the nill values to the end ones all the tokens have been processed. I don't even know if i have to sort?. can't i just put empty '' in those spots? Probably, but I still will have to iterate over them and the byte buffer will be larger then need be. But let's find out shall we ! :)

```
['<','h','1','>', ' ', 'c','o','n','t','e','t', ' ', '<', '/','h','1','>']
BECOMES
['#','','','', ' ', 'c','o','n','t','e','t', ' ', '/n', '','','','']
```

I read the go slices-intro. I will use the append function. I will go over the tokens. The html tokens I convert to markdown and append to the slice, content tokens I can just slice the input positions and append. I am not really sure if I need the closing tags? For now I will just place a linebreak when I encounter a closing tag (not taking into account the anchor tag).
The append function appends the elements x to the end of the slice s, and grows the slice if a greater capacity is needed.

```GO
a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}
```

To append one slice to another, use ... to expand the second argument to a list of arguments.

```GO
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

Another dump, maybe cleanup logs later (prob not)

### 14/09/2025

I had some time this afternoon, to program, but instead I spend the better part of an hour on tweaking my neovim lsp config... Got it working tough :D Somehow signature help didn't work (anymore?).
I found a bug in recoginzing html tags. I usualy look exact match for recognizing tags. but tags can have attributes, ... so exact match for a tag doesn't work. I found a strings.HasPrefix builtin method of go, and intead I'll create the relevant tokens if the tag starts with ....

I still have an issue with recognizing anchor tags, but is for tomorrow. Today I fixed my lsp. Which is also time well spend, seeing that I love using neovim.

Solved anchor tag issue. I think the simplest thing I can do now to fix the BUG where tags don't get recognized if they have attribtes, because now I look up a tag based on an exact match on `<h1>`, is by implementing a TRIE similar to the implementation of [TRIE in crafting interpretres chapter 16](https://craftinginterpreters.com/scanning-on-demand.html#tries-and-state-machines)
Place TRIE logic in consumeTag and refactor makeHtmltoken

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
