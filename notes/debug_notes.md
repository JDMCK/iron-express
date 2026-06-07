# Delve cheatsheet

break (b) - set a breakpoint
clear <breakpoint name/id> - clear a breakpoint
continue (c) - continue
next (n) - step to next source line
step (s) - step into function
stepout (so) - step out of a function
print - print a variable
display -a <var> - add variable to display list
display -d <var id> - remove a variable by id
rewind - run backwards until breakpoint
watch -r|w <var> - set data breakpoint (watch) on a variable


# Gotchas

To break at main(), do `b main.main`.
