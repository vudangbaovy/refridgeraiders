<h2>Frontend</h2>
<h3><p>What's changed:</p></h3>
  <p>1. New and improved home page</p>
  <p>2. Improved user log in/ register experience</p>
  <p>3. Working search bar and displaying API results</p>
  <p>4. User profile page</p>
 <h3><p>Tests:</p></h3>
   <p>1. Cypress test for the redirect links available on the site. The test clicks through the Login button, then the "First time at KitchenRescue?" link, then checks that it's at the register page.</p>
   <p>2. Cypress test to input "beef" as a parameter in the search bar and returns the correct recipes from the API.</p>
<h2>Backend API</h2>
<h3><p>Whats changed:</p></h3>
<h4><p>1. "name" has been changed to "user" in all calls to differentiate bettween username and fist/last names.</p></h4>
<p>2. /user/register now takes different parameters, no longer takes allergies, only first and last name(not required)</p>
<p>3. /user has two more calls, POST and PUT</p>
<p>4. Two new calls, /note and /note/create</p>
<p>5. Passwords are now hashed, and stored as such into the database for user account security</p>
<p>6. User authentication using sessions</p>
<br>
<h3>Localhost:3000/allergies </h3>
<h4> The /allergies calls need a JSON with the following: "name", "password" and "allergies".</p></h4>
Gives the user's data about their allergies so once we get tokens working may change. The allergies can be left blank in the case of POST.</p>
sends an empty json of same format if unsuccessful (will eventually be changed to error messages)</p>
</br>
<p><h5>1. Getting User's Allergies: POST</p><h5>
<p><h5>2. Updating User's Allergies: PUT</p><h5>
<br>
<h4><p>POST Example:<h4></p>
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
<p>&emsp;   "user": "",</p>
<p>&emsp;   "password": "",</p>
<p>&emsp;   "allergies": ""</h6></p>
<br>
<h4><p>PUT Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "NewAllergy,Pie"</h6></p>
<h4><p>Returns:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "NewAllergy,Pie"</h6></p>
<br>
<p><h3>Localhost:3000/user/register</h3>
<h4>/user/register calls need a JSON with the following: "name", "password", "firstN" and "lastN".</p>
<p>creates a new user entry in the database, lastN and firstN can be left blank</p>
sends an empty json of same format if unsuccessful (will eventually be changed to error messages)</h4></p>
<br>
<p><h5>1. Adding a new User: POST</p></h5>
</br>
<h4><p>POST Example:</h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "Nicholas",</p>
<p>&emsp;   "lastN": "Callahan"</h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "Nicholas",</p>
<p>&emsp;   "lastN": "Callahan"<h6></p>
<br>
<p><h3>Localhost:3000/user<h3>
<h4>/user calls need a JSON with the following: "user", "password", "firstN" and "lastN".</p>
returns username, password, first and last name</p>
sends an empty json of same format if unsuccessful (will eventually be changed to error messages)<h4></p>
<br>
<p><h5>1. Getting a User's First & Last name: POST</p><h5>
<p><h5>2. Changing a User's First & Last name: PUT</p><h5>
<p><h5>3. Deleting a new User: DELETE</p><h5>
<br>
<h4><p>POST Example:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "Nicholas",</p>
<p>&emsp;   "lastN": "Callahan"<h6></p>
<br>
<h4><p>PUT Example:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "NewFirstName",</p>
<p>&emsp;   "lastN": "NewLastName"<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "firstN": "NewFirstName",</p>
<p>&emsp;   "lastN": "NewLastName"<h6></p>
<br>
<p><h3>Localhost:3000/note/create<h3>
<h4>/note/create calls need a JSON with the following: "user", "password", "recipeName" and "note".</p>
user and password required for right now until tokens, recipeName is the name of the recipe, and note is the note</p>
sends an empty json of same format if unsuccessful (will eventually be changed to error messages)<h4></p>
<br>
<p><h5>1. Adding a new user note: POST</p><h5>
<br>
<h4><p>Example:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",<h6></p>
<br>
<p><h3>Localhost:3000/note<h3>
<h4>/note calls need a JSON with the following: "user", "password", "recipeName" and "note".</p>
same parameters as /note/create, note can be blank for POST</p>
sends an empty json of same format if unsuccessful (will eventually be changed to error messages)<h4></p>
<br>
<p><h5>1. Gets a user's note: POST</p><h5>
<p><h5>2. Updates a user's note: PUT</p><h5>
<p><h5>3. Deleting a user's note: DELETE</p><h5>
<br>
<h4><p>POST Example:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "user": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": "Pizza"</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",<h6></p>
<br>

  <h2>Backend Unit Tests</h2>
<br>
  <h3>User Search</h3>
<h4>A Test User is added to the database and then is seached for inside the databse.</p>
<p>The user is hard deleted after the test.</h4></p>
</br>
  <h3>User Add</h3>
<h4>Three Test Users are added to the database. The Test Users are then searched for if they aren't found the test fails.</p>
<p>The user is hard deleted after the test.</h4></p>
<br>
  <h3>User Login Api</h3>
<h4>A mock JSON file is sent to the router, the test then waits for the returned JSON signaling a successful login.</p>
<p>The test then compares it against the expected returning json.</h4></p>
</br>
  <h3>Note POST Test Api</h3>
<h4>Similar to the User Login test the Note Post tests both the router and the database by sending a mock json file to try</p>
to get a user's note from the database and then he test then compares it against the expected returning json</h4></p>
<br>
 <h3>Server Test Api</h3>
<h4>This test is relatively simple, all it does is test if the server is can read json files through the router,</p>
but its run every time the server starts.</h4></p>
   <h3> Correct Password Testing</h3>
 <h4> This test tests the password functions, which include hashing the passwords and comparing the real password with the hashed one. This test compares
   the correct password, that was previously declared and hashed, and compared with the stored hash password.
   <h3>Incorrect Password Testing</h3>
 <h4> Similar to the one above, this test hashes a password given and then is then compared with a different (wrong) password. 
