### State Pattern

you are tasked with building a DirectionService class for a navigation app. The class calculates the estimated time of arrival (ETA) and provides directions between two points. The ETA and direction differ based on the mode of transportation, which can be one of the following:

1. Walking 
2. Cycling
3. Car
4. Train

Train Coupling and Complex Conditional logic:

-> The DirectionService likely used conditional statements ( if-else or switch-case ) based on transportation mode enums to determine how to calculate ETA and provide directions

-> As the numbers of transportation modes increases, the conditional logic becomes more complex and harder to maintain

Violation of the Open/Closed Principle:

-> Adding new transportation modes (e.g Airplane, Boat) requires modifying the existing DirectionService class, which goes against the Open/Closed Principle ( classes should be open for extension but closed for modification)

Code duplication and Reduced Maintainability:

-> Similar code blocks for different transportation modes may lead to code duplication, making the system less maintainable and more error-prone

Scalability Issues:

-> As more features or transportation modes are added, the class becomes bulky, impacting scalability and readability

# State Pattern: Structure

Structure:

-> Context: Holds a ref to the current state.
-> State: Interface for state-specific behavior
-> Concrete State: Specific implementations of the State interface that represent a particular state of the context object

# State Pattern: Example

UI Navigation

-> Scenario: A mobile app UI where the navigation behavior changes based on whether used is logged in or not


Example:

    -> State: LoggenInState, LoggedOutState
    -> Context: The app's navigation system switches between these states

## State Pattern Use Cases

1. UI Components: Buttons that change behavior based on state (enabled, disabled, pressed)

2. Vending Machines: States like waiting for money, dispensing product, or out of stock

3. TCP Connections: Changing behavior based on connection state (listening, connected, closed)