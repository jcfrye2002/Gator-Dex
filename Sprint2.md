Google drive narrated video link: https://drive.google.com/file/d/1CmIoqSWsE9wzAfYM0LGDL5fxb45qLQzO/view?usp=sharing

Work completed in Sprint 2:

  **Front End** (Josh Frye)

      -Completely redid front end login screen to have CSS instead of SCSS
      -Created a new account page to pair with the backend functionality of creating accounts
      -Created a "user list" dummy page for future use
      -Created a toolbar for quick navigation through the website
      -Imported and used Angular Material for the login form, inputs, and toolbar
      -Used CSS to make the components of the website look good
      -Imported and used browser animation for the text boxes for login and new account forms
      -Added routing for navigation through website
      -Added wildcard routing to prevent unexpected behavior from entering a weird address
      -Changed favicon.ico to a small gator (go gators)
      -Changed the tab name to be Gator-Dex
      -Created and passed a simple Cypress test
      -Passed all Jasmine unit tests
      
  **Back End** (Sam St Jean and Juan Veliz)
  
  - connected the backend with the frontend
  - connected the backend with a new database with more features for users
  - got the create user function working in conjunction with the front end
  - got delete user working with boilerplate angular files
  - got getUser function working with boilerplate angular files
  - created a parent child sql database relationship between users and their card decks on mysql utilizing foreign & primary keys
  - created a login function to begin connecting with the front ends login
Testing:

  **Front End** (Josh Frye)


      **Cypress:** 
        A simple login form filling test
      **Unit Testing via Jasmine (Passed 9/9):**
        LoginComponent
          -should create
        UserDetailsComponent
          -should create
        PageNotFoundComponent
          -should create
        UserService
          -should create
        CreateUserService
          -should create
        AppComponent
          -should create the app
          -should have as title 'Gator-Dex'
          -should render title
        UserListComponent
          -should create
          
   **Back End** (Sam St Jean and Juan Veliz)
   - utalized Golangs built in testing to wite unit tests for working functions
   - unit test for createUser passed
   - unit test for getUser passed 
   - unit test for deleteUser passed

  **Back End** (API Documentation)
  
  CreateUser Function: 
  
  The CreateUser function allows a user to create his account within the database that we have linked to the front end. 
  This database stores the first name, last name, email and password given by a user. This function in conjection with our       mySQL database disallows duplicate values for emails. If the email is not found in the database, the aformentioned data is     stored in our local database, and is ready for usage by all the other functions. 
  
  GetUser Function: 
  
  The GetUser function allows a user to look for other users that are in the database. This function traverses through the       database and looks for the email given by the user doing the search. If it matches one in the database, we send the user's     info to the front end and it displays. This will eventually reveal all of the GatorDex that belong to the user being searched   for.
  
  DeleteUser Function: 
  
  The DeletUser function allows a user as well as ourselves as admin to delete an account from the database. The function         traverses through the database and finds the user that is asked to be deleted, and he is erased from the database. 
