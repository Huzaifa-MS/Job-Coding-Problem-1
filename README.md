# Job-Coding-Problem-1
Coding problem from fetch-rewards 

Solution:

Create Api skeleton
	Generating using given api.yml file
		Based on https://ldej.nl/post/generating-go-from-openapi-3/
	Did it to save time
	All options were similar: chose https://ogen.dev/ for simplicity
Have a simple rules engine - Array is enough for now
	Expand to have rules object? Rules depending on other rules
Process receipts using rules
Store in memory
Make processing independent from storage
* Maybe Store Data in hash format for id?
	For independent processing capabilities
Make test cases to guard against changes