Functinality

# Add RSS feeds from across the internet to be collected

Store the collected posts in a PostgreSQL database

Follow and unfollow RSS feeds that other users have added

View summaries of the aggregated posts in the terminal, with a link to the full post


# This a CLI application, and like many CLI applications, it has a set of valid commands. For example:

gator login - sets the current user in the config

gator register - adds a new user to the database

gator users - lists all the users in the database

# Connect the postgresql using Mac/Linux(WSL)

Mac OS with brew

brew install postgresql@16

Linux / WSL (Debian). Here are the docs from Microsoft, but simply:

sudo apt update

sudo apt install postgresql postgresql-contrib

Ensure the installation worked. The psql command-line utility is the default client for Postgres. Use it to make sure you're on version 16+ of Postgres:

psql --version

(Linux only) Update postgres password:

sudo passwd postgres


Enter a password, and be sure you won't forget it. You can just use something easy like postgres.


Start the Postgres server in the background

Mac: brew services start postgresql

Linux: sudo service postgresql start

Connect to the server. I recommend simply using the psql client. It's the "default" client for Postgres, and it's a great way to interact with the database. While it's not as user-friendly as a GUI like PGAdmin, 
it's a great tool to be able to do at least basic operations with.

Enter the psql shell:



Mac: psql postgres

Linux: sudo -u postgres psql


You should see a new prompt that looks like this:

postgres=#

Create a new database. I called mine gator:

CREATE DATABASE gator;

Connect to the new database:

\c gator

You should see a new prompt that looks like this:

gator=#

Set the user password (Linux only)

ALTER USER postgres PASSWORD 'postgres';

For simplicity, I used postgres as the password. Before, we altered the system user's password, now we're altering the database user's password.



Query the database

From here you can run SQL queries against the gator database. For example, to see the version of Postgres you're running, you can run:



SELECT version();

If everything is working, you can move on. You can type exit to leave the psql shell.

# Test your connection string by running psql, for example:

psql "protocol://username:password@host:port/database"

cd into the sql/schema directory and run:

goose postgres <connection_string> up

Run your migration! Make sure it works by using psql to find your newly created users table:

psql gator
\dt





