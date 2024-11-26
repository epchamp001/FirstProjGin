# FirstProjGin
The project demonstrates the capabilities of the [Gin](https://github.com/gin-gonic/gin) framework for developing web applications on Go. This simple application allows you to manage videos using the API, as well as visualize them on a web page.

## 🛠 Functionality

- Adding videos via POST requests.
- Getting a list of videos via GET requests.
- Authorization using Basic Auth.
- Data validation when sending POST requests.
- Logging requests and writing to a file.
- Displaying videos on a web page using HTML templates.

## 🏗 Project structure

```plaintext
learn_gin/
├── cmd/                    # Package to launch the application
│   └── myproj/             # The main executable file
│       └── server.go       # The entry point to the application
├── internal/               # Application logic (separation into packages)
│   ├── controller/         # Logic for processing requests (Video Controller)
│   │   └── video-controller.go
│   ├── entity/             # Definition of entities (Video, Person)
│   │   ├── person.go
│   │   └── video.go
│   ├── middleware/         # Middleware for logging and authorization
│   │   ├── basic-auth.go
│   │   └── logger.go
│   ├── service/            # The logic of working with data (VideoService)
│   │   └── video-service.go
│   └── validators/         # Custom validators
│       └── validators.go
├── templates/              # HTML templates and CSS
│   ├── css/                # Styles for the page
│   │   └── index.css
│   ├── footer.html
│   ├── header.html
│   └── index.html
├── gin.log                 # Request log file
├── go.mod                  # List of dependencies
```

## 📦 Installation
1. Make sure that you have Go version 1.23.3 or higher installed
2. Clone the repository:
```bash
git clone https://github.com/username/learn_gin.git
cd learn_gin
```
3. Install dependencies:
```bash
go mod tidy
```
4. Запустите сервер:
```bash
go run server.go
```
## 🔑 Authorization
The application uses Basic Auth. Use the following data to access:
* **Username**: admin
* **Password**: admin

## 🚀 Using
### Adding videos (POST /api/videos)
Send a POST request with the body in JSON format:
```json
{
  "title": "Гид по технологии Gin!",
  "description": "Введение в Gin Framework для Go",
  "url": "https://player.vimeo.com/video/76979871",
  "author": {
    "firstname": "Иван",
    "lastname": "Иванов",
    "age": 30,
    "email": "ivan.ivanov@example.com"
  }
}
```
Sample response:
```json
{
  "message": "Video Input is Valid!!!"
}
```

### Getting a list of videos (GET /api/videos)
Send a GET request to get a list of videos in JSON format
Sample response:
```json
[
  {
    "title": "Гид по технологии Gin!",
    "description": "Введение в Gin Framework для Go",
    "url": "https://player.vimeo.com/video/76979871",
    "author": {
      "firstname": "Иван",
      "lastname": "Иванов",
      "age": 30,
      "email": "ivan.ivanov@example.com"
    }
  }
]
```
### Viewing videos on a page (GET /view/videos)
Go to the address http://localhost:8080/view/videos to see the added videos on the webpage.

## 🌟 Screenshots

### Sending a POST request

### Authorization

### Sending an incorrect POST request

### Getting a list of videos

### A web page with a video

## 🔍 Notes
* Logging of requests is recorded in the `git.log` file.
* Example of custom validation: the `title` field must contain an exclamation mark (`!`).






