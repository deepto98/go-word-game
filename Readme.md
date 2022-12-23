## Learnings:
1. Setup multi root vs code workspaces. solves go tools issues <br>
https://code.visualstudio.com/docs/editor/multi-root-workspaces
2. -u flag in `go get -u github.com/gin-gonic/gin`:<br>
 The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.
3. Gin server graceful shutdown  <br>
https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
4. Ctrl+C - SIGINT, Ctrl + Z - SIGSTOP <br>
Read about different signals : https://qr.ae/pr13KC
5. For hot reload : https://github.com/cespare/reflex
6. To route a local url to server : edit `/etc/hosts` - add line `127.0.0.1 word-game.test`
7. Traefik : Reverse Proxy <br>Learn: https://www.youtube.com/watch?v=wLrmmh1eI94, https://www.youtube.com/watch?v=C6IL8tjwC5E
8. Grouping Routes: Useful for creating a parent route and sub routes, and define middleware once, for all the subroutes<br> https://chenyitian.gitbooks.io/gin-web-framework/content/docs/13.html
9. Architecture for go-account-app : https://github.com/bxcodec/go-clean-arch<br>
More about layered architecture: https://www.youtube.com/watch?v=V1a8msZ025k, https://blog.logrocket.com/flat-structure-vs-layered-architecture-structuring-your-go-app/,
https://www.youtube.com/watch?v=oZsUhHdC7H8 
Hex arch : https://www.youtube.com/watch?v=MpFog2kZsHk&t=2s

10. Examining errors with `errors.is (compares an error to a value)` and `errors.as (function tests whether an error is a specific type and if so, converts it to that type)`<br>
More : https://go.dev/blog/go1.13-errors, https://gosamples.dev/check-error-type/,<br> Great article about custom error handling : https://earthly.dev/blog/golang-errors/ 
11.  Context in Go : carries a deadline, a cancellation signal, and other values across API boundaries.<br>
Learn: https://www.youtube.com/watch?v=h2RdcrMLQAo, https://www.youtube.com/watch?v=kaZOXRqFPCw&t=505s, https://www.youtube.com/watch?v=mgJMIZsWfB4
12. Type assertion in Go : if a variable a is of type interface{} or any, and its value is expected to be of a certain type, say int, that type can be asserted using `x,ok := a.(int)` <br>
More : https://go.dev/ref/spec#Type_assertions. https://stackoverflow.com/a/24493004
13. `gin.H` - nothing but a shortcut for ` map[string]interface{}`
14. Testify(go get github.com/stretchr/testify) : Library for testing
15. Creating objects with value : `a := Car{} (a is of type Car, and &a has to be explicitly used to get reference)`<br>
    Creating objecta with reference : `a := new(Car) (a is of type *Car, can be directly passed to functions)`<br>
    More : https://medium.com/technofunnel/golang-object-oriented-programming-f2e6448b8f24
16. Running tests : `go test -v ./handler/tests/`. Files should end with `_test`. <br>
    Run all tests within subfolders: `go test -v ./...` Basics of tests: https://go.dev/doc/tutorial/add-a-test
17. For JSON struct tags go-playground validator is used in Gin, so the same rules for `validator:` can be added to `binding:`
18. Debate on storing access tokens: https://coolgk.medium.com/localstorage-vs-cookie-for-jwt-access-token-war-in-short-943fb23239ca<br>
    Different authentication methods : 
    1.     https://medium.com/@vivekmadurai/different-ways-to-authenticate-a-web-application-e8f3875c254a
    2.     https://blog.risingstack.com/web-authentication-methods-explained/
    3.     https://stackoverflow.com/questions/58339005/what-is-the-most-common-way-to-authenticate-a-modern-web-app
    4.     https://www.wallarm.com/what/oauth-vs-jwt-detailed-comparison
    5.     https://anil-pace.medium.com/json-web-tokens-vs-oauth-2-0-85dd0b32057`d`
19. Salting for passwords : https://www.pingidentity.com/en/resources/blog/post/encryption-vs-hashing-vs-salting.html<br>
Go example : https://go.dev/play/p/_Aw6WeWC42I, https://pkg.go.dev/golang.org/x/crypto/scrypt
20. How JWTs work : https://jwt.io/introduction