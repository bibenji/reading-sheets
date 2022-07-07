 docker run -it --rm --name test -v "$PWD":/usr/src/myapp -w /usr/src/myapp php:7.4-cli php test.php
 docker run -it --rm --name test -v "$PWD":/usr/src/myapp -w /usr/src/myapp php php test.php