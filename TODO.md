## TODO

- Setup bot env - check
- Create bot application commands - check
- Setup bot props - check
- Make bot run - check
- Modelate the DB - check
- Implement bot functionalities (/inscrever, /ta-pago, /ranking) - check

### Features

- /help: Should create a command to aux user to understand better another commands. check

- /inscrever: Should create a new table of user collecting discord data. check

- /ta-pago: Should increment the count table in DB. check

- /ranking: Should select and filter in the DB the name and the count table. check
  **Implement ranking service, exec queries in DB**
  **and return the value then format in discord embed**

### FIXES:

- /ta-pago: insert a new field in the DB to validate the day which the user submit
  the workout. check

- /ta-pago: when user submit a workout, the response message should be a custom message about
  mocking the user. check

- /ranking: fix the bug that when ranking is greater than 3 and less than 10, the string
  is repeating the numbers. check
