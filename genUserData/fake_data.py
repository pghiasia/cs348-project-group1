import duckdb 
from duckdb.typing import *
import faker

fake = faker.Faker()

def generate_user(seed):
    user = {
        'uID': fake.random_int(1, 1000),
        'name': fake.name,
        'DOB': fake.date_of_birth(),
        'language': fake.language_name(),
        'password': fake.password(length=12),
    }
    return user


def random_name(seed):
    return fake.unique.user_name()

def random_birth(seed):
    return fake.date_of_birth()

def random_lang(seed):
    return fake.language_name();

def random_pswd(seed, len):
    return fake.password(len);

con = duckdb.connect()

con.create_function(
    'random_name',
    random_name,
    [DOUBLE],
    VARCHAR
)

con.create_function(
    'random_DOB',
    random_birth,
    [DOUBLE],
    VARCHAR
)

con.create_function(
    'random_language',
    random_lang,
    [DOUBLE],
    VARCHAR
)

con.create_function(
    'random_pswd',
    random_pswd,
    [DOUBLE, INTEGER],
    VARCHAR
)


con.sql("""
CREATE TABLE users(
	uID INT PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE,
	DOB DATE NOT NULL,
	language VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
);

INSERT INTO users
SELECT * AS uID, random_name(uid) AS name, random_DOB(uid) AS DOB, random_language(uid) AS language, random_pswd(uid, 12) AS password FROM generate_series(1, 500);

COPY users TO 'usersProd.csv';
""")
