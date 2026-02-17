### Strategy Pattern

Lets consider a simple payment system where users can pay using different methods like credit cardes or paypal. Without the strategy pattern, you might use if-else conditions to handle the different payments methods, leading to less maintainable and flexible code 

# Problems in code 

-> The PaymentService class has multiple responsibilities ( deciding the payment type and processing it )

-> Adding a new payment method requires modifying the PaymentService class

-> The use of if-else conditions can make the code harder to maintain as more payment are added

With the strategy Pattern, the logic for each payment type is encapulated in seperate strategy classes, and the PaymentService (context class) delegates the task of payment processing to one of these strrategies at runtime

