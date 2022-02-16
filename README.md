# Process Manager
A process manager and process monitor.

## What is this project ?
Process Manager allows you to simulate the CPU process
managing. It allows you to monitor the processes and perform
operations on them.

## How to use this project ?
After cloning into the project:
```shell
git clone https://github.com/amirhnajafiz/PM.git
```

Go to the root directory and run the project with the following command:
```shell
cd ./PM
make start
```

By default, you can create upto 10 processes, but you can customize the limit
by setting the limit variable with following command:
```shell
make start limit=[number]
```

<p align="center">
    <img src="./assets/1.png" />
</p>

The following commands are supported for this application:
### New Process
```shell
new --task [task name] --delay [task delay] --burst [overhead time of process]
```

<p align="center">
    <img src="./assets/2.png" />
</p>

### Kill Process
```shell
kill --id [task PID]
```

<p align="center">
    <img src="./assets/5.png" />
</p>

### Stop a process
```shell
pause --id [task PID]
```

<p align="center">
    <img src="./assets/4.png" />
</p>

### Start a process
```shell
run --id [task PID]
```

<p align="center">
    <img src="./assets/3.png" />
</p>

### Exit
```shell
terminate
```

## Dependencies
- go 1.17
- go-pretty v6
- go-runewidth v0.0.13
- uniseg v0.2.0