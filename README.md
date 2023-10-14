# Me-My Encryption

## About the Algorithm (I'll try eplaining better when I get time)

let the key to be sent here be 
k = my_roll_me_roll_rand (eg. `210667_21xxxx_&89h9hKJbx`)
```mermaid
sequenceDiagram
Me ->> Server: Take my hearts {heart1, heart2, ...}
Me ->> Server: SHA(k1), enc(SHA(k1))<br/> SHA(k2), enc(SHA(k2))...
Note right of Server: On login
My-->>Server: Give me all enc & pubkeys
Server-->>My: enc1, enc2, ...
My-->>Server: heart1 is for me
My-->>Server: enc1 is mine
Server-->>My: How do I verify ?
My-->>Server: Here is SHA I got<br/> on decoding,
Note left of Server: SHA & enc match.<br/> So, assign the pair to user.
Note left of My: Now when "my" send his/her<br/>hearts, assigned pairs are also sent.<br/> Also, "my" gets a token for 10min.<br/>To send claimed heart again <br/>if not sent already. (quick fix for clash)
My->>Server: SHA(l1), enc(SHA(l1))<br/> SHA(l2), enc(SHA(l2))...<br/> + <br/> SHA(k1), enc(SHA(k1)) <br/>-enc with pubkey of receiver.
Server-->Me: Syncing data with Me...
Me->>Server: I got my heart back.
Server->>Me: How do I verify ?
Me->>Server: Here is my k(210667_21xxxx_&89h9hKJbx).<br/>Verify it with SHA.
Me->My: Matched
 ```
## What is stored and How ??

For user signup, a verification code will be sent to the registered iitk email id(as per OA portal) of the user. 

For the signup request user will generate an RSA public-private key pair and send the password(hashed), public key, AES encrypted private key, verification code which will be saved by the backend on successfull request.

On user login, the user will fetch data of the registered users to send heart to. 

When the user selects the person/s, they are saved in the user data which is sent to the server(in AES encrypted form) if the user wishes to save(not send) the choices.


How the user gets heart ?? 

Same as how it was earlier, if the user can decrypt the encrypted data(heart) then it is for the user. Which is then verified by the server by matching the hash. And, then added to the user profile.(only the hash is attached) 

How hearts are sent ??

User selects the peoples to which it want to send hearts to, or we can say that the user encrypts SHA256 hash of a string of the form "user-roll_smwn-roll_random-val5347fth" using RSA encryption of the public key of `smwn`.

Along with those, if(by any chance) the user has got any hearts, the hash of those hearts is treated same as hash described above and it is sent to all the choices ie. that(`*those` if multiple hearts received) hash is encrypted with public key of all the choices and sent in the request.

Wont it have edge cases causing collisions ??

It may have, but we'll take care of it by allowing the user to send the heart again(only return heart) if the user gets any heart within x minutes of sending the hearts. Thus, solving any problem that may arise due to syncronization issues at any side.

** in code implementation it will be much more complex than this(cuz of some other cases) but that's just to keep hekar peeps at a little distance from us 😊.

Since, only the encrypted hashes are sent which contain alphanumeric garbage random values so It should be sufficient to keep the user from knowing who sent them heart without sending it to them back.   

## How to setup for development

Clone the repo from the github link & use command `go mod tidy` to install all dependencies (make sure you've installed go).

Install(if not already) & Run the postgres server(refer [link](https://www.postgresql.org/download/)) on your device (depending the OS you're using).

Depending on the authentication details you've set up modify the `.env` file(refer envformat.txt) by filling in details of the postgres server.

> Format of the .env file -- Don't forget to disable SSL if not already.
```
host = localhost
port = 5432
password = 'password'
dbName = postgres
user = postgres
CfgAdminPass = something

AdminId = may_be_aleatoryfreak
AdminPass = you_can_just_guess_it

# Make sure both are different
UserjwtSigningKey = "something"
HeartjwtSigningKey = "something2"
``` 

Build the app by using `go mod build` & then run `./me-my_encryption` (the generated file).
OR you may directly run the main.go file as well.

Server should be up & listening at port 8080.