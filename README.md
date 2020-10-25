# go-bot
- Simple twitch bot made with Go - Study purpose
- I'm following the official [Twitch Chatbot Docs](https://dev.twitch.tv/docs/irc/guide)
### Packages
- [joho/godotenv](https://github.com/joho/godotenv)
	- Parse `.env` file
### Todo:
1. [x] Connect
	- [x] Connection via Golang's package [net](https://golang.org/pkg/net/) 
999. [ ] 100% Cli
	- This bot should be __100%__ configurable in the terminal
	- Maybe a JSON to store info:
	- 	```json
		{
			"!discord": {
				"timeout": 600, // 5 min,
				"access": 0, // 0 - everyone, 1 - VIP/Sub, 2 - Mod, 3 - Streamer,
				"message": "My discord channel invitation link is ..."
			},
			"!hug{:userTo}": {
				"timeout": 120, // 2 min,
				"access": 0,
				"message": "/me $(userFrom) hugged $(userTo)" 
			},
			"!count": {
				"timeout": 30, // 30 sec,
				"access": 1,
				"variable": {
					"counter": 124
				},
				"message": "This command have been used $(variable.counter)x times!"
			}
		}
		```
	- Or something like this
