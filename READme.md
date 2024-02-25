# FileOps: HTTP server & a command line client

A simple file store service (HTTP file server and a command line client)
that stores plain-text files in any programming language(Golang has been used here).

- The server would receive requests from clients to store, update, delete files, and perform operations on files stored in the server.

- The CLI acts as a client which can add, remove, edit and list files on the server.

- Add command should fail if the file already exists in the server.

- Client won't send a file if its content is already in the server(in some other file on the server) -> add a new file without having to send its contents



A CLI interface to an HTTP file server with multiple available commands. This CLI acts as a client which can add, remove, edit and list files on the server.

## Usage


- Clone the repo
- Run `go build`
- Use `./FileOps server`to run the server

### If you are using docker:

- `docker build -t fileops .`
- `docker run -e PORT=8080 -p 8080:8080 fileops`

### While developing, we can run

```
go run main.go
```

- Running `go run main.go`, you can directly use the CLI of the app
```
go run main.go server
```

- Running `go run main.go server` would run the main HTTP file server, showing if its working fine.

## Commands

#### `add` command - uploads the file(s) mentioned in the arguments to the server. 
```
./fileops add a.txt b.txt
```

#### `ls` command will list all the files present on the server

```
./fileops ls
```

#### `rm` command - deletes the file mentioned in the argument. 

```
./fileops rm a.txt
```

#### `update` command - updates contents of a file in server with the local file or creates a new file in server if it is absent

```
./fileops update a.txt
```

#### `wc` command - lists the count of all the words present in all the files

```
./fileops wc
```

#### `freq-words` command - lists top 10 most frequently used words by default. 
##### You can change the order and number of responses

```
./fileops freq-words [--limit|-n 10] [--order=dsc|asc]