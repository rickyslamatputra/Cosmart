# FOR INTERVIEW COSMART PURPOSE #
# SIMPLE LIBRARY FUNCTION #

# Available Services #
- Get list books based on genre/subject 
- Submit book pickup schedule
- Get list pickup schedule (additional since as librarian need to know list book pickup)

# How To Run #
after you extract the repository zip, you might need 
`go mod tidy`

run in terminal in repository 
`go run .\cmd\main.go`

if you're using vscode : 
you can open file `api.http` and klik `Send Request` in top each request

if you're using postman :
you can open `Cosmart Interview.postman_collection.json` in postman

or you can try open from web browser and hit `http://localhost:8080`
- `/api/books?subject=love` GET (subject can be changed ex: war,comedy,etc)
- `/api/schedule-pickup` POST
- `/api/schedule-pickup` GET



# More details 
the idea structure in this project is 
Handler - Service - Repository 

Handler manage all request response and validate the input
Service contains the logic ,if need data can go to repository / if need to hit services then do in service
Repository ideally contains the processing to databases or cache 