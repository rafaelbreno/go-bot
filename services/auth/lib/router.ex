defmodule Auth.Router do
  use Plug.Router 

  alias Fields.Plug.VerifyRequest

  plug :match
  plug :dispatch
  plug VerifyRequest, fields: ["username", "password"], paths: ["/auth/create", "/auth/login"]

  plug Plug.Parsers,
    parsers: [:json],
    pass: ["application/json"],
    json_decoder: Poison

  get "/test" do
    conn
    |> put_resp_content_type("application/json")
    |> send_resp(200, Poison.encode!(message()))
  end

  post "/create" do
    conn
    |> put_resp_content_type("application/json")
    |> send_resp(200, Poison.encode!(message()))
  end

  post "/login" do
    conn
    |> put_resp_content_type("application/json")
    |> send_resp(200, Poison.encode!(message()))
  end

  defp message do
    %{
      message: "Salve fi"
    }
  end

  match _ do
    send_resp(conn, 404, "Not Found!")
  end
end
