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
9. Architecture for go-account-app : https://github.com/bxcodec/go-clean-arch
10. Examining errors with `errors.is (compares an error to a value)` and `errors.as (function tests whether an error is a specific type and if so, converts it to that type)`<br>
More : https://go.dev/blog/go1.13-errors, https://gosamples.dev/check-error-type/,<br> Great article about custom error handling : https://earthly.dev/blog/golang-errors/ 

