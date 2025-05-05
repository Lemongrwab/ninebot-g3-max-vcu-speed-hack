@echo off
echo Starting memory dump...

openocd -s oocd\scripts ^
    -f oocd\scripts\interface\stlink.cfg ^
    -f oocd\scripts\target\at32.cfg ^
    -c "init" ^
	-c "reset halt" ^
	-c "flash probe 0" ^
	-c "stm32f1x unlock 0" ^
	-c "reset halt" ^
    -c "dump_image MEMORY_G3.bin 0x08000000 0x0803FFFF" ^
    -c "exit"
