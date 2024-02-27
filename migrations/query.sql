INSERT INTO customers (id, name, email, created_at, updated_at)
VALUES
    (1, 'John Doe', 'john@example.com', '2024-02-27 12:00:00', '2024-02-27 12:00:00'),
    (2, 'Alice Smith', 'alice@example.com', '2024-02-27 12:05:00', '2024-02-27 12:05:00'),
    (3, 'Bob Johnson', 'bob@example.com', '2024-02-27 12:10:00', '2024-02-27 12:10:00'),
    (4, 'Emily Brown', 'emily@example.com', '2024-02-27 12:15:00', '2024-02-27 12:15:00'),
    (5, 'Michael Lee', 'michael@example.com', '2024-02-27 12:20:00', '2024-02-27 12:20:00');

INSERT INTO subscriptions (id, name, code, description, price, duration, duration_unit, is_active, created_at, updated_at)
VALUES
    (1, 'Basic', 'BASIC', 'Basic subscription plan', 9.99, 30, 1, true, '2024-02-27 12:00:00', '2024-02-27 12:00:00'),
    (2, 'Premium', 'PREMIUM', 'Premium subscription plan', 19.99, 60, 1, true, '2024-02-27 12:05:00', '2024-02-27 12:05:00'),
    (3, 'Pro', 'PRO', 'Pro subscription plan', 29.99, 90, 1, true, '2024-02-27 12:10:00', '2024-02-27 12:10:00'),
    (4, 'Enterprise', 'ENTERPRISE', 'Enterprise subscription plan', 49.99, 180, 1, true, '2024-02-27 12:15:00', '2024-02-27 12:15:00'),
    (5, 'Free Trial', 'TRIAL', 'Free trial subscription plan', 0.00, 7, 1, true, '2024-02-27 12:20:00', '2024-02-27 12:20:00');

INSERT INTO contracts (id, customer_id, subscription_id, billing_interval, status, contract_start_date, contract_end_date, created_at, updated_at)
VALUES
    (1, 1, 1, 1, 1, '2024-02-27', '2024-03-27', '2024-02-27 12:00:00', '2024-02-27 12:00:00'),
    (2, 2, 2, 1, 1, '2024-02-28', '2024-03-28', '2024-02-27 12:05:00', '2024-02-27 12:05:00'),
    (3, 3, 3, 1, 1, '2024-02-29', '2024-03-29', '2024-02-27 12:10:00', '2024-02-27 12:10:00'),
    (4, 4, 4, 1, 1, '2024-03-01', '2024-03-30', '2024-02-27 12:15:00', '2024-02-27 12:15:00'),
    (5, 5, 5, 1, 1, '2024-03-02', '2024-03-31', '2024-02-27 12:20:00', '2024-02-27 12:20:00');
