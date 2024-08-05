show-license : show-license.go
	go build $?

make-license : make-license.go
	go build $?

get-jwt : get-jwt.go
	go build $?

check-php : check-php.go
	go build $?

clean :
	rm -rf show-license make-license get-jwt check-php
