<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Document</title>
  </head>
  <body>
    <form
      action="http://127.0.0.1:9090/upload"
      enctype="multipart/form-data"
      method="post"
    >
      <input type="file" name="uploadfile" />
      <input type="hidden" name="token" value="{{.}}" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
