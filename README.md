# ID-Generator 
creates custom id's which have a prefix followed by random characters, they are used at Stripe and allow for the easy identification of the ID's purpose. Documentation for how to use this project is below, you can either use the REST API or use RPC in a microservices architecture.

In the future I will add metric's to the app to track its state whilst its operating.




### ID-Generation Endpoint (REST API)

Used to generate an ID with a specific length and prefix

**URL** : `/id/:length/:prefix`

**Method** : `GET`

**Auth required** : NO

#### Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "id": "prefix_randomString"
}
```

#### Error Response

**Condition** : If you have a length greater than 30 characters or less than 1, 

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "message":  "Failure to generate ID, make sure the length is between 1 and 30 and the prefix is only 3 characters in length.",
}
```


**Condition** : If you provide a prefix which is NOT 3 characters in length, 
**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "message":  "Prefix not 3 characters",
}
```
