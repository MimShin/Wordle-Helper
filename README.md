# Wordle Helper
Helps you with solving **Wordle** and **Don't Wordle**. Enter the word and the color pattern from Wordle (**\*** for green, **+** for yellow, and **-** for grey*) and it gives you all remaining words you can use.

## Usage
```bash
% main <dictionary-file>
```
## Example
```
% go run main.go dict_5.txt
```

## Commands
```
% go run main.go dict_5.txt

Let's go
> h
Available commands:
h  (help)
a word key  (try new word)
  *: right char at right pos
  +: right char at wrong pos
  -: wrong char
  e.g.
  a style *--+-
u  (undo last word)
p  (print history)
w  (print words)
q  (quit)
```

## Playing
```
% go run main.go dict_5.txt

Let's go
> a style *--+-
filtering: word=style, key=*--+-
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-


Let's go
> a slime **++-
filtering: word=slime, key=**++-
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-
count:     0, filter: slime, key: **++-


Let's go
> u
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-


Let's go
> a slime **---
filtering: word=slime, key=**---
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-
count:    18, filter: slime, key: **---


Let's go
> w
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-
count:    18, filter: slime, key: **---

[slack slang slank slash sloan slock sloka slonk sloop slops slorp slosh slour slows slung slunk slurp slush]

Let's go
> a slang **++-
filtering: word=slang, key=**++-
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-
count:    18, filter: slime, key: **---
count:     1, filter: slang, key: **++-


Let's go
> w
count:  9972, filter: , key: 
count:    89, filter: style, key: *--+-
count:    18, filter: slime, key: **---
count:     1, filter: slang, key: **++-

[sloan]

Let's go
> 
```