# CSV Splitter

This Go program splits a CSV file into multiple CSV files based on the specified number of records in each file.

## Usage

To run the program, you need to have Go installed on your system. Follow these steps:

1. Clone or download the repository.

2. Open a command prompt or terminal and navigate to the directory containing the `main.go` file.

3. Run the following command:

```
go run main.go <csv_filename> <number_of_records_in_each_file>
```

Replace `<csv_filename>` with the name of your input CSV file and `<number_of_records_in_each_file>` with the desired number of records in each split file.
For example, to split a CSV file named `data.csv` into files with 100 records each, run:

```
go run main.go data.csv 100
```

The program will create an output directory named "outputs" next to the executable file and save the split CSV files inside it.

## Makefile

A `Makefile` is included in the repository to build executables for different platforms. You can use the following commands:

- To build the Windows executable, run:

```
make build-windows
```

This will create an executable named `output_windows.exe`.

- To build the macOS executable, run:

```
make build-macos
```

This will create an executable named `output_macos`.

- To clean up and remove the generated executables, run:

```
make clean
```

## License

This program is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

Feel free to modify and use this program according to your needs. If you have any questions or suggestions, please open an issue or submit a pull request.

Enjoy!
