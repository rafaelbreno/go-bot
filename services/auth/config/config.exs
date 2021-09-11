import Config

config :postgrex,
  :json_library, Poison

db_name = System.get_env("PGSQL_DBNAME", "project")
db_user = System.get_env("PGSQL_USER", "root")
db_password = System.get_env("PGSQL_PASSWORD", "root")
db_host = System.get_env("PGSQL_HOST", "localhost")

config :auth, Users.Repo,
  database: db_name,
  username: db_user,
  password: db_password,
  hostname: db_host

config :auth,
  ecto_repos: [Users.Repo]

config :auth, Auth.Guardian,
       issuer: "auth",
       secret_key: "zflZDuniaypJu4x/hJKHj3WipCV6ZJIquoarCx6k7qgGKd+fKZ7t2kiygUrL0OGb"
