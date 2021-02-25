package test
//To do 
// loop through each contests and collect all questions with participant type contestant
// Create JSON of following format
// { contest id:
// 	problems{
// 		Number : "A",
// 		Name :"ABC"
// 		tags:[],
// 		submission_count;
//       problem_rating
//          AC_TIME:
// 	}

// }
type Submission struct{
 result [] Result   `json: result`

} 
type Result struct {
     ID string `json: id`
	 ProgrammingLang  `programmingLanguage`
	
}



import (
	"fmt"
)

