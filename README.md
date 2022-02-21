<h1>Date difference in days Go lang CLI Application</h1>
<h2>Problem statement:</h2>

To calculate the distance in whole days between two dates, counting only the days in between those dates, For example 01/01/2001 to 03/01/2001 yields “1”. The valid date range is between 01/01/1900 and 31/12/2999, all other dates should be rejected.

<h2>How to run the application:</h2>

1.	Git clone the project on to local 
2.	Open terminal in the projects directory and run “go run main.go” command
3.	Application will prompt for user input and wait for the user to enter two different dates
4.	User has to enter two space separated dates of different values. There are validations put in place to check the user input before calculating the distance. 
Below is one example of the check that application does while taking the input.
![image](https://user-images.githubusercontent.com/26626292/155014704-68903043-31e3-4c14-b2e7-c366d28da979.png)
  So, make sure to enter the correct date range to proceed further.
5.	That’s it! Application will take care of calculating the number of days between the two inputted dates and outputs the result.


**Note: This requires Go lang already installed on the user's machine. Please follow this website for installing GO https://go.dev/doc/install**

<h2>Test outputs:</h2>
1.	For dates 02/06/1983 and 22/06/1983
 ![image](https://user-images.githubusercontent.com/26626292/155014717-beba96ad-663b-40b0-8eda-f2137d035da1.png)


2.	For dates 04/07/1984 and 25/12/1984
 ![image](https://user-images.githubusercontent.com/26626292/155014732-22e401b6-16b2-4d18-83d5-70ca83a7686a.png)


3.	For dates 03/08/1983 and 03/01/1989
 ![image](https://user-images.githubusercontent.com/26626292/155014744-928e5bbc-5c4d-4db4-b4ef-ddf7d365ca41.png)



