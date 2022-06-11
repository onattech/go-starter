# go-starter
This app simplifies starting a new GO project. After finding myself spending so much time just to start a new project along with its repository, I decided to write this app to automate the process.

It currently:
- Sets up a directory with bare minimum to get started on a GO project
- Initializes git in the project directory with a .gitignore file and commits the initial commit
- Sets up a GitHub repository and pushes the initial commit to it


## Configuration
You will have to update lines 17,18 and 19 in the main.go file to your needs then build if you want a standalone executable.
![image](https://user-images.githubusercontent.com/53895969/173196516-80ebe0bd-b2f3-44f3-80c3-21458fa25f36.png)

## Linux How-To
1. (Optional) Setup a new launcher on the desktop <br>![image](https://user-images.githubusercontent.com/53895969/173196777-1885846c-e360-4843-97f9-d36d685bd09d.png)

2. Run the app and add the project's name. Defaults to the value declared on line 17 in the main.go file.![image](https://user-images.githubusercontent.com/53895969/173196735-5887de67-4a19-44f4-bdec-22658a654708.png)

3. Select directory path. You may add more options by updating line 32 in the main.go file. <br>![image](https://user-images.githubusercontent.com/53895969/173196833-e18e1c3a-93bd-4f32-abb1-12ae0f8b138e.png)
4. If you want a GitHub repository hit enter, if 'n' is selected, program will start vscode and exit. <br>![image](https://user-images.githubusercontent.com/53895969/173196884-e7dd9f4f-873f-4ad1-a372-567348ca6b5b.png)
5. Enter desired repository name. If a repository with that name exists, you will be prompted to enter another name. <br>![image](https://user-images.githubusercontent.com/53895969/173196944-aa9eed8e-f095-43bc-97c8-399b73252508.png)
6. Enter a description if you like or leave blank if you don't. <br>![image](https://user-images.githubusercontent.com/53895969/173196969-1aba0afd-a3ec-437b-979c-19cf4b43821d.png)
7. Choose repository visibility. <br>![image](https://user-images.githubusercontent.com/53895969/173197009-9e657d93-7bb7-4957-9d91-c084cd2d6af7.png)
8. Type what the remote should be called. Defaults to origin on hitting enter. <br>![image](https://user-images.githubusercontent.com/53895969/173197040-5b1b04f2-f81e-49f9-b0bf-b08ddd73a903.png)
9. You will prompted on successful repository initialization. <br>![image](https://user-images.githubusercontent.com/53895969/173197073-454d6172-2e20-480e-a460-1c78f283374f.png)
10. If all goes well, a vscode instance should start with your new project directory, have git initialized, and its initial commit pushed to the new repository. <br>![image](https://user-images.githubusercontent.com/53895969/173197122-611d2f0c-1540-402c-91be-b738e61179c7.png)









