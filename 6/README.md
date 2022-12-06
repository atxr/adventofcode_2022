# Day 6

I used aarch64 (arm) asm for the day 6.

To emulate the arm env, I used `qemu`. Check [this link](https://azeria-labs.com/arm-on-x86-qemu-user/) if you want to setup your lab.
To compile the code and execute it:
```bash
aarch64-linux-gnu-as sol.s -o sol.o && aarch64-linux-gnu-ld sol.o -o sol
qemu-aarch64 ./sol | wc
```

You can modify the line 47 to modify the size of the packet.
