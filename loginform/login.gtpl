<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Login</title>
  </head>
  <body>
    <form action="/login?username=test" method="post">
      <input type="checkbox" name="intrest" value="football" />Football
      <input type="checkbox" name="intrest" value="basketball" />Basketball
      <input type="checkbox" name="intrest" value="tennis" />Tennis 
      Username: <input type="text" name="username" /> 
      Password: <input type="password" name="password" />
      <input type="hidden" name="token" value="{{.}}">
      <input type="submit" value="Login" />
    </form>
  </body>
</html>
