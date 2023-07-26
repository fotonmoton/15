# 15 puzzle
Simple 15 puzzle implementation with naive solver.

## Docker

```
docker run --rm -it $(docker build -q .) game
```
or 
```
docker run --rm -it $(docker build -q .) solve
```

## From source

```
cd cli && go run . solve
```
or
```
cd cli && go run . game
```
