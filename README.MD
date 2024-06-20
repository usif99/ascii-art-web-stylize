# Ascii Art Web

## Description
Ascii Art Web is a web application designed to enable users to craft ASCII art employing various banners (STANDARD, SHADOW, and THINKERTOY). Implemented in Go, the server furnishes a straightforward and interactive graphical user interface (GUI) for generating and showcasing ASCII art.

# HOW TO RUN
To run the program, execute

    go run .

    in the terminal. Then, open a web browser and enter [localhost:8080](localhost:8080) in the URL bar to access the web page, (or ctrl + click the link presented when running).
To interrupt the server Press CTRL + C

# inst
1. Enter your text: Input your text (up to 2000 characters) in the box on the left side of the page.
2. Choose font type: Select the font type (standard, shadow, or thinkertoy) for displaying your text.
3. Choose color: Pick a color for your text.
4. Submit: Click the "submit" button to see the graphical representation of your text.

# Implementation Details
**Algorithm:** The program uses HTTP Get and Post methods for handling user input and displaying the results.

Enjoy creating ASCII art from your text!

## Authors
- Tayma (tfakhar)
- Yusuf Yaser (ymaidan)
- Maryam (mazaman)

## Usage: How to Run
1. Clone the repository:
   git clone (repository_url)
2. Navigate to the repository
   cd ascii-art-web
3. Run the code
   go run main.go
   The website will run right away
   


## Implementation Details: Algorithm
The implementation of server-side logic for generating ASCII art is implemented using Go. Below is a broad overview of the algorithm employed:

1. HTTP Endpoints:
   * GET /:
      - Renders the main page containing a form for text input and banner selection.
   * POST /ascii-art:
      - Manages form submission, validates input, and generates ASCII art based on the chosen banner.

2. Text Processing:
   - Manages newline characters to ensure proper line breaks in the ASCII art.

3. Font Selection:
   - Accesses the chosen font file (e.g., Fonts/standard.txt) to fetch ASCII characters.

4. ASCII Art Generation:
   - Matches ASCII characters to their respective lines in the selected font.
   - Generates the final ASCII art by amalgamating lines for each character in the input text.

5. Error Handling:
   - Provides appropriate HTTP status codes for various scenarios (OK, Not Found, Bad Request, Internal Server Error).