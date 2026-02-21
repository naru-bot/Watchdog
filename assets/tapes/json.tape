# JSON Output for AI Agents
Output assets/json.gif

Set FontSize 18
Set Width 1200
Set Height 550
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
Type "upp check --json | jq ."
Enter
Sleep 4s

Type "upp status --json | jq '.[].uptime_percent'"
Enter
Sleep 3s

Sleep 1s
