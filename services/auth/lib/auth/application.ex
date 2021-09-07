defmodule Auth.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    unless Mix.env == :prod do
      Dotenv.load
      redis_host = System.get_env("REDIS_HOST")
      IO.puts(redis_host)
    end

    children = [
      {Plug.Cowboy, scheme: :http, plug: Auth.Endpoint, options: [port: 4001] },
      # Starts a worker by calling: Auth.Worker.start_link(arg)
      # {Auth.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [
      strategy: :one_for_one, 
      name: Auth.Supervisor
    ]
    Supervisor.start_link(children, opts)
  end
end
