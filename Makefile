

diywm: main.go
	go build -o diywm main.go

run: diywm
	echo "exec ./diywm" > xinitrc
	./run.sh

clean:
	rm ./diywm xinitrc
