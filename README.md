# Gramme
## The **Gr**ammar-based **M**odal **E**ditor

**Gramme** is a modal editing system designed around a formal language of text transformation.

### Getting started

Requires Go v1.15+

- Build `go build`
- Test `go test ./...`
    + Recursively runes tests in subpackages of the root module
- *Note*: Some code relies on generated files. If their sources are changed, run `go generate ./...` to re-synthesize files.

### Features

1. **Explorable**: Gramme's command system is designed from the ground up to make it easy for new and advanced users alike to learn and explore. It includes tools for searching for commands, breaking down commands to understand how they're composed, and discovering new commands similar to ones you already know.

2. **Consistent**: Gramme formalizes some of the concepts underlying other modal editors like Vim and Kakoune. The result is a more consistent and predictable editing experience for new users.

3. **Extensible**: Gramme's underlying principles make it easy to extend existing behaviors in a predictable way.

### Concepts

A **command** is a function which operates on the editor's state to produce a new state. A command is influenced only by the current editor state, and cannot take any parameters. They represent very specific, granular user interactions, and are generally not defined or invoked directly by users except in very simple cases.

> One of the simplest commands in Gramme is the `insert` command, which allows users to insert text without the command system interpreting incoming text as commands!

Instead, users compose commands from words. A **word** is simply a sub-component of a command. Alone, a word does not have enough information to meaningfully change the editor's state, but by combining several words users can form a complete command. A **phrase** is a combination of several words, which may or may not form a complete expression. An **expression** is a phrase which forms a complete command.

> The expression `delete inside sentence` defines a command which deletes all text within the sentence underneath the cursor. It is composed of three **words**: `delete`, `inside`, and `sentence`. If the word `sentence` is dropped, then we're left with the phrase `delete inside`. This is not a complete expression - what exactly are we deleting?

The meaning of an expression is composed of the meanings of the words that form it. Varying one word in an exprssion should produce a similar expression whose behavior varies according to the difference between the substituted words. Not all words can replace other words and produce a meaningful expression. Words which can be substituted for one another have the same **part of speech**.

> The *text-object* part of speech includes the word `sentence`, as well as other words like `paragraph`, `quotation`, `buffer`. They all define patterns that identify different components of a body of text. Just like the English sentence "Delete the text inside qutoation marks under the cursor" is very similar to "Delete the text inside of the sentence under the cursor.", the command `delete inside quotation` is very similar to `delete inside sentence`.

**Verbs** are the primary words which drive commands. Just like verbs in natural languages like English, they define what a command does. Unlike commands, verbs can take parameters in the form of other words, called the verb's **complement**. A complement refines the behavior of the verb, allowing users to alter the verb's behavior without having to define an entirely new word.

> In the expression `delete inside sentence`, the word `delete` is a verb. The phrase `inside sentence` is the verb's complement. The verb defines what we want to do (e.g. delete some text) and the complement refines that action (e.g. by specifying what to delete).

**Grammars** are what link exprsesions and commands together. Each verb is associated with a single grammar, which defines the expected pattern for that verb's complement as a sequence of parts of speech. Gramme expects every expression to *begin with a verb*. This allows Gramme to unambiguously parse user input by first finding a verb, then interpreting the remaining input according to that verb's complement, eventually producing a complete expression.

```
Parsing an Expression:

1. Parse the first word and interpret it as a verb
2. Find the grammar associated with that verb to determine the part of speech pattern for its complement
3. Parse a number of additional words based on the complement and interpret them according to their respective part of speech
4. Once a complete verb and complement are successfully parsed, combine them into a complete expression
5. Transform the expression into a parameterless command, which operates solely on the current editor state
```

It would be very tedious to type out the full name of each word to build up commands. Instead, each word may be bound or associated with a sigil. A **sigil** is a single character which uniquely identifies a word within a particular part of speech. A **binding** allows users to reference the corresponding word using only the sigil any time the words associated part of speech is referenced. The full list of possible sigils is given below:

```
abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-_=+[{]}|;'",<.>/?]`~
```

Essentially any character on the standard US keyboard except `:` and `\`.

> The words `insert` and `inside` are both bound to the sigil `i` to make it mnumonically easier for users to remember. This is possible because the words belong to different parts of speech: `insert` is a *verb*, whereas `inside` is a *preposition*. Any time `i` is used where a verb is expected, it will be interpreted as `insert`. Similarly, if a preposition is expected, it will be interpreted as `inside`. Similarly, the verb `delete` is bound to `d`, and the text object `sentence` is bound to `s`, allowing the expression `delete inside sentence` to be concisely input as `dis`.

The `:` character is used to reference a word directly by its unique identifier. An identifier consists of any number of alphanumeric characters, possibly separated by `-`, `_`, or `.` characters. If a `:` is encountered, then all input until a whitespace character is treated as a word identifier.

> For example, the `dis` expression is equivalent to `:delete :inside :sentence`, but also to `:delete is`, `di:sentence`, and so on.

The `\` character is used to represent a single literal character of text. For exmaple `\i` represents the character `i`, and not any word like `insert` or `inside`.

### Other important concepts

A **cursor** is a reference to a particular location within a body of text. A **span** of text is the space betwen two locations. A **region** is a span of text between two delimiting spans. Each delimiting span is defined by two locations, and the delimiters implicitly define a span between them:

```
hello[( )there( )]world!
------|-|-----|-|-------
      | |     | |
      a b     c d
```

The region here is indicated by the square bracket (`[...]`) characters. The two delimiters are indicated by the groups of parenthesis (`(..)`) on either side of the word `there`. Thus, the region is defined by 4 locations, labelled `a`, `b`, `c`, and `d` above.

A **selection** consists of a region and a cursor position within the region. A selection is said to be **forward** if the cursor is at the highest index within the region, **backward** if it is at the lowest index, and **indeterminate** otherwise.

### Tools

Gramme includes a number of tools to help users explore, learn, and extend the system.

1. **Phrase diagram**: Breaks down a phrase into consitutent parts, annotating each word with its part of speech and indicating expressions and incomplete phrases.

2. **Interactive command invocation**: Gramme can dynamically generate completion suggestions based on possible options for incomplete expressions.

3. **Interactive docs**: Gramme can display a human readable name, description, documentation, and examples for any word in the system.

4. **Complete the phrase**: Given a phrase with blank spots, Gramme can provide a list of words which complete that phrase by inspecting registered grammars.

5. **Unbound keys**: Gramme can display all bindings as well as all unbound keys for a part of speech

6. **FSM**: Gramme can generate a finite state machine diagram for parsing a single expression in the form of a DAG, which can be visualized.

