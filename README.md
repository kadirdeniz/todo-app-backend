# Todo Application Backend

This repository is a todo app sample with using Golang

## Table Of Contents
* [General info](#general-info)
* [What is Test Driven Development](#what-is-test-driven-development)
* [Clone the project](#clone-the-project)
* [Setup](#setup)
* [Test](#test)

## General info
This project is a basic todo application created with ```Golang```, in this project i try to follow ```Acceptance Test Driven Development``` short for ```ATDD```. ATDD is like combination of BDD and TDD. In this project there is 1 feature which is ```Create Todo```, you can find it feature in below.
```
Given Empty ToDo list
When I write "buy some milk" to <text box> and click to <add button>
Then I should see "buy some milk" item in ToDo list
``` 
In the application,```Pact``` were used together as contract testing. In the contract test, application should answer requests that consumer create correctly, so first testing case is contract test.

Second type of test is ```Integration Test```. In integration test i try to test handler and service methods. While testing in the handler and service layers, most things were mocked up to create isolated tests, resulting in much more accurate test results.

## What is Test Driven Development
Test driven development was first introduced by Kent Beck as a software development method. In the simplest terms, we can think of tdd as Test -> Code -> Refactor, the author thinks that tests should be written before the code and supports this with very strong ideas. According to the book, when we add a new feature to the application, we need to look at the application from the outside before starting the code directly, it says that what kind of extreme cases are, what we can encounter and finally we need to decide clearly what we want to do, then we need to test these thoughts and then put them in the code. He says it's better for us to start. When we write the test before the code phase, when we come to the starting point of the code, we will know very well what we want to do, and with this, we will have a "todo list" of tests, we can focus on our development very comfortably without stress, ensuring that the application works successfully. In the book, he talks about a lot of things about how the written tests should be, and I would like to briefly touch on them. The tests are expected to be done in a completely isolated manner, they should be independent from each other as much as possible because when an error occurs, we want to instantly detect where the error is, we can do this with isolated tests, I think that mock is very important because we have to mock certain values in the application so that we can write isolated tests. .

## Clone the project

```
$ git clone https://github.com/kadirdeniz/todo-app-backend.git
$ cd todo-app-backend
```

## Setup
##### There are two ways to start this application, first one is dockerize the application then run it and second way is run in your local machine via using makefile scripts

#### Dockerize
```
$ make dockerize
$ make docker-run
```

#### Basic Start
```
 $ make run
````

 ## Test
 ```
 $ make test
 ```
