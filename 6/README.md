# Day 6

I tried the day 6 in aarch64 asm.
I manage to code a solution for the first start, but not for the 2nd.
I added a short python script to get the second star :)

To run the arm program, I used qemu. Check [this link](https://azeria-labs.com/arm-on-x86-qemu-user/) if you want to setup your lab.
To compile the code and get the solution for the first star:
```bash
aarch64-linux-gnu-as sol.s -o sol.o && aarch64-linux-gnu-ld sol.o -o sol
qemu-aarch64 ./sol | wc
```
