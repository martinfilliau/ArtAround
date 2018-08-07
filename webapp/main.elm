import Html exposing (..)
import Html.Events exposing (onClick)
import Html.Events exposing (..)
import Http
import Json.Decode as Decode

main =
  Html.program
    { init = init 
    , view = view
    , update = update
    , subscriptions = subscriptions
    }

getStatus: Cmd Msg
getStatus =
  let
    url =
      "http://127.0.0.1:8080/ping"

    request =
      Http.get url decodeStatus 
  in
    Http.send ReceivedUpdatedStatus (Http.get url decodeStatus)

decodeStatus: Decode.Decoder String
decodeStatus=
  Decode.at ["message"] Decode.string

type alias Model =
  { message: String
  }

init : (Model, Cmd Msg)
init =
  (Model "initial", Cmd.none)


-- UPDATE

type Msg = UpdateStatusMessage | ReceivedUpdatedStatus (Result Http.Error String)

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    UpdateStatusMessage ->
      (model, getStatus)
    ReceivedUpdatedStatus (Ok message) ->
        ({ model | message = message}, Cmd.none)
    ReceivedUpdatedStatus (Err _) ->
        ({model| message = "Error" }, Cmd.none)

-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none


-- VIEW

view : Model -> Html Msg
view model =
  div []
    [ h2 [] [text model.message]
    , button [ onClick UpdateStatusMessage ] [ text "Check"]
    ]