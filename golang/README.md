# The Bowling Game

Exercise from Uncle Bob's website: http://butunclebob.com/ArticleS.UncleBob.TheBowlingGameKata.

I cover the following test cases in this kata

- A Gutter Game will always return a score of 0
- Knock down one pin each roll, (score = 20)
- Spare in first roll, one pin down in each other roll, (score = 11(spare) + 18 + 1 = 30)
- Spare in last frame, one pin down in each other roll, (score = 18 + 10  = 28)
- Strike in first roll, one pin down in each other roll, (score = 12 + 19 = 31)
- Strike in last frame, one pin down in each other roll, (score = 18 + 10 = 28)
- A Double Strike, then one pin down in each other roll, (score = 22 + 12 + 17 = 51)
- A Turkey, then one pin down in each other roll, (score = 30 + 22 + 12 + 15 = 79)
- A Four Bagger, then one pin down in each other roll, (score = 30 + 30 + 22 + 12 + 13 = 107)
- Strikes in each of the first nine frames, and two in the tenth (score = 300)

What's next? I'd like to introduce a mutation testing example to this kata.

## Summary

- [Prerequisites](#Prerequisites)
- [Runing the tests](#running-the-tests)
- [Authors](#authors)

### Prerequisites

Install Go on your environment. For Mac users, via:

    brew install go

## Running the tests

    go test


## Authors

- [Apastasalad](https://github.com/apastasalad)
