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





