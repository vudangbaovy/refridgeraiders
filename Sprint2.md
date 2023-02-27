<h1>Sprint 2 Documentation</h1>
<h2>Backend API</h2>
<h3>Localhost:8080/User<h3>
<h4> Each of the /User calls needs a json with the following identifiers: "name", "password", "adminlevel" and "allergies".</p>
These can be left blank for the calls where its not applicable. Each call returns the data in the same format. </p>
If there there is something wrong with what is sent, i.e. wrong password, we currently have it send an empty json of the same format.<h4></p>
<h6><p>1. Adding a New User: Post</p>
<p>2. Updating a Current User's Information: Put</p>
<p>3. Logging in to a Current User: Get</p>
<p>4. Deleting a Current User: Delete</p><h5>
<h5><p>Here's a Get Example using postman:<h6></p>
<h6>
<p>&emsp;   "name": "test22Nick",</p>
<p>&emsp;    "password": "pa442rd",</p>
<p>&emsp;    "adminLevel": 0,</p>
<p>&emsp;    "allergies": ""<h6></p>
<h5><p>Should Return through postman:<h5></p>
<h6>
<p>&emsp;   "name": "test22Nick",</p>
<p>&emsp;    "password": "pa442rd",</p>
<p>&emsp;    "adminLevel": 2,</p>
<p>&emsp;    "allergies": "Peanuts,Milk"<h6></p>
