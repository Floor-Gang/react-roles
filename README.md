# Role reaction
Role reactions - dynamicly create a "react to get role" message.

[Trello card - claimed by @Elian0213](https://trello.com/c/RxG3znVR)

# Usage
The commands to use this service are as follows:

```
# Add a role-reaction to a message
.[syntax command] add (channel*) [message id] [emoij] [role]

# Remove role-reaction from a message
.[syntax command] remove (channel*) [message id] [emoij]
```

# Database layout

This will make the rules as dynamic as possible.

| id | guild_id | channel_id | message_id | reaction | role |
|----|----------|------------|------------|----------|------|
| 1  | guild    | channel    | message    | emoij    | role |

# Setup
Download golang if you haven't already at https://golang.org/dl/ after that install the packages 

```
$ go get
$ go build 
```