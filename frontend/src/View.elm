module View exposing (..)

import Bootstrap.Button as Button exposing (attrs, onClick)
import Bootstrap.Form as Form
import Bootstrap.Form.Input as Input
import Html exposing (Html, div, text)
import Html.Attributes exposing (class)
import Model exposing (..)


view : Model -> Html Msg
view model =
    div []
        [ Form.form [ class "flex-container" ]
            [ Form.group [ Form.attrs [ class "form-group" ] ]
                [ emailGroup
                , passwordGroup
                , submitBtn
                ]
            ]
        ]


emailGroup : Html Msg
emailGroup =
    div []
        [ Form.label [] [ text "Email" ]
        , Input.email
            [ Input.id "email"
            , Input.small
            , Input.onInput EmailInputMsg
            ]
        ]


passwordGroup : Html Msg
passwordGroup =
    div []
        [ Form.label [] [ text "Password" ]
        , Input.text
            [ Input.id "password"
            , Input.small
            , Input.onInput PasswordInputMsg
            ]
        ]


submitBtn : Html Msg
submitBtn =
    Button.button
        [ Button.primary
        , attrs [ class "button" ]
        , onClick SubmitMsg
        ]
        [ text "Submit" ]
