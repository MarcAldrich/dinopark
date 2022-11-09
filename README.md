# Start here
## What is this?

A sample project used to test out some architecture ideas when starting a generic go project. The sample logic in the repository implements a basic restful API built to satisfy the contrived usage case and business requirements included in this readme below.

## Why does it exist?

I was interviewing and doing code tests that all had similar business requirements in the problem prompts. In order to speed up these processes I decided to build a common publicly viewable repository to have some code to discuss. To make the effort more fun for myself I decided to try some new greenfield architecture ideas since I had no previous code to work with.

## Things I was reading leading up to creating this project

REF: https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/

REF: https://github.com/eliben/modlib

REF: https://dave.cheney.net/2016/08/20/solid-go-design

REF: http://www.catb.org/~esr/writings/taoup/html/

REF: https://cloud.google.com/apis/design


---
# Project how to

## How to run
TODO: Fill in

## How to test project/code
TODO: Fill in 


---
# Contrived buisness situation
## Problem Statement and context
It's 1993 and you're appling to become the lead software developer for the new Dino Park! The current software development team would like to see an example of how you would implement software to solve the following project:
```
Park operations needs a system to keep track of the different cages around the park and the different dinosaurs in each one. You'll need to integrate physical cage control and implement a dinosaur registration system into an API for the internal tools team. They are writing a web frontend client to allow the park construction workers to create new cages as dinosaurs are hatched. The Park's science and health teams will also require the ability to view and update the statuses of dinosaurs and cages.
```


## Business Requirements

### Notes on integrating with physical cages

- Cages must have a maximum capacity for how many dinosaurs it can hold.
- Cages know how many dinosaurs are contained.
- Must be able to query a listing of dinosaurs in a specific cage.
- Cages have a power status of ACTIVE or DOWN.
- Cages cannot be powered off if they contain dinosaurs.
- Dinosaurs cannot be moved into a cage that is powered down.
- **!SAFETY REQUIREMENT! Herbivores cannot be in the same cage as carnivores.**
- **!SAFETY REQUIREMENT! Carnivores can only be in a cage with other dinosaurs of the same species.**

### Notes on implementing the dinosaur registration system

- Each dinosaur must have a name.
- Each dinosaur must have a species.
- Each dinosaur is considered an herbivore or a carnivore, depending on its species.
- The science team has been very specific about ensuring we model the dinosaurs following their classification system exactly for easy integration with their existing paper-based tracking system:
  - Carnivores
    - Tyrannosaurus 
    - Velociraptor 
    - Spinosaurus 
    - Megalosaurus
  - Herbivores:
    - Brachiosaurus 
    - Stegosaurus 
    - Ankylosaurus 
    - Triceratops


### Notes on code quality requirements
- When querying dinosaurs or cages they should be filterable on their attributes
  - Use-case example: "list all powered cages"
  - Use-case exmaple: "list all herbivores"
- Pay attention to  HTTP status codes and a response if necessary, representing either the success or error conditions.
  - If unsure on error type follow advice in https://cloud.google.com/apis/design
