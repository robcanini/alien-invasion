![ufo](assets/images/ufo_640.png)

# Alien invasion

Mad aliens are about to invade the planet X and this application simulates the invasion.

## Build

### Prerequisites

Before proceeding with the build step, ensure to have installed Go in your system.
The project has been developed and tested with: `go1.18 linux/amd64`.

### Build

To build the executables:
```bash
$ make clean build
```
This command generates three executable inside `bin` folder for Linux, Darwin and Win platform.

### Run
The program requires two cli arguments for launching the invasion simulation:
- `filePath`: the file containing the grid configuration (<i>see `examples` folder</i>)
- `aliens`: the number of aliens which are going to be spawned in the grid

So an execution example would be:
```bash
$ ./bin/invade-linux --filePath examples/data.db --aliens 5
```

## Simulation result

If you have specified valid input data, the program should print out the grid in the same format as the input file.
You can see that the grid is missing of one or more city, which has been destroyed by two alien collisions.

The program prints out also events log about cities destruction and the involved aliens.

For instance, given grid below as input file:
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```
if Bar were destroyed the grid would now be something like:
```
Foo west=Baz south=Qu-ux
```

## License

Alien invasion is released under [MIT License](LICENSE.md).
