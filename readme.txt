pre-requisites:

1. install redis on your local server.


To start the service

1. Download and unzip the file.
2. Run webook binary file

Note:
    Maximum Number of attributes and traits can be configured.
    Minimum number of attributes and traits is 1.
    To configure open config.json file and change server port, host, redis config.
    
    To use it your local browser, after starting the service just visit the url 
        http://localhost:9191/swagger/index.html
    click payload module to expand and click on the route to expand, so that payload can be send

    Webook URL where the data is send from webhook app
    https://webhook.site/#!/18fe2e34-338c-4bda-881e-acfe7520d482/72ea7726-134f-4d95-b847-2b314b1c7bf1/1


    After the payload send to the url 
       1. http://localhost:9191/api/v1/uplink
    
    Sample request body:
{
"ev": "contact_form_submitted",
"et": "form_submit",
"id": "cl_app_id_001",
"uid": "cl_app_id_001-uid-001",
"mid": "cl_app_id_001-uid-001",
"t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"l": "en-US",
"sc": "1920 x 1080",
"atrk1": "form_varient",
"atrv1": "red_top",
"atrt1": "string",
"atrk2": "ref",
"atrv2": "XPOWJRICW993LKJD",
"atrt2": "string",
"uatrk1": "name",
"uatrv1": "iron man",
"uatrt1": "string",
"uatrk2": "email",
"uatrv2": "ironman@avengers.com",
"uatrt2": "string",
"uatrk3": "age",
"uatrv3": "32",
"uatrt3": "integer"
}
