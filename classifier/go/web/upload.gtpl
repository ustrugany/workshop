<html>
  <head>
         <title>Test Upload a File</title>
  </head>
  <body>
<form enctype="multipart/form-data" action="http://localhost:8080/classify" method="post">
          {{/* 1. File input */}}
          <input type="file" name="uploadfile" />

          {{/* 2. Submit button */}}
          <input type="submit" value="upload file" />
      </form>

  </body>
  </html>
