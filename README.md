# GROUPIE-TRACKER-search-bar
* `kkanat`

#### Groupie-Tracker-Search-Bar consists on making your site

    Enhanced in terms of visual appeal, interactivity, and user-friendliness. 
    It provides more comprehensive feedback to users. The system utilizes a 
    JSON file containing the necessary data for display.
    
   
    Furthermore, improvements have been made to the search bar, allowing users to input strings.
    The search bar now suggests strings based on the JSON file, including options such as artist/band 
    name, members, locations, first album date, and creation date.

    If the entered substring matches any of the parameters in the search bar, the system dynamically 
    displays a list of possible suggestions for the user. Additionally, the main page dynamically 
    showcases only those bands and/or artists that are part of the suggested list, providing a more 
    tailored and focused user experience.


#### Description:

* The text/template package was employed for handling and responding to GET requests.
* GET /: This route sends an HTML response, serving as the main page.
* The system utilizes a JSON file containing the data necessary for displaying content on the pages.
* An external CSS file has been linked to ensure an aesthetically pleasing design.
* To implement dynamic changes to the bands and the list of suggestions, JavaScript (JS) was utilized on the frontend.



## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@git.01.alem.school:kkanat/groupie-tracker-search-bar.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd groupie-tracker-search-bar
```
Run a Server:
```CMD/Terminal 
go run main.go
```

Follow the link on the terminal:
```CMD/Terminal 
Starting server got testing... http://127.0.0.1:7000 
```

you can play with the page by choosing the music band's image and so on.
As well as seaching the strings



## HTTP status codes
* OK (200), if everything went without errors.
* Not Found, if the wrong URL was provided.
* Bad Request, for incorrect requests. for exaple, the id is out of range.
* Internal Server Error, for unhandled errors.
* If string is not a substring of any items which were mentioned in the first paragraph, 
then the page will be displaying "NOT FOUND"


## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.
