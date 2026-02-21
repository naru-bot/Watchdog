# Uptime Monitoring with Sparklines
Output assets/uptime.gif

Set FontSize 18
Set Width 1200
Set Height 500
Set LineHeight 1.4
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
Type "upp check"
Enter
Sleep 4s

Type "upp status"
Enter
Sleep 4s

Sleep 1s
