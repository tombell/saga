# Saga

**This project has been discontinued**

API server for track information from [Serato DJ Pro][serato].

[serato]: https://serato.com

## Installation

Coming soon.

## Usage

Run `saga` with either the `--sesion-file` or `--session-dir` flags. The
`--session-file` can point to an existing session file, or `--session-dir` can
specify the directory Saga should watch for a new session file when you start
Serato. This will start the websocket API server on `localhost:8080`.

You can write a front end using JavaScript to connect to the websocket server,
and listen for messages that will send the track history data.
