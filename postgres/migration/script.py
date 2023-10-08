# pip install faker
from faker import Faker
import random

# Create a Faker instance
fake = Faker()

# Number of rows to generate
num_rows = 1000

# Define the output file name
output_file = '2_data.sql'

existing_names = set()

# Open the output file in write mode
with open(output_file, 'a') as file:
    # Generate and write SQL INSERT statements for each row
    for i in range(num_rows):
        name = fake.name().replace("'", "''")  # Ensure single quotes are escaped
        if name in existing_names:
            continue

        values = f"'{name}'"
        if i == num_rows - 1:
            sql_insert = f"({i+2}, {values});\n"  # Replace 'your_table_name' with your actual table name
        else:
            sql_insert = f"({i+2}, {values}),\n"
        file.write(sql_insert)

print(f"{num_rows} rows of fake data written to {output_file}")