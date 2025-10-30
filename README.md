# Calculator Thing

It is just an educational excercise.
Built a basic calculator that can do `2 + 3 * 4` and `(2 + 3) * 4` correctly.

## Quick overview

Three steps:
1. **Lexer** - text to tokens (`"2 + 3"` becomes `[2, PLUS, 3]`)
2. **Parser** - tokens to tree (this is the hard part!)
3. **Eval** - walk the tree, do the math

## Why trees?

Example: `2 + 3 * 4`

We want `*` to happen first, so build a tree where it's deeper:
```
      +
     / \
    2   *
       / \
      3   4
```

Eval just walks the tree recursively:
- `+` needs left (2) and right (`*`)
- `*` needs left (3) and right (4) → returns 12
- `+` gets 2 + 12 → returns 14

Done!

## The tricky part: building the tree

How to go from `[2, +, 3, *, 4]` to the correct tree?

**My approach:** Each thing has a priority (numbers highest, then `*`/`/`, then `+`/`-`)

Process tokens left-to-right. When you hit an operator:
1. Start at current position
2. Climb UP the tree until you find something with lower priority
3. Insert the operator there, "stealing" the right child

**Example walkthrough:** `2 + 3 * 4`

Start: `2` → just make it root

See `+`: Lower priority than 2, so `+` becomes root, 2 goes left
```
   +
  /
 2
```

See `3`: Just add to right
```
   +
  / \
 2   3
```

See `*`: Higher priority than `+`! 
- Climb up from 3 → hit `+` 
- `+` has lower priority, so insert `*` here
- Steal the 3, make it `*`'s left child
```
   +
  / \
 2   *
    /
   3
```

See `4`: Add to right of `*`
```
   +
  / \
 2   *
    / \
   3   4
```

## Parentheses: the depth trick

Parentheses just mean "higher priority". So I use a depth counter:
- `(` → depth++
- `)` → depth--

When comparing priorities: higher depth = higher priority (regardless of operator type)

Example: `(2 + 3) * 4`
```
Tokens: ( 2 + 3 ) * 4
Depth:  1 1 1 1 0 0 0
```

The `+` is at depth 1, the `*` is at depth 0. Even though `*` normally wins, depth 1 > depth 0, so `+` gets higher priority and stays deeper in the tree.

Result:
```
      *
     / \
    +   4
   / \
  2   3
```

That's it! Took me a while to figure out I needed to DECREMENT on `)` not increment lol.

## Main takeaway

Parser is the hard part. Getting the tree construction right took some thinking, especially the parentheses depth thing.

---

If i have some free time i think that i can expand this simple AST with variable assignment, loops and iteratios, function definitions etc.
