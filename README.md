# Performance Measuring
Simple performance measuring tool.

## Setup
- Install the logging mod found in the [log-producer](log-producer) directory and add it to your playset
- Copy [log-collector.exe](log-collector/log-collector.exe) (or [log-collector](log-collector/log-collector) on linux) into your Victoria 3 logs directory (`<Documents>/Paradox Interactive/Victoria 3/logs/`)
- Start the game (Main Menu)
- Start the `log-collector.exe`
- Start an observer game
- Results can be found in the `performance` directory which was created in the Victoria 3 logs directory (`<Documents>/Paradox Interactive/Victoria 3/logs/performance/`)

## How To Build
This repository contains a prebuilt binary already, but if you want to build it yourself then you need to follow these steps:

First download and install the Go SDK:
- https://go.dev/doc/install

Next, open the [log-collector](log-collector) folder in a terminal (e.g. cmd) and run the following command:
```
go build
```

That is it. There should be an executable in the project folder now.