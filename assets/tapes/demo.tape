# Hero Demo: add → check → status
Output assets/demo.gif

Set FontSize 18
Set Width 1200
Set Height 500
Set LineHeight 1.4
Set Padding 20
Set Theme "Catppuccin Mocha"
Set TypingSpeed 40ms

Sleep 500ms
Type "upp add https://httpbin.org/get --name 'API Monitor'"
Enter
Sleep 2.5s

Type "upp add https://mosets.com --name 'Mosets'"
Enter
Sleep 2s

Type "upp check"
Enter
Sleep 4s

Type "upp status"
Enter
Sleep 4s

Sleep 1s
