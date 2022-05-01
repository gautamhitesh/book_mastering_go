# book_mastering_go

Repo for code and explanation of the book Mastering Go by Mihalis Tsoukalos

The repo is divided into sections for respective chapters in the book with questions and code.

## Topics

### 1. Go and Operating System

#### 1.1. Preprocessor

#### 1.2. Godoc Utility

#### 1.3. Compiling and Executing Go Code

`For DSA and sample problems create one folder and for all go programs in it, declare main package with main function and run the program`

#### 1.4. Go rules --> 2

#### 1.5. Downloading GO packages

#### 1.6. Stdin, Stdout, Stderr

#### 1.7. Output : Standard, Print

#### 1.8. User input

#### 1.9. :=


#### 1.10. Command line arguments


#### 1.11. Error Output


#### 1.12. Writing Logs



#### 1.13. Error handling
    
    a. import "errors
    b. errors.New() creates a new variable of errors type
    c. errors.Error() is used to compare error with a string
    d. err != nil
    e. Error handling is different from printing errors to console or a logger. Printing to a logger is one of the ways to handle errors (not an ideal one). Error handling is a way to control the execution of program if it faces a condition which results in a unpredictable scenario.
    f. Try sending all errors to a log service.
    `log.Println(err)
    os.Exit(10)`
    g. Panic function that stops the program. Recover function is used to save from panic. Read more in chapter 2.
    h. 

