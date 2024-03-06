CREATE DATABASE TastyBitesDB;

USE TastyBitesDB;

CREATE TABLE Employees (
    EmployeeID INT AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR(20) NOT NULL,
    LastName VARCHAR(20) NOT NULL,
    Position VARCHAR(20) NOT NULL
);

INSERT INTO Employees(FirstName, LastName, Position)
VALUES ('John', 'Doe', 'Waiter');

CREATE TABLE MenuItems (
    ItemID INT AUTO_INCREMENT PRIMARY KEY,
    ItemName VARCHAR(20) NOT NULL,
    ItemDescription TEXT NOT NULL,
    Price DECIMAL(10,2) NOT NULL,
    Category VARCHAR(20) NOT NULL
);

INSERT INTO MenuItems(ItemName, ItemDescription, Price, Category)
VALUES ('Steak', 'Grilled sirloin steak', 25.99, 'Main Course');

CREATE TABLE Orders (
    OrderID INT AUTO_INCREMENT PRIMARY KEY,
    TableNumber INT NOT NULL,
    EmployeeID INT,
    OrderDate DATE NOT NULL,
    OrderStatus VARCHAR(20) NOT NULL,
    FOREIGN KEY (EmployeeID) REFERENCES Employees(EmployeeID)
);

INSERT INTO Orders(TableNumber, EmployeeID, OrderDate, OrderStatus)
VALUES (5, 1, '2023-08-04', 'Pending');

CREATE TABLE OrderItems (
    OrderItemID INT AUTO_INCREMENT PRIMARY KEY,
    Quantity INT NOT NULL,
    SubTotal DECIMAL(10,2) NOT NULL,
    OrderID INT,
    ItemID INT,
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ItemID) REFERENCES MenuItems(ItemID)
);

INSERT INTO OrderItems(Quantity, SubTotal, OrderID, ItemID)
VALUES (2, 51.98, 1, 1);

CREATE TABLE Payments (
    PaymentID INT AUTO_INCREMENT PRIMARY KEY,
    OrderID INT,
    PaymentDate DATE NOT NULL,
    PaymentMethod VARCHAR(20),
    TotalAmount DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID)
);

INSERT INTO Payments (OrderID, PaymentDate, PaymentMethod, TotalAmount)
VALUES (1, '2023-08-04', 'Credit Card', 51.98);

-- Update 
UPDATE Employees
SET Salary = 11000000
WHERE EmployeeID = 1;

-- Delete 
DELETE FROM OrderItems
WHERE OrderItemID = 1;

-- Conditional Insertion
INSERT INTO Employees (EmployeeID, FirstName, LastName, Position)
SELECT * FROM (SELECT 123, 'John', 'Doe', 'Manager') AS tmp
WHERE NOT EXISTS (
    SELECT EmployeeID FROM Employees WHERE EmployeeID = 123
) LIMIT 1;

-- Data update with Join and Aggregation
UPDATE Orders
JOIN (
    SELECT OrderID, SUM(SubTotal) AS TotalSubTotal
    FROM OrderItems
    GROUP BY OrderID
) AS SubTotals ON Orders.OrderID = SubTotals.OrderID
SET Orders.TotalAmount = SubTotals.TotalSubTotal;

-- Soft Delete
ALTER TABLE Employees
ADD COLUMN IsDeleted BOOLEAN DEFAULT FALSE;

UPDATE Employees
SET IsDeleted = TRUE
WHERE EmployeeID = 1;

SELECT *
FROM Employees
WHERE IsDeleted = FALSE;

UPDATE Employees
SET IsDeleted = FALSE
WHERE EmployeeID = 1;

-- Bulk Update with Case statement
UPDATE Employees
SET Position = 
    CASE 
        WHEN EmployeeID = 1 THEN 'Manager'
        WHEN EmployeeID = 2 THEN 'Supervisor'
        ELSE 'Staff'
    END
WHERE EmployeeID IN (1, 2);

-- Data De-duplication 
DELETE FROM Employees
WHERE EXISTS (
    SELECT 1
    FROM Employees AS innerEmployees
    WHERE innerEmployees.EmployeeID < Employees.EmployeeID
    AND innerEmployees.FirstName = Employees.FirstName
    AND innerEmployees.LastName = Employees.LastName
    AND innerEmployees.Position = Employees.Position
);