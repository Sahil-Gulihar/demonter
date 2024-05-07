# System Crash Attempt

## Description
This program is designed to attempt to crash your system by exhausting available disk space, consuming high CPU resources, and hogging RAM. It repeatedly performs the following actions:
- Writes large files to the disk until the disk space is fully utilized.
- Utilizes CPU resources intensively.
- Consumes a large amount of RAM.

## Usage
1. Ensure you have the Go programming language installed on your system.
2. Clone this repository or download the `main.go` file.
3. Navigate to the directory containing `main.go` in your terminal.
4. Run the following command to compile and execute the program:
    ```
    go run main.go
    ```
5. Monitor your system's disk space usage, CPU usage, and RAM usage while the program is running.

## Disclaimer
**Caution:** Running this program may lead to data loss, system instability, or other unintended consequences. It is intended for educational purposes only. Use it at your own risk.

