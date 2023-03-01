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
