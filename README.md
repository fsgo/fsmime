# fsmime

Auto register extension mime types if not exists.  
If your go program runs on a very old OS, this will be useful.


Data from:
 1. https://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types

## Usage

```go
import _ "github.com/fsgo/fsmime" // auth register extension mime types
```