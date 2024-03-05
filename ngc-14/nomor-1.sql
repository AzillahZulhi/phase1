CREATE TABLE products (
    productID INT PRMARY KEY,
    productName VARCHAR(200),
    productCode VARCHAR(20),
    price DECIMAL(10,2)
);

INSERT INTO products (productID, productName, productCode, price)
VALUES
    (1, 'Laptop', 'LT001', 799.99),
    (2, 'SmartPhone', 'SP002', 499.99),
    (3, 'Tablet', 'TB003', 299.99),
    (4, 'HandPhone', 'HP004', 149.99),
    (5, 'Monitor', 'MN005', 399.99);

CREATE TABLE productSales (
    saleID INT PRMARY KEY,
    productID INT,
    saleDate DATE,
    quantitySold INT,
    FOREIGN KEY (productID) REFERENCES products(productID)
);

INSERT INTO productSales (saleID, productID, saleDate, quantitySold)
VALUES
    (1,1, '2023-08-01',10),
    (2,1, '2023-08-02',5),
    (3,2, '2023-08-03',8),
    (4,3, '2023-08-03',12),
    (5,3, '2023-08-04',6),
    (6,4, '2023-08-04',20),
    (7,5, '2023-08-05',15);

-- nomor 1
SELECT 
    CASE 
        WHEN price < 300 THEN 'Low Price'
        WHEN price >= 300 AND price <= 600 THEN 'Medium Price'
        ELSE 'High Price'
    END AS Price_Category,
    COUNT(*) AS Product_Count
FROM 
    products
GROUP BY 
    CASE 
        WHEN price < 300 THEN 'Low Price'
        WHEN price >= 300 AND price <= 600 THEN 'Medium Price'
        ELSE 'High Price'
    END;

-- nomor 2
SELECT 
    p.productName,
    SUM(ps.quantitySold) AS Total_Quantity_Sold
FROM 
    products p
JOIN 
    productSales ps ON p.productID = ps.productID
GROUP BY 
    p.productID
ORDER BY 
    Total_Quantity_Sold DESC
LIMIT 3;

-- nomor 3
SELECT 
    p.productName,
    SUM(CASE WHEN ps.saleDate >= DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH) THEN ps.quantitySold ELSE 0 END) AS Current_Period_Quantity,
    SUM(CASE WHEN ps.saleDate >= DATE_SUB(DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH), INTERVAL 1 MONTH) AND ps.saleDate < DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH) THEN ps.quantitySold ELSE 0 END) AS Previous_Period_Quantity,
    ROUND(
        IFNULL(
            100 * (
                (SUM(CASE WHEN ps.saleDate >= DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH) THEN ps.quantitySold ELSE 0 END) /
                NULLIF(SUM(CASE WHEN ps.saleDate >= DATE_SUB(DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH), INTERVAL 1 MONTH) AND ps.saleDate < DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH) THEN ps.quantitySold ELSE 0 END), 0)) - 1
            ), 0
        ),
        2
    ) AS Sales_Growth_Percentage
FROM 
    products p
LEFT JOIN 
    productSales ps ON p.productID = ps.productID
GROUP BY 
    p.productID, p.productName;

-- nomor 4
SELECT 
    p1.productID,
    p1.productName,
    p1.price,
    AVG(p1.price - COALESCE(p2.price, 0)) AS Avg_Price_Difference
FROM 
    products p1
LEFT JOIN 
    products p2 ON p1.productID = p2.productID + 1
WHERE 
    p1.productID != (SELECT MIN(productID) FROM products)
GROUP BY 
    p1.productID, p1.productName, p1.price;

-- nomor 5
SELECT 
    productName,
    productCode
FROM 
    products
WHERE 
    price IS NULL;
