INSERT INTO customers (id, name, email, created_at, updated_at)
VALUES
    (1, 'John Doe', 'john@example.com', '2024-02-27 12:00:00', '2024-02-27 12:00:00'),
    (2, 'Alice Smith', 'alice@example.com', '2024-02-27 12:05:00', '2024-02-27 12:05:00'),
    (3, 'Bob Johnson', 'bob@example.com', '2024-02-27 12:10:00', '2024-02-27 12:10:00'),
    (4, 'Emily Brown', 'emily@example.com', '2024-02-27 12:15:00', '2024-02-27 12:15:00'),
    (5, 'Michael Lee', 'michael@example.com', '2024-02-27 12:20:00', '2024-02-27 12:20:00');

INSERT INTO subscriptions (id, name, code, description, price, is_active, created_at, updated_at)
VALUES
    (1, 'Basic', 'BASIC', 'Basic subscription plan', 9.99, true, '2024-02-27 12:00:00', '2024-02-27 12:00:00'),
    (2, 'Premium', 'PREMIUM', 'Premium subscription plan', 19.99, true, '2024-02-27 12:05:00', '2024-02-27 12:05:00'),
    (3, 'Pro', 'PRO', 'Pro subscription plan', 29.99, true, '2024-02-27 12:10:00', '2024-02-27 12:10:00'),
    (4, 'Enterprise', 'ENTERPRISE', 'Enterprise subscription plan', 49.99, true, '2024-02-27 12:15:00', '2024-02-27 12:15:00'),
    (5, 'Free Trial', 'TRIAL', 'Free trial subscription plan', 0.00, true, '2024-02-27 12:20:00', '2024-02-27 12:20:00');

INSERT INTO contracts (id, customer_id, subscription_id, billing_interval, installment_amount, duration, duration_unit, status, contract_start_date, contract_end_date)
VALUES
    (1, 1, 2, 3, 46.53, 5, 2, 1, '2024-12-15', '2025-05-15'),
    (2, 1, 4, 2, 98.91, 7, 1, 2, '2025-01-04', '2025-08-04'),
    (3, 2, 5, 4, 20.84, 2, 1, 1, '2025-02-13', '2025-04-13'),
    (4, 3, 1, 2, 12.53, 11, 1, 3, '2024-12-03', '2025-11-03'),
    (5, 5, 3, 4, 92.30, 9, 2, 2, '2024-12-26', '2025-09-26');
