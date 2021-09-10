defmodule Cache do
  use GenServer
  require Logger

  def start_link(url) do
    GenServer.start_link(__MODULE__, {url})
  end

  def init({url}) do
    Logger.info("Connecting to Redis at: #{url}")
    case Redix.start_link(url) do
      {:ok, conn} -> {:ok, conn}
      {:error, err} -> {:error, err}
    end
  end

  def set(conn, key, value) do
    GenServer.call(conn, {:set, key, value})
  end

  def handle_call({:set, key, value}, _from, state ) do
    reply = Redix.command(state, ["SET", key, value])
    {:reply, {:ok, reply}, state}
  end

  def get(conn, key) do
    GenServer.call(conn, {:get, key})
  end

  def handle_call({:get, key}, _from, state ) do
    reply = Redix.command(state, ["GET", key])
    {:reply, {:ok, reply}, state}
  end
end
