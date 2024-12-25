-- +goose Up

-- +goose StatementBegin
CREATE TABLE UserRoles (
 Id CHAR(36) PRIMARY KEY,
 Name VARCHAR(255) NOT NULL,
 Description MEDIUMTEXT NOT NULL,
 Role INT NOT NULL UNIQUE,
 CreatedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 ModifiedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

 INDEX(CreatedOn)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Users (
 Id CHAR(36) PRIMARY KEY,
 FirstName VARCHAR(255) NOT NULL,
 LastName VARCHAR(255) NOT NULL,
 RoleId CHAR(36) NOT NULL,
 Email VARCHAR(50) NOT NULL UNIQUE,
 PhoneNumber VARCHAR(20),
 Password VARCHAR(255) NOT NULL,
 IsDeprecated TINYINT NOT NULL,
 CreatedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 ModifiedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 
 INDEX(CreatedOn)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Categories (
    Id CHAR(36) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Type INT NOT NULL UNIQUE,
    IsDeprecated TINYINT NOT NULL,
    CreatedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ModifiedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX(CreatedOn)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Products (
    Id CHAR(36) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    Stock INT NOT NULL,
    CategoryId CHAR(36) NOT NULL,
    ImageUrl VARCHAR(255) NULL,
    Status INT NOT NULL,
    IsDeprecated TINYINT NOT NULL,
    CreatedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ModifiedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX(CreatedOn)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Orders (
    Id CHAR(36) PRIMARY KEY,
    UserId CHAR(36) NOT NULL,
    Status INT NOT NULL,
    TotalAmount DECIMAL(10, 2) NOT NULL,
    TaxAmount DECIMAL(10, 2),
    IsDeprecated TINYINT NOT NULL,
    CreatedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ModifiedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX(CreatedOn)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE OrderItems (
    Id CHAR(36) PRIMARY KEY,
    OrderId CHAR(36) NOT NULL,
    ProductId CHAR(36) NOT NULL,
    Quantity INT NOT NULL,
    UnitPrice DECIMAL(10, 2) NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL,
    IsDeprecated TINYINT NOT NULL,
    CreatedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ModifiedOn DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (OrderId) REFERENCES Orders(Id),
    FOREIGN KEY (ProductId) REFERENCES Products(Id),

    INDEX(CreatedOn)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX Idx_Product_Category ON Products(CategoryId)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX Idx_Orders_User ON Orders(UserId)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX Idx_Order_Items_Product ON OrderItems(ProductId)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX Idx_Order_Items_Order ON OrderItems(OrderId)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS UserRoles;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS OrderItems;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS Orders;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS Products;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS Categories;
-- +goose StatementEnd


