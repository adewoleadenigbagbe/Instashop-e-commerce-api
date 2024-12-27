INSERT INTO `categories` (`Id`, `Name`, `Description`, `Type`, `IsDeprecated`, `CreatedOn`, `ModifiedOn`) VALUES
	('019407bd-7504-757e-a6c3-7dc98279e434', 'Beauty', 'Beauty and Lifestyle', 3, 0, '2024-12-27 11:50:44', '2024-12-27 11:50:44');
	
INSERT INTO `orders` (`Id`, `UserId`, `Status`, `TotalAmount`, `TaxAmount`, `IsDeprecated`, `CreatedOn`, `ModifiedOn`) VALUES
	('01940781-8c39-7228-9003-2ab60d5783e7', '123e4567-e89b-12d3-a456-426614174000', 1, 119.98, 24.00, 0, '2024-12-27 10:45:18', '2024-12-27 10:45:18'),
	('01940786-1e5d-76a6-bf10-1b1fed84c6e1', '123e4567-e89b-12d3-a456-426614174000', 5, 119.98, 24.00, 0, '2024-12-27 10:50:18', '2024-12-27 14:13:47');

INSERT INTO `orderitems` (`Id`, `OrderId`, `ProductId`, `Quantity`, `UnitPrice`, `TotalPrice`, `IsDeprecated`, `CreatedOn`, `ModifiedOn`) VALUES
	('01940781-8c3d-78e3-b14e-593382e3c500', '01940781-8c39-7228-9003-2ab60d5783e7', '01940364-66ce-75ac-a5d1-4b392b56a7d5', 2, 59.99, 119.98, 0, '2024-12-27 10:45:18', '2024-12-27 10:45:18'),
	('01940786-1e66-7dcc-8e99-e3fce65e6df0', '01940786-1e5d-76a6-bf10-1b1fed84c6e1', '01940364-66ce-75ac-a5d1-4b392b56a7d5', 2, 59.99, 119.98, 0, '2024-12-27 10:50:18', '2024-12-27 10:50:18');

=INSERT INTO `products` (`Id`, `Name`, `Description`, `Price`, `Stock`, `CategoryId`, `ImageUrl`, `Status`, `IsDeprecated`, `CreatedOn`, `ModifiedOn`) VALUES
	('01940364-66ce-75ac-a5d1-4b392b56a7d5', 'Sample Product 2', 'This is a sample product description.', 59.99, 48, 'd5b28b8f-6d4f-4c90-b1a6-f23456789abc', 'https://example.com/images/sample-product.jpg', 1, 0, '2024-12-26 15:34:59', '2024-12-27 14:13:46');

INSERT INTO `userroles` (`Id`, `Name`, `Description`, `Role`, `CreatedOn`, `ModifiedOn`) VALUES
	('01940863-eecb-7b36-87a8-e44abd8e50c5', 'EndUser', 'user role permission', 1, '2024-12-27 14:52:35', '2024-12-27 14:52:35'),
	('01940864-6764-7368-852a-14862f167783', 'Admin', 'admin role permission', 2, '2024-12-27 14:53:05', '2024-12-27 14:53:05');

INSERT INTO `users` (`Id`, `FirstName`, `LastName`, `RoleId`, `Email`, `PhoneNumber`, `Password`, `IsDeprecated`, `CreatedOn`, `ModifiedOn`) VALUES
	('01940870-4f70-7a9b-ae39-54898da8bb78', 'John', 'Doe', '01940863-eecb-7b36-87a8-e44abd8e50c5', 'john.doe@example.com', '+1-202-555-0173', '$2a$10$L7hxBJHQ7uotALDk5xc1WeHp3epDKd1Rt9uUdhUFmxu.LuE2awt6e', 0, '2024-12-27 15:06:06', '2024-12-27 15:06:06'),
	('01940872-7b99-74fc-9343-34e25401554f', 'Jane', 'Smith', '01940864-6764-7368-852a-14862f167783', 'jane.smith@example.com', '+44-789-456-1234', '$2a$10$iEWuvsNlLyNFmrwdowQgXOyiJ846k8m4oOnRRXeX95o2c1ULGAZq2', 0, '2024-12-27 15:08:28', '2024-12-27 15:08:28');
