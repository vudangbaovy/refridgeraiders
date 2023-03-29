<h2>Backend API</h2>
<h3>Localhost:3000/allergies <h3>
<h4> The /allergies calls need a JSON with the following: "name", "password" and "allergies".</p>
The idea behind the call is to give the user's data about their allergies so once we get tokens working may change.</p>
The allergies can be left blank, the only reason its included is because the call returns the data in the same format. (can be changed) </p>
In the case of wrong password, we currently have it send an empty JSON of the same format. (will be changed to error codes)<h4></p>
<br>
<p><h5>1. Getting User's Allergies: POST</p><h5>
<p><h5>2. Updating User's Allergies: PUT</p><h5>
<br>
<h4><p>POST Example:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": ""<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "Pie"<h6></p>
<h4><p>If Unsuccessful Returns:<h4></p>
<h6>
<p>&emsp;   "name": "",</p>
<p>&emsp;   "password": "",</p>
<p>&emsp;   "allergies": ""<h6></p>
<br>
<p><h3>Localhost:3000/user/register<h3>
<h4>/user/register calls need a JSON with the following: "name", "password" and "allergies".</p>
If there there is something wrong with what is sent, i.e. username taken, is sent empty JSON.<h4></p>
<br>
<p><h5>1. Adding a new User: POST</p><h5>
<br>
<h4><p>Example:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "Pie"<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "Pie"<h6></p>
<br>
<p><h3>Localhost:3000/user<h3>
<h4>/user calls need a JSON with the following: "name", "password" and "allergies".</p>
will be changed later to not include allergies returns user's name, password and allergies if successful(used for testing)</p>
sends an empty json of same format if unsuccessful login</p>
will eventually be changed to error messages<h4></p>
<br>
<p><h5>1. Deleting a new User: DELETE</p><h5>
<br>
<h4><p>Example:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": ""<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe2",</p>
<p>&emsp;   "allergies": "Pie"<h6></p>
<br>
<p><h3>Localhost:3000/note/create<h3>
<h4>/note/create calls need a JSON with the following: "name", "password", "recipeName" and "note".</p>
name and password required for right now until tokens, recipeName is the name of the recipe, and note is the note</p>
sends an empty json of same format if unsuccessful login</p>
will eventually be changed to error messages<h4></p>
<br>
<p><h5>1. Adding a new user note: POST</p><h5>
<br>
<h4><p>Example:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": ""</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": ""</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",<h6></p>
<br>
<p><h3>Localhost:3000/note<h3>
<h4>/note calls need a JSON with the following: "name", "password", "recipeName" and "note".</p>
name and password required for right now until tokens, recipeName is the name of the recipe, and note is the note</p>
sends an empty json of same format if unsuccessful login</p>
will eventually be changed to error messages<h4></p>
<br>
<p><h5>1. Gets a user's note: POST</p><h5>
<p><h5>2. Updates a user's note: PUT</p><h5>
<p><h5>3. Deleting a user's note: DELETE</p><h5>
<br>
<h4><p>POST Example:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": ""</p>
<p>&emsp;   "note": "",<h6></p>
<h4><p>Returns:<h4></p>
<h6>
<p>&emsp;   "name": "Nick",</p>
<p>&emsp;   "password": "Pwe3",</p>
<p>&emsp;   "recipeName": ""</p>
<p>&emsp;   "note": "TOO MUCH CHEESE",<h6></p>
<br>

  <h2>Backend Unit Tests</h2>
<br>
  <h3>User Search<h3>
<h4>A Test User is added to the database and then is seached for inside the databse.</p>
<p>The user is hard deleted after the test.<h4></p>
<br>
  <h3>User Add<h3>
<h4>Three Test Users are added to the database. The Test Users are then searched for if they aren't found the test fails.</p>
<p>The user is hard deleted after the test.<h4></p>
<br>
  <h3>User Login Api<h3>
<h4>A mock JSON file is sent to the router, the test then waits for the returned JSON signaling a successful login.</p>
<p>The test then compares it against the expected returning json.<h4></p>
<br>
  <h3>Note POST Test Api<h3>
<h4>Similar to the User Login test the Note Post tests both the router and the database by sending a mock json file to try</p>
to get a user's note from the database and then he test then compares it against the expected returning json<h4></p>
   <h3>Server Test Api<h3>
<h4>This test is relatively simple, all it does is test if the server is can read json files through the router,</p>
but its run every time the server starts.<h4></p>
