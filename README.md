# advent-of-code-2022

This project is inspired by https://github.com/timkelleher/aocgen

## Setup

```shell
cp .env.example .env
```

Copy your cookie from adventofcode.com and paste it in your `.env` as the `AOC_SESSION` environment variable.

## Generate a new day

```shell
./aocgen gen -y 2022 -d <day>
```

## Run a day

```shell
./aocgen run -y 2022 -d <day>
```

## List input of a certain day

```shell
./aocgen input -y 2022 -d <day>
```