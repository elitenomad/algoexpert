# Python programming language cheat sheet

## DATA TYPES

1. Int (1, 2, 3000 etc...)
2. Float (1.0, 3.14 etc...)
3. Bool (True or False)
4. String ("Name", "Text"...)

## COMMENTS

Two ways we can add comments one is using ``` and another is """


'''
THIS IS A COMMENT
'''

"""
THIS IS A COMMENT AS WELL
"""

## VARIABLES

1. should not start with number
2. should not contain special chars except `underscore` (_)
3. should not contain any kind of spaces


## CONSOLE OUTPUT

If you have to log information onto the console, language provides a `print` statement.

```
name = input("what is your name ?")
print(name)
```

## ARITHMETIC OPERATORS

> Available operators in the language are +, -, *, /, //, %, **

1. Integer division (//)
2. Modulus Operation (%)
3. Concatenaion (+ operator is polymorphic, if operands are integers it provides sum of the values and if operators are strings it provides concatenated string)
4. Float division (/)
5. Exponent (**)
6. Multiplication (*)

## TYPE CONVERSION

One data type can be converted to another

```
a = "42"
b = int(a)

print(a, type(a))
print(b, type(b))

print(bool(""))


int("pranava") # Will raise an exception 
```

## CONDITIONS

- Conditional is an expression which evaluates to True or False
- `>, <, >=, <=, == != ` are the operators used to evaluate the expressions
- e.g 2 == 2 (True),  3 < 2 (False)
- ASCII, ord('a') provides ascii value, char(97) returns the ascii char

## COMPOUND CONDITIONS

- And, Or, Not

## CONDITIONALS

- IF
- ELIF
- Else

## DATA STRUCTURES

### LIST

- DS that stores ordered collection of elements. Can have elements that belong to different data types.
- `in` keyword is used to verify if an element is present in the collection.

### STRINGS

### TUPLES
- Once the tuple if created, we cannot mutate the collection and doing so will raise an exception

### SET
- Is not a ordered collection
- Unique elements

### Dict
- ds with `key=value` pairs.


## FOR LOOP
- `for name in ["pranava", "balugari"]:`
- `for i in range(start, end, step):`
  - start is inclusive
  - end is exclusive

## WHILE LOOPS
- `while a < 2:` 
- Another form of loop to make life easier

## SLICING

```
student_names = ["Tim", "Antoine", "Clement", "Simon", "Alex"]
good_students = student_names[:2]
bad_students = student_names[2:]

reversed_bad_students = bad_students[::-1]

print(f"Good students: {good_students}")
print(f"Bad students: {bad_students}")
print(f"Reversed Bad students: {reversed_bad_students}")

best_student = student_names[0]
worst_student = student_names[len(student_names) - 1]

print(f"Best student: {best_student}")
print(f"Worst student: {worst_student}")
```

## EXCEPTIONS
```
# Welcome to our Python playground!

digits = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

try:
    if len(digits) != 10:
        raise Exception("There should be exactly 10 digits!")
except Exception as e:
    print(e)
finally:
    print("Goodbye!")
```

## FUNCTIONS
- Arguments and Parameters

## MUTABILITY
- Immutable
  - Int
  - Float
  - String
  - Bool
  - Set
- Mutable
  - List
  - Set
  - Dict

## MATH module
- Look into `os, math and threading` modules
- Look into `random` module
  - random.randint(start, stop)
  - random.randrange(start, stop, step)
  - random.choice(iterable)

## SORTING
- strings.sort() - Mutates original collection
- sorted(strings, reverse = True) - Will not mutate original collection

## MISCELLANIOUS

```
# Welcome to our Python playground!

digits = [i for i in range(10)]
print(digits)

even_numbers = [i for i in range(20) if i % 2 == 0]
print(even_numbers)

divisible_by_3 = {i: i % 3 == 0 for i in range(30)}
print(divisible_by_3)

x, y = 13, "hello"
print(f"x = {x}, y = {y}")

x = y = 10
print(x, y)
```

