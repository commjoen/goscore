# goscore
Simple vulnerable webserver in Go
<!-- Keep your project structure simple and flat. Especially in the beginning. No strong conventions like in the Java world. 
test-files have <lala>_test.go naming convention and are put in the same directory (package) as the sources itself
go.mod is the way to go.
Give me all call when you need more (edited) 

Roy Bos:gopherdance:  3 hours ago
What I would do and is pretty common in the Go world:
If it is a library/package, start with everything in the root of the repository.
If it is a app/api/webservice use a little bit structure like https://github.com/golang-standards/project-layout. Most important directories you will probably need: cmd , internal , pkg . If you have multiple commands, you will directly know why this structure is useful :wink:-->