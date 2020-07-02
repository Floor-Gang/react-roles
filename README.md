# Role reaction
Role reactions - dynamicly create a "react to get role" message.

[Trello card - claimed by @Elian0213](https://trello.com/c/RxG3znVR)

# Usage
The commands to use this service are as follows:

```
# Add a role-reaction to a message
.[syntax command] add (channel*) [message id] | [emoij] [role]

# Remove role-reaction from a message
.[syntax command] remove (channel*) [message id] | [emoij]
```

# Database layout

This will make the rules as dynamic as possible.

| id | channel_id         | message_id         | reaction           | role               | created_at |
|----|--------------------|--------------------|--------------------|--------------------|------------|
| 1  | 727571348297089086 | 727571348297089086 | 728026421494022266 | 725544860303622164 | 07-02-2020 |
| 2  | 727571365531484291 | 727571348297089086 | 728026421494022266 | 725544860303622164 | 07-02-2020 |
| 3  | 727571366567739473 | 727571348297089086 | 728026421494022266 | 725544860303622164 | 07-02-2020 |

# Setup
Download golang if you haven't already at https://golang.org/dl/ after that install the packages 

```
$ go get
$ go build 
```