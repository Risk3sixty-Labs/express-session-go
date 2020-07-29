# express-session-go

Functions and HTTP middleware to check validity and retrieve session information that was populated from [express-session](https://github.com/expressjs/session).

## Why

If you use [express-session](https://github.com/expressjs/session) in a NodeJS/express app to store session data, it's fine and dandy if all your APIs and services live in the same express app with all requests flowing through the express-session middleware. With this package, you can also build APIs and services in go and implement a similar HTTP middleware to retrieve and use session info as is being used in your express app.

## Usage

See `/examples` folder

## Stores

- [memory store](https://github.com/whatl3y/express-session-go/blob/master/store/memory.go) is the default store, same as [express-session](https://github.com/expressjs/session) (note: should [not be used](https://github.com/expressjs/session#sessionoptions) in production)
- [redis store](https://github.com/Risk3sixty-Labs/express-redis), same implementation as [connect-redis](https://github.com/tj/connect-redis) in NodeJS

## TODO

- Create example session store retrieval of session that can be used as `sessionParser` replacement

## Special Thanks

Special thanks to [Chris Tomich](https://github.com/chris-tomich) and his [blog post](https://mymemorysucks.wordpress.com/2016/05/26/sharing-an-expressjs-connect-passportjs-session-with-golang-part-1/)
for helping get started on this project.