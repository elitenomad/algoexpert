import LinkedListNode from "./LinkedListNode";

export default class LinkedList {
  constructor() {
    this.head = null; // head is of type LinkedListNode
    this.tail = null; // tail is of type LinkedListNode
  }

  prepend(value) {
    // Create a node with given value and next node as current head
    let newNode = new LinkedListNode(value, this.head)

    // Set the node created as NEW head
    this.head = newNode

    // At this time if you don't have a tail set the tail to newNode
    if(!this.tail){
      this.tail = newNode
    }

    return this;
  }

  append(value) {
    // Create a node with give value and next node as null
    let newNode = new LinkedListNode(value, null)

    if(!this.head){
      this.head = newNode
      this.tail = newNode
      return this
    }

    this.tail.next = newNode
    this.tail = newNode

    return this;
  }

  delete(value){
    if(!this.head){
      return null
    }

    let deleteNode = null

    // If value is equal to head value
    if(this.head && this.head.value === value){
      deleteNode = this.head
      this.head = this.head.next
    }

    let currentNode = this.head;
    if(currentNode !== null){
      while(currentNode.next){
        if(currentNode.next.value === value){
          deleteNode = currentNode.next
          currentNode.next = currentNode.next.next
        }else{
          currentNode = currentNode.next
        }
      }
    }

    // If value is equal to tail value
    if (this.tail.value === value) {
      this.tail = currentNode;
    }

    return deleteNode
  }

  deleteTail(){
    const deletedTail = this.tail

    if(this.head === this.tail) {
      this.head = null
      this.tail = null

      return deletedTail
    }

    let currentNode = this.head
    while(currentNode){
      if(currentNode.next === deletedTail) {
        currentNode.next = null
        this.tail = currentNode
      }else{
        currentNode = currentNode.next
      }
    }

    /*
      while (currentNode.next) {
        if (!currentNode.next.next) {
          currentNode.next = null;
        } else {
          currentNode = currentNode.next;
        }
      }
      this.tail = currentNode
    */

    return deletedTail
  }

  deleteHead() {
    if(!this.head){
      return null
    }

    const deletedHead = this.head

    if(this.head === this.tail){
      this.head = null
      this.tail = null

      return deletedHead
    }

    this.head = deletedHead.next

    return deletedHead
  }

  find(value) {
    if (!this.head) {
      return null;
    }

    let currentNode = this.head;

    while (currentNode) {
      if (value !== undefined && currentNode.value === value) {
        return currentNode;
      }

      currentNode = currentNode.next;
    }

    return null;
  }

  toArray() {
    let nodes = [];

    let currentNode = this.head
    while(currentNode){
      nodes.push(currentNode)
      currentNode = currentNode.next
    }
  }

  toString(callback){
    this.toArray().map((node) => node.toString(callback)).toString();
  }
}