# golang
Golang Tutorial
<!-- ----------------------------------- -->
When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path.

before : run the *go mod init "foldername"* created the go.mod file for initilization
<!-- -------------------------------- -->

To run the gofile 
--> go run : go run "filenamewithextension" 
for example file is variales.go , than the command will be ----------- *go run variables.go*

<!--  -->
Note " in golang you can use := operator inside the method but we can't use this outside the methjod/func. for more refence check variables line number 23

in Golang the **public variable** can be make by capital letter of the 1st variable 
for example *const LoginToken string= "jbekjejvrv"*  
<!--  -->

<!--  -->
1.**ReadWrite** To take the input the package which helps is *buffio*
2.**Conversion** to convert the valure from one data type to another , package is *strconv*
<!--  -->