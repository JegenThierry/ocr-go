# OCR GO
A server that can OCR pdfs using https://github.com/ocrmypdf/OCRmyPDF

This is a very basic implementation of a OCR Backend.

Run this command to run the Backend:
```
docker-compose up --build
```

## Known-Issues

- The first request when running the server in a docker container will fail
- The UI does not inform about the progress of conversion
- UI Feedback is lacking
