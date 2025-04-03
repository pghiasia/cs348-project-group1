-- User Signup.
-- Assume the user we are signing up is called hahafhaha.

-- Check if user exist.
SELECT EXISTS(SELECT * FROM users WHERE name = 'r10b-user') as user_exists;

-- Get new user id and sign up, here we assign a random user_id but in production in would be an incremental new user id.
SELECT COUNT(*) as new_id FROM users;

INSERT INTO users(uid, name, dob, password, language) VALUES (
    54321, 
    'r10b-user', 
    '2000-01-01', 
    'random', 
    'English');

-- Check new user
SELECT * FROM users WHERE name = 'r10b-user';

-- User login with right password
SELECT EXISTS(SELECT * FROM users WHERE name = 'r10b-user' AND password = 'random') AS authenticated;

-- User login with wrong password
SELECT EXISTS(SELECT * FROM users WHERE name = 'r10b-user' AND password = 'random123') AS authenticated;