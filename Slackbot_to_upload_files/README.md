1- Need working slack account with admin privileges.  
  
2- go to api.slack.com/apps  
  
3- Click on create an app -> From scratch -> give any name to app name and select your workspace and create app.  
  
4- Go on socket mode on left and enable it -> it will ask for token name, give any and generate your socket token.
      
5- Go to OAuth & Permissions on left -> Add following Oauth scopes  
  
```a) chat:write  
b) chat:write:public  
c) channels:read  
d) im:read  
e) im:write  
f) mpim:read  
g) mpim:write
h) files:read
i) files:write
j) remote_files:read
k) remote_files:write
l) remote_files:share```
  
6- Now Click on Install to workspace above in OAuth & Permissions -> Allow the permissions being asked  
  
7- We will get a bot Oauth user token-  

8- We need Channel ID for the channel where we will use this bot- hence go to the channel in slack and right click on it and click on open channel details and in the bottom copy the channel ID.  
  
9- When we start the app, it uploads the file on Slack Channel, it can give error if the bot is not added to channel prehand  