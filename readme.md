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
  - [ ] outbound
    - [ ] profile
    - [ ] timeline
    - [ ] notifications
    - [ ] ...
  - [ ] inbound

## network?
- [x] cloudflare tunnel
  - [ ] does it cost money?
  - [ ] request by proxy.

frontend is in another repository
