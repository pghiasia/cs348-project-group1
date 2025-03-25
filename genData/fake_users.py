import duckdb 
from duckdb.typing import *
import faker

fake = faker.Faker()

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
	uID INT,
	name VARCHAR,
	DOB DATE,
	language VARCHAR,
	password VARCHAR,
);

INSERT INTO users
SELECT * AS uID, random_name(uid) AS name, random_DOB(uid) AS DOB, random_language(uid) AS language, random_pswd(uid, 12) AS password 
FROM generate_series(1, 150_000);

COPY users TO '../bigData/usersProd.csv';
""")
