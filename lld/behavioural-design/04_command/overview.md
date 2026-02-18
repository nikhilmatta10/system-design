### Command Pattern

Imagine you are developing a basic text editor with buttons for bold, italic and underline text formatiing

Without the command pattern, the buttons directly interact with the TextEditor class, and you haved end up hardcoding behavior into the UI classes, making them tightly coupled 

## Problems in code

-> Each button class is tightly coupled with the TextEditor 

-> Its harder to extend with new commands or add functionality such as undo/redo or logging

### Command Pattern

By introducing the Command Pattern, we can decouple the actions ( bold, italic, underline) from the UI components ( buttons ), making the design more flexible and maintainable. The buttons no longer need to know about the editor directly instead work with generic Command Objects 

### Command Pattern Structure

Structure:

    - Command: Interface for executing operations
    - Invoker: Sends the command
    - Receiver: Performs the operation

### Command Pattern Benefits

    - Decoupling of Invoker and Receiver: The button (invoker) doesnt know the details of the TextEditor ( receiver), making the system more flexible and resusable

    - Command History and Undo: Commands can be logged for undo/redo functionality

    - Task Queeuing: Commands can be stored in a queue and executed later, making it useful for task scheduling

    - Extensibility: New commands can be added easily without modifying existing code. For example, adding a ChangeColorCommand only requires creating a new command class

### Command Pattern Usecases

    - Gui Applications:
        Commands can be associated with buttons, menus and keyboard shortcuts in applications like text editors, spreadsheets, or drawing software

    - Task Scheduling:
        Commands can be placed in a queue and executed later, useful in batch processing or deffered task execution

    - Undo/Redo Functionality:
        Commands can be stored and rolled back to provide undo and redo capabilities, especially in applications like IDEs, word processors, or graphics software

    - Macro Recording:
        Actions performed by the user can be recorded as a series of commands, which can then be played back as macros