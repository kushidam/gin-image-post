
## curl での使い方:

単一ファイル
```
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
```

複数ファイル
```
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```



[https://gin-gonic.com/ja/docs/examples/upload-file](https://gin-gonic.com/ja/docs/examples/upload-file/)