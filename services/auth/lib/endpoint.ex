defmodule Auth.Endpoint do
  use Plug.Router

  plug :match
  plug :dispatch

  get "/hello" do
    send_resp(conn, 200, "world!")
  end

  match _ do
    send_resp(conn, 404, "Not Found!")
  end

  def child_spec(opts) do
    %{
      id: __MODULE__,
      start: {__MODULE__, :start_link, [opts]}
    }
  end

  def start_link(_opts),
    do: Plug.Adapters.Cowboy.http(__MODULE__, [])
  
  # it's the same $HIT as
  #def start_link(_opts) do
    #Plug.Adapters.Cowboy.http(__MODULE__, [])
  #end
end
