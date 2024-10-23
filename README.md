# grop

`grop` is a small utility to search for a query string in text files.  It's essentially a less feature-complete `grep`.  It was created to continue learning Go.  

## Features
- Searches each file in the target destination in a separate `goroutine`
- Highlights the queried word in the output

## Usage
Run the executable with 2 flags: `d` (folder to search), and `q` (query)

## Example
<img width="1421" alt="Screenshot 2024-10-23 at 3 19 49 PM" src="https://github.com/user-attachments/assets/4d648f27-1212-4b3f-93da-460ba51b3d5b">

## Disclaimer
This project is for educational purposes and not intended to be used IRL. There are a variety of known bugs I probably wont fix
