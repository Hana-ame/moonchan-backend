# mooonchan's backend

- [ ] ActivePub api
  - [x] webfinger
  - [ ] httpsign
    - [x] inbound
      - [ ] did it check the body's hash?
      - [ ] db: make a cache
    - [ ] outbound
    - [ ] hint: write some test cases...? or connect to a fedi server
    - [ ] actor, should be a special user(misskey) or actor(mastodon)?
  - [ ] get user info
  - [ ] follow
    - [ ] check if follow succeeded
  - [ ] unfollow
    - [ ] undo
  - [ ] post message
    - [ ] delete message
    - [ ] edit message
  - [ ] notifycation
  - [ ] block
  - [ ] vote
    - [ ] what it does?
  - [ ] outbound
    - [ ] TIMEOUT!!!
    - [ ] httpsign
    - [ ] get user
    - [ ] post messages
      - [ ] post
      - [ ] follow
      - [ ] block?
- [ ] user interface api
  - [ ] account
    - [ ] create account
    - [ ] login
    - [ ] annonymous?  
  - [ ] profile
  - [ ] timeline
  - [ ] notifications
  - [ ] ...
  - [ ] posts
- [ ] db
  - [ ] gorm
    - go get -u gorm.io/gorm
    - ...
  - [ ] user
    - [ ] create a user
      - [ ] username, passwd, email, private key, user_object.

## network?

- [x] cloudflare tunnel
  - [x] does it cost money?
    - no
  - [ ] request by proxy.

frontend is in another repository
