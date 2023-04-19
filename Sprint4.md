**We worked in the Sprint4 Branch for this portion of the project**

**Our fourth group and second angular team member dropped this class**

Sprint 4 Video Link:

Overview Video Link: 

**Work Completed**

  Front End (Josh Frye):
    
    -Integrated the new angular front end so that users, decks, and cards can be created while looking decent through angular material
    
    -Added new front end file deck.service.ts to handle communication from front end to back end
    
    -Added new dex-info component so that dex information could be collected before cards were created
    
    -Small aesthetic changes with angular material
    
    -Added unit tests for new components
    
    -Aestetic changes with user list
    
    
    
 Back End (Juan Veliz and Sam St Jean):
   
   - 
   
**Unit Tests**

Front End- used the integrated unit tests with Karma/Jasmine
  
    HomePageComponent
      should create
    DexInfoComponent
      should create
    PageNotFoundComponent
      should create
    LoginComponent
      should create
    UserService
      should be created
    CreateUserComponent
      should create
    ViewDexComponent
      should create
    CreateDexComponent
      should create
     AppComponent
      should have as title 'Gator-Dex'
      should create the app
    UserListComponent
      should create
    UserDetailsComponent
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

updateUser function: 
The updateUser function allows a user to update his own information (First name, Last name, email, password). This function traverses through the database and looks for the id of the user that has logged in and wants to update his information. Once it is found, the data is overriden with the data newly inputed by the user, and the change is reflected in the database. 

CreateDeck Function:
The CreateGatorDeck function allows a user to create decks of cards within the database. This database stores the name of the deck and the class code for the deck given by a user. This function in conjunction with our mySQL database, has a foreign key relationship with User where the User ID dictates that a particular user has a deck of cards. This “has multiple ” relationship is established in our backend Go and our mysql table structure. 

GetDeck Function:
The GetDeck function allows a user to look for decks that are in the database. This function traverses through the database and looks for the deckname given by the user doing the search. If it matches one in the database, we send the deck's info to the front end and it displays. This will eventually reveal all of the gator cards that belong to the deck being searched for.

DeleteDeck Function:
The DeletDeck function allows a user as well as ourselves as admin to delete a deck from the database. The function traverses through the database and finds the deck that is asked to be deleted, and he is erased from the database.
