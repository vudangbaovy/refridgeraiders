<h2>Frontend</h2>
<h3><p>What's changed:</p></h3>
<p>1. Allowed displaying of the recipes</p>
<p>2. Integrated recipe bookmarking</p>
<p>3. Profile page now displays the current logged in user's name and username</p>

<p>Front-end is currently using POST request to /login for the login function and /user to get user's profile information from the databse.</p>

<h3>Tests</h3>
<p> 1. Cypress test for functionality of home page and the user registration routing from Home > Log In > "First time at KitchenRescue?" > Register</p>
<p>2. Cypress test to verify that the POST request to backend database returns the correct username of the current logged in user</p>

<h2>Backend</h2>
<h3><p>What's changed:</p></h3>
<p>1. Backend issues with sessions have been fixed</p>
<p>2. All previous and current api calls have the ablilty to create sessions(in the case one already isnt up and a username and password is provided) or use a preexisting session if already authorized</p>
<p>3. /bookmark is now available (Post, Put and delete)</p>
<p>4. Fixed bugs with front and backend connection</p>
<br>
<h2>Backend API</h2>
<h3>Localhost:3000/login </h3>
<p> Validates user login credentials and returns user profile information if true
<p>&emsp;   "user": "Maria123",</p>
<p>&emsp;   "password": "pass",</p>
<h4><p>Returns:</h4></p></h6>
<h6>
<p>&emsp;   "UserNotes": [
    {
      "UserRef": 1234,
      "User": "Maria123",
      "RecipeName": "gigi hadid pasta ",
      "Note": "delicious!"
    },
    {
      "UserRef": 5678,
      "User": "Maria123",
      "RecipeName": "Colombian arepas",
      "Note": "add one more cup of flour next time"
    }
    // more UserNotes objects if present
  ],</p>
<p>&emsp;   "userbookmarks": "fettuccine alfredo, brownie and iceream mix",</p>
<p>&emsp;   "user": "Maria123",</p>
<p>&emsp;   "firstn": "Maria",</p>
<p>&emsp;   "lastn": "Morales",</p>
<p>&emsp;   "password": "pass",</p>
<p>&emsp;   "allergies": "Pie"</h6></p>
<h4><p>If Unsuccessful Returns:</h4>
 </h6>


<h3>Localhost:3000/allergies </h3>
<p>The /allergies calls need a JSON with the following: "name", "password" and "allergies".</p>
Gives the user's data about their allergies so once we get tokens working may change. The allergies can be left blank in the case of POST.</p>
status code 201 if unsuccessful</p>
</br>
<p><h5>1. Getting User's Allergies: POST</p>
<p><h5>2. Updating User's Allergies: PUT</p>
<br>
<h4><p>POST Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
  <h4><p>Returns:</h4></p></h6>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "Pie"</h6></p>
<h4><p>If Unsuccessful Returns:</h4></p>
<h6>
<p>Status Code: 201</h6></p>
<br>
<h4><p>PUT Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "NewAllergy"</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": ",preExisting,NewAllergy"</h6></p>
<br>
<p><h3>Localhost:3000/user/register</h3>
<p>/user/register calls need a JSON with the following: "name", "password", "firstN" and "lastN".</p>
<p>creates a new user entry in the database, lastN and firstN can be left blank</p>
status code 201 if unsuccessful</p>
<br>
<p>1. Adding a new User: POST</p>
</br>
<h4><p>POST Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "Nicholas",</p>
<p>&emsp;   "lastN": "Callahan"</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "Nicholas",</p>
<p>&emsp;   "lastN": "Callahan"</h6></p>
<br>
<p><h3>Localhost:3000/user</h3>
/user calls need a JSON with the following: "user", "password", "firstN" and "lastN".</p>
returns username, password, first and last name</p>
status code 201 if unsuccessful</p>
<br>
<p>1. Getting a User's First & Last name: POST</p>
<p>2. Changing a User's First & Last name: PUT</p>
<p>3. Deleting a new User: DELETE</p>
<br>
<h4><p>POST Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "Nicholas",</p>
<p>&emsp;   "lastN": "Callahan"</h6></p>
<br>
<h4><p>PUT Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "NewFirstName",</p>
<p>&emsp;   "lastN": "NewLastName"</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "NewFirstName",</p>
<p>&emsp;   "lastN": "NewLastName"</h6></p>
<br>
<p><h3>Localhost:3000/note/create</h3>
/note/create calls need a JSON with the following: "user", "password", "recipeName" and "note".</p>
user and password required for right now until tokens, recipeName is the name of the recipe, and note is the note</p>
status code 201 if unsuccessful</p>
<br>
<p>1. Adding a new user note: POST</p>
<br>
<h4><p>Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",</h6></p>
<br>
<p><h3>Localhost:3000/note</h3>
/note calls need a JSON with the following: "user", "password", "recipeName" and "note".</p>
same parameters as /note/create, note can be blank for POST</p>
status code 201 if unsuccessful</p>
<br>
<p>1. Gets a user's note: POST</p>
<p>2. Updates a user's note: PUT</p>
<p>3. Deleting a user's note: DELETE</p>
<br>
<h4><p>POST Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",</h6></p>
<br>

<p><h3>Localhost:3000/bookmark</h3>
/bookmark calls need a JSON with the following: "user", "password" and "bookmark".</p>
status code 201 if unsuccessful</p>
<br>
<p>1. Gets a user's bookmarks: POST</p>
<p>2. adds a user's bookmark: PUT</p>
<p>3. Deleting a user's bookmark: DELETE</p>
<br>
<h4><p>PUT Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "bookmark": "Pizza"</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": ",preExisting,Pizza"</p>
<br>
<h4><p>POST Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "bookmark": ""</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": ",preExisting,Pizza"</p>
<br>

<h2>Backend Unit Tests</h2>
<br>
<h3>User Search DB Test</h3>
<p>A Test User is added to the database and then is seached for inside the databse. The user is hard deleted after the test.</p>
<p>This requires the DB respond to the test.</p>
<br>
<h3>User Add DB Test</h3>
<p>Three Test Users are added to the database. The Test Users are then searched for if they aren't found the test fails. The user is hard deleted after the test.</p>
<p>This requires the DB respond to the test.</p>
<br>
<h3>NotesPost API Test</h3>
<p>The NotesPost API Test tests the backend's ablity to receive and interpret a json file sent as a http post request from /note. A mock JSON file is sent to the router, then the test waits for the correct response.</p>
<p>A mock json is created requesting a pre-existing note from a pre-existing user and then the test waits for the server response. The recived json is tested against the known correct response.</p>
<p>This requires the server, router and DB to receive, interpret and respond to the test.</p>
<br>
<h3>AllergiesPost API Test</h3>
<p>The AllergiesPost API Test tests backend's ablity to receive and interpret a json file sent as a http post request from /allergies. A mock JSON file is sent to the router, then the test waits for the correct response.</p>
<p>The test is similar in design to the previous API tests, A mock json is created requesting a pre-existing allergies from a pre-existing user and then the test waits for the server response. The recived json is tested against the known correct response.</p>
<p>This requires the server, router and DB to receive, interpret and respond to the test.</p>
<br>
<h3>UserPost API Test</h3>
<p>The UserPost API Test tests the backend's ablity to receive and interpret a json file sent as a http post request from /user. A mock JSON file is sent to the router, then the test waits for the correct response.</p>
<p>The test is similar in design to the previous API tests, A mock json is created requesting a pre-existing First and last name from a pre-existing user and then the test waits for the server response. The recived json is tested against the known correct response.</p>
<p>This requires the server, router and DB to receive, interpret and respond to the test.</p>
<br>
<h3>UserPut API Test</h3>
<p>The UserPost API Test tests the backend's ablity to receive and interpret a json file sent as a http put request from /user. A mock JSON file is sent to the router, then the test waits for the correct response.</p>
<p>A mock json is created requesting to change a pre-existing user's first and last name and then the test waits for the server response. The recived json is tested against the known correct response. After reciving the correct response the test then sends a http post request on /user to ensure that the change was saved to the data base.</p>
<p>This requires the server, router and DB to receive, interpret and respond to the test.</p>
<br>
<h3>UserDelete API Test</h3>
<p>The UserPost API Test tests the backend's ablity to receive and interpret a json file sent as a http delete request from /user. A mock JSON file is sent to the router, then the test waits for the data base to respond.</p>
<p>The test is similar in design to the previous API tests, however the test does not check the response from the server it instread checks the data base directly for the if the deleted user is still there.</p>
<p>This requires the server, router and DB to receive, interpret and react to the test.</p>
<br>
<h3>Server API Test </h3>
<p>This test is relatively simple, it tests if the server is can read json files through the router this test runs on server start up.</p>
<p>This requires the server and router to receive interpret and react to the test.</p>
<br>
<h3>Bookmark API Test </h3>
<p>This test is the lengthiest test out of them all, and represents a complete end to end test of the bookmark command, the test first calls the server through a POST to /bookmark to check that there are no preExisting bookmarks, then adds two bookmarks (KeyLimePie and Pie) individually and ensures their response is correct, then the test checks both bookmarks were added by calling the POST command (using the opened session) again and finally the Pie bookmark is deleted and the KeyLimePie is check to ensure that delete properly deleted the correct bookmark. </p>
<p>This requires the server and router to receive interpret and react to the test.</p>
<br>
<h3> Correct Password Testing</h3>
<p>This test tests the password functions, which include hashing the passwords and comparing the real password with the hashed one. This test compares the correct password, that was previously declared and hashed, and compared with the stored hash password.</p>
<h3>Incorrect Password Testing</h3>
<p>Similar to the one above, this test hashes a password given and then is then compared with a different (wrong) password.</p>
