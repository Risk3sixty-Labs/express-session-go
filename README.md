# express-session-go

Functions and HTTP middleware to check validity and retrieve session information 
that was populated from [express-session](https://github.com/expressjs/session).

## Why

At [risk3sixty](https://risk3sixty.com/) our primary web application authenticates and
persists session data using [Passport](https://github.com/jaredhanson/passport) and [express-session](https://github.com/expressjs/session).
We are starting to build APIs and services using Go, so to prevent from
rearchitecting the way we authenticate from day one we needed a way to
authenticate users and populate session data in our Go services.

## Usage

See `/examples` folder

## TODO

- Create example session store retrieval of session that can be used as `sessionParser` replacement

## Special Thanks

Special thanks to [Chris Tomich](https://github.com/chris-tomich) and his [blog post](https://mymemorysucks.wordpress.com/2016/05/26/sharing-an-expressjs-connect-passportjs-session-with-golang-part-1/)
for helping get started on this project.