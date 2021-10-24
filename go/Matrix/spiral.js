let top = 0
let bottom = r - 1
let left = 0
let right = c - 1
let numbers = []


while(left <= right && top <= bottom){
    for(let i = left; i <= right;i++){
        numbers.push(matrix[top][i])
    }
    top += 1
    
    for(let i = top; i <= bottom; i++){
        numbers.push(matrix[i][right])
    }
    
    right -= 1
    
    if(top <= bottom){
        for(let i = right; i >= left; i-- ){
            numbers.push(matrix[bottom][i])
        }
        
         bottom -= 1
    }
    
   
    if(left <= right){
         for(let i = bottom; i>= top;i--){
            numbers.push(matrix[i][left])
        }
        left += 1
    }
   
}

return numbers  