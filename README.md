# Segway G3 Max: VCU Modification (Speed Limits and Region Change)

## ⚠️ Warning
**The scooter's warranty will be voided!**  
You perform these actions at your own risk. Possible consequences:
- VCU controller failure
- Unstable scooter performance

## 📋 Required Components
1. **Hardware**:
   - ST-Link v2 programmer
   - DUPON cables
   - Soldering iron (for TP13V3 contact)

2. **Software**:
   - [STLink Driver]
   - [Modification Tool](https://github.com/Sharkboy-j/ninebot-g3-max-vcu-speed-hack/releases/latest)

## 🛠 Preparation
1. Install the driver (`dpinst_amd64.exe`)
2. Extract the software archive to a folder
3. Key files:
   - `dump_memory.bat` - memory dump tool
   - `fix_vcu.exe` - modification tool
   - `flash_memory_patched.bat` - firmware flash tool

## 🔌 ST-Link Connection
Before connecting, carefully check your ST-Link pinout as it may vary depending on your model.  
![Pinout](https://github.com/Sharkboy-j/ninebot-g3-max-vcu-speed-hack/raw/main/img/stlink.png)

**Important!** Disconnect the VCU from the scooter before starting.

| No. | ST-Link | VCU       |
|:---:|:-------:|:----------|
| 1 | SWDIO   | DIO       |
| 2 | SWCLK   | CLK       |
| 3 | GND     | GND       |
| 4 | 3.3V    | TP13V3*   |
| 5 | GND**   | C45 (temporary) |

\*Soldering required (fragile contact!)  
\*\*Used only for shorting to C45  

![Contacts](https://github.com/Sharkboy-j/ninebot-g3-max-vcu-speed-hack/raw/main/img/pins.png)

## 🔄 Modification Process

### 1. Creating a Dump
1. Short pin 5 (GND) to C45
2. Run `dump_memory.bat`
3. When the line `oocd\scripts/mem_helper.tcl", line 37` appears, disconnect the contact
4. Verify the `MEMORY_G3.bin` file (128 KB) exists

### 2. Parameter Modification
Run `fix_vcu.exe` and enter:  
Speed (1-99): <u>**99**</u>  
>the tool accepts 1-255, but values above 99 don't work correctly  

Change region? (Y/N): <u>**Y**</u>  

### 3. Flashing
1. Short GND to C45 again
2. Run `flash_memory_patched.bat`
3. Disconnect when `oocd\scripts/mem_helper.tcl", line 37` appears
5. The transfer is complete when the console shows:  
`wrote 131072 bytes from file MEMORY_G3.bin.patched.bin to flash bank 0 at offset 0x00000000 in 3.471415s (36.873 KiB/s)`

### 4. Post-Flash Setup
1. Unbind the scooter in the app
2. Reconnect the VCU to the scooter (disconnect STLink first)
3. Rebind and activate **both** toggles:  
![Settings](https://github.com/Sharkboy-j/ninebot-g3-max-vcu-speed-hack/raw/main/img/ninebotsettings1.png)

### 5. Final Step
Repeat steps 1-3 to activate changes.  
**Important!** Do not touch the sliders afterward!

## ❌ Possible Errors
- `open failed` → Check STLink connection to PC  
- `init mode failed` → VCU connection error (check contacts)  
