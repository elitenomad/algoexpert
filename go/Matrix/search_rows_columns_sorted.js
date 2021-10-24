/**
 * @param {number[][]} matrix
 * @param {number} n
 * @param {number} m
 * @param {number} x
 * @returns {boolean}
*/
class Solution
{
    //Function to search a given number in row-column sorted matrix.
    search(matrix, n, m, x)
    {
        // Use case 1
        if(x < matrix[0][0]){
            return 0
        }
        
        // Use case 2
        if(x > matrix[n-1][m-1]) {
            return 0
        }
        
        let r = n - 1
        let c = 0
        
    	let bottom = matrix[r][0]
    	
    	while(r >= 0 && c <= m - 1){
    	    if(matrix[r][c] < x){
    	        c +=1
    	    }else if(matrix[r][c] > x){
    	        r -= 1
    	    }else{
    	        return 1
    	    }
    	}
    }
}