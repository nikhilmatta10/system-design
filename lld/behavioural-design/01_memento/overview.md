### Memento Pattern

Problem: How to provide undo/redo functionality or state restoration without exposing the object internal state and breaking encapsulation.

Solution: The Memento Pattern captures the internal state of an object in a memento, allowing the object to restore its state later on without revealing internal details

Components:

1. Originator: The object whose state needs to be saved and restored (Editor)

2. Memento: Captures and stores the internal state of the originator (Editor Memento)

3. Caretaker: Manages and stores the mementos, without modifying them ( Caretaker )

# Memento Pattern Applications

Undo/Redo in Applications: Commonly used in text editors, drawing applications, or any system that requires history management

Store Restoration: Used in scenarios where you need to periodically save system states (e.g. , games, data recovery) and allow users to return to previous states

Use Cases:
    - Games: Saving the game state for load/reload functionality
    - Document Editors: Undo/redo functionality to navigate through document changes