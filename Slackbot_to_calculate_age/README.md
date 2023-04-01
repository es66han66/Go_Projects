1- Need working slack account with admin privileges.  
2- go to api.slack.com/apps  
3- Click on create an app -> From scratch -> give any name to app name and select your workspace and create app.  
4- Go on socket mode on left and enable it -> it will ask for token name, give any and generate your socket token.  
5- Go to Event subscriptions on left and enable it -> subscribe to following events so that our bot can listen to those.  
    a) app_mention  
    b) message.im  
    c) im_history_changed  
    d) message.channels  
    e) message.mpim  
6- Go to OAuth & Permissions on left -> Add following Oauth scopes  
    a) chat:write  
    b) chat:write:public  
    c) channels:read  
    d) im:read  
    e) im:write  
    f) mpim:read  
    g) mpim:write  
7- Now Click on Install to workspace above in OAuth & Permissions -> Allow the permissions being asked  
8- We will get a bot Oauth user token-  
9- when we run our app, after that head to slack and in some group type @age-bot my year of birth is 1999 and it will give answer  