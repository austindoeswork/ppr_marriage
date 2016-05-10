# PHI PHI RHO HOUSING SCRIPT
Will run the stable marriage algorithm with the tenants as the choosers  
The Tenant's preferences are the order in which the rooms are listed in the corresponding CSV  
The Room's preferences are the offers each tenant placed on it ordered from highest to lowest  

#### hello.go:
contains the main func
#### lib.go:
contains the classes and member functions
#### marry.go:
contains the code for the stable marriage algorithm
#### parse.go:
parses the csvs in csvs/ into two maps, one map of Tenants, and one map of Rooms

#### csvs in format: 
room numbers should be in order of preference

roomnumber,offer  
roomnumber,offer  
roomnumber,...  
ex.  
2,380  
1,400  
4,500 ...  

