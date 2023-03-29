**We worked in the Sprint3 Branch for this portion of the project**

Video Link: https://drive.google.com/file/d/15Cnv87NAlCh8WVMc9zrZNVHw1VBro4Gu/view?usp=sharing

**Work Completed**

  Front End (Josh Frye):
    
    -Integrated the new angular component with the backend so that users can be created
    
    -Updated login, create user, app components with angular material for more aesthetic text entries and a overall more visually appealing website
    
    -Complete overhaul of color/layout of website
    
    -Redid Angular Material Toolbar for easier navigation between pages with buttons
    
    -Worked on User List that can modify and delete users that are created
    
    -Created a homepage component where users can land when they log in, integrated a photo
    
    -Created a "Create Dex" component where users are able to create Gator-Dex and Cards, text boxes are fillable and check for empty portions before form completion
    
    -Created a "View Dex" component where users will eventually be able to view their Dex
    
    -Created a "Page Not Found" component that appears when a user tries to go to an invalid link and allows for navigation back to the homepage
    
    -Simplified and adjusted routing for all components
    
   Back End (Juan Veliz and Sam St Jean):
   
   - Update user function was finished and implemented 
 
- Get users function was updated and implemented 
 
- Added structs for cards and gators decks 

- Created foreign key relationship between user and gator-decks, and gator-decks and card in MySQL and in our backend

- Created deck tables in MySQL to hold users decks and cards

- Created create Deck function to make decks
 
- Created getDeck to retrieve a specific deck

- Created delete deck function to delete a deck

- Created update deck function to update decks
 
- Connected frontend with the functions that did not previously work (get users and update users) 
   
**Unit Tests**

Front End- used the integrated unit tests with Karma/Jasmine
  
    -CreateDexComponent
      should create
    -HomePageComponent
      should create
    -CreateUserComponent
      should create
    -PageNotFoundComponent
      should create
    -LoginComponent
      should create
    -UserService
      should be created
    -UserListComponent
      should create
    -UserDetailsComponent
      should create
    -AppComponent
      should create the app
      should have as title 'Gator-Dex'
    -ViewDexComponent
      should create
    
Back End

-Utilized GoLang built in testing to write unit test for working functions

-Unit test for get users passed

-Unit test for update passed

-Unit test for create user passed 

-Unit test for delete user passed

-Unit test for get user passed

-unit test for login user passed



Back End (API Documentation)

CreateUser Function:
The CreateUser function allows a user to create his account within the database that we have linked to the front end. This database stores the first name, last name, email and password given by a user. This function in conjunction with our mySQL database disallows duplicate values for emails. If the email is not found in the database, the aforementioned data is stored in our local database, and is ready for usage by all the other functions.

GetUser Function:
The GetUser function allows a user to look for other users that are in the database. This function traverses through the database and looks for the email given by the user doing the search. If it matches one in the database, we send the user's info to the front end and it displays. This will eventually reveal all of the GatorDex that belong to the user being searched for.

DeleteUser Function:
The DeletUser function allows a user as well as ourselves as admin to delete an account from the database. The function traverses through the database and finds the user that is asked to be deleted, and he is erased from the database.

userLogIn function:
The login function allows a user to log into their account if they use the correct email and password. This function traverses through the database and looks for the email given by the user doing the search. If it matches one in the database and the password provided is the same and the one in the database that correlates with the user, then the user may be logged in. 

CreateUser Function:
The CreateGatorDeck function allows a user to create decks of cards within the database. This database stores the name of the deck and the class code for the deck given by a user. This function in conjunction with our mySQL database, has a foreign key relationship with User where the User ID dictates that a particular user has a deck of cards. This “has multiple ” relationship is established in our backend Go and our mysql table structure. 

