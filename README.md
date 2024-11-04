```
                      _   _                           _                                _      
                     (_) (_)                         | |                              | |     
  __ _   ___    ___   _   _   ______    __ _   _ __  | |_   ______  __      __   ___  | |__   
 / _` | / __|  / __| | | | | |______|  / _` | | '__| | __| |______| \ \ /\ / /  / _ \ | '_ \  
| (_| | \__ \ | (__  | | | |          | (_| | | |    \ |_            \ V  V /  |  __/ | |_) | 
 \__,_| |___/  \___| |_| |_|           \__,_| |_|     \__|            \_/\_/    \___| |_.__/  
                                                                                              
                                                                                              
```

# ascii-art-web
A web version of the ascii-artwith GUI for easier usability.
The website is host on the local machine on port 8080.
Only `localhost:8080/` and `localhost:8080/ascii-art`are valid URL.
When the user keys in the valid URL, a GET request is sent to the server.
And a webpage with text field with selection buttons for the user to define the parameters will be display.
When the user clicks on the `Generate ASCII Art` button, a POST request is sent to the server.
The user defined parameters will be grab and passed to the GenArt() function to generate ASCII art.

```mermaid
flowchart TB
    subgraph server [SERVER]
        start{"go run main.go
        checkRequired()
        define static file server route
        http.ListenAndServe(8080)"}
        homeFn("homeHandler()")
        rForm("getFormInputs()")
        genArt("GenArt()")
        template("getTemplate()")
        handlePost("handlePost()")
    end

    subgraph client [CLIENT Browser Application]
        invalidURL["404 not found"]
        inValidRq["405 method not allowed"]
        tempErr["404 not found "]
        rFormErr["400 bad request"]
        missing["500 internal server error"]
        index(("display index.html
        200 OK"))
        pressBtn((Press 'Generate ASCII Art'))
    end

    server<-->|HTTP Request and Response| client;
   
    start--
    Valid URL
    localhost:8080/ || 
    localhost:8080/ascii-art
    -->homeFn;
    start--Invalid URL
    redirect to /error-->invalidURL;
    
    homeFn--GET request-->
    template-->
    index--User input-->
    pressBtn;
    
    homeFn--POST request-->handlePost-->
    rForm-->
    genArt--Art/Error generaterd-->
    template;

    pressBtn--
    Post request to
    localHost:8080/ascii-art
    -->homeFn;
    
    homeFn--Invalid Request
    redirect to /error-->inValidRq;
    rForm--Uable to grab data
    redirect to /error-->rFormErr;
    genArt--Missing files
    redirect to /error-->missing;
    template--No file or permission
    redirect to /error-->tempErr;
```
## Requirement
- start and run a server (DONE)
- web GUI for ascii-art (DONE)
- Must allow the use of the 3 banners (DONE)
- Implement HTTP endpoints: (DONE)
    - GET "/": go templates
    - POST "/ascii-art": use form to make post request
- Display result of POST in home page. (DONE)
- main page must have: (DONE)
    - text input
    - radio buttons
    - button
- HTTP status code - make redirecting work properly
    - 200 OK
    - 404 Not found
    - 405 Bad request
    - 500 Inernal Server Error
- Include in README.md
    - descriptons
    - Authors
    - usage
    - implementation details

## Tasks
- main.go
    - P1 - redirecting not working
    - P2 - sanitize input. get rid of starting and ending newlines
    - P3 - reorganize code for imrpove readability
- P4 - Improve README.md
    - descriptons
    - Authors
    - usage
    - implementation details
- Design index.html and css
    - more functionality?

- Learn HTTP protocol, handlers and pattern
    - focus on HTTP status code 
- Learn server and client
- Implement more ASCII art functionality?
    - ANSI color won't work

## Optionals
- export output
- stylize with css
- dockerize