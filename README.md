# advent-of-code-2022

This project uses https://github.com/timkelleher/aocgen

## Setup

```shell
cp .env.example .env
```

Copy your cookie from adventofcode.com and paste it in your `.env` as the `AOC_SESSION` environment variable.

## List input of a certain day

```shell
./aocgen input -y <year> -d <day>
```

## Generate a new day

```shell
./aocgen gen -y <year> -d <day>
```

## Run a day

```shell
./aocgen run -y <year> -d <day>
```