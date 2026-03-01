# Jobtrack

Jobtrack is a simple command-line income tracker. Start the timer when you begin working, check your total at any time, and save the results to a text file.

## Run

```bash
go run jobtrack.go
```

## Commands

- `help` shows the available commands
- `start` begins tracking
- `pause` stops tracking without exiting
- `status` prints the current total
- `save` writes the total to `track.txt` and exits

## Adjust the earning rate

Update the `time.Sleep` duration inside the `addcash()` function to change how quickly earnings accumulate.
