BEGIN;

INSERT INTO public.users (id, username, email, phone_number, password, created_at, updated_at, deleted_at) 
    VALUES ('user-3e54661e-91ff-4294-8fd6-108da6f3bcf6', 'fahmyabida', 'fahmyabida@gmail.com', '087767618662', '$2a$10$AAkaymNSzoloVfPMpiy9H.GAgAMcNgH.686MnztOi7kbY2U89a37C', '2024-05-28 05:18:11.901069', '2024-05-28 05:18:11.901069', '2024-05-28 05:18:11.901069');

INSERT INTO products ("name", "description", "price", "available_quantity", "reserved_quantity", "sold_quantity", "created_at") 
    VALUES 
    ('Product 1', 'Description for Product 1', 19.99, 100, 0, 0, NOW() - INTERVAL '1 day'),
    ('Product 2', 'Description for Product 2', 29.99, 150, 0, 0, NOW() - INTERVAL '2 day'),
    ('Product 3', 'Description for Product 3', 39.99, 200, 0, 0, NOW() - INTERVAL '3 day'),
    ('Product 4', 'Description for Product 4', 49.99, 120, 0, 0, NOW() - INTERVAL '4 day'),
    ('Product 5', 'Description for Product 5', 59.99, 80, 0, 0, NOW() - INTERVAL '5 day'),
    ('Product 6', 'Description for Product 6', 69.99, 90, 0, 0, NOW() - INTERVAL '6 day'),
    ('Product 7', 'Description for Product 7', 79.99, 110, 0, 0, NOW() - INTERVAL '7 day'),
    ('Product 8', 'Description for Product 8', 89.99, 130, 0, 0, NOW() - INTERVAL '8 day'),
    ('Product 9', 'Description for Product 9', 99.99, 140, 0, 0, NOW() - INTERVAL '9 day'),
    ('Product 10', 'Description for Product 10', 109.99, 160, 0, 0, NOW() - INTERVAL '10 day'),
    ('Product 11', 'Description for Product 11', 119.99, 180, 0, 0, NOW() - INTERVAL '11 day'),
    ('Product 12', 'Description for Product 12', 129.99, 190, 0, 0, NOW() - INTERVAL '12 day'),
    ('Product 13', 'Description for Product 13', 139.99, 170, 0, 0, NOW() - INTERVAL '13 day'),
    ('Product 14', 'Description for Product 14', 149.99, 150, 0, 0, NOW() - INTERVAL '14 day'),
    ('Product 15', 'Description for Product 15', 159.99, 140, 0, 0, NOW() - INTERVAL '15 day'),
    ('Product 16', 'Description for Product 16', 169.99, 130, 0, 0, NOW() - INTERVAL '16 day'),
    ('Product 17', 'Description for Product 17', 179.99, 120, 0, 0, NOW() - INTERVAL '17 day'),
    ('Product 18', 'Description for Product 18', 189.99, 110, 0, 0, NOW() - INTERVAL '18 day'),
    ('Product 19', 'Description for Product 19', 199.99, 100, 0, 0, NOW() - INTERVAL '19 day'),
    ('Product 20', 'Description for Product 20', 209.99, 90, 0, 0, NOW() - INTERVAL '20 day'),
    ('Product 21', 'Description for Product 21', 219.99, 80, 0, 0, NOW() - INTERVAL '21 day'),
    ('Product 22', 'Description for Product 22', 229.99, 70, 0, 0, NOW() - INTERVAL '22 day'),
    ('Product 23', 'Description for Product 23', 239.99, 60, 0, 0, NOW() - INTERVAL '23 day'),
    ('Product 24', 'Description for Product 24', 249.99, 50, 0, 0, NOW() - INTERVAL '24 day'),
    ('Product 25', 'Description for Product 25', 259.99, 40, 0, 0, NOW() - INTERVAL '25 day'),
    ('Product 26', 'Description for Product 26', 269.99, 30, 0, 0, NOW() - INTERVAL '26 day'),
    ('Product 27', 'Description for Product 27', 279.99, 20, 0, 0, NOW() - INTERVAL '27 day'),
    ('Product 28', 'Description for Product 28', 289.99, 10, 0, 0, NOW() - INTERVAL '28 day'),
    ('Product 29', 'Description for Product 29', 299.99, 5, 0, 0, NOW() - INTERVAL '29 day'),
    ('Product 30', 'Description for Product 30', 309.99, 8, 0, 0, NOW() - INTERVAL '30 day'),
    ('Product 31', 'Description for Product 31', 319.99, 11, 0, 0, NOW() - INTERVAL '31 day'),
    ('Product 32', 'Description for Product 32', 329.99, 13, 0, 0, NOW() - INTERVAL '32 day'),
    ('Product 33', 'Description for Product 33', 339.99, 15, 0, 0, NOW() - INTERVAL '33 day'),
    ('Product 34', 'Description for Product 34', 349.99, 17, 0, 0, NOW() - INTERVAL '34 day'),
    ('Product 35', 'Description for Product 35', 359.99, 19, 0, 0, NOW() - INTERVAL '35 day');



COMMIT;