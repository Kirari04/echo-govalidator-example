# Example of using govalidator with echo

## Result Examples

```md
GET http://127.0.0.1:1323/
{"message":"name: Name cant be empty, name needs to be between 2 and 10 characters"}


GET http://127.0.0.1:1323/?name=A
{"message":"name: name needs to be between 2 and 10 characters"}

GET http://127.0.0.1:1323/?name=ABC
Hi ABC
```