### Template Method Pattern

Consider a scenario where you have different data parsers (e.g CSV, XML, and JSON). Each parser follows the same steps: open file, parse data, and close file

Without the template Method Pattern, you might end up duplicating the common steps in each parser class

## Problems in our code 

-> Code duplication: The openFile() and closeFile() methods are duplicated in both parsers.

-> Any changes to the common logic would require changes in every parser, ciolating the DRY (Dont Repeat Yourself) principle

## Template Method Pattern

Problem: Different part of an algorithm may need to vary in subclasses, but the overall strcuture should remain consistent

Solution: The Template Method Pattern defines the skeleton of an algorithm in a base class and lets subclasses override specific steps

Structure:
    - Abstract Class: Defines the algorithm skeleton.
    - Concrete Subclasses: Override specific steps of the algorithm