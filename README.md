## shebl - SHEll By Line

Runs shell scripts line by line for live demos and trainings. Unit tests for
your talk's live demo.

The first rule of live demos is: Never do a live demo. If you must then you are
better off scripting it.

###Install
```bash
go get github.com/eosrei/shebl
```

###Demo
```bash
# Run the script normally, the "unit test"
./tests/basic.sh
# Run the script line by line, the "line by line live demo"
shebl tests/basic.sh
```

###Use
Press enter to display a command, enter again to run it, repeat.

###Todo:
Lots! This is my first pass on the idea.

1. Use: https://github.com/kr/pty to control a full interactive shell
2. Support for Python, Ruby, PHP, etc.
3. Use: https://github.com/nsf/termbox-go for a better UI
4. Scroll through the input to create "breakpoints"
5. Optionally display line numbers
6. Allow user to enter/edit commands
7. Watch variables
