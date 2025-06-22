DELETE FROM sessions;

INSERT INTO sessions (
    id, user_id, device_type, device_name, app_type, app_version,
    os, os_version, device_model, ip_address, city, country,
    is_active, lifetime, last_active_at, created_at, updated_at
) VALUES
-- Мобильные Android сессии (device_type = 1 - Mobile, app_type = 1 - ChesshubMobile)
(gen_random_uuid(), 1, 1, 'Samsung Galaxy S23', 1, '2.1.0', 'android', '14', 'SM-S918B', '192.168.1.10', 'Moscow', 'Russia', true, '30 days', NOW() - INTERVAL '5 minutes', NOW() - INTERVAL '2 days', NOW() - INTERVAL '5 minutes'),
(gen_random_uuid(), 1, 1, 'Xiaomi Redmi Note 12', 1, '2.0.5', 'android', '13', 'Redmi Note 12', '192.168.1.11', 'Moscow', 'Russia', true, '7 days', NOW() - INTERVAL '1 hour', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 hour'),
(gen_random_uuid(), 2, 1, 'Google Pixel 7', 1, '2.1.0', 'android', '14', 'Pixel 7', '10.0.0.15', 'Saint Petersburg', 'Russia', true, '14 days', NOW() - INTERVAL '30 minutes', NOW() - INTERVAL '3 hours', NOW() - INTERVAL '30 minutes'),
(gen_random_uuid(), 3, 1, 'OnePlus 11', 1, '2.0.8', 'android', '13', 'CPH2449', '172.16.0.5', 'Kazan', 'Russia', true, '30 days', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '1 day', NOW() - INTERVAL '2 hours'),

-- Мобильные iOS сессии (device_type = 1 - Mobile, app_type = 1 - ChesshubMobile)
(gen_random_uuid(), 2, 1, 'iPhone 15 Pro', 1, '2.1.0', 'ios', '17.2', 'iPhone16,1', '192.168.0.25', 'New York', 'USA', true, '30 days', NOW() - INTERVAL '15 minutes', NOW() - INTERVAL '6 hours', NOW() - INTERVAL '15 minutes'),
(gen_random_uuid(), 4, 1, 'iPhone 14', 1, '2.0.9', 'ios', '17.1', 'iPhone15,3', '10.1.1.20', 'Los Angeles', 'USA', true, '14 days', NOW() - INTERVAL '3 hours', NOW() - INTERVAL '2 days', NOW() - INTERVAL '3 hours'),

-- Планшетные сессии (device_type = 3 - Tablet, app_type = 3 - ChesshubTablet)
(gen_random_uuid(), 5, 3, 'iPad Pro', 3, '2.1.0', 'ios', '17.2', 'iPad14,3', '192.168.2.30', 'London', 'UK', true, '7 days', NOW() - INTERVAL '1 day', NOW() - INTERVAL '3 days', NOW() - INTERVAL '1 day'),
(gen_random_uuid(), 15, 3, 'Samsung Galaxy Tab S9', 3, '2.1.0', 'android', '14', 'SM-X910', '192.168.3.40', 'Zurich', 'Switzerland', true, '14 days', NOW() - INTERVAL '5 hours', NOW() - INTERVAL '3 days', NOW() - INTERVAL '5 hours'),
(gen_random_uuid(), 16, 3, 'iPad Air', 3, '2.0.8', 'ios', '17.1', 'iPad13,16', '10.3.3.25', 'Brussels', 'Belgium', true, '7 days', NOW() - INTERVAL '1 day', NOW() - INTERVAL '4 days', NOW() - INTERVAL '1 day'),

-- Web сессии (device_type = 0 - Web, app_type = 0 - ChesshubWeb)
(gen_random_uuid(), 1, 0, 'Chrome', 0, '2.1.0', 'windows', '11', NULL, '203.0.113.45', 'Berlin', 'Germany', true, '1 day', NOW() - INTERVAL '10 minutes', NOW() - INTERVAL '4 hours', NOW() - INTERVAL '10 minutes'),
(gen_random_uuid(), 3, 0, 'Firefox', 0, '2.0.7', 'macos', '14.2', NULL, '198.51.100.10', 'Paris', 'France', true, '3 days', NOW() - INTERVAL '45 minutes', NOW() - INTERVAL '1 day', NOW() - INTERVAL '45 minutes'),
(gen_random_uuid(), 6, 0, 'Safari', 0, '2.1.0', 'macos', '14.3', NULL, '203.0.113.100', 'Tokyo', 'Japan', true, '7 days', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '12 hours', NOW() - INTERVAL '2 hours'),
(gen_random_uuid(), 7, 0, 'Edge', 0, '2.0.9', 'windows', '10', NULL, '192.0.2.50', 'Sydney', 'Australia', true, '14 days', NOW() - INTERVAL '6 hours', NOW() - INTERVAL '2 days', NOW() - INTERVAL '6 hours'),
(gen_random_uuid(), 12, 0, 'Opera', 0, '2.1.0', 'linux', 'Ubuntu 22.04', NULL, '192.0.2.90', 'Prague', 'Czech Republic', true, '3 days', NOW() - INTERVAL '1 hour', NOW() - INTERVAL '6 hours', NOW() - INTERVAL '1 hour'),
(gen_random_uuid(), 18, 0, 'Brave', 0, '2.1.0', 'freebsd', '14.0', NULL, '37.139.20.5', 'Helsinki', 'Finland', true, '1 day', NOW() - INTERVAL '12 hours', NOW() - INTERVAL '2 days', NOW() - INTERVAL '12 hours'),

-- Desktop сессии (device_type = 2 - Desktop, app_type = 2 - ChesshubDesktop)
(gen_random_uuid(), 8, 2, 'ChessHub Desktop', 2, '2.1.0', 'windows', '11', NULL, '203.0.113.60', 'Berlin', 'Germany', true, '30 days', NOW() - INTERVAL '20 minutes', NOW() - INTERVAL '1 day', NOW() - INTERVAL '20 minutes'),
(gen_random_uuid(), 9, 2, 'ChessHub Desktop', 2, '2.0.8', 'macos', '14.2', NULL, '198.51.100.15', 'Paris', 'France', true, '14 days', NOW() - INTERVAL '1 hour', NOW() - INTERVAL '6 hours', NOW() - INTERVAL '1 hour'),
(gen_random_uuid(), 10, 2, 'ChessHub Desktop', 2, '2.1.0', 'linux', 'Ubuntu 22.04', NULL, '203.0.113.70', 'Stockholm', 'Sweden', true, '7 days', NOW() - INTERVAL '3 hours', NOW() - INTERVAL '2 days', NOW() - INTERVAL '3 hours'),

-- Неактивные сессии разных типов
(gen_random_uuid(), 1, 1, 'Samsung Galaxy S21', 1, '1.9.5', 'android', '12', 'SM-G991B', '192.168.1.12', 'Moscow', 'Russia', false, '7 days', NOW() - INTERVAL '10 days', NOW() - INTERVAL '15 days', NOW() - INTERVAL '10 days'),
(gen_random_uuid(), 2, 1, 'iPhone 13', 1, '1.8.3', 'ios', '16.5', 'iPhone14,5', '10.0.0.16', 'Saint Petersburg', 'Russia', false, '14 days', NOW() - INTERVAL '20 days', NOW() - INTERVAL '25 days', NOW() - INTERVAL '20 days'),
(gen_random_uuid(), 4, 0, 'Chrome', 0, '1.9.0', 'windows', '10', NULL, '203.0.113.46', 'Berlin', 'Germany', false, '1 day', NOW() - INTERVAL '5 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '5 days'),
(gen_random_uuid(), 5, 2, 'ChessHub Desktop', 2, '1.8.5', 'windows', '10', NULL, '203.0.113.80', 'Vienna', 'Austria', false, '30 days', NOW() - INTERVAL '45 days', NOW() - INTERVAL '50 days', NOW() - INTERVAL '45 days'),

-- Старые версии приложений
(gen_random_uuid(), 8, 1, 'Huawei P40', 1, '1.5.2', 'android', '10', 'ELS-NX9', '172.16.1.25', 'Barcelona', 'Spain', true, '30 days', NOW() - INTERVAL '1 hour', NOW() - INTERVAL '30 days', NOW() - INTERVAL '1 hour'),
(gen_random_uuid(), 9, 1, 'iPhone 12 Mini', 1, '1.7.1', 'ios', '15.8', 'iPhone13,1', '10.2.2.15', 'Rome', 'Italy', true, '14 days', NOW() - INTERVAL '4 hours', NOW() - INTERVAL '20 days', NOW() - INTERVAL '4 hours'),
(gen_random_uuid(), 11, 0, 'Firefox', 0, '1.6.0', 'windows', '8.1', NULL, '203.0.113.90', 'Madrid', 'Spain', true, '3 days', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '60 days', NOW() - INTERVAL '2 hours'),

-- Экзотические устройства и локации
(gen_random_uuid(), 10, 1, 'Sony Xperia 1 V', 1, '2.1.0', 'android', '13', 'XQ-DQ54', '198.51.100.75', 'Stockholm', 'Sweden', true, '30 days', NOW() - INTERVAL '20 minutes', NOW() - INTERVAL '8 hours', NOW() - INTERVAL '20 minutes'),
(gen_random_uuid(), 11, 1, 'Nothing Phone 2', 1, '2.0.8', 'android', '14', 'A065', '203.0.113.80', 'Amsterdam', 'Netherlands', true, '7 days', NOW() - INTERVAL '3 hours', NOW() - INTERVAL '1 day', NOW() - INTERVAL '3 hours'),
(gen_random_uuid(), 13, 1, 'Motorola Edge 40', 1, '2.0.9', 'android', '13', 'XT2301-4', '85.140.1.50', 'Warsaw', 'Poland', true, '14 days', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '5 days', NOW() - INTERVAL '2 hours'),
(gen_random_uuid(), 14, 1, 'Oppo Find X6', 1, '2.1.0', 'android', '14', 'CPH2449', '91.198.174.100', 'Vienna', 'Austria', true, '30 days', NOW() - INTERVAL '30 minutes', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '30 minutes'),

-- Планшеты Android
(gen_random_uuid(), 17, 3, 'Lenovo Tab P12 Pro', 3, '2.0.7', 'android', '13', 'TB-Q706F', '46.105.57.169', 'Copenhagen', 'Denmark', true, '30 days', NOW() - INTERVAL '8 hours', NOW() - INTERVAL '1 week', NOW() - INTERVAL '8 hours'),
(gen_random_uuid(), 19, 3, 'Amazon Fire HD 10', 3, '2.0.5', 'android', '11', 'KFTRWI', '203.0.113.120', 'Dublin', 'Ireland', true, '14 days', NOW() - INTERVAL '12 hours', NOW() - INTERVAL '3 days', NOW() - INTERVAL '12 hours'),

-- Desktop приложения разных OS
(gen_random_uuid(), 20, 2, 'ChessHub Desktop', 2, '2.1.0', 'windows', '11', NULL, '192.0.2.100', 'Oslo', 'Norway', true, '30 days', NOW() - INTERVAL '30 minutes', NOW() - INTERVAL '4 hours', NOW() - INTERVAL '30 minutes'),
(gen_random_uuid(), 21, 2, 'ChessHub Desktop', 2, '2.0.9', 'macos', '14.3', NULL, '142.250.191.78', 'Reykjavik', 'Iceland', true, '14 days', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '1 day', NOW() - INTERVAL '2 hours'),
(gen_random_uuid(), 22, 2, 'ChessHub Desktop', 2, '2.1.0', 'linux', 'Fedora 39', NULL, '17.253.144.10', 'Lisbon', 'Portugal', true, '7 days', NOW() - INTERVAL '1 hour', NOW() - INTERVAL '8 hours', NOW() - INTERVAL '1 hour'),

-- Очень старые сессии
(gen_random_uuid(), 19, 1, 'Samsung Galaxy S20', 1, '1.2.0', 'android', '11', 'SM-G980F', '192.168.4.15', 'Oslo', 'Norway', false, '7 days', NOW() - INTERVAL '60 days', NOW() - INTERVAL '90 days', NOW() - INTERVAL '60 days'),
(gen_random_uuid(), 20, 0, 'Internet Explorer', 0, '1.0.0', 'windows', '7', NULL, '203.0.113.200', 'Reykjavik', 'Iceland', false, '1 day', NOW() - INTERVAL '365 days', NOW() - INTERVAL '400 days', NOW() - INTERVAL '365 days'),

-- Недавно созданные сессии
(gen_random_uuid(), 21, 1, 'Google Pixel 8 Pro', 1, '2.1.0', 'android', '14', 'Pixel 8 Pro', '142.250.191.78', 'Dublin', 'Ireland', true, '30 days', NOW() - INTERVAL '1 minute', NOW() - INTERVAL '5 minutes', NOW() - INTERVAL '1 minute'),
(gen_random_uuid(), 22, 1, 'iPhone 15', 1, '2.1.0', 'ios', '17.2', 'iPhone15,4', '17.253.144.10', 'Lisbon', 'Portugal', true, '14 days', NOW() - INTERVAL '30 seconds', NOW() - INTERVAL '2 minutes', NOW() - INTERVAL '30 seconds'),

-- Разные lifetime периоды
(gen_random_uuid(), 23, 1, 'Asus ROG Phone 7', 1, '2.0.9', 'android', '13', 'AI2205', '61.216.2.164', 'Budapest', 'Hungary', true, '1 hour', NOW() - INTERVAL '10 minutes', NOW() - INTERVAL '30 minutes', NOW() - INTERVAL '10 minutes'),
(gen_random_uuid(), 24, 0, 'Vivaldi', 0, '2.1.0', 'macos', '14.3', NULL, '185.199.108.153', 'Ljubljana', 'Slovenia', true, '12 hours', NOW() - INTERVAL '2 hours', NOW() - INTERVAL '8 hours', NOW() - INTERVAL '2 hours'),

-- Разные user_id для тестирования группировки
(gen_random_uuid(), 1, 1, 'Realme GT 3', 1, '2.1.0', 'android', '14', 'RMX3708', '192.168.5.50', 'Bratislava', 'Slovakia', true, '30 days', NOW() - INTERVAL '3 hours', NOW() - INTERVAL '1 day', NOW() - INTERVAL '3 hours'),
(gen_random_uuid(), 1, 0, 'Chrome', 0, '2.1.0', 'windows', '11', NULL, '192.168.5.51', 'Bratislava', 'Slovakia', true, '1 day', NOW() - INTERVAL '1 hour', NOW() - INTERVAL '6 hours', NOW() - INTERVAL '1 hour'),
(gen_random_uuid(), 1, 2, 'ChessHub Desktop', 2, '2.1.0', 'windows', '11', NULL, '192.168.5.52', 'Bratislava', 'Slovakia', true, '30 days', NOW() - INTERVAL '5 hours', NOW() - INTERVAL '12 hours', NOW() - INTERVAL '5 hours'),
(gen_random_uuid(), 25, 1, 'Vivo X90 Pro', 1, '2.0.8', 'android', '13', 'V2242A', '123.125.114.144', 'Zagreb', 'Croatia', true, '14 days', NOW() - INTERVAL '4 hours', NOW() - INTERVAL '2 days', NOW() - INTERVAL '4 hours');
