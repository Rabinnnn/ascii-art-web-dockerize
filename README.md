# ASCII ART WEB STYLIZE

## Description
This is a simple web-based application developed using Go, HTML, CSS and a little Javascript. It takes input text from the user then converts it to the appropriate graphical representation using the chosen ascii art banner file (or font). The user interacts with it through a graphical interface.

## Usage
1. **Clone Repository**: Clone this repository to your local machine.
   ```bash
   $ git clone https://learn.zone01kisumu.ke/git/hshikuku/ascii-art-web-stylize.git
   ```
2. **Install go**: Ensure you have Go installed on your machine.(Go to official Go website and follow instruction(based on your OS))
3. **Run Program**: To get the server up and running, navigate to the directory where you've cloned the repo then use the following command:
    ```bash
    $ cd /ascii-art-web-stylize
    $ go run . 
    ```
    This will automatically open your browser and display the GUI.
4.  **Launch GUI**: If your browser doesn't open automatically after executing the command above, open your web browser and on the address bar type this:
    ```
    http://localhost:8080 
    ```
If the GUI doesn't open check your firewall settings and grant access to the program.

5. **Choose banner** : Select your banner of choice in the list provided.

6. **Input text**: In the text area labeled input text, type the text that you would like to be converted.
7. **Execute**: Click on the 'produce art' button. The program will execute and the results returned. If for example your input text is 'hello' and your banner of choice is 'standard' then the output will be:
    ```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
    ``` 
## Implementation Details

### GET request
- This type of request method is implemented when fetching home page template. 
- The response from the http server to the user delivers the home page which displays a form that expects user input. 
- This type of request method is also implemented when displaying the about ascii and about us pages.

### POST request
- The implementation of post method is carried out when sending form data to the server for implementation of the corresponding ascii art.
- Once the form data is parsed in the ASCIIArthandler and deemed to fulfill all the necessary conditions then the ascii art program is called to convert to its corresponding ascii art.
- This data in form of a string is then injected in a template and served back to the user in a readable format.

### Error handling
- The server uses logging functionalities as defined in package log to handle various errors. 
- In addition, the endpoints implement http methods such as http.StatusBadRequest, http.StatusNotFound, and http.StatusMethodNotAllowed that returns the appropriate status codes.

## File System
Within the repo there are 2 directories and other files.
1. **server.go**

    This is the entry point of the program. It registers the handler functions, serves the static files and also sets up the server.
     
2. **ascii-art**

    This directory contains the banner files functions used to create the graphical representation, and their test files.
-  **banners**: This sub-directory has the banner files that 
contain the ascii-art characters.
-  **art_map_creator.go**: Creates a map of ascii art characters from a chosen banner file.
-  **get_banner_file.go**: Gets the appropriate banner file that matches the choice. Example; if choice is 'standard' the function returns 'standard.txt'
-  **art_maker.go**: Converts the text and returns the output.
-  **integrity_checker.go**: Ensures that the banner files are available and in original condition.
-  **get_banner_file.go**: Retrieves the specific banner file according to the argument passed.

3. **handlers**
-  **handlers.go**: The file contains the handler functions used in the program.

4. **static**

    This directory contains static files such as images, css and js files.

5. **templates**

This contains various html templates used by the GUI.
-  **index.html**:  Displays the home page.
-  **result.html**: This is where the output is displayed.
-  **aboutus.html**:  This page contains a the name of the authors and thier github links.
-  **aboutascii.html**:  This page contains information about the concept of ascii art.
-  **400.html**: Displays http status code 400
-  **404.html**: Displays http status code 404
-  **500.html**: Displays https status code 500

6. **testCases**
  Contains the expected output used in TestHandleASCIIArt

7. **handlers_test.go**
   Tests the functionality of the handler functions.

## Authors

-   **[Hezborn Shikuku](https://github.com/Mania124/ascii-art-web)**
-   **[Rogers Ragwel](https://github.com/oragwelr/ascii-art-web)**
-   **[Rabin Otieno](https://github.com/Rabinnnn/ascii-art-web)**