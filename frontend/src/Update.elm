module Update exposing (..)

import Browser.Navigation exposing (load)
import Model exposing (Model, Msg(..))


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        EmailInputMsg email ->
            ( { model | email = email }, Cmd.none )

        PasswordInputMsg password ->
            ( { model | password = password }, Cmd.none )

        SubmitMsg ->
            ( model, load "http://localhost:5173/?code=123-123-123" )
