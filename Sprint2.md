<h1>Sprint 2 Documentation</h1>

<h2>Completed Tasks</h2>
<p>Connect to external Database API named edemam. User types in food they desire to consume, click a button, and recieve a JSON file describing different recipes containing desired ingredient</p>
<p>Front end tests: Button click test ensuring the button that the user clicks to enter ingredient works and fetchData() test which ensures function that grabs data from API regarding users input, works</p>
<p>Created Unit Tests that cover a wide range of the Databse and internal API functionality without being overly redundent.</p>
<p>Designed and implimented a basic internal API that can preform actions on the Database.</p>
<p>Connected front and backend through a user login and register page.</p>


<h2>Backend API</h2>
<h3>Localhost:3000/User<h3>
<h4> Each of the /User calls needs a JSON with the following: "name", "password" and "allergies".</p>
These can be left blank for the calls where its not applicable. Each call returns the data in the same format. </p>
If there there is something wrong with what is sent, i.e. wrong password, we currently have it send an empty JSON of the same format.<h4></p>
<p>1. Updating a Current User's Information: Put</p>
<p>2. Logging in to a Current User: Post</p>
<p>3. Deleting a Current User: Delete</p><h5>
<br>
<p><h3>Localhost:3000/User/Register<h3>
<h4>/User/Register calls needs a JSON with the following: "name", "password" and "allergies".</p>
If there there is something wrong with what is sent, i.e. username taken, is sent empty JSON.<h4></p>
<p>1. Adding a new User: Post</p>
<br>
<h3><p>Here's a /User Post Example:<h3></p>
<h6>
<p>&emsp;   "name": "test22Nick",</p>
<p>&emsp;    "password": "pa442rd",</p>
<p>&emsp;    "allergies": ""<h6></p>
<h3><p>Should Return:<h3></p>
<h6>
<p>&emsp;   "name": "test22Nick",</p>
<p>&emsp;    "password": "pa442rd",</p>
<p>&emsp;    "allergies": "Peanuts,Milk"<h6></p>
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
<p>The test then compares it against the expected returning JSON.<h4></p>
