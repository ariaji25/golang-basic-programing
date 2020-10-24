
<h3 align="center">Basic Programing with Golang</h3>

---

<p align="center"> Basic programing to learn golang, this repo will try 4 case to resolve with golang.
    <br> 
</p>

## CASE 0 : POLINDROM NUMBERS
  This case is about to find how many polindrom numbers between range of two number. The polindrom numbers is a number if we read it from left or right it
  always be the same number. Minimum value of input is 1 and maximum value of input is 10^9 Example 101, 222, etc. 
  Example.<br/>
  Input : 1 10<br/>
  Output : 9<br/>

## CASE 1 : MANGE LIBRARY 
  This case is about how to amange books in a bookshelft at a library. The books have a code, the first index in code is the book category, the second index of
  code is the book name (First Character of Book Name), and the rest of code is the book size. In the a bookshelft, the books will sorted by it categoies, and then
  books in the same categories will sorted by it size ( Descending), size of books is greater equal to 13 cm and less equal to 14 cm. At a bookshelft, just allowed 2 same books. Two books is same if the books is in the same category and has the same name. 
  <br/>
  List Of Categories :<br/>
  6 : Applied Sciences (600) <br/>
  7 : Arts (700) <br/>
  0 : General (000)<br/>
  9 : Geography, History (900)<br/>
  4 : Language (400)<br/>
  8 : Literature (800)<br/>
  1 : Philosophy, Psychology (100)<br/>
  2 : Religion (200)<br/>
  5 : Science, Math (500)<br/>
  3 : Social Sciences (300)<br/>

  Example.<br/>
  Input : 3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14<br>
  Output : 0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13

## CASE 2 : LOST NUMBER
  This case is about find a lost number from a given string that contain a sorted list of numbers without sperators. 
  The minimum numbers in list is 1 and the maximum numbers in list is 10^6. Minimum total of numbers in list is 3 and maximum total of numbers is 1000.
  <br/>
  Example.<br/>
  Input : 134<br/>
  Output : 2<br/>

## CASE 3 : CREATE WEB SERVICE
  This case is about create a web service for program of (case 0 or case 1 or case 2).<br/>
  In this case we use the case 2 as the function of web service. Available method is GET and POST
  <br/>
  command with curl "curl -X POST http://localhost:8000/?{numbers=[sorted_list_of_numbers]}"<br/>
  example : <br/>
  curl -X POST http://localhost:8000/?numbers=1235 <br/>

## CASE 4 : DOCKER SERVICE
  Create a docker service to deploy the web service to docker. First we create the Dockerfile to config the building method to create docker image for web service.
  And then create docker-compose.yml to manage the docker service. <br?>
  command : <br/>
  #### docker-compose up
  the command function is to build and deploy our docker service to docker container

## ⛏️ Built Using <a name = "built_using"></a>

- [Golang](https://www.golang.org/) - Golang
- [Docker](https://www.docker.com/) - Docker

## ✍️ Authors <a name = "authors"></a>

- [@ariaji25](https://github.com/ariaji25) - IAri Purnama Aji
