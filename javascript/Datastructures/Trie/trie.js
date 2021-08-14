class TrieNode{
  constructor(){ 
      this.children = new Array(26);
      this.children.fill(null);
      // isEndOfWord is True if node represent the end of the word 
      this.isEndOfWord = false
  }
}  
class Trie{
  // Trie data structure class 
  constructor(){ 
      this.root =new TrieNode()
  }
}

class Solution 
{
  //Function to insert string into TRIE.
  insert(root, key) 
  {
    let currentNode = root
    
    for(let i = 0; i < key.length; i++){
        let index = key.charCodeAt(i) - "a".charCodeAt()
        if(currentNode.children[index] === null){
            currentNode.children[index] = new TrieNode()
        }
        
        currentNode = currentNode.children[index]
    }
    
    currentNode.isEndOfWord = true
  }
  
  //Function to use TRIE data structure and search the given string.
  search(root, key) 
  { 
      let currentNode = root

      for(let i = 0; i < key.length; i++){
        let index = key.charCodeAt(i) - "a".charCodeAt()

        if(currentNode.children[index] === null){
            return 0
        }
        
        currentNode = currentNode.children[index]
    }
    
    if(currentNode !== null && currentNode.isEndOfWord) {
        return 1
    }
    
    return 0
  }
}

function main() {
  let words = ["the","a","there","answer","any","by","bye","their"]
  let obj = new Solution();
  let t = new Trie();
  console.log(t.root)
  for(let i=0;i< words.length;i++)
    {   
        let key = words[i]
        obj.insert(t.root,key);

        let abc = key;   
        let ans = obj.search(t.root, abc);
        if(ans){
            ans=1;
        }else{
            ans=0;
        }
        console.log("Answer is: ",ans);
    }
}