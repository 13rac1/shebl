## shebl - SHEll By Line

Runs shell scripts line by line for live demos and trainings. Unit tests for
your talk's live demo.

*Of course, you can fill your shell script with `read`, but this is better.*

###Install
```bash
go get github.com/eosrei/shebl
```

###Demo
```bash
shebl tests/basic.sh
# Press 'Enter' many times to step through the script
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
6. Allow user to enter their own commands
7. Watch variables
