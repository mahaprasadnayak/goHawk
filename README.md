# goHawk
Search with precision, soar with speed

## Features
```sh
  1. Basic Text Search
  2. Recursive Search
  3. Case-Insensitive Search
  4. Display Line Numbers
  5. Count Matches
  6. Limit Output
```2.201
## CMD
- For recursive search:
  ```sh
  goHawk.exe --pattern <text_to_search> --file <file_directory> -r
  ```
- For Single File Search
   ```sh
   goHawk.exe --pattern <text_to_search> --file <file_path>
   ```
## Example-1
 ```sh
  goHawk.exe --pattern email --file ./imp -r
 ```
## Example-2
 ```sh
  goHawk.exe --pattern 700 --file ./imp/test.txt
 ```

