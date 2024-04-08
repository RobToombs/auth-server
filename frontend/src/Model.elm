module Model exposing (..)


type alias Model =
    { email : String
    , password : String
    }


defaultModel : Model
defaultModel =
    Model "" ""


type Msg
    = EmailInputMsg String
    | PasswordInputMsg String
    | SubmitMsg
