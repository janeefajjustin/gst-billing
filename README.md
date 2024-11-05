# gst-billing

## Signup API - POST Request : 
For signing up using email and password.
Email ID should be unique for every user.

URL:http://localhost:8080/signup

Expected Input Example:

{
        "Email": "xyz@gmail.com",
        "Password": "bhn"
    }


## Login API - POST Request : 
For logging in using a registered email and password.
Returns token on successful login

URL:http://localhost:8080/login

Expected Input Example:

{
        "Email": "xeryz@gmail.com",
        "Password": "bhvn"
    }



## Add Product API - POST Request:
To add a new product.
Each product should have a unique code.
Only authorized users can access this handler.

URL:http://localhost:8080/products

Example for Input:

{
"Code":7734,
"Name":"Cup Cakes",
"Price":100.56,
"GST":1.3
}



## Search Product API - GET Request

To search a product.
A Product can be searched using its name or its code.
Only authorized users can access this handler.

Example URLs:
URL:http://localhost:8080/products/Cake or URL:http://localhost:8080/products/9734

## Generate Bill API - GET Request

To generate Bill.
Only authorized users can access this handler.
Bill details will be saved each time a bill is generated.

Example URL:http://localhost:8080/billing/Cake/10


## Billing API -GET Request
To get all the billing data.
Only authorized users can access this handler.

URL:http://localhost:8080/billing







