module Quiz2 exposing (main)

-- Bring everything from Html Like import namespace
import Html exposing (..)
import Html.Attributes exposing (class, src)
import Html.Events exposing (onClick)
import Browser


-- Create a Type Msg for Colored and uncolored for star button
type Msg =
    Colored | Uncolor

-- create and Initialize Model to be false
intialModel : { filled : Bool }
intialModel =
    {
        filled = False
    }

-- Create model
viewStar : {  filled : Bool} -> Html Msg
viewStar model =  
    let
        buttonType = 
            if model.filled then "star" else "star_border"
        msg = 
            if model.filled then Uncolor else Colored
    in

     div [ ]
            [
                span [class "material-icons md-100", onClick msg ]
                    [ text buttonType ]
            ]


view : {  filled : Bool} -> Html Msg 
view model =

    div [class "center"]
    [     
        viewStar model
    ]


update : Msg -> {  filled : Bool} -> {  filled : Bool} 
update msg model =
    case msg of
        Colored ->
           {model | filled = True}
        Uncolor ->
           { model | filled = False} 
main : Program () {  filled : Bool} Msg
main =
    Browser.sandbox
    {
        init = intialModel
        ,view = view
        ,update = update
    }
