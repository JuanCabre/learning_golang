shared:
	rm -f libkeyboard.so
	rm -f ../TestApp/libkeyboard.so
	gcc -c -fPIC keyboard.c
	gcc -shared -W -lncurses -o libkeyboard.so keyboard.o
	rm -f keyboard.o
	cp libkeyboard.so ../TestApp/libkeyboard.so
