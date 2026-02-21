# TUI Dashboard Demo
Output assets/tui.gif

Set FontSize 18
Set Width 1200
Set Height 600
Set LineHeight 1.3
Set Padding 20
Set Theme "Catppuccin Mocha"
Set TypingSpeed 40ms

# Seed data first
Hide
Type "upp add https://httpbin.org/get --name 'API Monitor'"
Enter
Sleep 1s
Type "upp add https://mosets.com --name 'Mosets'"
Enter
Sleep 1s
Type "upp check"
Enter
Sleep 5s
Show

Sleep 500ms
Type "upp tui"
Enter
Sleep 3s

Down
Sleep 800ms
Down
Sleep 800ms
Up
Sleep 800ms

Enter
Sleep 2.5s

Type "c"
Sleep 3s

Escape
Sleep 1s

Type "q"
Sleep 1s
